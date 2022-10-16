package service

import (
	"context"
	"log"

	"github.com/kiaedev/kiae/api/app"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	vela "github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AppStatusService struct {
	rClient     client.Client
	velaClients *vela.Clientset

	appSvc *AppService
}

func NewAppStatusService(rClient client.Client, velaClients *vela.Clientset, appSvc *AppService) *AppStatusService {
	return &AppStatusService{rClient: rClient, velaClients: velaClients, appSvc: appSvc}
}

func (s *AppStatusService) OnAdd(obj interface{}) {
	va := obj.(*v1beta1.Application)
	log.Printf("app %s added: %v", va.Name, va.Status.Phase)

	s.updateStatus(va)
}

func (s *AppStatusService) OnUpdate(oldObj, newObj interface{}) {
	va := newObj.(*v1beta1.Application)
	log.Printf("app %s updated: %v", va.Name, va.Status.Phase)

	s.updateStatus(va)
}

func (s *AppStatusService) OnDelete(obj interface{}) {
	va := obj.(*v1beta1.Application)
	log.Printf("app status update: %v", va.Status.Phase)
}

func (s *AppStatusService) updateStatus(va *v1beta1.Application) {
	ctx := context.Background()
	ap, err := s.appSvc.daoApp.GetByName(ctx, va.Name)
	if err != nil {
		return
	}

	// fixme: 停止时误把stopped状态改为了running
	// 先更新的component，后改的数据库，这里的判定大概率判定不到
	if ap.Status == app.Status_STATUS_STOPPED {
		return
	}

	if _, err := s.appSvc.updateStatus(ctx, ap, buildAppStatus(va)); err != nil {
		return
	}
}

func buildAppStatus(va *v1beta1.Application) app.Status {
	statusMap := map[common.ApplicationPhase]app.Status{
		common.ApplicationRunning:   app.Status_STATUS_RUNNING,
		common.ApplicationUnhealthy: app.Status_STATUS_UNHEALTHY,
	}

	apStatus, ok := statusMap[va.Status.Phase]
	if !ok {
		apStatus = app.Status_STATUS_DEPLOYING
	}

	// todo maybe update more status

	return apStatus
}
