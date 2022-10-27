package service

import (
	"context"
	"fmt"
	"time"

	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/provider"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/pkg/klient"
	"github.com/kiaedev/kiae/pkg/gitp"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/oauth2"
	bb "golang.org/x/oauth2/bitbucket"
	gh "golang.org/x/oauth2/github"
	gl "golang.org/x/oauth2/gitlab"
	"google.golang.org/protobuf/types/known/emptypb"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

type ProviderService struct {
	daoProvider *dao.ProviderDao

	kubeSecret v1.SecretInterface
}

func NewProviderService(daoProvider *dao.ProviderDao, kClients *klient.LocalClients) *ProviderService {
	return &ProviderService{
		daoProvider: daoProvider,

		kubeSecret: kClients.K8sCs.CoreV1().Secrets("kiae-builder"),
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
	// for _, path := range in.GetUpdateMask().Paths {
	// 	if path == "status" {
	// 		existedProvider.Status = payload.Status
	// 	}
	// }
}

func (s *ProviderService) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.daoProvider.Delete(ctx, in.Id)
}

func (s *ProviderService) Config(ctx context.Context, providerName string) (*oauth2.Config, error) {
	return s.GetOauth2Config(ctx, providerName)
}

func (s *ProviderService) TokenF(ctx context.Context, providerName string, token *oauth2.Token) error {
	return s.saveToken(ctx, providerName, token)
}

func (s *ProviderService) GetOauth2Config(ctx context.Context, providerName string) (*oauth2.Config, error) {
	pvd, err := s.daoProvider.GetByName(ctx, providerName)
	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		ClientID:     pvd.ClientId,
		ClientSecret: pvd.ClientSecret,
		Endpoint:     oauth2Endpoint(pvd),
		Scopes:       pvd.Scopes,
	}, nil
}

func oauth2Endpoint(pvd *provider.Provider) oauth2.Endpoint {
	return oauth2.Endpoint{AuthURL: pvd.AuthorizeUrl, TokenURL: pvd.TokenUrl}
}

func (s *ProviderService) ListRepos(ctx context.Context, in *provider.ListReposRequest) (*provider.ListReposResponse, error) {
	pv, err := s.getProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}

	results, err := pv.ListRepos(ctx)
	return &provider.ListReposResponse{Items: results, Total: int64(len(results))}, err
}

func (s *ProviderService) ListBranches(ctx context.Context, in *provider.ListBranchesRequest) (*provider.ListBranchesResponse, error) {
	pv, err := s.getProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}

	results, err := pv.ListBranches(ctx, in.RepoName)
	return &provider.ListBranchesResponse{Items: results, Total: int64(len(results))}, err
}

func (s *ProviderService) ListTags(ctx context.Context, in *provider.ListTagsRequest) (*provider.ListTagsResponse, error) {
	pv, err := s.getProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}

	results, err := pv.ListTags(ctx, in.RepoName)
	return &provider.ListTagsResponse{Items: results, Total: int64(len(results))}, err
}

func (s *ProviderService) getProvider(ctx context.Context, providerName string) (gitp.Provider, error) {
	ot, err := s.getProviderToken(ctx, providerName)
	if err != nil {
		return nil, err
	}

	return gitp.Select(providerName, ot.AccessToken)
}

func (s *ProviderService) getProviderToken(ctx context.Context, providerName string) (*oauth2.Token, error) {
	secret, err := s.kubeSecret.Get(ctx, TokenSecretName(ctx, providerName), metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	token := secret2Token(secret)
	if token.Expiry.IsZero() {
		return token, nil
	}

	if token.Expiry.Before(time.Now()) {
		if err := s.refreshToken(ctx, providerName, token); err != nil {
			return nil, err
		}
	}

	return token, nil
}

func (s *ProviderService) refreshToken(ctx context.Context, providerName string, token *oauth2.Token) error {
	cfg, err := s.GetOauth2Config(ctx, providerName)
	if err != nil {
		return err
	}

	newToken, err := cfg.TokenSource(ctx, token).Token()
	if err != nil {
		return err
	}

	return s.saveToken(ctx, providerName, newToken)
}

func (s *ProviderService) saveToken(ctx context.Context, providerName string, token *oauth2.Token) (err error) {
	// fetch the latest username
	username, err := s.getUsername(ctx, providerName, token)
	if err != nil {
		return err
	}
	secret := token2Secret(username, token)
	secret.SetName(TokenSecretName(ctx, providerName))
	secret.Annotations = map[string]string{
		"kiae.dev/git-provider": providerName,
	}

	// upsert the secret
	_, err = s.kubeSecret.Get(ctx, secret.Name, metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return
	} else if errors.IsNotFound(err) {
		_, err = s.kubeSecret.Create(ctx, secret, metav1.CreateOptions{})
		return
	}

	_, err = s.kubeSecret.Update(ctx, secret, metav1.UpdateOptions{})
	return
}

func (s *ProviderService) getUsername(ctx context.Context, providerName string, token *oauth2.Token) (string, error) {
	pvd, err := gitp.Select(providerName, token.AccessToken)
	if err != nil {
		return "", err
	}

	return pvd.AuthedUser(ctx)
}

func TokenSecretName(ctx context.Context, gitProvider string) string {
	userid := MustGetUserid(ctx)
	if userid == "" {
		userid = "1000"
	}

	return fmt.Sprintf("%s-%s", userid, gitProvider)
}

func secret2Token(secret *corev1.Secret) *oauth2.Token {
	m := secret.Data
	expiry, _ := time.Parse(time.Layout, string(m["expires_at"]))
	return &oauth2.Token{
		AccessToken:  string(m["access_token"]),
		RefreshToken: string(m["refresh_token"]),
		Expiry:       expiry,
	}
}

func token2Secret(username string, ot *oauth2.Token) *corev1.Secret {
	expiresAt := ""
	if !ot.Expiry.IsZero() {
		expiresAt = ot.Expiry.Format(time.Layout)
	}
	return &corev1.Secret{
		Type: "kubernetes.io/basic-auth",
		StringData: map[string]string{
			"username":      username,
			"password":      ot.AccessToken,
			"access_token":  ot.AccessToken,
			"refresh_token": ot.RefreshToken,
			"expires_at":    expiresAt,
		},
	}
}
