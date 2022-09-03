package kiaeutil

import "github.com/kiaedev/kiae/api/project"

func ConfigsMerge(projConfigs, appConfigs []*project.Configuration) []*project.Configuration {
	configs := make([]*project.Configuration, 0, len(projConfigs)+len(appConfigs))
	setConfigLevel(projConfigs, project.ConfigLevel_CONFIG_LEVEL_PROJECT)
	setConfigLevel(appConfigs, project.ConfigLevel_CONFIG_LEVEL_APP)
	configs = append(configs, projConfigs...)
	configs = append(configs, appConfigs...)
	return configs
}

func setConfigLevel(configs []*project.Configuration, level project.ConfigLevel) {
	for _, config := range configs {
		config.Level = level
	}
}
