package dao

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewApp,
	NewEgressDao,
	NewEntryDao,
	NewClusterDao,
	NewMiddlewareClaimDao,
	NewMiddlewareInstanceDao,
	NewProject,
	NewProjectImageDao,
	NewProviderDao,
	NewProviderTokenDao,
	NewRouteDao,
	NewDeploymentDao,
	NewBuilderDao,
	NewRegistryDao,
)
