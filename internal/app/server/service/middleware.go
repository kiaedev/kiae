package service

import (
	"context"

	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/middleware"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	mw_provider "github.com/kiaedev/kiae/internal/pkg/mw-provider"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MiddlewareService struct {
	middleware.UnimplementedMiddlewareServiceServer

	rc client.Client
	kc *kubernetes.Clientset

	daoMwInstance *dao.MiddlewareInstance
}

func NewMiddlewareService(cs *Service) *MiddlewareService {
	return &MiddlewareService{
		kc:            cs.K8sClient,
		rc:            cs.RuntimeClient,
		daoMwInstance: dao.NewMiddlewareInstanceDao(cs.DB),
	}
}

func (s *MiddlewareService) List(ctx context.Context, in *middleware.ListRequest) (*middleware.ListResponse, error) {
	query := make(bson.M)
	if in.Type != "" {
		query["type"] = in.Type
	}

	results, total, err := s.daoMwInstance.List(ctx, query)
	return &middleware.ListResponse{Items: results, Total: total}, err
}

func (s *MiddlewareService) Create(ctx context.Context, in *middleware.Instance) (*middleware.Instance, error) {
	secretName := "db-admin-conn-" + in.Name
	secret := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: secretName}, StringData: in.Properties}
	_, err := s.kc.CoreV1().Secrets("kiae-system").Create(ctx, secret, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	object := mw_provider.BuildConfig(in.Type, in.Name, secretName)
	if err := s.rc.Create(ctx, object, &client.CreateOptions{}); err != nil {
		return nil, err
	}

	return s.daoMwInstance.Create(ctx, in)
}

func (s *MiddlewareService) Update(ctx context.Context, in *middleware.Instance) (*middleware.Instance, error) {
	return s.daoMwInstance.Update(ctx, in)
}

func (s *MiddlewareService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	mwc, err := s.daoMwInstance.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	secretName := "db-admin-conn-" + mwc.Name
	if err := s.kc.CoreV1().Secrets("kiae-system").Delete(ctx, secretName, metav1.DeleteOptions{}); err != nil {
		return nil, err
	}

	object := mw_provider.BuildConfig(mwc.Type, mwc.Name, secretName)
	if err := s.rc.Delete(ctx, object, &client.DeleteOptions{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, s.daoMwInstance.Delete(ctx, in.Id)
}
