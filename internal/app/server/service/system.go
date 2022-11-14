package service

import (
	"context"

	"github.com/kiaedev/kiae/api/system"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type System struct {
	daoProvider *dao.ProviderDao
	daoRegistry *dao.ImageRegistryDao
	daoBuilder  *dao.BuilderDao
}

func NewSystem(daoProvider *dao.ProviderDao, daoRegistry *dao.ImageRegistryDao, daoBuilder *dao.BuilderDao) *System {
	return &System{daoProvider: daoProvider, daoRegistry: daoRegistry, daoBuilder: daoBuilder}
}

func (s *System) GetStatus(ctx context.Context, empty *emptypb.Empty) (*system.SystemStatus, error) {
	_, providerNum, _ := s.daoProvider.List(ctx, bson.M{})
	_, registryNum, _ := s.daoRegistry.List(ctx, bson.M{})
	_, builderNum, _ := s.daoBuilder.List(ctx, bson.M{})
	return &system.SystemStatus{
		Ready: providerNum > 0 && registryNum > 0 && builderNum > 0,
	}, nil
}
