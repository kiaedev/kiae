package oauth2

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type OidcConfig struct {
	Endpoint     string `yaml:"endpoint"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

type OIDC struct {
	sync.Map
	Oauth2
	cfg *OidcConfig

	hookIDToken  func(ctx context.Context, token string) http.HandlerFunc
	hookUserInfo func(ctx context.Context, userInfo *oidc.UserInfo) error
}

func NewOIDC(cfg *OidcConfig) *OIDC {
	o := &OIDC{
		cfg: cfg,
		hookIDToken: func(ctx context.Context, token string) http.HandlerFunc {
			return func(writer http.ResponseWriter, request *http.Request) {}
		},
		hookUserInfo: func(ctx context.Context, userInfo *oidc.UserInfo) error {
			return nil
		},
	}
	o.SetOauth2ConfigBuilder(o.Oauth2Cfg)
	o.SetCallbackHook(o.oauth2TokenHandler)
	return o
}

func (s *OIDC) SetupIDTokenHook(hook func(ctx context.Context, token string) http.HandlerFunc) {
	s.hookIDToken = hook
}

func (s *OIDC) SetupUserInfoHook(hook func(ctx context.Context, userInfo *oidc.UserInfo) error) {
	s.hookUserInfo = hook
}

func (s *OIDC) Oauth2Cfg(ctx context.Context, req *http.Request) (*oauth2.Config, error) {
	provider, err := s.getProvider(ctx)
	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		ClientID:     s.cfg.ClientID,
		ClientSecret: s.cfg.ClientSecret,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "email", "profile"},
	}, nil
}

func (s *OIDC) oauth2TokenHandler(ctx context.Context, token *oauth2.Token, w http.ResponseWriter, r *http.Request) {
	if _, err := s.verifyIDTokenFromOToken(ctx, token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo, err := s.getUserInfo(ctx, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := s.hookUserInfo(ctx, userInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.hookIDToken(ctx, getIdTokenFromOauth2Token(token)).ServeHTTP(w, r)
}

func (s *OIDC) verifyIDTokenFromOToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("No id_token field in oauth2 token.")
	}

	return s.VerifyIDToken(ctx, rawIDToken)
}

func (s *OIDC) VerifyIDToken(ctx context.Context, rawIDToken string) (*oidc.IDToken, error) {
	provider, err := s.getProvider(ctx)
	if err != nil {
		return nil, err
	}

	return provider.Verifier(&oidc.Config{ClientID: s.cfg.ClientID}).Verify(ctx, rawIDToken)
}

func (s *OIDC) getUserInfo(ctx context.Context, token *oauth2.Token) (*oidc.UserInfo, error) {
	provider, err := s.getProvider(ctx)
	if err != nil {
		return nil, err
	}

	return provider.UserInfo(ctx, oauth2.StaticTokenSource(token))
}

func (s *OIDC) getProvider(ctx context.Context) (*oidc.Provider, error) {
	issuer := s.cfg.Endpoint
	v, ok := s.Load(issuer)
	if ok {
		return v.(*oidc.Provider), nil
	}

	op, err := oidc.NewProvider(oidc.InsecureIssuerURLContext(ctx, issuer), "http://kiae-dex:5556/dex")
	if err != nil {
		return nil, err
	}

	s.Store(issuer, op)
	return op, nil
}

func getIdTokenFromOauth2Token(token *oauth2.Token) string {
	rawIDToken, _ := token.Extra("id_token").(string)
	return rawIDToken
}
