package render

import (
	"github.com/kiaedev/kiae/internal/pkg/render/components"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Middlewares      []common.ApplicationComponent
// Depends          []string

func NewApplication(name string, components ...components.Component) *v1beta1.Application {
	oApp := &v1beta1.Application{ObjectMeta: metav1.ObjectMeta{Name: name}}
	for _, component := range components {
		oApp.Spec.Components = append(oApp.Spec.Components, common.ApplicationComponent{
			Name:       component.GetName(),
			Type:       component.GetType(),
			Properties: util.Object2RawExtension(component),
			// Traits:     buildTraits(traits),
			// DependsOn:  nil,
		})
	}
	return oApp
}

func NewApplicationWith(oApp *v1beta1.Application, components ...components.Component) *v1beta1.Application {
	oApp.Spec.Components = make([]common.ApplicationComponent, 0, len(components))
	for _, component := range components {
		oApp.Spec.Components = append(oApp.Spec.Components, common.ApplicationComponent{
			Name:       component.GetName(),
			Type:       component.GetType(),
			Properties: util.Object2RawExtension(component),
			// Traits:     buildTraits(traits),
			// DependsOn:  nil,
		})
	}
	return oApp
}
