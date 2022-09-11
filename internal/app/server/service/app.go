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
	"github.com/kiaedev/kiae/internal/pkg/render"
	"github.com/kiaedev/kiae/internal/pkg/render/components"
	"github.com/kiaedev/kiae/pkg/kiaeutil"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	"github.com/saltbo/gopkg/strutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type AppService struct {
	app.UnimplementedAppServiceServer

	daoProj   *dao.ProjectDao
	daoApp    *dao.AppDao
	daoEntry  *dao.EntryDao
	daoRoute  *dao.RouteDao
	k8sClient *kubernetes.Clientset
	oamClient *versioned.Clientset
}

func NewAppService(cs *Service) *AppService {
	return &AppService{
		daoProj:   dao.NewProject(cs.DB),
		daoApp:    dao.NewApp(cs.DB),
		daoEntry:  dao.NewEntryDao(cs.DB),
		daoRoute:  dao.NewRouteDao(cs.DB),
		k8sClient: cs.K8sClient,
		oamClient: cs.OamClient,
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

	// traits := []*kiae.Trait{}

	in.Replicas = 2
	in.Name = strings.ToLower(fmt.Sprintf("%s-%s", proj.Name, strutil.RandomText(4)))
	ns := kiaeutil.BuildAppNs(in.Env)
	oApp, err := s.buildOApp(ctx, in, proj)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	if _, err := s.oamClient.CoreV1beta1().Applications(ns).Create(ctx, oApp, metav1.CreateOptions{}); err != nil {
		return nil, status.Errorf(codes.Internal, "creating app failed: %v", err)
	}

	return s.daoApp.Create(ctx, in)
}

func (s *AppService) List(ctx context.Context, req *app.ListRequest) (*app.ListResponse, error) {
	proj, err := s.daoProj.Get(ctx, req.Pid)
	if err != nil {
		return nil, err
	}

	results, total, err := s.daoApp.List(ctx, bson.M{"pid": req.Pid})
	for _, rt := range results {
		rt.Configs = kiaeutil.ConfigsMerge(proj.Configs, rt.Configs)
	}

	return &app.ListResponse{Items: results, Total: total}, err
}

func (p *AppService) Read(ctx context.Context, in *kiae.IdRequest) (*app.Application, error) {
	kApp, err := p.daoApp.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	proj, err := p.daoProj.Get(ctx, kApp.Pid)
	if err != nil {
		return nil, err
	}

	kApp.Configs = kiaeutil.ConfigsMerge(proj.Configs, kApp.Configs)
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
	if err := s.UpdateAllComponents(ctx, existedApp); err != nil {
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
	if err := s.UpdateAllComponents(ctx, model.NewAppAction(existedApp).Do(in.Action)); err != nil {
		return nil, err
	}

	return s.daoApp.Update(ctx, existedApp)
}

func (s *AppService) buildOApp(ctx context.Context, app *app.Application, proj *project.Project) (*v1beta1.Application, error) {
	return render.NewApplication(app.Name, components.NewKWebservice(app, proj)), nil
}

func (s *AppService) rebuildOApp(ctx context.Context, app *app.Application) (*v1beta1.Application, error) {
	proj, err := s.daoProj.Get(ctx, app.Pid)
	if err != nil {
		return nil, err
	}

	entries, _, err := s.daoEntry.List(ctx, bson.M{"appid": app.Id, "status": kiae.OpStatus_OP_STATUS_ENABLED})
	if err != nil {
		return nil, err
	}
	routes, _, err := s.daoRoute.List(ctx, bson.M{"appid": app.Id, "status": kiae.OpStatus_OP_STATUS_ENABLED})
	if err != nil {
		return nil, err
	}

	oApp, err := s.oamClient.CoreV1beta1().Applications(kiaeutil.BuildAppNs(app.Env)).Get(ctx, app.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	coms := make([]render.Component, 0)
	coms = append(coms, components.NewKWebservice(app, proj))
	if len(entries) > 0 || len(routes) > 0 {
		coms = append(coms, components.NewRouteComponent(app.Name, entries, routes))
	}

	fmt.Println(coms)
	return render.NewApplicationWith(oApp, coms...), nil
}

func (s *AppService) UpdateAllComponents(ctx context.Context, app *app.Application) error {
	oApp, err := s.rebuildOApp(ctx, app)
	if err != nil {
		return err
	}

	if _, err := s.oamClient.CoreV1beta1().Applications(kiaeutil.BuildAppNs(app.Env)).Update(ctx, oApp, metav1.UpdateOptions{}); err != nil {
		return status.Errorf(codes.Internal, "update the application failed: %v", err)
	}

	return nil
}

func (s *AppService) UpdateAllComponentsByAppid(ctx context.Context, appid string) error {
	aa, err := s.daoApp.Get(ctx, appid)
	if err != nil {
		return err
	}

	return s.UpdateAllComponents(ctx, aa)
}
