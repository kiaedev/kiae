package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/internal/app/server/dao"
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

const (
	APPSTORE_REPO_PATH = "../appstore/apps"
)

type App struct {
	app.UnimplementedAppServiceServer

	daoApp    *dao.AppDao
	daoProj   *dao.ProjectDao
	k8sClient *kubernetes.Clientset
	oamClient *versioned.Clientset
}

func NewAppStore(cs *Service) *App {
	return &App{
		daoApp:    dao.NewApp(cs.DB),
		daoProj:   dao.NewProject(cs.DB),
		k8sClient: cs.K8sClient,
		oamClient: cs.OamClient,
	}
}

func (s *App) Create(ctx context.Context, in *app.Application) (*app.Application, error) {
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

func (s *App) List(ctx context.Context, req *app.ListRequest) (*app.ListResponse, error) {
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

func (s *App) Update(ctx context.Context, in *app.UpdateRequest) (*app.Application, error) {
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
	oApp, err := s.rebuildOApp(ctx, existedApp)
	if err != nil {
		return nil, err
	}

	if _, err := s.oamClient.CoreV1beta1().Applications(kiaeutil.BuildAppNs(existedApp.Env)).Update(ctx, oApp, metav1.UpdateOptions{}); err != nil {
		return nil, status.Errorf(codes.Internal, "update the application failed: %v", err)
	}

	return s.daoApp.Update(ctx, existedApp)
}

func (s *App) handlePatch(in *app.UpdateRequest, existedApp *app.Application) {
	payload := in.Payload
	for _, path := range in.GetUpdateMask().Paths {
		// 只允许运行状态时进行修改，避免状态调整为0又被改回去
		if path == "replicas" && existedApp.Replicas != 0 {
			existedApp.Replicas = payload.Replicas
		}

		if path == "size" {
			existedApp.Size = payload.Size
		}

		if path == "status" {
			if existedApp.Status == app.Status_STATUS_RUNNING && payload.Status == app.Status_STATUS_STOPPED {
				// 停止逻辑: 实例数调到0
				existedApp.PreviousReplicas = existedApp.Replicas
				existedApp.Replicas = 0
			} else if existedApp.Status == app.Status_STATUS_STOPPED && payload.Status == app.Status_STATUS_RUNNING {
				// 启动逻辑：实例数调回停止前
				existedApp.Replicas = existedApp.PreviousReplicas
				existedApp.PreviousReplicas = 0
			} else if existedApp.Status == app.Status_STATUS_RESTARTING {
				// todo 重启逻辑：设置一个Annotation

			}
			existedApp.Status = payload.Status
		}
	}
}

func (s *App) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
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

func (s *App) buildOApp(ctx context.Context, app *app.Application, proj *project.Project) (*v1beta1.Application, error) {
	return render.NewApplication(components.NewKWebservice(app, proj)), nil
}

func (s *App) rebuildOApp(ctx context.Context, app *app.Application) (*v1beta1.Application, error) {
	proj, err := s.daoProj.Get(ctx, app.Pid)
	if err != nil {
		return nil, err
	}

	oApp, err := s.oamClient.CoreV1beta1().Applications(kiaeutil.BuildAppNs(app.Env)).Get(ctx, app.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return render.NewApplicationWith(components.NewKWebservice(app, proj), oApp), nil
}
