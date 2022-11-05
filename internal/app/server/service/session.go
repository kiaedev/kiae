package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/kiaedev/kiae/pkg/oauth2"
)

type CtxKey string

const (
	CtxUserid CtxKey = "kiae-userid"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type Session struct {
	oidc *oauth2.OIDC

	userSvc *UserSvc
}

func NewSession(oidc *oauth2.OIDC, userSvc *UserSvc) *Session {
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
	router.HandleFunc("/api/v1/session", s.logout)
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

func (s *Session) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "kiae-session",
		Path:    "/",
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
	})
}

func (s *Session) userInfoHook(ctx context.Context, userInfo *oidc.UserInfo) error {
	return s.userSvc.saveFromOidcUserInfo(ctx, userInfo)
}

func (s *Session) Middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/oauth2") {
				next.ServeHTTP(w, r)
				return
			}

			if !strings.Contains(r.URL.Path, "oauth2") && !strings.HasPrefix(r.URL.Path, "/api") {
				next.ServeHTTP(w, r)
				return
			}

			// Get a session. We're ignoring the error resulted from decoding an
			// existing session: Get() always returns a session, even if empty.
			session, _ := store.Get(r, "kiae-session")

			// check login
			idToken := session.Values["idToken"]
			if idToken == nil {
				w.Header().Set("Location", "/oauth2/authorize")
				http.Error(w, "not login", http.StatusUnauthorized)
				// http.Redirect(w, r, "/oauth2/authorize", http.StatusFound)
				return
			}

			ctx := r.Context()
			idt, err := s.oidc.VerifyIDToken(ctx, idToken.(string))
			if err != nil {
				log.Println("Error verifying token: ", err)
				w.Header().Set("Location", "/oauth2/authorize")
				http.Error(w, "not login", http.StatusUnauthorized)
				// http.Redirect(w, r, "/oauth2/authorize", http.StatusFound)
				return
			}

			u, err := s.userSvc.userDao.GetByOuterId(ctx, idt.Subject)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, CtxUserid, u.Id)))
		})
	}
}

func MustGetUserid(ctx context.Context) string {
	userid, ok := ctx.Value(CtxUserid).(string)
	if !ok {
		panic("not found the userid from context")
	}

	return userid
}
