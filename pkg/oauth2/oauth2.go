package oauth2

import (
	"context"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

type (
	O2CBuilderConstructor   func(ctx context.Context, req *http.Request) (*oauth2.Config, error)
	CallbackHookConstructor func(ctx context.Context, token *oauth2.Token, w http.ResponseWriter, r *http.Request)
)

type Oauth2 struct {
	buildOauth2Config O2CBuilderConstructor
	callbackHook      CallbackHookConstructor
}

func (s *Oauth2) Setup(handleFunc func(path string, h http.HandlerFunc)) {
	handleFunc("/oauth2/authorize", s.authorize)
	handleFunc("/oauth2/callback", s.callback)
}

func (s *Oauth2) SetOauth2ConfigBuilder(buildOauth2Config O2CBuilderConstructor) {
	s.buildOauth2Config = func(ctx context.Context, req *http.Request) (*oauth2.Config, error) {
		o2c, err := buildOauth2Config(ctx, req)
		if err != nil {
			return nil, err
		}

		if o2c.RedirectURL == "" {
			o2c.RedirectURL = buildDefaultRedirectURL(req)
		}

		return o2c, err
	}
}

func (s *Oauth2) SetCallbackHook(callbackHook CallbackHookConstructor) {
	s.callbackHook = callbackHook
}

func (s *Oauth2) authorize(w http.ResponseWriter, r *http.Request) {
	o2c, err := s.buildOauth2Config(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, o2c.AuthCodeURL("test"), 302)
}

func (s *Oauth2) callback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	o2c, err := s.buildOauth2Config(ctx, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := o2c.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.callbackHook(ctx, token, w, r)
	location := r.FormValue("callback")
	if location == "" {
		location = "/"
	}

	http.Redirect(w, r, location, 302)
}

func buildDefaultRedirectURL(r *http.Request) string {
	u := &url.URL{Scheme: "http", Host: r.Host}
	u.Path = "/oauth2/callback"
	return u.String()
}
