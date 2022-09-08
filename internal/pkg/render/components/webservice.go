package components

import (
	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/internal/pkg/render/utils"
	"github.com/kiaedev/kiae/pkg/kiaeutil"
	v1 "k8s.io/api/core/v1"
)

type KWebservice struct {
	Name             string                   `json:"name"`
	Labels           map[string]string        `json:"labels,omitempty"`
	Annotations      map[string]string        `json:"annotations,omitempty"`
	Envs             map[string]string        `json:"envs,omitempty"`
	Image            string                   `json:"image"`
	ImagePullPolicy  string                   `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets []string                 `json:"imagePullSecrets,omitempty"`
	Ports            []*project.Port          `json:"ports"`
	Replicas         uint32                   `json:"replicas"`
	Resources        v1.ResourceRequirements  `json:"resources"`
	Configs          []*project.Configuration `json:"configs,omitempty"`
	LivenessProbe    *project.HealthProbe     `json:"livenessProbe,omitempty"`
	ReadinessProbe   *project.HealthProbe     `json:"readinessProbe,omitempty"`

	// Traits           []common.ApplicationTrait
	// Middlewares      []common.ApplicationComponent
	// Depends          []string
}

func NewKWebservice(kApp *app.Application, proj *project.Project) *KWebservice {
	return &KWebservice{
		Name:        kApp.Name,
		Labels:      map[string]string{"kiae.dev/test": "test"},
		Annotations: map[string]string{"kiae.dev/test": "test"},
		Image:       kApp.Image,
		// ImagePullSecrets: "",
		Replicas:  kApp.Replicas,
		Ports:     kApp.Ports,
		Envs:      map[string]string{},
		Resources: utils.BuildResources(kApp.Size, 0.5),
		Configs:   kiaeutil.ConfigsMerge(proj.Configs, kApp.Configs),

		// Traits:      templates.buildTraits(traits),
		// Middlewares: templates.buildMiddlewares(proj.Middlewares),
	}
}

func (c *KWebservice) GetName() string {
	return c.Name
}

func (c *KWebservice) GetType() string {
	return "k-webservice"
}
