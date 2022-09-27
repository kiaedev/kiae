package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/egress"
	"github.com/kiaedev/kiae/api/entry"
	"github.com/kiaedev/kiae/api/graph"
	"github.com/kiaedev/kiae/api/graph/generated"
	"github.com/kiaedev/kiae/api/middleware"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/api/route"
	"github.com/kiaedev/kiae/internal/app/server/service"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
	runtimeClient, err := client.New(config, client.Options{})
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
			DB:            dbClient.DB.Database("kiae"),
			K8sClient:     k8sClientSet,
			RuntimeClient: runtimeClient,
			OamClient:     oamClientSet,
		},
	}, nil
}

func (s *Server) Run() error {
	port := 8888
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gs := grpc.NewServer()
	app.RegisterAppServiceServer(gs, service.NewAppService(s.ss))
	go func() {
		log.Printf("grpc server listening at %v", lis.Addr())
		if err := gs.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	resolver := graph.NewResolver(s.ss.K8sClient)
	if err := resolver.Run(context.Background()); err != nil {
		return err
	}

	s.runGraphql(resolver)
	return s.runGateway()
}

func (s *Server) runGraphql(resolver *graph.Resolver) {
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	http.Handle("/api/graphql", srv)
	http.Handle("/graphql", playground.Handler("My GraphQL App", "/api/graphql"))
}

func (s *Server) runGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	_ = project.RegisterProjectServiceHandlerServer(ctx, mux, service.NewProjectService(s.ss))
	_ = project.RegisterImageServiceHandlerServer(ctx, mux, service.NewProjectImageSvc(s.ss))
	_ = app.RegisterAppServiceHandlerServer(ctx, mux, service.NewAppService(s.ss))
	_ = egress.RegisterEgressServiceHandlerServer(ctx, mux, service.NewEgressService(s.ss))
	_ = entry.RegisterEntryServiceHandlerServer(ctx, mux, service.NewEntryService(s.ss))
	_ = route.RegisterRouteServiceHandlerServer(ctx, mux, service.NewRouteService(s.ss))

	_ = middleware.RegisterMiddlewareServiceHandlerServer(ctx, mux, service.NewMiddlewareService(s.ss))
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// err := app.RegisterAppServiceHandlerFromEndpoint(ctx, mux, "localhost:8888", opts)
	// if err != nil {
	// 	return err
	// }

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.Handle("/", mux)
	log.Printf("http server listening at %v", 8081)
	return http.ListenAndServe(":8081", nil)
}
