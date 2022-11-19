package service

import (
	"context"

	"github.com/kiaedev/kiae/api/gateway"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/pkg/klient"
	"github.com/kiaedev/kiae/internal/pkg/velarender/components"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned/typed/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Gateway struct {
	daoGateway *dao.Gateway

	velaApp v1beta1.ApplicationInterface
}

func NewGateway(daoGateway *dao.Gateway, kClients *klient.LocalClients) *Gateway {
	return &Gateway{
		daoGateway: daoGateway,

		velaApp: kClients.VelaCs.CoreV1beta1().Applications("kiae-system"),
	}
}

func (s *Gateway) List(ctx context.Context, in *gateway.ListRequest) (*gateway.ListResponse, error) {
	results, total, err := s.daoGateway.List(ctx, bson.M{})
	return &gateway.ListResponse{Items: results, Total: total}, err
}

func (s *Gateway) Create(ctx context.Context, in *gateway.Gateway) (*gateway.Gateway, error) {
	vap, err := s.velaApp.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	vap.SetName(in.Name)
	istioGateway := components.NewIstioGateway(in)
	vap.Spec.Components = append(vap.Spec.Components, common.ApplicationComponent{
		Name:       istioGateway.GetName(),
		Type:       istioGateway.GetType(),
		Properties: util.Object2RawExtension(istioGateway),
	})

	if _, err := s.velaApp.Create(ctx, vap, metav1.CreateOptions{}); err != nil {
		return nil, err
	}

	return s.daoGateway.Create(ctx, in)
}

func (s *Gateway) Update(ctx context.Context, in *gateway.UpdateRequest) (*gateway.Gateway, error) {
	// todo implement me

	return s.daoGateway.Update(ctx, in.Payload)
}

func (s *Gateway) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	gatewayBp, err := s.daoGateway.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if err := s.velaApp.Delete(ctx, gatewayBp.Name, metav1.DeleteOptions{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, s.daoGateway.Delete(ctx, in.Id)
}
