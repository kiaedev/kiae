package components

import (
	"strings"

	"github.com/kiaedev/kiae/api/builder"
	"github.com/kiaedev/kiae/api/image"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
)

type KpackBuilder struct {
	ImageNs    string          `json:"imageNs"`
	Packs      []*builder.Pack `json:"packs"`
	StackID    string          `json:"stackId"`
	BuildImage string          `json:"buildImage"`
	RunImage   string          `json:"runImage"`
}

func NewKpackBuilder(builderPb *builder.Builder, imgRegistry *image.Registry) Component {
	imageNs := "kiae-builders"
	if strings.Contains(imgRegistry.Server, "docker.io") {
		imageNs = imgRegistry.Username
	}

	return &KpackBuilder{
		ImageNs:    imageNs,
		Packs:      builderPb.Packs,
		StackID:    builderPb.StackId,
		BuildImage: builderPb.BuildImage,
		RunImage:   builderPb.RunImage,
	}
}

func (m *KpackBuilder) GetName() string {
	return "builder"
}

func (m *KpackBuilder) GetType() string {
	return "k-pack-builder"
}

func (m *KpackBuilder) GetTraits() []common.ApplicationTrait {
	return nil
}
