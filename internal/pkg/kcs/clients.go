package kcs

import (
	vela "github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	kpack "github.com/pivotal/kpack/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KubeClients struct {
	K8sCs   *kubernetes.Clientset
	VelaCs  *vela.Clientset
	KpackCs *kpack.Clientset

	RuntimeClient client.Client
}

func NewKubeClients(config *rest.Config) (*KubeClients, error) {
	k8sCs, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	runtimeClient, err := client.New(config, client.Options{})
	if err != nil {
		return nil, err
	}

	velaCs, err := vela.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	kpackCs, err := kpack.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &KubeClients{
		K8sCs:   k8sCs,
		VelaCs:  velaCs,
		KpackCs: kpackCs,

		RuntimeClient: runtimeClient,
	}, nil
}
