package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/middleware"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/app/server/model"
	"github.com/kiaedev/kiae/internal/pkg/render/components"
	"github.com/kiaedev/kiae/internal/pkg/render/traits"
	"github.com/kiaedev/kiae/pkg/kiaeutil"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"github.com/saltbo/gopkg/strutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type AppService struct {
	app.UnimplementedAppServiceServer

	daoProj       *dao.ProjectDao
	daoApp        *dao.AppDao
	daoEntry      *dao.EntryDao
	daoRoute      *dao.RouteDao
	daoMwInstance *dao.MiddlewareInstance
	daoMwClaim    *dao.MiddlewareClaim
	daoEgress     *dao.EgressDao

	k8sClient *kubernetes.Clientset
	oamClient *versioned.Clientset
}

func NewAppService(daoProj *dao.ProjectDao, daoApp *dao.AppDao, daoEntry *dao.EntryDao, daoRoute *dao.RouteDao, daoMwInstance *dao.MiddlewareInstance, daoMwClaim *dao.MiddlewareClaim, daoEgress *dao.EgressDao, k8sClient *kubernetes.Clientset, oamClient *versioned.Clientset) *AppService {
	return &AppService{daoProj: daoProj, daoApp: daoApp, daoEntry: daoEntry, daoRoute: daoRoute, daoMwInstance: daoMwInstance, daoMwClaim: daoMwClaim, daoEgress: daoEgress, k8sClient: k8sClient, oamClient: oamClient}
}

func (s *AppService) Create(ctx context.Context, in *app.Application) (*app.Application, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, err
	}

	proj, err := s.daoProj.Get(ctx, in.Pid)
	if err == mongo.ErrNoDocuments {
		return nil, status.Errorf(codes.NotFound, "project not found by the pid %v", in.Pid)
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	if _, count, _ := s.daoApp.List(ctx, bson.M{"env": in.Env, "pid": proj.Id}); count > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "该环境已存在")
	}

	s.fillAppDefaultValue(proj, in)
	if err := s.createAppComponent(ctx, in, proj); err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	return s.daoApp.Create(ctx, in)
}

func (s *AppService) List(ctx context.Context, req *app.ListRequest) (*app.ListResponse, error) {
	userid := MustGetUserid(ctx)
	fmt.Println(userid)

	query := make(bson.M)
	if req.Pid != "" {
		query["pid"] = req.Pid
	}

	results, total, err := s.daoApp.List(ctx, query)
	return &app.ListResponse{Items: results, Total: total}, err
}

func (p *AppService) Read(ctx context.Context, in *kiae.IdRequest) (*app.Application, error) {
	kApp, err := p.daoApp.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return kApp, nil
}

func (s *AppService) Update(ctx context.Context, in *app.UpdateRequest) (*app.Application, error) {
	existedApp, err := s.daoApp.Get(ctx, in.Payload.Id)
	if err != nil {
		return nil, err
	}

	if in.UpdateMask == nil {
		existedApp = in.Payload
	} else {
		s.handlePatch(in, existedApp)
	}

	// update the application
	if err := s.updateAppComponent(ctx, existedApp); err != nil {
		return nil, err
	}

	return s.daoApp.Update(ctx, existedApp)
}

func (s *AppService) handlePatch(in *app.UpdateRequest, existedApp *app.Application) {
	payload := in.Payload
	for _, path := range in.GetUpdateMask().Paths {
		if path == "replicas" {
			existedApp.Replicas = payload.Replicas
			// 当在停止状态调整副本数时将状态置为启动
			if existedApp.Replicas > 0 && existedApp.Status == app.Status_STATUS_STOPPED {
				existedApp.Status = app.Status_STATUS_RUNNING
			}
		}

		if path == "size" {
			existedApp.Size = payload.Size
		}
	}
}

