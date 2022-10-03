package watcher

import (
	"context"
	"time"

	"github.com/kiaedev/kiae/internal/pkg/kcs"
	velaInformers "github.com/oam-dev/kubevela-core-api/pkg/generated/client/informers/externalversions"
	kpackInformers "github.com/pivotal/kpack/pkg/client/informers/externalversions"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/client-go/informers"
)

type Watcher struct {
	*kcs.KubeClients
	db *mongo.Database

	k8sInformer    informers.SharedInformerFactory
	velaInformers  velaInformers.SharedInformerFactory
	kpackInformers kpackInformers.SharedInformerFactory
}

func NewWatcher(db *mongo.Database, kcs *kcs.KubeClients) (*Watcher, error) {
	return &Watcher{
		db:             db,
		KubeClients:    kcs,
		k8sInformer:    informers.NewSharedInformerFactory(kcs.K8sCs, time.Hour),
		velaInformers:  velaInformers.NewSharedInformerFactory(kcs.VelaCs, time.Hour),
		kpackInformers: kpackInformers.NewSharedInformerFactory(kcs.KpackCs, time.Hour),
	}, nil
}

func (w *Watcher) Run(ctx context.Context) error {
	w.SetupEventHandler(ctx)
	go w.k8sInformer.Core().V1().Pods().Informer().Run(ctx.Done())
	go w.velaInformers.Core().V1beta1().Applications().Informer().Run(ctx.Done())
	go w.kpackInformers.Kpack().V1alpha2().Images().Informer().Run(ctx.Done())
	return nil
}

func (w *Watcher) SetupEventHandler(ctx context.Context) {
	// w.k8sInformer.Core().V1().Pods().Informer().AddEventHandler()
	// w.velaInformers.Core().V1beta1().Applications().Informer().AddEventHandler()
	iw := NewImageWatcher(w.db, w.KubeClients)
	go iw.checkNotDoneStatus(ctx)
	w.kpackInformers.Kpack().V1alpha2().Images().Informer().AddEventHandler(iw)
}
