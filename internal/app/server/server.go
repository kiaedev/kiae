package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
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
	"github.com/kiaedev/kiae/api/image"
	"github.com/kiaedev/kiae/api/middleware"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/api/provider"
	"github.com/kiaedev/kiae/api/route"
	"github.com/kiaedev/kiae/internal/app/server/service"
	"github.com/kiaedev/kiae/internal/app/server/watcher"
	"github.com/kiaedev/kiae/internal/pkg/kcs"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"github.com/koding/websocketproxy"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	db            *mongo.Database
	kcs           *kcs.KubeClients
	watcher       *watcher.Watcher
	graphResolver *graph.Resolver
}

func NewServer(kClients *kcs.KubeClients) (*Server, error) {
	dbClient, err := mongoutil.New(viper.GetString("dsn"))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}

	db := dbClient.DB.Database("kiae")
	w, err := watcher.NewWatcher(kClients)
	if err != nil {
		return nil, err
	}

	appPodSvc := service.NewAppPodsService(w, db, kClients)
	w.SetupPodsEventHandler(appPodSvc)
	w.SetupImagesEventHandler(service.NewImageWatcher(db, kClients))
	return &Server{
		db:  db,
		kcs: kClients,

		graphResolver: graph.NewResolver(appPodSvc),
		watcher:       w,
	}, nil
}

func (s *Server) Run() error {
	port := 8888
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gs := grpc.NewServer()
	app.RegisterAppServiceServer(gs, service.NewAppService(s.db, s.kcs))
	go func() {
		log.Printf("grpc server listening at %v", lis.Addr())
		if err := gs.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	s.runWatcher()
	s.runGraphql()
	return s.runGateway()
}

func (s *Server) runWatcher() {
	_ = s.watcher.Run(context.Background())
}

func (s *Server) runGraphql() {
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: s.graphResolver}))
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
	mux := runtime.NewServeMux(runtime.WithUnescapingMode(runtime.UnescapingModeAllExceptReserved))
	_ = provider.RegisterProviderServiceHandlerServer(ctx, mux, service.NewProviderService(s.db, s.kcs))
	_ = project.RegisterProjectServiceHandlerServer(ctx, mux, service.NewProjectService(s.db, s.kcs))
	_ = image.RegisterImageServiceHandlerServer(ctx, mux, service.NewProjectImageSvc(s.db, s.kcs))
	_ = app.RegisterAppServiceHandlerServer(ctx, mux, service.NewAppService(s.db, s.kcs))
	_ = egress.RegisterEgressServiceHandlerServer(ctx, mux, service.NewEgressService(s.db, s.kcs))
	_ = entry.RegisterEntryServiceHandlerServer(ctx, mux, service.NewEntryService(s.db, s.kcs))
	_ = route.RegisterRouteServiceHandlerServer(ctx, mux, service.NewRouteService(s.db, s.kcs))

	_ = middleware.RegisterMiddlewareServiceHandlerServer(ctx, mux, service.NewMiddlewareService(s.db, s.kcs))
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// err := app.RegisterAppServiceHandlerFromEndpoint(ctx, mux, "localhost:8888", opts)
	// if err != nil {
	// 	return err
	// }

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.Handle("/", mux)
	service.NewOauth2Service(s.db, s.kcs).SetupHandler()
	u, _ := url.Parse("ws://localhost:3100") // todo get loki url from config
	websocketproxy.DefaultUpgrader.CheckOrigin = func(req *http.Request) bool { return true }
	http.Handle("/proxies/loki/api/v1/tail", http.StripPrefix("/proxies", websocketproxy.NewProxy(u)))

	log.Printf("http server listening at %v", 8081)
	return http.ListenAndServe(":8081", nil)
}
