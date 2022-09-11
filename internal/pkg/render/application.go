package render

import (
	"github.com/kiaedev/kiae/internal/pkg/render/components"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = components.KWebservice{}
var _ = components.KCronjob{}

type Component interface {
	GetName() string
	GetType() string
}

func NewApplication(name string, components ...Component) *v1beta1.Application {
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

func NewApplicationWith(oApp *v1beta1.Application, components ...Component) *v1beta1.Application {
	existComponents := make(map[string]int)
	for idx, c := range oApp.Spec.Components {
		existComponents[c.Type] = idx
	}

	for _, component := range components {
		if idx, ok := existComponents[component.GetType()]; ok {
			oApp.Spec.Components[idx].Name = component.GetName()
			oApp.Spec.Components[idx].Properties = util.Object2RawExtension(component)
		} else {
			oApp.Spec.Components = append(oApp.Spec.Components, common.ApplicationComponent{
				Name:       component.GetName(),
				Type:       component.GetType(),
				Properties: util.Object2RawExtension(component),
				// Traits:     buildTraits(traits),
				// DependsOn:  nil,
			})
		}
	}
	return oApp
}
