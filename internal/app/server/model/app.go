package model

import (
	"strconv"
	"time"

	"github.com/kiaedev/kiae/api/app"
)

type App struct {
}

type AppAction struct {
	*app.Application
}

func NewAppAction(a *app.Application) *AppAction {
	if a.Annotations == nil {
		a.Annotations = make(map[string]string)
	}
	return &AppAction{a}
}

func (aa *AppAction) PreviousReplicas() uint32 {
	replicas, _ := strconv.ParseInt(aa.Annotations["app.kiae.dev/previous-replicas"], 10, 64)
	return uint32(replicas)
}

func (aa *AppAction) SetPreviousReplicas(replicas uint32) {
	aa.Annotations["app.kiae.dev/previous-replicas"] = strconv.FormatInt(int64(replicas), 10)
}

func (aa *AppAction) SetRestartAt(restartAt string) {
	aa.Annotations["app.kiae.dev/restartAt"] = restartAt
}

func (aa *AppAction) Do(action app.ActionPayload_Action) *app.Application {
	switch action {
	case app.ActionPayload_START:
		// 启动逻辑：实例数调回停止前
		aa.Replicas = aa.PreviousReplicas()
		aa.SetPreviousReplicas(0)
		aa.Status = app.Status_STATUS_RUNNING
	case app.ActionPayload_STOP:
		// 停止逻辑: 实例数调到0
		aa.SetPreviousReplicas(aa.Replicas)
		aa.Replicas = 0
		aa.Status = app.Status_STATUS_STOPPED
	case app.ActionPayload_RESTART:
		aa.SetRestartAt(time.Now().String())
	}

	return aa.Application
}
