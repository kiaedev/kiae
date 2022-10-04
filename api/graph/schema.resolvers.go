package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kiaedev/kiae/api/graph/generated"
	"github.com/kiaedev/kiae/api/graph/model"
)

// Pods is the resolver for the pods field.
func (r *queryResolver) Pods(ctx context.Context, ns string, app *string) ([]*model.Pod, error) {
	return r.appPodsSvc.Pods(ctx, ns, *app)
}

// Pods is the resolver for the pods field.
func (r *subscriptionResolver) Pods(ctx context.Context, ns string, app *string) (<-chan []*model.Pod, error) {
	return r.appPodsSvc.SubPods(ctx, ns, *app), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
