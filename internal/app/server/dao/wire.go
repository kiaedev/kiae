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
	NewRouteDao,
	NewDeploymentDao,
	NewBuilderDao,
	NewRegistryDao,
	NewUserDao,
	NewGateway,
)
