package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/pkg/templates"
	"github.com/kiaedev/kiae/pkg/kiaeutil"
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	"github.com/saltbo/gopkg/strutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "k8s.io/api/core/v1"
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

	traits := []*kiae.Trait{}

	in.Replicas = 2
	in.Name = strings.ToLower(fmt.Sprintf("%s-%s", proj.Name, strutil.RandomText(4)))
	oApp, err := templates.NewApplication(in, proj, traits)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "rendering app failed: %v", err)
	}

	ns := kiaeutil.BuildAppNs(in.Env)
	if _, err := s.k8sClient.CoreV1().ConfigMaps(ns).Create(ctx, buildConfigs(proj, in), metav1.CreateOptions{}); err != nil {
		return nil, status.Errorf(codes.Internal, "creating app-config failed: %v", err)
	}

	if _, err := s.oamClient.CoreV1beta1().Applications(ns).Create(ctx, oApp, metav1.CreateOptions{}); err != nil {
		return nil, status.Errorf(codes.Internal, "creating app failed: %v", err)
	}

	return s.daoApp.Create(ctx, in)
}

func buildConfigs(proj *project.Project, app *app.Application) *v1.ConfigMap {
	configs := kiaeutil.ConfigsMerge(proj.Configs, app.Configs)
	data := make(map[string]string)

	for _, config := range configs {
		data[config.Filename] = config.Content
	}
	// todo 考虑一下是否有必要把不同MountPath的配置创建单独的ConfigMap

	cm := &v1.ConfigMap{
		Data: data,
	}
	cm.SetName("kiaeapp-" + app.Name)
	return cm
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

// func (as *App) Start(ctx context.Context, req *app.AppStatusRequest) (*app.AppStatusReply, error) {
// 	result := new(app.Application)
// 	if err := as.collection.FindOneAndDelete(ctx, bson.M{"id": req.Id}).Decode(result); err != nil {
// 		return nil, err
// 	}
//
// 	m, err := templates.NewApplication(result, nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	_, err = as.oamClientSet.CoreV1beta1().Applications(result.Env).Update(ctx, m, metav1.UpdateOptions{})
// 	return &app.AppStatusReply{}, err
// }
//
// func (as *App) Stop(ctx context.Context, req *app.AppStatusRequest) (*app.AppStatusReply, error) {
// 	rt := new(app.Application)
// 	if err := as.collection.FindOneAndDelete(ctx, bson.M{"id": req.Id}).Decode(rt); err != nil {
// 		return nil, err
// 	}
//
// 	_, err := as.oamClientSet.CoreV1beta1().Applications(rt.Env).Update(ctx, &v1beta1.Application{}, metav1.UpdateOptions{})
// 	return &app.AppStatusReply{}, err
// }

func (s *App) Delete(ctx context.Context, in *kiae.DeleteRequest) (*emptypb.Empty, error) {
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
