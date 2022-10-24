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
	*BuilderSvc
	*ImageRegistrySvc
	*Session
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
	NewProjectService,
	NewProjectImageSvc,
	NewDeploymentService,
	NewProviderService,
	NewProviderOauth2Svc,
	NewRouteService,
	NewClusterService,
	NewBuilderSvc,
	NewImageRegistrySvc,
	NewUserSvc,
	NewSession,
	wire.Struct(new(ServiceSets), "*"),
)
