package graph

import (
	"github.com/kiaedev/kiae/internal/app/server/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AppPodsSvc   *service.AppPodsService
	AppEventsSvc *service.AppEventService
}
