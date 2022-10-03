package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/kiaedev/kiae/api/image"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	alpha2 "github.com/pivotal/kpack/pkg/client/clientset/versioned/typed/build/v1alpha2"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ProjectImageSvc struct {
	image.UnimplementedImageServiceServer

	daoProjImg  *dao.ProjectImageDao
	kPackClient alpha2.KpackV1alpha2Interface
	daoProj     *dao.ProjectDao
}

func NewProjectImageSvc(s *Service) *ProjectImageSvc {
	return &ProjectImageSvc{
		daoProj:     dao.NewProject(s.DB),
		daoProjImg:  dao.NewProjectImageDao(s.DB),
		kPackClient: s.KpackClient.KpackV1alpha2(),
	}
}

func (p *ProjectImageSvc) List(ctx context.Context, in *image.ImageListRequest) (*image.ImageListResponse, error) {
	results, total, err := p.daoProjImg.List(ctx, bson.M{"pid": in.Pid})
	return &image.ImageListResponse{Items: results, Total: total}, err
}

func (p *ProjectImageSvc) Create(ctx context.Context, in *image.Image) (*image.Image, error) {
	proj, err := p.daoProj.Get(ctx, in.Pid)
	if err != nil {
		return nil, err
	}

	_, total, err := p.daoProjImg.List(ctx, bson.M{"pid": in.Pid, "image": in.Image})
	if err != nil {
		return nil, err
	} else if total > 0 {
		return nil, fmt.Errorf("image already exists: %s", in.Image)
	}

	if in.Image == "" {
		in.Image = fmt.Sprintf("%s:%s", proj.ImageRepo, in.Commit[:7])
	}

	imageItems := strings.Split(in.Image, ":")
	tag := "latest"
	if len(imageItems) == 2 {
		tag = imageItems[1]
	}

	in.Tag = tag
	in.Name = fmt.Sprintf("%s-%s", proj.Name, tag)
	imgCli := p.kPackClient.Images("default")
	kImage, err := imgCli.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	} else if err == nil {
		return nil, fmt.Errorf("image %s is already exist", in.Name)
	}

	kImage.SetName(in.Name)
	kImage.Spec.ServiceAccountName = "tutorial-service-account"
	kImage.Spec.Builder.Kind = "Builder"
	kImage.Spec.Builder.Name = "my-builder"
	kImage.Spec.Tag = in.Image
	kImage.Spec.Source.Git = &v1alpha1.Git{URL: ssh2https(proj.GitRepo), Revision: in.Commit}
	if _, err := imgCli.Create(ctx, kImage, metav1.CreateOptions{}); err != nil {
		return nil, err
	}

	return p.daoProjImg.Create(ctx, in)
}

func (p *ProjectImageSvc) Update(ctx context.Context, in *image.Image) (*image.Image, error) {
	return p.daoProjImg.Update(ctx, in)
}

func (s *ProjectImageSvc) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	img, err := s.daoProjImg.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	imgCli := s.kPackClient.Images("default")
	if err := imgCli.Delete(ctx, img.Name, metav1.DeleteOptions{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, s.daoProjImg.Delete(ctx, in.Id)
}

func ssh2https(gitssh string) string {
	gitssh = strings.Replace(gitssh, ":", "/", -1)
	return strings.Replace(gitssh, "git@", "https://", -1)
}
