package service

import (
	"context"

	"github.com/kiaedev/kiae/api/egress"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/pkg/kcs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EgressService struct {
	egress.UnimplementedEgressServiceServer

	appSvc    *AppService
	daoEgress *dao.EgressDao
}

func NewEgressService(db *mongo.Database, kClients *kcs.KubeClients) *EgressService {
	return &EgressService{
		appSvc:    NewAppService(db, kClients),
		daoEgress: dao.NewEgressDao(db),
	}
}

func (s *EgressService) List(ctx context.Context, in *egress.ListRequest) (*egress.ListResponse, error) {
	results, total, err := s.daoEgress.List(ctx, bson.M{"appid": in.Appid})
	return &egress.ListResponse{Items: results, Total: total}, err
}

func (s *EgressService) Create(ctx context.Context, in *egress.Egress) (*egress.Egress, error) {
	_, err := s.appSvc.daoApp.Get(ctx, in.Appid)
	if err != nil {
		return nil, err
	}

	eg, err := s.daoEgress.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	return eg, s.appSvc.updateAppComponentById(ctx, in.Appid)
}

func (s *EgressService) Update(ctx context.Context, in *egress.Egress) (*egress.Egress, error) {
	return s.daoEgress.Update(ctx, in)
}

func (s *EgressService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	eg, err := s.daoEgress.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if err := s.daoEgress.Delete(ctx, in.Id); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, s.appSvc.updateAppComponentById(ctx, eg.Appid)
}
