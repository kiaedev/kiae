package kiaeutil

import (
	"fmt"

	"github.com/saltbo/gopkg/strutil"
)

const (
	NsPrefix = "kiae-"
)

var stdNs = []string{"dev", "test", "stage", "prod"}

func IsStdEnv(env string) bool {
	return strutil.StrInSlice(env, stdNs)
}

func SystemNs() string {
	return NsPrefix + "system"
}

func BuildAppNs(env string) string {
	if !IsStdEnv(env) {
		env = "custom"
	}

	return fmt.Sprintf("%sapp-%s", NsPrefix, env)
}
