package graph

import (
	"context"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	cs *kubernetes.Clientset

	podInformer v1.PodInformer
}

func NewResolver(cs *kubernetes.Clientset) *Resolver {
	informerFactory := informers.NewSharedInformerFactory(cs, time.Hour*24)
	podInformer := informerFactory.Core().V1().Pods()
	return &Resolver{
		cs:          cs,
		podInformer: podInformer,
	}
}

func (r *Resolver) Run(ctx context.Context) error {
	go r.podInformer.Informer().Run(ctx.Done())
	return nil
}
