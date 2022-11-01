package service

import (
	"context"
	"fmt"

	"github.com/kiaedev/kiae/api/image"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/project"
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

type ProjectImageSvc struct {
	daoProj    *dao.ProjectDao
	daoProjImg *dao.ProjectImageDao
	daoBuilder *dao.BuilderDao

	svcRegistry *ImageRegistrySvc
	svcProvider *ProviderService
	velaApp     v1beta1.ApplicationInterface
}

func NewProjectImageSvc(daoProj *dao.ProjectDao, daoProjImg *dao.ProjectImageDao, daoBuilder *dao.BuilderDao,
	svcRegistry *ImageRegistrySvc, svcProvider *ProviderService, kClients *klient.LocalClients) *ProjectImageSvc {
	return &ProjectImageSvc{
		daoProj:     daoProj,
		daoProjImg:  daoProjImg,
		daoBuilder:  daoBuilder,
		svcRegistry: svcRegistry,
		svcProvider: svcProvider,

		velaApp: kClients.VelaCs.CoreV1beta1().Applications("kiae-builder"),
	}
}

func (s *ProjectImageSvc) List(ctx context.Context, in *image.ImageListRequest) (*image.ImageListResponse, error) {
	query := bson.M{"pid": in.Pid}
	if in.Status > image.Image_UNSPECIFIED {
		query["status"] = in.Status
	}

	results, total, err := s.daoProjImg.List(ctx, query)
	return &image.ImageListResponse{Items: results, Total: total}, err
}

func (s *ProjectImageSvc) Create(ctx context.Context, in *image.Image) (*image.Image, error) {
	proj, err := s.daoProj.Get(ctx, in.Pid)
	if err != nil {
		return nil, err
	}

	_, total, err := s.daoProjImg.List(ctx, bson.M{"pid": in.Pid, "image": in.Image})
	if err != nil {
		return nil, err
	} else if total > 0 {
		return nil, fmt.Errorf("image already exists: %s", in.Image)
	}

	// for import a image from external hub
	if in.Image != "" {
		in.SetImage(in.Image)
		goto DIRECT
	}

	// create a new image from the git source repository
	if err := s.buildNewImage(ctx, in, proj); err != nil {
		return nil, err
	}

DIRECT:
	return s.daoProjImg.Create(ctx, in)
}

func (s *ProjectImageSvc) buildNewImage(ctx context.Context, in *image.Image, proj *project.Project) error {
	if _, err := s.svcProvider.getProviderToken(ctx, proj.GitProvider); err != nil {
		return err
	}

	builder, err := s.daoBuilder.Get(ctx, proj.BuilderId)
	if err != nil {
		return fmt.Errorf("failed to get build builder: %v", err)
	}

	imgReg, err := s.svcRegistry.imageRegistryDao.Get(ctx, builder.RegistryId)
	if err != nil {
		return err
	}

	in.SetImage(imgReg.BuildImageWithTag(proj.Name, in.CommitId[:7]))
	vap, err := s.velaApp.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	} else if err == nil {
		return fmt.Errorf("image %s is already exist", in.Name)
	}

	vap.SetName(in.Name)
	tokenSecretName := TokenSecretName(ctx, proj.GitProvider)
	kpackImage := components.NewKpackImage(in, builder.Name, proj.GitHTTPSUrl(), imgReg.GetSecretName(), tokenSecretName)
	vap.Spec.Components = append(vap.Spec.Components, common.ApplicationComponent{
		Name:       kpackImage.GetName(),
		Type:       kpackImage.GetType(),
		Properties: util.Object2RawExtension(kpackImage),
		Traits:     kpackImage.GetTraits(),
	})
	_, err = s.velaApp.Create(ctx, vap, metav1.CreateOptions{})
	return err
}

func (s *ProjectImageSvc) Update(ctx context.Context, in *image.Image) (*image.Image, error) {
	return s.daoProjImg.Update(ctx, in)
}

func (s *ProjectImageSvc) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	img, err := s.daoProjImg.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if err := s.velaApp.Delete(ctx, img.Name, metav1.DeleteOptions{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, s.daoProjImg.Delete(ctx, in.Id)
}

func (s *ProjectImageSvc) UpdateStatus(ctx context.Context, name string, status image.Image_Status) (*image.Image, error) {
	img, err := s.daoProjImg.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	img.Status = status
	return s.daoProjImg.Update(ctx, img)
}

func (s *ProjectImageSvc) ListNotDoneStatus(ctx context.Context) ([]*image.Image, error) {
	results, _, err := s.daoProjImg.List(ctx, bson.M{"$nor": bson.A{
		bson.M{"status": image.Image_PUBLISHED}, bson.M{"status": image.Image_FAILED}},
	})
	return results, err
}
