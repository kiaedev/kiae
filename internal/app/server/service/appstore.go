package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/api/trait"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/pkg/templates"
	"github.com/kiaedev/kiae/pkg/kiae"
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	"github.com/saltbo/gopkg/strutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	APPSTORE_REPO_PATH = "../appstore/apps"
)

type AppStore struct {
	app.UnimplementedAppServiceServer

	daoApp    *dao.AppDao
	daoProj   *dao.ProjectDao
	k8sClient *kubernetes.Clientset
	oamClient *versioned.Clientset
}

func NewAppStore(cs *Service) *AppStore {
	return &AppStore{
		daoApp:    dao.NewApp(cs.DB),
		daoProj:   dao.NewProject(cs.DB),
		k8sClient: cs.K8sClient,
		oamClient: cs.OamClient,
	}
}

func (as *AppStore) Create(ctx context.Context, in *app.Application) (*app.Application, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, err
	}

	proj, err := as.daoProj.Get(ctx, in.Pid)
	if err == mongo.ErrNoDocuments {
		return nil, status.Errorf(codes.NotFound, "project not found by the pid %v", in.Pid)
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	if _, count, _ := as.daoApp.List(ctx, bson.M{"env": in.Env, "pid": proj.Id}); count > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "该环境已存在")
	}

	traits := []*trait.Trait{}

	in.Name = strings.ToLower(fmt.Sprintf("%s-%s", proj.Name, strutil.RandomText(4)))
	oApp, err := templates.NewApplication(in, proj, traits)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "rendering app failed: %v", err)
	}

	ns := kiae.BuildAppNs(in.Env)
	if _, err := as.k8sClient.CoreV1().ConfigMaps(ns).Create(ctx, buildConfigs(proj, in), metav1.CreateOptions{}); err != nil {
		return nil, status.Errorf(codes.Internal, "creating app-config failed: %v", err)
	}

	if _, err := as.oamClient.CoreV1beta1().Applications(ns).Create(ctx, oApp, metav1.CreateOptions{}); err != nil {
		return nil, status.Errorf(codes.Internal, "creating app failed: %v", err)
	}

	return as.daoApp.Create(ctx, in)
}

func buildConfigs(proj *project.Project, app *app.Application) *v1.ConfigMap {
	configs := kiae.ConfigsMerge(proj.Configs, app.Configs)
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

func (as *AppStore) List(ctx context.Context, req *app.ListRequest) (*app.ListResponse, error) {
	results, total, err := as.daoApp.List(ctx, bson.M{"pid": req.Pid})
	return &app.ListResponse{Items: results, Total: total}, err
}

func (as *AppStore) Install(ctx context.Context, req *app.AppOpRequest) (*app.AppOpReply, error) {
	// item, exist := as.apps.Load(req.Name)
	// if !exist {
	// 	return nil, fmt.Errorf("app %s not exist", req.Name)
	// }

	// appProto := item.(*app.Application)
	// buf, err := template.Render("app", appProto)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// var application v1beta1.Application
	// if err := yaml.Unmarshal(buf.Bytes(), &application); err != nil {
	// 	return nil, err
	// }

	// _, err = as.cs.CoreV1beta1().Applications("kiae").Create(ctx, &application, metav1.CreateOptions{})
	return &app.AppOpReply{}, nil
}

// func (as *AppStore) Start(ctx context.Context, req *app.AppStatusRequest) (*app.AppStatusReply, error) {
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
// func (as *AppStore) Stop(ctx context.Context, req *app.AppStatusRequest) (*app.AppStatusReply, error) {
// 	rt := new(app.Application)
// 	if err := as.collection.FindOneAndDelete(ctx, bson.M{"id": req.Id}).Decode(rt); err != nil {
// 		return nil, err
// 	}
//
// 	_, err := as.oamClientSet.CoreV1beta1().Applications(rt.Env).Update(ctx, &v1beta1.Application{}, metav1.UpdateOptions{})
// 	return &app.AppStatusReply{}, err
// }

func (as *AppStore) Delete(ctx context.Context, req *app.AppOpRequest) (*app.AppOpReply, error) {
	// rt := new(app.Application)
	// if err := as.collection.FindOneAndDelete(ctx, bson.M{"id": req.Id}).Decode(rt); err != nil {
	// 	return nil, err
	// }

	// err := as.oamClientSet.CoreV1beta1().Applications(rt.Env).Delete(ctx, rt.Name, metav1.DeleteOptions{})
	return &app.AppOpReply{}, nil
}
