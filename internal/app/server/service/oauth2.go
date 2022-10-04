package service

import (
	"context"
	"net/http"
	"net/url"

	"github.com/kiaedev/kiae/api/provider"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Oauth2 struct {
	projSvc *ProjectService
	pvdSvc  *ProviderService

	daoToken *dao.ProviderTokenDao
}

func NewOauth2Service(projSvc *ProjectService, pvdSvc *ProviderService, daoToken *dao.ProviderTokenDao) *Oauth2 {
	return &Oauth2{projSvc: projSvc, pvdSvc: pvdSvc, daoToken: daoToken}
}

func (s *Oauth2) SetupHandler() {
	http.HandleFunc("/oauth2/authorize", s.authorize)
	http.HandleFunc("/oauth2/callback", s.callback)
}

func (s *Oauth2) authorize(w http.ResponseWriter, r *http.Request) {
	providerName := r.FormValue("provider")

	ctx := context.Background()
	_, cfg, err := s.pvdSvc.GetProvider(ctx, providerName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.setupRedirectURL(cfg, r)
	http.Redirect(w, r, cfg.AuthCodeURL("test", oauth2.AccessTypeOnline), 302)
}

func (s *Oauth2) callback(w http.ResponseWriter, r *http.Request) {
	providerName := r.FormValue("provider")
	callback := r.FormValue("callback")

	ctx := context.Background()
	pvd, cfg, err := s.pvdSvc.GetProvider(ctx, providerName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.setupRedirectURL(cfg, r)
	token, err := cfg.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// todo remove exist token if it exists
	pt := &provider.Token{
		// Userid:       ctx.Value("userid").(string),
		Provider:     pvd.Name,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    timestamppb.New(token.Expiry),
		CreatedAt:    timestamppb.Now(),
		UpdatedAt:    timestamppb.Now(),
	}
	if _, err := s.daoToken.Upsert(ctx, pt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, callback, 302)
}

func (s *Oauth2) setupRedirectURL(cfg *oauth2.Config, r *http.Request) {
	query := make(url.Values)
	query.Set("provider", r.URL.Query().Get("provider"))
	query.Set("callback", r.URL.Query().Get("callback"))
	u := &url.URL{Scheme: "http", Host: r.Host, RawQuery: query.Encode()}
	u.Path = "/oauth2/callback"
	cfg.RedirectURL = u.String()
}
