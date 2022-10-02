package service

import (
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	kpack "github.com/pivotal/kpack/pkg/client/clientset/versioned"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Service struct {
	DB            *mongo.Database
	K8sClient     *kubernetes.Clientset
	OamClient     *versioned.Clientset
	RuntimeClient client.Client
	KpackClient   *kpack.Clientset
}
