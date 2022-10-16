package service

import (
	"context"

	"github.com/kiaedev/kiae/api/builder"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/pkg/klient"
	alpha2 "github.com/pivotal/kpack/pkg/client/clientset/versioned/typed/build/v1alpha2"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BuilderSvc struct {
	daoBuilder *dao.BuilderDao

	kPackClient alpha2.KpackV1alpha2Interface
}

func NewBuilderSvc(daoProjImg *dao.BuilderDao, kClients *klient.LocalClients) *BuilderSvc {
	return &BuilderSvc{
		daoBuilder:  daoProjImg,
		kPackClient: kClients.KpackCs.KpackV1alpha2(),
	}
}

func (s *BuilderSvc) List(ctx context.Context, in *builder.BuilderListRequest) (*builder.BuilderListResponse, error) {
	query := bson.M{}
	results, total, err := s.daoBuilder.List(ctx, query)
	return &builder.BuilderListResponse{Items: results, Total: total}, err
}

func (s *BuilderSvc) Create(ctx context.Context, in *builder.Builder) (*builder.Builder, error) {

	// kpBuilder := &v1alpha2.Builder{}
	// imgCli := s.kPackClient.Builders("default")
	// if _, err := imgCli.Create(ctx, kpBuilder, metav1.CreateOptions{}); err != nil {
	// 	return nil, err
	// }

	return s.daoBuilder.Create(ctx, in)
}

func (s *BuilderSvc) Update(ctx context.Context, in *builder.Builder) (*builder.Builder, error) {
	return s.daoBuilder.Update(ctx, in)
}

func (s *BuilderSvc) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	// img, err := s.daoBuilder.Get(ctx, in.Id)
	// if err != nil {
	// 	return nil, err
	// }

	// imgCli := s.kPackClient.Builders("default")
	// if err := imgCli.Delete(ctx, img.Name, metav1.DeleteOptions{}); err != nil {
	// 	return nil, err
	// }

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
