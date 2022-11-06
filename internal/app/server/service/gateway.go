package service

import (
	"context"

	"github.com/kiaedev/kiae/api/gateway"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Gateway struct {
	daoGateway *dao.Gateway
}

func NewGateway(daoGateway *dao.Gateway) *Gateway {
	return &Gateway{daoGateway: daoGateway}
}

func (s *Gateway) List(ctx context.Context, in *gateway.ListRequest) (*gateway.ListResponse, error) {
	results, total, err := s.daoGateway.List(ctx, bson.M{})
	return &gateway.ListResponse{Items: results, Total: total}, err
}

func (s *Gateway) Create(ctx context.Context, in *gateway.Gateway) (*gateway.Gateway, error) {
	// TODO: create a Gateway CR

	return s.daoGateway.Create(ctx, in)
}

func (s *Gateway) Update(ctx context.Context, in *gateway.UpdateRequest) (*gateway.Gateway, error) {
	return s.daoGateway.Update(ctx, in.Payload)
}

func (s *Gateway) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	_, err := s.daoGateway.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	// TODO: remove the Gateway

	return &emptypb.Empty{}, s.daoGateway.Delete(ctx, in.Id)
}
