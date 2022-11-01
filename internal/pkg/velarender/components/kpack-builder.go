package components

import (
	"fmt"
	"path/filepath"

	"github.com/kiaedev/kiae/api/builder"
	"github.com/kiaedev/kiae/api/image"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
)

type KpackBuilder struct {
	Name          string          `json:"name"`
	ImageTag      string          `json:"imageTag"`
	ImageRegistry string          `json:"imageRegistry"`
	StackID       string          `json:"stackId"`
	BuildImage    string          `json:"buildImage"`
	RunImage      string          `json:"runImage"`
	Packs         []*builder.Pack `json:"packs"`
}

func NewKpackBuilder(builderPb *builder.Builder, registry *image.Registry) *KpackBuilder {
	imageName := filepath.Join("kiae-builders", builderPb.Name)
	if registry.IsDockerhub() {
		imageName = fmt.Sprintf("kiae-builder-%s", builderPb.Name)
	}

	return &KpackBuilder{
		ImageTag:      registry.BuildImage(imageName),
		ImageRegistry: registry.GetSecretName(),
		Name:          builderPb.Name,
		StackID:       builderPb.StackId,
		BuildImage:    builderPb.BuildImage,
		RunImage:      builderPb.RunImage,
		Packs:         builderPb.Packs,
	}
}

func (m *KpackBuilder) GetName() string {
	return m.Name
}

func (m *KpackBuilder) GetType() string {
	return "k-kpack-builder"
}

func (m *KpackBuilder) GetTraits() []common.ApplicationTrait {
	return nil
}
