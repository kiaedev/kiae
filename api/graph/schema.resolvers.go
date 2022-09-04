package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kiaedev/kiae/api/graph/generated"
	"github.com/kiaedev/kiae/api/graph/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
