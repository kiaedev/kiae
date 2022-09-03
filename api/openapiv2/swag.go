package openapiv2

import "embed"

var (
	//go:embed *.json
	EmbedFs embed.FS
)
