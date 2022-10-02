package service

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kiaedev/kiae/api/image"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	alpha2 "github.com/pivotal/kpack/pkg/client/clientset/versioned/typed/build/v1alpha2"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ProjectImageSvc struct {
	image.UnimplementedImageServiceServer

	daoProjImg  *dao.ProjectImageDao
	kPackClient alpha2.KpackV1alpha2Interface
}

func NewProjectImageSvc(s *Service) *ProjectImageSvc {
	return &ProjectImageSvc{
		daoProjImg:  dao.NewProjectImageDao(s.DB),
		kPackClient: s.KpackClient.KpackV1alpha2(),
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

	// todo create an Image resource
	kImage := &v1alpha2.Image{
		Spec:   v1alpha2.ImageSpec{},
		Status: v1alpha2.ImageStatus{},
	}
	p.kPackClient.Images("kiae-system").Create(ctx, kImage, metav1.CreateOptions{})

	return p.daoProjImg.Create(ctx, in)
}

func (p *ProjectImageSvc) Update(ctx context.Context, in *image.Image) (*image.Image, error) {
	return p.daoProjImg.Update(ctx, in)
}

func (s *ProjectImageSvc) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.daoProjImg.Delete(ctx, in.Id)
}
