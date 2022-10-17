package service

import (
	"context"

	"github.com/kiaedev/kiae/api/builder"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/pkg/klient"
	"github.com/kiaedev/kiae/internal/pkg/render/components"
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

	valaClient v1beta1.CoreV1beta1Interface
}

func NewBuilderSvc(daoProjImg *dao.BuilderDao, kClients *klient.LocalClients) *BuilderSvc {
	return &BuilderSvc{
		daoBuilder: daoProjImg,
		valaClient: kClients.VelaCs.CoreV1beta1(),
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

	cli := s.valaClient.Applications("kiae-system")
	vap, err := cli.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	vap.SetName(in.Name)
	coreComponent := components.NewKpackBuilder(in, registry)
	vap.Spec.Components = append(vap.Spec.Components, common.ApplicationComponent{
		Name:       coreComponent.GetName(),
		Type:       coreComponent.GetType(),
		Properties: util.Object2RawExtension(coreComponent),
		Traits:     coreComponent.GetTraits(),
	})

	if _, err := cli.Create(ctx, vap, metav1.CreateOptions{}); err != nil {
		return nil, err
	}

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

	cli := s.valaClient.Applications("kiae-system")
	if err := cli.Delete(ctx, builderBp.Name, metav1.DeleteOptions{}); err != nil {
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
