package service

import (
	"github.com/google/wire"
)

type ServiceSets struct {
	*AppService
	*AppPodsService
	*AppStatusService
	*AppEventService
	*ClusterService
	*EgressService
	*EntryService
	*ImageWatcher
	*MiddlewareService
	*Oauth2
	*ProjectService
	*ProjectImageSvc
	*ProviderService
	*RouteService
	*DeploymentService
}

var ProviderSet = wire.NewSet(
	NewAppService,
	NewAppPodsService,
	NewAppStatusService,
	NewAppEventService,
	NewEgressService,
	NewEntryService,
	NewImageWatcher,
	NewMiddlewareService,
	NewOauth2Service,
	NewProjectService,
	NewProjectImageSvc,
	NewDeploymentService,
	NewProviderService,
	NewRouteService,
	NewClusterService,
	wire.Struct(new(ServiceSets), "*"),
)
