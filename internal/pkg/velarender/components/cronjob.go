package components

import (
	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/internal/pkg/velarender/utils"
	v1 "k8s.io/api/core/v1"
)

type KCronjob struct {
	Name             string                  `json:"name"`
	Labels           map[string]string       `json:"labels,omitempty"`
	Annotations      map[string]string       `json:"annotations,omitempty"`
	Envs             map[string]string       `json:"envs,omitempty"`
	Image            string                  `json:"image"`
	ImagePullPolicy  string                  `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets []string                `json:"imagePullSecrets,omitempty"`
	Replicas         uint32                  `json:"replicas"`
	Resources        v1.ResourceRequirements `json:"resources"`
	// Configs          []*app.Configuration    `json:"configs,omitempty"`

	// Traits           []common.ApplicationTrait
	// Middlewares      []common.ApplicationComponent
	// Depends          []string
}

func NewKCronjob(kApp *app.Application, proj *project.Project) *KCronjob {
	return &KCronjob{
		Name:        kApp.Name,
		Labels:      map[string]string{"kiae.dev/test": "test"},
		Annotations: map[string]string{"kiae.dev/test": "test"},
		Image:       kApp.Image,
		// ImagePullSecrets: "",
		Replicas:  kApp.Replicas,
		Envs:      map[string]string{},
		Resources: utils.BuildResources(kApp.Size, 0.5),
		// Configs:   kApp.Configs,
	}
}
