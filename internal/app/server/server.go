package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/graph"
	"github.com/kiaedev/kiae/api/graph/generated"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/internal/app/server/service"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Server struct {
	ss *service.Service
}

func NewServer(kubeconfig string) (*Server, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	k8sClientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	oamClientSet, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	dbClient, err := mongoutil.New(viper.GetString("dsn"))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}
	return &Server{
		ss: &service.Service{
			DB:        dbClient.DB.Database("kiae"),
			K8sClient: k8sClientSet,
			OamClient: oamClientSet,
		},
	}, nil
}

func (s *Server) Run() error {
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	panic(err.Error())
	// }

	resolver := graph.NewResolver(s.ss.K8sClient)
	if err := resolver.Run(context.Background()); err != nil {
		return err
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	http.Handle("/graphql", srv)
	http.Handle("/graphiql", playground.Handler("My GraphQL App", "/graphql"))

	port := 8888
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gs := grpc.NewServer()
	app.RegisterAppServiceServer(gs, service.NewAppStore(s.ss))
	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := gs.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return s.runGateway()
}

func (s *Server) runGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	_ = app.RegisterAppServiceHandlerServer(ctx, mux, service.NewAppStore(s.ss))
	_ = project.RegisterProjectServiceHandlerServer(ctx, mux, service.NewProjectService(s.ss))
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// err := app.RegisterAppServiceHandlerFromEndpoint(ctx, mux, "localhost:8888", opts)
	// if err != nil {
	// 	return err
	// }

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.Handle("/", mux)
	return http.ListenAndServe(":8081", nil)
}
