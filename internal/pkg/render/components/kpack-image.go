package components

import (
	"github.com/kiaedev/kiae/api/image"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
)

type KpackImage struct {
	Name          string `json:"name"`
	BuilderName   string `json:"builderName"`
	ImageTag      string `json:"imageTag"`
	GitUrl        string `json:"gitUrl"`
	GitCommit     string `json:"gitCommit"`
	GitRepoSecret string `json:"gitRepoSecret"`
	ImgRegSecret  string `json:"imgRegSecret"`
}

func NewKpackImage(img *image.Image, builderName, gitUrl, gitRepoSecret, imgRegSecret string) *KpackImage {
	return &KpackImage{
		Name:          img.Name,
		ImageTag:      img.Image,
		BuilderName:   builderName,
		GitUrl:        gitUrl,
		GitCommit:     img.CommitId,
		GitRepoSecret: gitRepoSecret,
		ImgRegSecret:  imgRegSecret,
	}
}

func (m *KpackImage) GetName() string {
	return m.Name
}

func (m *KpackImage) GetType() string {
	return "k-kpack-image"
}

func (m *KpackImage) GetTraits() []common.ApplicationTrait {
	return nil
}
