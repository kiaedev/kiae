package service

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/kiaedev/kiae/pkg/oauth2"
	oauth22 "golang.org/x/oauth2"
)

type Oauth2 struct {
	oauth2.Oauth2

	pvdSvc *ProviderService
}

func NewProviderOauth2Svc(pvdSvc *ProviderService) *Oauth2 {
	return &Oauth2{pvdSvc: pvdSvc}
}

func (s *Oauth2) SetupEndpoints(router *mux.Router) {
	s.SetOauth2ConfigBuilder(func(ctx context.Context, r *http.Request) (*oauth22.Config, error) {
		o2c, err := s.pvdSvc.GetOauth2Config(ctx, r.FormValue("provider"))
		if err != nil {
			return nil, err
		}

		o2c.RedirectURL = buildRedirectURL(r)
		return o2c, err
	})
	s.SetCallbackHook(func(ctx context.Context, token *oauth22.Token, w http.ResponseWriter, r *http.Request) {
		if err := s.pvdSvc.saveToken(ctx, r.FormValue("provider"), token); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	s.Oauth2.Setup(func(path string, h http.HandlerFunc) {
		router.HandleFunc("/provider"+path, h)
	})
}

func buildRedirectURL(r *http.Request) string {
	query := make(url.Values)
	query.Set("provider", r.URL.Query().Get("provider"))
	query.Set("callback", r.URL.Query().Get("callback"))
	u := &url.URL{Scheme: "http", Host: r.Host, RawQuery: query.Encode()}
	u.Path = "/provider/oauth2/callback"
	return u.String()
}
