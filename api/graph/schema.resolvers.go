package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/kiaedev/kiae/api/graph/generated"
	"github.com/kiaedev/kiae/api/graph/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Pod, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// Pods is the resolver for the pods field.
func (r *queryResolver) Pods(ctx context.Context, ns string) ([]*model.Pod, error) {
	results, err := r.cs.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{})
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
func (r *subscriptionResolver) Pods(ctx context.Context, ns string) (<-chan []*model.Pod, error) {
	channel := make(chan []*model.Pod, 1)
	latestPods := func(obj interface{}) {
		selector := labels.NewSelector()
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

			pods = append(pods, &model.Pod{
				Name:      pod.Name,
				Namespace: pod.Namespace,
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
