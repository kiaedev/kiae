package service

import (
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/client-go/kubernetes"
)

type Service struct {
	DB        *mongo.Database
	K8sClient *kubernetes.Clientset
	OamClient *versioned.Clientset
}
