package render

import (
	"github.com/kiaedev/kiae/internal/pkg/render/components"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)

var _ = components.KWebservice{}
var _ = components.KCronjob{}

type Component interface {
	GetName() string
	GetType() string
}

func NewApplication(component Component) *v1beta1.Application {
	oApp := new(v1beta1.Application)
	oApp.SetName(component.GetName())
	oApp.Spec.Components = []common.ApplicationComponent{
		{
			Name:       component.GetName(),
			Type:       component.GetType(),
			Properties: util.Object2RawExtension(component),
			// Traits:     buildTraits(traits),
			// DependsOn:  nil,
		},
	}
	return oApp
}

func NewApplicationWith(component Component, oApp *v1beta1.Application) *v1beta1.Application {
	for idx, c := range oApp.Spec.Components {
		if c.Type == component.GetType() {
			c.Name = component.GetName()
			oApp.Spec.Components[idx].Properties = util.Object2RawExtension(component)
		}
	}
	return oApp
}
