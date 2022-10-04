package service

import (
	"context"
	"log"

	"github.com/kiaedev/kiae/api/graph/model"
	"github.com/kiaedev/kiae/internal/app/server/watch"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type AppPodsService struct {
	w *watch.Watcher

	podsChan map[string]chan []*model.Pod // TODO: lock?
}

func NewAppPodsService(w *watch.Watcher) *AppPodsService {
	return &AppPodsService{w: w, podsChan: make(map[string]chan []*model.Pod)}
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

	pods, err := s.Pods(nil, ns, appName)
	if err != nil {
		log.Printf("failed to list pods: %v", err)
	}

	subCh <- pods
}

func (s *AppPodsService) Pods(ctx context.Context, ns, appName string) ([]*model.Pod, error) {
	rt, err := s.w.Pods(ns, map[string]string{"kiae.dev/component": appName})
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
			containers = append(containers, &model.Container{
				Name:  container.Name,
				Image: container.Image,
				// Status: containerStatus.State,

				RestartCount: int(containerStatus.RestartCount),
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

func buildAppSelector(app *string) (labels.Selector, error) {
	matchLabels := make(map[string]string)
	if app != nil {
		matchLabels["kiae.dev/component"] = *app
	}

	return metav1.LabelSelectorAsSelector(&metav1.LabelSelector{MatchLabels: matchLabels})
}
