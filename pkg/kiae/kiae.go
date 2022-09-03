package kiae

import "github.com/kiaedev/kiae/api/project"

func ConfigsMerge(projConfigs, appConfigs []*project.Configuration) []*project.Configuration {
	configs := make([]*project.Configuration, 0, len(projConfigs)+len(appConfigs))
	configs = append(configs, projConfigs...)
	configs = append(configs, appConfigs...)
	return configs
}
