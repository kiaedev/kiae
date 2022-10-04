package service

import (
	"github.com/google/wire"
)

type ServiceSets struct {
	*AppService
	*AppPodsService
	*EgressService
	*EntryService
	*ImageWatcher
	*MiddlewareService
	*Oauth2
	*ProjectService
	*ProjectImageSvc
	*ProviderService
	*RouteService
}

var ProviderSet = wire.NewSet(
	NewAppService,
	NewAppPodsService,
	NewEgressService,
	NewEntryService,
	NewImageWatcher,
	NewMiddlewareService,
	NewOauth2Service,
	NewProjectService,
	NewProjectImageSvc,
	NewProviderService,
	NewRouteService,
	wire.Struct(new(ServiceSets), "*"),
)
