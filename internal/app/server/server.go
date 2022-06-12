package server

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"github.com/openkos/openkos/api/app"
	"github.com/openkos/openkos/api/settings"
	"github.com/openkos/openkos/internal/app/server/service"
)

func Run() {
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	panic(err.Error())
	// }

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "conf.d/cce.dev.me.conf"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	oamClientSet, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	port := 8888
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	app.RegisterAppServer(s, service.NewAppStore(oamClientSet))
	settings.RegisterSettingsServer(s, service.NewSettings(clientset))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
