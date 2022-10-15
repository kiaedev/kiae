package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kiaedev/kiae/api/graph/model"
	"github.com/kiaedev/kiae/internal/app/server/watch"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AppPodsService struct {
	mci *watch.MultiClusterInformers

	podsChan map[string]chan []*model.Pod // TODO: lock?
}

func NewAppPodsService(mci *watch.MultiClusterInformers) *AppPodsService {
	return &AppPodsService{mci: mci, podsChan: make(map[string]chan []*model.Pod)}
}

func (s *AppPodsService) OnAdd(obj interface{}) {
	pod := obj.(*corev1.Pod)
	s.pubLatestPods(pod.Namespace, pod.Labels["kiae.dev/component"])
}

func (s *AppPodsService) OnUpdate(oldObj, newObj interface{}) {
	pod := newObj.(*corev1.Pod)
	s.pubLatestPods(pod.Namespace, pod.Labels["kiae.dev/component"])
}

func (s *AppPodsService) OnDelete(obj interface{}) {
	pod := obj.(*corev1.Pod)
	s.pubLatestPods(pod.Namespace, pod.Labels["kiae.dev/component"])
}

func (s *AppPodsService) pubLatestPods(ns, appName string) {
	subCh, ok := s.podsChan[ns+appName]
	if !ok {
		return
	}

	pods, err := s.Pods(context.Background(), ns, appName)
	if err != nil {
		log.Printf("failed to list pods: %v", err)
	}

	subCh <- pods
}

func (s *AppPodsService) Pods(ctx context.Context, ns, appName string) ([]*model.Pod, error) {
	rt, err := s.mci.Pods(ns, map[string]string{"kiae.dev/component": appName})
	if err != nil {
		return nil, err
	}

	pods := make([]*model.Pod, 0)
	for _, pod := range rt {
		containerStatusMap := make(map[string]corev1.ContainerStatus)
		for _, status := range pod.Status.ContainerStatuses {
			containerStatusMap[status.Name] = status
		}

		containers := make([]*model.Container, 0)
		for _, container := range pod.Spec.Containers {
			containerStatus := containerStatusMap[container.Name]
			status, errMsg := formatContainerStatus(containerStatus.State)
			restartReason, restartRrrMsg := formatContainerStatus(containerStatus.LastTerminationState)
			containers = append(containers, &model.Container{
				Name:          container.Name,
				Image:         container.Image,
				Status:        status,
				ErrMsg:        errMsg,
				RestartCount:  int(containerStatus.RestartCount),
				RestartReason: restartReason,
				RestartErrMsg: restartRrrMsg,
				StartedAt:     formatContainerStartedAt(containerStatus).Format(time.RFC3339),
			})
		}
		pods = append(pods, &model.Pod{
			Name:       pod.Name,
			Namespace:  pod.Namespace,
			Containers: containers,
			Status:     string(pod.Status.Phase),
			PodIP:      pod.Status.PodIP,
			NodeIP:     pod.Status.HostIP,
			CreatedAt:  pod.CreationTimestamp.String(),
			// RestartCount: .RestartCount,
		})
	}

	return pods, nil
}

func formatContainerStatus(state corev1.ContainerState) (string, string) {
	if state.Running != nil {
		return "Running", ""
	}

	if terminated := state.Terminated; terminated != nil {
		return terminated.Reason, fmt.Sprintf("exitCode: %d\r\nerrMsg:%s", terminated.ExitCode, terminated.Message)
	}

	if waiting := state.Waiting; waiting != nil {
		return waiting.Reason, waiting.Message
	}

	return "Unknown", "None"
}

func formatContainerStartedAt(status corev1.ContainerStatus) metav1.Time {
	if running := status.State.Running; running != nil {
		return running.StartedAt
	}

	if terminated := status.State.Terminated; terminated != nil {
		return terminated.StartedAt
	}

	if lastTerminated := status.LastTerminationState.Terminated; lastTerminated != nil {
		return lastTerminated.FinishedAt
	}

	return metav1.Time{}
}

func (s *AppPodsService) SubPods(ctx context.Context, ns string, app string) chan []*model.Pod {
	log.Printf("sub[%s%s] enter", ns, app)
	s.podsChan[ns+app] = make(chan []*model.Pod, 1)
	go func() {
		<-ctx.Done()
		log.Printf("sub[%s%s] exited", ns, app)
		delete(s.podsChan, ns+app)
	}()

	s.pubLatestPods(ns, app)
	return s.podsChan[ns+app]
}