func (s *AppService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	rt, err := s.daoApp.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	ns := kiaeutil.BuildAppNs(rt.Env)
	if err := s.oamClient.CoreV1beta1().Applications(ns).Delete(ctx, rt.Name, metav1.DeleteOptions{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, s.daoApp.Delete(ctx, in.Id)
}

func (s *AppService) DoAction(ctx context.Context, in *app.ActionPayload) (*app.Application, error) {
	existedApp, err := s.daoApp.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if in.Action == app.ActionPayload_START && existedApp.Status != app.Status_STATUS_STOPPED {
		return nil, fmt.Errorf("can not start for the not running application - %s", in.Id)
	} else if in.Action == app.ActionPayload_STOP && existedApp.Status != app.Status_STATUS_RUNNING {
		return nil, fmt.Errorf("can not stop for the not stoped application - %s", in.Id)
	} else if in.Action == app.ActionPayload_RESTART && existedApp.Status != app.Status_STATUS_RUNNING {
		return nil, fmt.Errorf("can not restart for the not running application - %s", in.Id)
	}

	// update the application
	if err := s.updateAppComponent(ctx, model.NewAppAction(existedApp).Do(in.Action)); err != nil {
		return nil, err
	}

	return s.daoApp.Update(ctx, existedApp)
}

func (s *AppService) updateStatus(ctx context.Context, ap *app.Application, status app.Status) (*app.Application, error) {
	ap.Status = status
	return s.daoApp.Update(ctx, ap)
}

func (s *AppService) createAppComponent(ctx context.Context, appPb *app.Application, proj *project.Project) error {
	appCli := s.oamClient.CoreV1beta1().Applications(kiaeutil.BuildAppNs(appPb.Env))
	oApp, err := appCli.Get(ctx, appPb.Name, metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	oApp.SetName(appPb.Name)
	coreComponent := components.NewKWebservice(appPb, proj)
	oApp.Spec.Components = append(oApp.Spec.Components, common.ApplicationComponent{
		Name:       coreComponent.GetName(),
		Type:       coreComponent.GetType(),
		Properties: util.Object2RawExtension(coreComponent),
		Traits:     coreComponent.GetTraits(),
		// DependsOn:  nil,
	})

	if _, err := appCli.Create(ctx, oApp, metav1.CreateOptions{}); err != nil {
		return status.Errorf(codes.Internal, "creating app failed: %v", err)
	}

	return nil
}

func (s *AppService) updateAppComponent(ctx context.Context, ap *app.Application) error {
	proj, err := s.daoProj.Get(ctx, ap.Pid)
	if err != nil {
		return err
	}

	entries, _, err := s.daoEntry.List(ctx, bson.M{"appid": ap.Id, "status": kiae.OpStatus_OP_STATUS_ENABLED})
	if err != nil {
		return err
	}
	routes, _, err := s.daoRoute.List(ctx, bson.M{"appid": ap.Id, "status": kiae.OpStatus_OP_STATUS_ENABLED})
	if err != nil {
		return err
	}

	kAppComponent := components.NewKWebservice(ap, proj)
	if len(entries) > 0 || len(routes) > 0 {
		kAppComponent.SetupTrait(traits.NewRouteTrait(ap.Name, entries, routes))
	}
	if len(ap.Configs) > 0 {
		kAppComponent.SetupTrait(traits.NewConfigsTrait(ap.Configs))
	}

	mwClaims, _, err := s.daoMwClaim.List(ctx, bson.M{"appid": ap.Id, "status": kiae.OpStatus_OP_STATUS_ENABLED})
	if err != nil {
		return err
	}

	for _, dd := range mwClaims {
		if dd.Status == middleware.Claim_BOUND {
			kAppComponent.SetupTrait(traits.NewSecret2File(dd.Name))
		}
	}

	egresses, _, err := s.daoEgress.List(ctx, bson.M{"appid": ap.Id})
	if err != nil {
		return err
	}
	if len(egresses) > 0 {
		kAppComponent.SetupTrait(traits.NewSidecar(egresses))
		kAppComponent.SetupTrait(traits.NewServiceEntry(egresses))
	}

	return s.updateComponent(ctx, ap.Id, kAppComponent)
}

func (s *AppService) updateAppComponentById(ctx context.Context, appid string) error {
	aa, err := s.daoApp.Get(ctx, appid)
	if err != nil {
		return err
	}

	return s.updateAppComponent(ctx, aa)
}

func (s *AppService) addComponent(ctx context.Context, appid string, com components.Component) error {
	appPb, err := s.daoApp.Get(ctx, appid)
	if err != nil {
		return err
	}

	appCli := s.oamClient.CoreV1beta1().Applications(kiaeutil.BuildAppNs(appPb.Env))
	oApp, err := appCli.Get(ctx, appPb.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	for _, component := range oApp.Spec.Components {
		if component.Type == com.GetType() && component.Name == com.GetName() {
			return fmt.Errorf("component [%s]%s already exists", com.GetType(), com.GetName())
		}
	}

	oApp.Spec.Components = append(oApp.Spec.Components, common.ApplicationComponent{
		Name:       com.GetName(),
		Type:       com.GetType(),
		Properties: util.Object2RawExtension(com),
		Traits:     com.GetTraits(),
		// DependsOn:  nil,
	})
	_, err = appCli.Update(ctx, oApp, metav1.UpdateOptions{})
	return err
}

func (s *AppService) updateComponent(ctx context.Context, appid string, com components.Component) error {
	appPb, err := s.daoApp.Get(ctx, appid)
	if err != nil {
		return err
	}

	appCli := s.oamClient.CoreV1beta1().Applications(kiaeutil.BuildAppNs(appPb.Env))
	oApp, err := appCli.Get(ctx, appPb.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	for idx, component := range oApp.Spec.Components {
		if component.Type == com.GetType() && component.Name == com.GetName() {
			oApp.Spec.Components[idx] = common.ApplicationComponent{
				Name:       com.GetName(),
				Type:       com.GetType(),
				Properties: util.Object2RawExtension(com),
				Traits:     com.GetTraits(),
				// DependsOn:  nil,
			}
			break
		}
	}

	_, err = appCli.Update(ctx, oApp, metav1.UpdateOptions{})
	return err
}

func (s *AppService) removeComponent(ctx context.Context, appid string, com components.Component) error {
	appPb, err := s.daoApp.Get(ctx, appid)
	if err != nil {
		return err
	}

	appCli := s.oamClient.CoreV1beta1().Applications(kiaeutil.BuildAppNs(appPb.Env))
	oApp, err := appCli.Get(ctx, appPb.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	for idx, component := range oApp.Spec.Components {
		if component.Type == com.GetType() && component.Name == com.GetName() {
			oApp.Spec.Components = append(oApp.Spec.Components[:idx], oApp.Spec.Components[idx+1:]...)
			_, err = appCli.Update(ctx, oApp, metav1.UpdateOptions{})
			return err
		}
	}

	return fmt.Errorf("not found component [%s]%s", com.GetType(), com.GetName())
}

func (s *AppService) fillAppDefaultValue(proj *project.Project, in *app.Application) {
	in.Replicas = 2
	in.Name = strings.ToLower(fmt.Sprintf("%s-%s", proj.Name, strutil.RandomText(4)))
	in.Environments = defaultEnvironments(proj, in)
	defaultPort := uint32(8000)
	if in.Ports != nil && len(in.Ports) > 0 {
		defaultPort = in.Ports[0].Port
	}
	in.ReadinessProbe = defaultHealthProbe(defaultPort, "/healthz")
	in.LivenessProbe = defaultHealthProbe(defaultPort, "/healthz")
}

func defaultEnvironments(proj *project.Project, in *app.Application) []*app.Environment {
	if in.Environments == nil {
		in.Environments = make([]*app.Environment, 0)
	}

	systemEnvs := map[string]string{
		"KIAE_ENV":       in.Env,
		"KIAE_HOME":      "/kiae/home/",
		"KIAE_LOG_PATH":  "/kiae/logs/",
		"KIAE_CFG_PATH":  "/kiae/configs/",
		"KIAE_PROJ_NAME": proj.Name,
		"KIAE_APP_NAME":  in.Name,
	}

	for k, v := range systemEnvs {
		in.Environments = append(in.Environments, &app.Environment{
			Type:      app.Environment_SYSTEM,
			Name:      k,
			Value:     v,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		})
	}

	return in.Environments
}

func defaultHealthProbe(port uint32, path string) *app.HealthProbe {
	return &app.HealthProbe{
		Port:                port,
		Path:                path,
		PeriodSeconds:       30,
		TimeoutSeconds:      3,
		SuccessThreshold:    1,
		FailureThreshold:    3,
		InitialDelaySeconds: 5,
	}
}
