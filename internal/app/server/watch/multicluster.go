package watch

import (
	"context"
	"sync"

	"github.com/kiaedev/kiae/api/cluster"
	"github.com/kiaedev/kiae/internal/pkg/klient"
	multicluster "github.com/oam-dev/cluster-gateway/pkg/apis/cluster/transport"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

type ClusterEventType string

const (
	ClusterEventAddon   ClusterEventType = "addon"
	ClusterEventRemoved ClusterEventType = "removed"
)

type ClusterEvent struct {
	Type        ClusterEventType
	ClusterName string
}

type MultiClusterInformers struct {
	sync.Map

	kubeconfig *rest.Config

	fetchStockCluster func(ctx context.Context) []*cluster.Cluster
	clusterEvents     chan ClusterEvent

	podEventHandler cache.ResourceEventHandler
}

func NewMultiClusterInformers(kubeconfig *rest.Config) *MultiClusterInformers {
	return &MultiClusterInformers{kubeconfig: kubeconfig, clusterEvents: make(chan ClusterEvent)}
}

func (w *MultiClusterInformers) ClusterClients(clusterName string) (*klient.ClusterClients, error) {
	cfg := *w.kubeconfig
	newCfg := &cfg
	if clusterName != "local" {
		newCfg.Wrap(multicluster.NewProxyPathPrependingClusterGatewayRoundTripper(clusterName).NewRoundTripper)
	}

	return klient.NewClusterClients(newCfg)
}

func (w *MultiClusterInformers) Pods(ns string, matchLabels map[string]string) ([]*corev1.Pod, error) {
	selector, _ := buildAppSelector(matchLabels)

	pods := make([]*corev1.Pod, 0)
	w.Range(func(key, value any) bool {
		kcc := value.(*klient.ClusterClients)
		_pods, _ := kcc.Informer().Core().V1().Pods().Lister().Pods(ns).List(selector)
		pods = append(pods, _pods...)
		return true
	})

	return pods, nil
}

func (w *MultiClusterInformers) SetupPodsEventHandler(h cache.ResourceEventHandler) {
	w.podEventHandler = h
}

func (w *MultiClusterInformers) SetupStockClusterFetcher(fetchStockCluster func(ctx context.Context) []*cluster.Cluster) {
	w.fetchStockCluster = fetchStockCluster
}

func (w *MultiClusterInformers) ClusterEvent(eventType ClusterEventType, cluster *cluster.Cluster) {
	w.clusterEvents <- ClusterEvent{eventType, cluster.Name}
}

func (w *MultiClusterInformers) Start(ctx context.Context) {
	storeClusterClients := func(clusterName string) {
		kubeClientSet, err := w.ClusterClients(clusterName)
		if err != nil {
			return
		}

		podInformer := kubeClientSet.Informer().Core().V1().Pods().Informer()
		podInformer.AddEventHandler(w.podEventHandler)
		kubeClientSet.InformerStart(ctx)
		w.Store(clusterName, kubeClientSet)
	}

	for _, clusterPb := range w.fetchStockCluster(ctx) {
		storeClusterClients(clusterPb.Name)
	}

	go func() {
		for event := range w.clusterEvents {
			switch event.Type {
			case ClusterEventAddon:
				storeClusterClients(event.ClusterName)
			case ClusterEventRemoved:
				w.Delete(event.ClusterName)
			}
		}
	}()
}
