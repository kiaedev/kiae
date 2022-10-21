package service

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

type Oauth2 struct {
	projSvc *ProjectService
	pvdSvc  *ProviderService
}

func NewOauth2Service(projSvc *ProjectService, pvdSvc *ProviderService) *Oauth2 {
	return &Oauth2{projSvc: projSvc, pvdSvc: pvdSvc}
}

func (s *Oauth2) SetupHandler(router *mux.Router) {
	router.HandleFunc("/oauth2/authorize", s.authorize)
	router.HandleFunc("/oauth2/callback", s.callback)
}

func (s *Oauth2) authorize(w http.ResponseWriter, r *http.Request) {
	providerName := r.FormValue("provider")

	ctx := context.Background()
	cfg, err := s.pvdSvc.GetOauth2Config(ctx, providerName)
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
	cfg, err := s.pvdSvc.GetOauth2Config(ctx, providerName)
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

	if err := s.pvdSvc.saveToken(ctx, providerName, token); err != nil {
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
