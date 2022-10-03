package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/kiaedev/kiae/api/graph/generated"
	"github.com/kiaedev/kiae/api/graph/model"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// Pods is the resolver for the pods field.
func (r *queryResolver) Pods(ctx context.Context, ns string, app *string) ([]*model.Pod, error) {
	selector, err := buildAppSelector(app)
	if err != nil {
		return nil, err
	}

	results, err := r.cs.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{LabelSelector: selector.String()})
	if err != nil {
		return nil, err
	}

	pods := make([]*model.Pod, 0)
	for _, item := range results.Items {
		pods = append(pods, &model.Pod{
			Name:      item.Name,
			Namespace: item.Namespace,
		})
	}
	return pods, nil
}

// Pods is the resolver for the pods field.
func (r *subscriptionResolver) Pods(ctx context.Context, ns string, app *string) (<-chan []*model.Pod, error) {
	selector, err := buildAppSelector(app)
	if err != nil {
		return nil, err
	}

	channel := make(chan []*model.Pod, 1)
	latestPods := func(obj interface{}) {
		rt, err := r.podInformer.Lister().List(selector)
		if err != nil {
			log.Println(err)
			return
		}

		pods := make([]*model.Pod, 0)
		for _, pod := range rt {
			if pod.Namespace != ns {
				continue
			}

			containerStatusMap := make(map[string]v1.ContainerStatus)

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
		channel <- pods
	}

	r.podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: latestPods,
		UpdateFunc: func(oldObj, newObj interface{}) {
			latestPods(newObj)
		},
		DeleteFunc: latestPods,
	})

	return channel, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func buildAppSelector(app *string) (labels.Selector, error) {
	matchLabels := make(map[string]string)
	if app != nil {
		matchLabels["kiae.dev/component"] = *app
	}

	return metav1.LabelSelectorAsSelector(&metav1.LabelSelector{MatchLabels: matchLabels})
}
