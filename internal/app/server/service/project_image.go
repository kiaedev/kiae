package service

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kiaedev/kiae/api/image"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProjectImageSvc struct {
	image.UnimplementedImageServiceServer

	daoProjImg *dao.ProjectImageDao
}

func NewProjectImageSvc(s *Service) *ProjectImageSvc {
	return &ProjectImageSvc{
		daoProjImg: dao.NewProjectImageDao(s.DB),
	}
}

func (p *ProjectImageSvc) List(ctx context.Context, in *image.ImageListRequest) (*image.ImageListResponse, error) {
	results, total, err := p.daoProjImg.List(ctx, bson.M{"pid": in.Pid})
	return &image.ImageListResponse{Items: results, Total: total}, err
}

func (p *ProjectImageSvc) Create(ctx context.Context, in *image.Image) (*image.Image, error) {
	_, total, err := p.daoProjImg.List(ctx, bson.M{"pid": in.Pid, "image": in.Image})
	if err != nil {
		return nil, err
	} else if total > 0 {
		return nil, fmt.Errorf("image already exists: %s", in.Image)
	}

	imageItems := strings.Split(in.Image, ":")
	tag := "latest"
	if len(imageItems) == 2 {
		tag = imageItems[1]
	}

	in.Tag = tag
	in.Name = filepath.Base(imageItems[0])
	return p.daoProjImg.Create(ctx, in)
}

func (p *ProjectImageSvc) Update(ctx context.Context, in *image.Image) (*image.Image, error) {
	return p.daoProjImg.Update(ctx, in)
}

func (s *ProjectImageSvc) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.daoProjImg.Delete(ctx, in.Id)
}
