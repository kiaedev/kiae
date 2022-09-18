package mw_provider

import (
	"github.com/crossplane-contrib/provider-sql/apis"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	utilruntime.Must(apis.AddToScheme(scheme.Scheme))
}
