package service

import (
	"context"
	"net/http"
	"os"
	"strings"

	oidc2 "github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/kiaedev/kiae/pkg/oidc"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type Session struct {
	oidc *oidc.Oauth2

	userSvc *UserSvc
}

func NewSession(oidc *oidc.Oauth2, userSvc *UserSvc) *Session {
	return &Session{
		oidc: oidc,

		userSvc: userSvc,
	}
}

func (s *Session) SetupEndpoints(router *mux.Router) {
	s.oidc.SetupIDTokenHook(s.idTokenHook)
	s.oidc.SetupUserInfoHook(s.userInfoHook)
	s.oidc.Setup(func(path string, h http.HandlerFunc) {
		router.HandleFunc(path, h)
	})
}

func (s *Session) idTokenHook(ctx context.Context, idTokenStr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "kiae-session")
		session.Values["idToken"] = idTokenStr
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Session) userInfoHook(ctx context.Context, userInfo *oidc2.UserInfo) error {
	return s.userSvc.saveFromOidcUserInfo(ctx, userInfo)
}

func (s *Session) Middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get a session. We're ignoring the error resulted from decoding an
			// existing session: Get() always returns a session, even if empty.
			session, _ := store.Get(r, "kiae-session")

			// check login
			idToken := session.Values["idToken"]
			if !strings.HasPrefix(r.URL.Path, "/oauth2") && idToken == nil {
				http.Redirect(w, r, "/oauth2/authorize", http.StatusFound)
				return
			}

			ctx := r.Context()
			idt, err := s.oidc.VerifyIDToken(ctx, idToken.(string))
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			u, err := s.userSvc.userDao.GetByOuterId(ctx, idt.Subject)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "kiae-userid", u.Id)))
		})
	}
}

func MustGetUserid(ctx context.Context) string {
	userid, ok := ctx.Value("kiae-userid").(string)
	if !ok {
		panic("not found the userid from context")
	}

	return userid
}
