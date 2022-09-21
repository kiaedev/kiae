package traits

import (
	"github.com/kiaedev/kiae/api/app"
	"github.com/saltbo/gopkg/strutil"
)

type ConfigsTrait struct {
	Configs []*app.Configuration `json:"configs"`
}

func NewConfigsTrait(configs []*app.Configuration) *ConfigsTrait {
	return &ConfigsTrait{configs}
}

func (m *ConfigsTrait) GetName() string {
	return "config-" + strutil.RandomText(5)
}

func (m *ConfigsTrait) GetType() string {
	return "k-config"
}
