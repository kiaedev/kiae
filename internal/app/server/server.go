package server

import (
	"context"
	"log"
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
	"github.com/kiaedev/kiae/internal/app/server/watch"
	"github.com/kiaedev/kiae/internal/pkg/kcs"
	"github.com/koding/websocketproxy"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/client-go/rest"
)

type Server struct {
	db            *mongo.Database
	kcs           *kcs.KubeClients
	watcher       *watch.Watcher
	graphResolver *graph.Resolver

	svcSets *service.ServiceSets
}

func NewServer(config *rest.Config) (*Server, error) {
	return buildInjectors(config)
}

func (s *Server) Run(ctx context.Context) error {
	s.watcher.Start(ctx)

	s.setupGraphQLEndpoints()
	return s.runHTTPServer(ctx)
}

func (s *Server) setupProxiesEndpoints() {
	// service.NewOauth2Service(s.db, s.kcs).SetupHandler()

	u, _ := url.Parse("ws://localhost:3100") // todo get loki url from config
	websocketproxy.DefaultUpgrader.CheckOrigin = func(req *http.Request) bool { return true }
	http.Handle("/proxies/loki/api/v1/tail", http.StripPrefix("/proxies", websocketproxy.NewProxy(u)))
}

func (s *Server) setupGraphQLEndpoints() {
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

func (s *Server) runHTTPServer(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := []runtime.ServeMuxOption{
		runtime.WithUnescapingMode(runtime.UnescapingModeAllExceptReserved),
	}
	mux := runtime.NewServeMux(opts...)
	s.setupEndpoints(ctx, mux)
	http.Handle("/", mux)

	log.Printf("http server listening at %v", 8081)
	return http.ListenAndServe(":8081", nil)
}

func (s *Server) setupEndpoints(ctx context.Context, mux *runtime.ServeMux) {
	_ = provider.RegisterProviderServiceHandlerServer(ctx, mux, s.svcSets.ProviderService)
	_ = project.RegisterProjectServiceHandlerServer(ctx, mux, s.svcSets.ProjectService)
	_ = image.RegisterImageServiceHandlerServer(ctx, mux, s.svcSets.ProjectImageSvc)
	_ = app.RegisterAppServiceHandlerServer(ctx, mux, s.svcSets.AppService)
	_ = egress.RegisterEgressServiceHandlerServer(ctx, mux, s.svcSets.EgressService)
	_ = entry.RegisterEntryServiceHandlerServer(ctx, mux, s.svcSets.EntryService)
	_ = route.RegisterRouteServiceHandlerServer(ctx, mux, s.svcSets.RouteService)
	_ = middleware.RegisterMiddlewareServiceHandlerServer(ctx, mux, s.svcSets.MiddlewareService)

	s.watcher.SetupPodsEventHandler(s.svcSets.AppPodsService)
	s.watcher.SetupApplicationsEventHandler(s.svcSets.AppStatusService)
	s.watcher.SetupImagesEventHandler(s.svcSets.ImageWatcher)
}
