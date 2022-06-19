package service

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"

	"github.com/openkos/openkos/api/gen/go/app"
)

const (
	APPSTORE_REPO_PATH = "../appstore/apps"
)

type AppStore struct {
	app.UnimplementedAppServiceServer

	cs *versioned.Clientset
}

func NewAppStore(cs *versioned.Clientset) *AppStore {
	return &AppStore{cs: cs}
}

func (as *AppStore) Sync(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	// git.PlainClone()
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}

func (as *AppStore) List(ctx context.Context, req *app.ListRequest) (*app.ListResponse, error) {
	f, err := os.Open(APPSTORE_REPO_PATH)
	dirs, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	appItems := make([]*app.Application, 0)
	for _, dirName := range dirs {
		fc, err := ioutil.ReadFile(filepath.Join(APPSTORE_REPO_PATH, dirName, "app.json"))
		if err != nil {
			log.Println(err)
			continue
		}

		var appItem app.Application
		if err := json.Unmarshal(fc, &appItem); err != nil {
			log.Println(err)
			continue
		}
		appItems = append(appItems, &appItem)
	}

	return &app.ListResponse{
		Items: appItems,
	}, nil
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
