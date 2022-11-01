package service

import (
	"context"

	"github.com/kiaedev/kiae/api/builder"
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

type BuilderSvc struct {
	daoBuilder       *dao.BuilderDao
	imageRegistrySvc *ImageRegistrySvc

	valaApp v1beta1.ApplicationInterface
}

func NewBuilderSvc(daoProjImg *dao.BuilderDao, imageRegistrySvc *ImageRegistrySvc, kClients *klient.LocalClients) *BuilderSvc {
	return &BuilderSvc{
		daoBuilder:       daoProjImg,
		imageRegistrySvc: imageRegistrySvc,

		valaApp: kClients.VelaCs.CoreV1beta1().Applications("kiae-builder"),
	}
}

func (s *BuilderSvc) List(ctx context.Context, in *builder.BuilderListRequest) (*builder.BuilderListResponse, error) {
	query := bson.M{}
	results, total, err := s.daoBuilder.List(ctx, query)
	return &builder.BuilderListResponse{Items: results, Total: total}, err
}

func (s *BuilderSvc) Create(ctx context.Context, in *builder.Builder) (*builder.Builder, error) {
	registry, err := s.imageRegistrySvc.imageRegistryDao.Get(ctx, in.RegistryId)
	if err != nil {
		return nil, err
	}

	vap, err := s.valaApp.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	vap.SetName(in.Name)
	kpackBuilder := components.NewKpackBuilder(in, registry)
	vap.Spec.Components = append(vap.Spec.Components, common.ApplicationComponent{
		Name:       kpackBuilder.GetName(),
		Type:       kpackBuilder.GetType(),
		Properties: util.Object2RawExtension(kpackBuilder),
		Traits:     kpackBuilder.GetTraits(),
	})

	if _, err := s.valaApp.Create(ctx, vap, metav1.CreateOptions{}); err != nil {
		return nil, err
	}

	in.Artifact = kpackBuilder.ImageTag
	return s.daoBuilder.Create(ctx, in)
}

func (s *BuilderSvc) Update(ctx context.Context, in *builder.Builder) (*builder.Builder, error) {
	return s.daoBuilder.Update(ctx, in)
}

func (s *BuilderSvc) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	builderBp, err := s.daoBuilder.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if err := s.valaApp.Delete(ctx, builderBp.Name, metav1.DeleteOptions{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, s.daoBuilder.Delete(ctx, in.Id)
}

func (s *BuilderSvc) SuggestedStacks(ctx context.Context, empty *emptypb.Empty) (*builder.SuggestedStackListResponse, error) {
	return &builder.SuggestedStackListResponse{Items: suggestedStacks}, nil
}

var suggestedStacks = []*builder.SuggestedStack{
	{
		Name:       "Paketo Base",
		Intro:      "A minimal Paketo stack based on Ubuntu 18.04",
		StackId:    "io.buildpacks.stacks.bionic",
		BuildImage: "paketobuildpacks/build:base-cnb",
		RunImage:   "paketobuildpacks/run:base-cnb",
	},
	{
		Name:       "Paketo Full",
		Intro:      "A large Paketo stack based on Ubuntu 18.04",
		StackId:    "io.buildpacks.stacks.bionic",
		BuildImage: "paketobuildpacks/build:full-cnb",
		RunImage:   "paketobuildpacks/run:full-cnb",
	},
	{
		Name:       "Paketo Tiny",
		Intro:      "A tiny Paketo stack based on Ubuntu 18.04, similar to distroless",
		StackId:    "io.paketo.stacks.tiny",
		BuildImage: "paketobuildpacks/build:tiny-cnb",
		RunImage:   "paketobuildpacks/run:tiny-cnb",
	},
	{
		Name:       "Heroku",
		Intro:      "The official Heroku stack based on Ubuntu 20.04",
		StackId:    "heroku-20",
		BuildImage: "heroku/heroku:20-cnb-build",
		RunImage:   "heroku/heroku:20-cnb",
	},
}
