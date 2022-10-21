package image

import (
	"path/filepath"
	"strings"
)

func (img *Image) SetImage(image string) {
	img.Name, img.Tag = splitNameTag(image)
	img.Image = image
}

func splitNameTag(image string) (string, string) {
	imageItems := strings.Split(image, ":")
	tag := "latest"
	if len(imageItems) == 2 {
		tag = imageItems[1]
	}

	return filepath.Base(imageItems[0]), tag
}
