package service

import (
	"context"

	"github.com/kiaedev/kiae/api/image"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ImageRegistrySvc struct {
	imageRegistryDao *dao.ImageRegistryDao
}

func NewImageRegistrySvc(imageRegistryDao *dao.ImageRegistryDao) *ImageRegistrySvc {
	return &ImageRegistrySvc{imageRegistryDao: imageRegistryDao}
}

func (s *ImageRegistrySvc) List(ctx context.Context, in *image.RegistryListRequest) (*image.RegistryListResponse, error) {
	query := bson.M{}
	results, total, err := s.imageRegistryDao.List(ctx, query)
	return &image.RegistryListResponse{Items: results, Total: total}, err
}

func (s *ImageRegistrySvc) Create(ctx context.Context, in *image.Registry) (*image.Registry, error) {

	return s.imageRegistryDao.Create(ctx, in)
}

func (s *ImageRegistrySvc) Update(ctx context.Context, in *image.Registry) (*image.Registry, error) {
	return s.imageRegistryDao.Update(ctx, in)
}

func (s *ImageRegistrySvc) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	_, err := s.imageRegistryDao.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, s.imageRegistryDao.Delete(ctx, in.Id)
}
