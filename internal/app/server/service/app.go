package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/kiae"
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
	k8sClient     *kubernetes.Clientset
	oamClient     *versioned.Clientset
	daoDepend     *dao.DependDao
}

func NewAppService(cs *Service) *AppService {
	return &AppService{
		daoProj:       dao.NewProject(cs.DB),
		daoApp:        dao.NewApp(cs.DB),
		daoDepend:     dao.NewDependDao(cs.DB),
		daoEntry:      dao.NewEntryDao(cs.DB),
		daoRoute:      dao.NewRouteDao(cs.DB),
		daoMwInstance: dao.NewMiddlewareInstanceDao(cs.DB),
		k8sClient:     cs.K8sClient,
		oamClient:     cs.OamClient,
	}
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

	in.Replicas = 2
	in.Name = strings.ToLower(fmt.Sprintf("%s-%s", proj.Name, strutil.RandomText(4)))
	if err := s.createAppComponent(ctx, in, proj); err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	return s.daoApp.Create(ctx, in)
}

func (s *AppService) List(ctx context.Context, req *app.ListRequest) (*app.ListResponse, error) {
	results, total, err := s.daoApp.List(ctx, bson.M{"pid": req.Pid})
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
		// Traits:     buildTraits(traits),
		// DependsOn:  nil,
	})

	if _, err := appCli.Create(ctx, oApp, metav1.CreateOptions{}); err != nil {
		return status.Errorf(codes.Internal, "creating app failed: %v", err)
	}

	return nil
}

func (s *AppService) updateAppComponent(ctx context.Context, app *app.Application) error {
	proj, err := s.daoProj.Get(ctx, app.Pid)
	if err != nil {
		return err
	}

	entries, _, err := s.daoEntry.List(ctx, bson.M{"appid": app.Id, "status": kiae.OpStatus_OP_STATUS_ENABLED})
	if err != nil {
		return err
	}
	routes, _, err := s.daoRoute.List(ctx, bson.M{"appid": app.Id, "status": kiae.OpStatus_OP_STATUS_ENABLED})
	if err != nil {
		return err
	}

	kAppComponent := components.NewKWebservice(app, proj)
	if len(entries) > 0 || len(routes) > 0 {
		kAppComponent.SetupTrait(traits.NewRouteTrait(app.Name, entries, routes))
	}

	return s.updateComponent(ctx, app.Id, kAppComponent)
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
		// Traits:     buildTraits(traits),
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
				// Traits:     buildTraits(traits),
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
