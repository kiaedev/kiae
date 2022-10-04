package watch

import (
	"context"
	"time"

	"github.com/kiaedev/kiae/internal/pkg/kcs"
	velaInformers "github.com/oam-dev/kubevela-core-api/pkg/generated/client/informers/externalversions"
	kpackInformers "github.com/pivotal/kpack/pkg/client/informers/externalversions"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type Watcher struct {
	*kcs.KubeClients

	k8sInformer    informers.SharedInformerFactory
	velaInformers  velaInformers.SharedInformerFactory
	kpackInformers kpackInformers.SharedInformerFactory
}

func NewWatcher(kcs *kcs.KubeClients) (*Watcher, error) {
	return &Watcher{
		KubeClients:    kcs,
		k8sInformer:    informers.NewSharedInformerFactory(kcs.K8sCs, time.Hour),
		velaInformers:  velaInformers.NewSharedInformerFactory(kcs.VelaCs, time.Hour),
		kpackInformers: kpackInformers.NewSharedInformerFactory(kcs.KpackCs, time.Hour),
	}, nil
}

func (w *Watcher) Start(ctx context.Context) {
	go w.k8sInformer.Core().V1().Pods().Informer().Run(ctx.Done())
	go w.velaInformers.Core().V1beta1().Applications().Informer().Run(ctx.Done())
	go w.kpackInformers.Kpack().V1alpha2().Images().Informer().Run(ctx.Done())
}

func (w *Watcher) Pods(ns string, matchLabels map[string]string) ([]*corev1.Pod, error) {
	selector, _ := buildAppSelector(matchLabels)
	return w.k8sInformer.Core().V1().Pods().Lister().Pods(ns).List(selector)
}

func (w *Watcher) SetupPodsEventHandler(h cache.ResourceEventHandler) {
	w.k8sInformer.Core().V1().Pods().Informer().AddEventHandler(h)
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
