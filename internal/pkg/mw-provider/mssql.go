package mw_provider

import (
	"github.com/crossplane-contrib/provider-sql/apis/mssql/v1alpha1"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func MsSQLConfig(instanceName, secretName string) client.Object {
	return &v1alpha1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceName,
		},
		Spec: v1alpha1.ProviderConfigSpec{
			Credentials: v1alpha1.ProviderCredentials{
				Source: v1alpha1.CredentialsSourceMSSQLConnectionSecret,
				ConnectionSecretRef: &xpv1.SecretReference{
					Name:      secretName,
					Namespace: "kiae-system",
				},
			},
		},
	}
}
