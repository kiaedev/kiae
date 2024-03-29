package klient

import (
	"github.com/google/wire"
	vela "github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	kpack "github.com/pivotal/kpack/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrClient "sigs.k8s.io/controller-runtime/pkg/client"
)

type LocalClients struct {
	K8sCs   *kubernetes.Clientset
	VelaCs  *vela.Clientset
	KpackCs *kpack.Clientset

	RuntimeClient ctrClient.Client
}

var ProviderSet = wire.NewSet(
	kubernetes.NewForConfig,
	vela.NewForConfig,
	kpack.NewForConfig,
	CtrRuntimeClient,
	NewProxy,
	wire.Struct(new(LocalClients), "*"),
)

func CtrRuntimeClient(config *rest.Config) (ctrClient.Client, error) {
	return ctrClient.New(config, ctrClient.Options{})
}
