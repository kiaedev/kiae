package templates

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"github.com/openkos/openkos/api/app"
	"github.com/openkos/openkos/api/project"
)

type Application struct {
	Name        string
	Image       string
	Ports       []*project.Port
	Configs     []*project.Configuration
	Traits      []common.ApplicationTrait
	Middlewares []common.ApplicationComponent
}

func NewApplication(app *app.Application, project *project.Project, traits []common.ApplicationTrait) (*v1beta1.Application, error) {
	appTmplModel := &Application{
		Name:        app.Name,
		Image:       app.Image,
		Ports:       project.Ports,
		Configs:     project.Configs,
		Traits:      traits,
		Middlewares: buildMiddlewares(project.Middlewares),
	}

	var oam v1beta1.Application
	err := New("app").Render(appTmplModel, &oam)
	if err != nil {
		return nil, err
	}

	return &oam, nil
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
