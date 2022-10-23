package oidc

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Oauth2 struct {
	cfg *Config

	hookIDToken  func(ctx context.Context, token string) http.HandlerFunc
	hookUserInfo func(ctx context.Context, userInfo *oidc.UserInfo) error
}

func New(cfg *Config) *Oauth2 {
	return &Oauth2{
		cfg: cfg,
		hookIDToken: func(ctx context.Context, token string) http.HandlerFunc {
			return func(writer http.ResponseWriter, request *http.Request) {}
		},
		hookUserInfo: func(ctx context.Context, userInfo *oidc.UserInfo) error {
			return nil
		},
	}
}

func (s *Oauth2) Setup(handleFunc func(path string, h http.HandlerFunc)) {
	handleFunc("/oauth2/authorize", s.Authorize)
	handleFunc("/oauth2/callback", s.Callback)
}

func (s *Oauth2) SetupIDTokenHook(hook func(ctx context.Context, token string) http.HandlerFunc) {
	s.hookIDToken = hook
}

func (s *Oauth2) SetupUserInfoHook(hook func(ctx context.Context, userInfo *oidc.UserInfo) error) {
	s.hookUserInfo = hook
}

func (s *Oauth2) Authorize(w http.ResponseWriter, r *http.Request) {
	o2c, err := s.Oauth2Cfg(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, o2c.AuthCodeURL("test"), 302)
}

func (s *Oauth2) Callback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	o2c, err := s.Oauth2Cfg(ctx, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := o2c.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := s.verifyIDTokenFromOToken(ctx, token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.hookIDToken(ctx, getIdTokenFromOauth2Token(token)).ServeHTTP(w, r)

	userInfo, err := s.getUserInfo(ctx, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := s.hookUserInfo(ctx, userInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.FormValue("callback"), 302)
}

func getIdTokenFromOauth2Token(token *oauth2.Token) string {
	rawIDToken, _ := token.Extra("id_token").(string)
	return rawIDToken
}

func buildRedirectURL(r *http.Request) string {
	query := make(url.Values)
	// query.Set("provider", r.URL.Query().Get("provider"))
	// query.Set("callback", r.URL.Query().Get("callback"))
	u := &url.URL{Scheme: "http", Host: r.Host, RawQuery: query.Encode()}
	u.Path = "/oauth2/callback"
	return u.String()
}

func (s *Oauth2) Oauth2Cfg(ctx context.Context, r *http.Request) (*oauth2.Config, error) {
	provider, err := s.getProvider(ctx)
	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		ClientID:     s.cfg.ClientID,
		ClientSecret: s.cfg.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  buildRedirectURL(r),
		Scopes:       []string{oidc.ScopeOpenID, "email", "profile"},
	}, nil
}

func (s *Oauth2) verifyIDTokenFromOToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("No id_token field in oauth2 token.")
	}

	return s.VerifyIDToken(ctx, rawIDToken)
}

func (s *Oauth2) VerifyIDToken(ctx context.Context, rawIDToken string) (*oidc.IDToken, error) {
	provider, err := s.getProvider(ctx)
	if err != nil {
		return nil, err
	}

	return provider.Verifier(&oidc.Config{ClientID: s.cfg.ClientID}).Verify(ctx, rawIDToken)
}

func (s *Oauth2) getUserInfo(ctx context.Context, token *oauth2.Token) (*oidc.UserInfo, error) {
	provider, err := s.getProvider(ctx)
	if err != nil {
		return nil, err
	}

	return provider.UserInfo(ctx, oauth2.StaticTokenSource(token))
}

func (s *Oauth2) getProvider(ctx context.Context) (*oidc.Provider, error) {
	return oidc.NewProvider(ctx, s.cfg.Endpoint)
}
