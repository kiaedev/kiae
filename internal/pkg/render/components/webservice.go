package components

import (
	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/internal/pkg/render/traits"
	"github.com/kiaedev/kiae/internal/pkg/render/utils"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	v1 "k8s.io/api/core/v1"
)

type KWebservice struct {
	Name             string                  `json:"name"`
	Labels           map[string]string       `json:"labels,omitempty"`
	Annotations      map[string]string       `json:"annotations,omitempty"`
	Envs             map[string]string       `json:"envs,omitempty"`
	Image            string                  `json:"image"`
	ImagePullPolicy  string                  `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets []string                `json:"imagePullSecrets,omitempty"`
	Ports            []*project.Port         `json:"ports"`
	Replicas         uint32                  `json:"replicas"`
	Resources        v1.ResourceRequirements `json:"resources"`
	LivenessProbe    *app.HealthProbe        `json:"livenessProbe,omitempty"`
	ReadinessProbe   *app.HealthProbe        `json:"readinessProbe,omitempty"`

	traits []common.ApplicationTrait
}

func NewKWebservice(ap *app.Application, proj *project.Project) *KWebservice {
	ts := make([]common.ApplicationTrait, 0)
	if len(ap.Configs) > 0 {
		ts = append(ts, common.ApplicationTrait{
			Type: "k-config", Properties: util.Object2RawExtension(map[string]interface{}{"configs": ap.Configs}),
		})
	}

	envs := make(map[string]string)
	for _, env := range ap.Environments {
		envs[env.Name] = env.Value
	}

	return &KWebservice{
		Name: ap.Name,
		// Labels:      map[string]string{"kiae.dev/test": "test"},
		Annotations: ap.Annotations,
		Image:       ap.Image,
		// ImagePullSecrets: "",
		Replicas:  ap.Replicas,
		Ports:     ap.Ports,
		Resources: utils.BuildResources(ap.Size, 0.5),
		Envs:      envs,

		traits: ts,
	}
}

func (c *KWebservice) GetName() string {
	return c.Name
}

func (c *KWebservice) GetType() string {
	return "k-webservice"
}

func (c *KWebservice) GetTraits() []common.ApplicationTrait {
	return c.traits
}

func (c *KWebservice) SetupTrait(trait traits.Trait) {
	c.traits = append(c.traits, common.ApplicationTrait{Type: trait.GetType(), Properties: util.Object2RawExtension(trait)})
}
