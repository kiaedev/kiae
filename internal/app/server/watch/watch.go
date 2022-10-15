package watch

import (
	"context"
	"time"

	"github.com/kiaedev/kiae/internal/pkg/klient"
	velaInformers "github.com/oam-dev/kubevela-core-api/pkg/generated/client/informers/externalversions"
	kpackInformers "github.com/pivotal/kpack/pkg/client/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type Watcher struct {
	*MultiClusterInformers

	velaInformers  velaInformers.SharedInformerFactory
	kpackInformers kpackInformers.SharedInformerFactory
}

func NewWatcher(kcs *klient.LocalClients, mci *MultiClusterInformers) (*Watcher, error) {
	return &Watcher{
		MultiClusterInformers: mci,

		velaInformers:  velaInformers.NewSharedInformerFactory(kcs.VelaCs, time.Hour),
		kpackInformers: kpackInformers.NewSharedInformerFactory(kcs.KpackCs, time.Hour),
	}, nil
}

func (w *Watcher) Start(ctx context.Context) {
	go w.MultiClusterInformers.Start(ctx)
	go w.velaInformers.Start(ctx.Done())
	go w.kpackInformers.Start(ctx.Done())
}

func (w *Watcher) SetupPodsEventHandler(h cache.ResourceEventHandler) {
	w.MultiClusterInformers.SetupPodsEventHandler(h)
}

func (w *Watcher) SetupImagesEventHandler(h cache.ResourceEventHandler) {
	w.kpackInformers.Kpack().V1alpha2().Images().Informer().AddEventHandler(h)
}

func (w *Watcher) SetupApplicationsEventHandler(h cache.ResourceEventHandler) {
	w.velaInformers.Core().V1beta1().Applications().Informer().AddEventHandler(h)
}

func buildAppSelector(matchLabels map[string]string) (labels.Selector, error) {
	return metav1.LabelSelectorAsSelector(&metav1.LabelSelector{MatchLabels: matchLabels})
}
