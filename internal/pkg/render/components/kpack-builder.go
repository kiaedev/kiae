package components

import (
	"fmt"
	"path/filepath"
	"strings"

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

func NewKpackBuilder(builderPb *builder.Builder) *KpackBuilder {
	return &KpackBuilder{
		Name:       builderPb.Name,
		StackID:    builderPb.StackId,
		BuildImage: builderPb.BuildImage,
		RunImage:   builderPb.RunImage,
		Packs:      builderPb.Packs,
	}
}

func (m *KpackBuilder) SetupRegistry(imgRegistry *image.Registry, imgRegistrySecret string) {
	imageTag := filepath.Join(imgRegistry.Server, "kiae-builders", m.Name)
	if strings.Contains(imgRegistry.Server, "docker.io") {
		imageTag = filepath.Join(imgRegistry.Username, fmt.Sprintf("kiae-builder-%s", m.Name))
	}

	m.ImageTag = imageTag
	m.ImageRegistry = imgRegistrySecret
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
