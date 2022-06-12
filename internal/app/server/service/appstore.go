package service

import (
	"context"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openkos/openkos/api/app"
)

type AppStore struct {
	app.UnimplementedAppServer

	cs *versioned.Clientset
}

func NewAppStore(cs *versioned.Clientset) *AppStore {
	return &AppStore{cs: cs}
}

func (as *AppStore) List(ctx context.Context, req *app.AppListRequest) (*app.AppList, error) {
	_, err := as.cs.CoreV1beta1().Applications("openkos").List(ctx, metav1.ListOptions{})
	return &app.AppList{}, err
}

func (as *AppStore) Install(ctx context.Context, req *app.AppOpRequest) (*app.AppOpReply, error) {
	// app := model.NewApplication()
	// app.Render()
	_, err := as.cs.CoreV1beta1().Applications("openkos").Create(ctx, &v1beta1.Application{}, metav1.CreateOptions{})
	return &app.AppOpReply{}, err
}

func (as *AppStore) Uninstall(ctx context.Context, req *app.AppOpRequest) (*app.AppOpReply, error) {
	name := ""
	err := as.cs.CoreV1beta1().Applications("openkos").Delete(ctx, name, metav1.DeleteOptions{})
	return &app.AppOpReply{}, err
}

func (as *AppStore) Start(ctx context.Context, req *app.AppStatusRequest) (*app.AppStatusReply, error) {
	_, err := as.cs.CoreV1beta1().Applications("openkos").Update(ctx, &v1beta1.Application{}, metav1.UpdateOptions{})
	return &app.AppStatusReply{}, err
}

func (as *AppStore) Stop(ctx context.Context, req *app.AppStatusRequest) (*app.AppStatusReply, error) {
	_, err := as.cs.CoreV1beta1().Applications("openkos").Update(ctx, &v1beta1.Application{}, metav1.UpdateOptions{})
	return &app.AppStatusReply{}, err
}
