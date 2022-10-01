package service

import (
	"context"

	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/provider"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/bitbucket"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/gitlab"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProviderService struct {
	provider.UnimplementedProviderServiceServer

	daoProvider *dao.ProviderDao
}

func NewProviderService(s *Service) *ProviderService {
	return &ProviderService{
		daoProvider: dao.NewProviderDao(s.DB),
	}
}
func (s *ProviderService) Prepare(context.Context, *emptypb.Empty) (*provider.PreparesResponse, error) {
	items := []*provider.Prepare{
		{Name: "github", AuthorizeUrl: github.Endpoint.AuthURL, TokenUrl: github.Endpoint.TokenURL, Scopes: []string{"repo", "admin:repo_hook"}},
		{Name: "gitlab", AuthorizeUrl: gitlab.Endpoint.AuthURL, TokenUrl: gitlab.Endpoint.TokenURL, Scopes: []string{"api"}},
		{Name: "bitbucket", AuthorizeUrl: bitbucket.Endpoint.AuthURL, TokenUrl: bitbucket.Endpoint.TokenURL},
	}

	return &provider.PreparesResponse{Items: items}, nil
}

func (s *ProviderService) List(ctx context.Context, in *provider.ListRequest) (*provider.ListResponse, error) {
	results, total, err := s.daoProvider.List(ctx, bson.M{})
	return &provider.ListResponse{Items: results, Total: total}, err
}

func (s *ProviderService) Create(ctx context.Context, in *provider.Provider) (*provider.Provider, error) {
	return s.daoProvider.Create(ctx, in)
}

func (s *ProviderService) Update(ctx context.Context, in *provider.UpdateRequest) (*provider.Provider, error) {
	existedProvider, err := s.daoProvider.Get(ctx, in.Payload.Id)
	if err != nil {
		return nil, err
	}

	if in.UpdateMask == nil {
		existedProvider = in.Payload
	} else {
		s.handlePatch(in, existedProvider)
	}

	return s.daoProvider.Update(ctx, existedProvider)
}

func (s *ProviderService) handlePatch(in *provider.UpdateRequest, existedProvider *provider.Provider) {
	// payload := in.Payload
	for _, path := range in.GetUpdateMask().Paths {
		if path == "status" {
			// existedProvider.Status = payload.Status
		}
	}
}

func (s *ProviderService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.daoProvider.Delete(ctx, in.Id)
}

func (s *ProviderService) GetProvider(ctx context.Context, providerName string) (*provider.Provider, *oauth2.Config, error) {
	pvd, err := s.daoProvider.GetByName(ctx, providerName)
	if err != nil {
		return nil, nil, err
	}

	return pvd, &oauth2.Config{
		ClientID:     pvd.ClientId,
		ClientSecret: pvd.ClientSecret,
		Endpoint:     oauth2Endpoint(pvd),
		Scopes:       pvd.Scopes,
	}, nil
}

func oauth2Endpoint(pvd *provider.Provider) oauth2.Endpoint {
	return oauth2.Endpoint{AuthURL: pvd.AuthorizeUrl, TokenURL: pvd.TokenUrl}
}
