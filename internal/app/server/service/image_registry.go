package service

import (
	"context"
	"fmt"

	"github.com/kiaedev/kiae/api/image"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/pkg/klient"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

type ImageRegistrySvc struct {
	imageRegistryDao *dao.ImageRegistryDao

	kubeSecret v1.SecretInterface
}

func NewImageRegistrySvc(imageRegistryDao *dao.ImageRegistryDao, kClients *klient.LocalClients) *ImageRegistrySvc {
	return &ImageRegistrySvc{
		imageRegistryDao: imageRegistryDao,

		kubeSecret: kClients.K8sCs.CoreV1().Secrets("kiae-system"),
	}
}

func (s *ImageRegistrySvc) List(ctx context.Context, in *image.RegistryListRequest) (*image.RegistryListResponse, error) {
	query := bson.M{}
	results, total, err := s.imageRegistryDao.List(ctx, query)
	return &image.RegistryListResponse{Items: results, Total: total}, err
}

func (s *ImageRegistrySvc) Create(ctx context.Context, in *image.Registry) (*image.Registry, error) {
	registrySecret := buildRegistrySecret(in)
	_, err := s.kubeSecret.Create(ctx, registrySecret, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return s.imageRegistryDao.Create(ctx, in)
}

func (s *ImageRegistrySvc) Update(ctx context.Context, in *image.Registry) (*image.Registry, error) {
	registrySecret := buildRegistrySecret(in)
	_, err := s.kubeSecret.Update(ctx, registrySecret, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return s.imageRegistryDao.Update(ctx, in)
}

func (s *ImageRegistrySvc) Delete(ctx context.Context, in *kiae.IdRequest) (*emptypb.Empty, error) {
	_, err := s.imageRegistryDao.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, s.imageRegistryDao.Delete(ctx, in.Id)
}

func buildRegistrySecret(in *image.Registry) *corev1.Secret {
	registrySecret := &corev1.Secret{}
	registrySecret.SetName(RegistrySecretName(in))
	registrySecret.SetAnnotations(map[string]string{
		"kpack.io/docker": in.Server,
	})
	registrySecret.Type = "kubernetes.io/basic-auth"
	registrySecret.StringData = map[string]string{
		"username": in.Username,
		"password": in.Password,
	}
	return registrySecret
}

func RegistrySecretName(reg *image.Registry) string {
	return fmt.Sprintf("kpack-reg-%s", reg.Name)
}
