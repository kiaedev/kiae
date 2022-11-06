package authz

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiaedev/kiae/internal/app/server/service"
	"github.com/storyicon/grbac"
	"github.com/storyicon/grbac/pkg/meta"
	"gopkg.in/yaml.v3"
)

//go:embed rbac.yml
var embedRules []byte

func Middleware() mux.MiddlewareFunc {
	rules := make(meta.Rules, 0)
	if err := yaml.Unmarshal(embedRules, &rules); err != nil {
		log.Fatalln(err)
	}

	ctrl, err := grbac.New(grbac.WithRules(rules))
	if err != nil {
		log.Fatalln(err)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			roles, ok := r.Context().Value(service.CtxUserRoles).([]string)
			if !ok {
				roles = []string{"anonymous"}
			}

			state, err := ctrl.IsRequestGranted(r, roles)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if !state.IsGranted() {
				http.Error(w, "access deny", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
