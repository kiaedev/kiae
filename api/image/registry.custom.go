package image

import (
	"fmt"
	"path/filepath"
	"strings"
)

func (reg *Registry) GetSecretName() string {
	return fmt.Sprintf("kpack-reg-%s", reg.Name)
}

func (reg *Registry) BuildImage(name string) string {
	imageTag := filepath.Join(reg.Server, name)
	if reg.IsDockerhub() {
		imageTag = filepath.Join(reg.Username, name)
	}

	return imageTag
}

func (reg *Registry) BuildImageWithTag(name, tag string) string {
	return reg.BuildImage(name) + ":" + tag
}

func (reg *Registry) IsDockerhub() bool {
	return strings.Contains(reg.Server, "docker.io")
}
