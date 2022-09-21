package kiaeutil

// func ConfigsMerge(projConfigs, appConfigs []*app.Configuration) []*app.Configuration {
// 	configs := make([]*app.Configuration, 0, len(projConfigs)+len(appConfigs))
// 	setConfigLevel(projConfigs, project.ConfigLevel_CONFIG_LEVEL_PROJECT)
// 	setConfigLevel(appConfigs, project.ConfigLevel_CONFIG_LEVEL_APP)
// 	configs = append(configs, projConfigs...)
// 	configs = append(configs, appConfigs...)
// 	return configs
// }
//
// func setConfigLevel(configs []*app.Configuration, level project.ConfigLevel) {
// 	for _, config := range configs {
// 		config.Level = level
// 	}
// }
