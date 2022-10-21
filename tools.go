//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/favadi/protoc-go-inject-tag"
	_ "github.com/google/wire/cmd/wire"
)
