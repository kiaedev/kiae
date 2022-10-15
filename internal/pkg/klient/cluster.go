package klient

import (
	"context"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ClusterClients struct {
	clientSet *kubernetes.Clientset
	informer  informers.SharedInformerFactory
}

func NewClusterClients(cfg *rest.Config) (*ClusterClients, error) {
	kubeClients, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &ClusterClients{
		clientSet: kubeClients,
		informer:  informers.NewSharedInformerFactory(kubeClients, time.Hour),
	}, nil
}

func (cc *ClusterClients) ClientSet() *kubernetes.Clientset {
	return cc.clientSet
}

func (cc *ClusterClients) Informer() informers.SharedInformerFactory {
	return cc.informer
}

func (cc *ClusterClients) InformerStart(ctx context.Context) {
	cc.informer.Start(ctx.Done())
}
