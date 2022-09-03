package templates

import (
	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/pkg/kiaeutil"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"github.com/saltbo/gopkg/strutil"
)

type Application struct {
	Name        string
	Image       string
	Ports       []*project.Port
	ConfigPaths []string
	Traits      []common.ApplicationTrait
	Middlewares []common.ApplicationComponent
}

func NewApplication(app *app.Application, proj *project.Project, traits []*kiae.Trait) (*v1beta1.Application, error) {
	appTmplModel := &Application{
		Name:        app.Name,
		Image:       app.Image,
		Ports:       proj.Ports,
		ConfigPaths: buildMountPaths(kiaeutil.ConfigsMerge(proj.Configs, app.Configs)),
		Traits:      buildTraits(traits),
		Middlewares: buildMiddlewares(proj.Middlewares),
	}

	var oam v1beta1.Application
	err := New("app").Render(appTmplModel, &oam)
	if err != nil {
		return nil, err
	}

	return &oam, nil
}

func buildMountPaths(configs []*project.Configuration) []string {
	paths := make([]string, 0, len(configs))
	for _, cfg := range configs {
		if strutil.StrInSlice(cfg.MountPath, paths) {
			continue
		}

		paths = append(paths, cfg.MountPath)
	}

	return paths
}

func buildTraits(traits []*kiae.Trait) []common.ApplicationTrait {
	finalTraits := make([]common.ApplicationTrait, 0)
	for _, traitItem := range traits {
		finalTraits = append(finalTraits, common.ApplicationTrait{
			Type:       traitItem.Type,
			Properties: util.Object2RawExtension(traitItem.Properties),
		})

	}
	return finalTraits
}

func buildMiddlewares(middlewares []*project.Middleware) []common.ApplicationComponent {
	var res []common.ApplicationComponent
	for _, middleware := range middlewares {
		res = append(res, common.ApplicationComponent{
			Name:       middleware.Name,
			Type:       middleware.Type,
			Properties: util.Object2RawExtension(middleware.Properties),
		})

	}
	return res
}
