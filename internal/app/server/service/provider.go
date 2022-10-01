package service

import (
	"context"

	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/provider"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/pkg/gitp"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/oauth2"
	bb "golang.org/x/oauth2/bitbucket"
	gh "golang.org/x/oauth2/github"
	gl "golang.org/x/oauth2/gitlab"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProviderService struct {
	provider.UnimplementedProviderServiceServer

	daoProvider      *dao.ProviderDao
	daoProviderToken *dao.ProviderTokenDao
}

func NewProviderService(s *Service) *ProviderService {
	return &ProviderService{
		daoProvider:      dao.NewProviderDao(s.DB),
		daoProviderToken: dao.NewProviderTokenDao(s.DB),
	}
}
func (s *ProviderService) Prepare(context.Context, *emptypb.Empty) (*provider.PreparesResponse, error) {
	items := []*provider.Prepare{
		{Name: "github", AuthorizeUrl: gh.Endpoint.AuthURL, TokenUrl: gh.Endpoint.TokenURL, Scopes: []string{"repo", "admin:repo_hook"}},
		{Name: "gitlab", AuthorizeUrl: gl.Endpoint.AuthURL, TokenUrl: gl.Endpoint.TokenURL, Scopes: []string{"api"}},
		{Name: "bitbucket", AuthorizeUrl: bb.Endpoint.AuthURL, TokenUrl: bb.Endpoint.TokenURL},
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

func (s *ProviderService) RepoList(ctx context.Context, in *provider.RepoListRequest) (*provider.RepoListResponse, error) {
	pvt, err := s.daoProviderToken.GetByProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}

	pv, err := gitp.Select(pvt.Provider, pvt.AccessToken)
	if err != nil {
		return nil, err
	}

	results, err := pv.List(ctx)
	return &provider.RepoListResponse{Items: results}, err
}
