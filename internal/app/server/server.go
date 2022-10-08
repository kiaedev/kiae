package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"net/http/httputil"
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

	s.setupProxiesEndpoints()
	s.setupGraphQLEndpoints()
	return s.runHTTPServer(ctx)
}

func (s *Server) setupProxiesEndpoints() {
	s.svcSets.Oauth2.SetupHandler()

	u, _ := url.Parse("ws://localhost:3100")    // todo get loki url from config
	u2, _ := url.Parse("http://localhost:3100") // todo get loki url from config
	websocketproxy.DefaultUpgrader.CheckOrigin = func(req *http.Request) bool { return true }
	http.Handle("/proxies/loki/api/v1/tail", http.StripPrefix("/proxies", websocketproxy.NewProxy(u)))
	http.Handle("/proxies/loki/api/v1/query_range", http.StripPrefix("/proxies", httputil.NewSingleHostReverseProxy(u2)))

	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6Ikl3WGM1ek9ZTm5FN3R3OWFOV1NVZnVLb1lJTm1YQTlCWFdzc0FnaHc2Q28ifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJraW5kbmV0LXRva2VuLTJ4Nmc0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImtpbmRuZXQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiIyNmU5ZDUwMy00MzQ1LTRlNzktOTY1MS1kZjE4Mzg3MTE0ZTAiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06a2luZG5ldCJ9.mUDCJsMe_QzAKOs6pQ3MlMttBqKLo5xv-mVeOgl8N9sZVQStAasxekTsLKTNDXbgthsOjXCaS6HRJosNSYBl80qbLicAWh8hKkMzRQBpOBl_LnIjgIAT-qIkD57XbLV_f065HuSdArRcOq_MMzRm9EQzZB4Jjg1zqmLx4GPRiunzHEC7CasEBC6MVoLKUF2Dgzf4JK-sr3nSfijSjmrHgTmsHr83JBWzpxmZT2GX82zmwLlWarNmkVhUJBEM5fWAjxGFCPfl2n8df3ePXrVXrLUr71uwTQ_oiKo8YTuTPZZMAYLGItQUC4vHTtS1tezfuH2k9bNOoZyEKG4Hg_mk7g"
	rootPEM := `-----BEGIN CERTIFICATE-----
MIIC/jCCAeagAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTIyMDkyOTE0NTAyM1oXDTMyMDkyNjE0NTAyM1owFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAM9B
6L9uia74LcWxbdIURWS8V/FLZ8SsTw1STpj562eRdYWQb1eeT5x6Z0UB64IF7LMu
qegFif/AWL9caA6gYfky4X01HRUpzQz605uBfgV+eYJbL+Voir9IHqA/gsa3oEs0
pABWeUD8xe7axmQ0ajeyoBDmC9EXD9e+TKTRn7hrLNDyY3PI0LcjvZEsmtV3T3qA
j76LC7A4pSSdt7GOyGKx3A3vvKv4dQ4hrnXqdjy99iysFb0zNyVne2Jfso9pWz5q
P7oW0NtxESU6CCqMLnp33+W8Go3+H3YlFRtshL/J2vq5GBQxMnLdTOqaxCK5b1CY
jzh4LZ0uQzTvkhrcZ2ECAwEAAaNZMFcwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wHQYDVR0OBBYEFKXHwhnenI+34bkqOZp/oOccuJkvMBUGA1UdEQQO
MAyCCmt1YmVybmV0ZXMwDQYJKoZIhvcNAQELBQADggEBAJLWb8CW5qcArn/NAYqB
9qcdYL9K9+bw7R/vE4F1wHdRvZU/Vs5lK4W8gwMlgGU4n074GsM/ziuwPeLrb2BT
lP/zIX2LYv61+fLIw8N3/85JADT6yI7hsQFf/x0buJVQV0Yax5wrALOoIUqraa8U
Wvrlmq7ADMJNnvsBxDMisz95iAkKiEX90KyH17mFiLRqm22kuNw5Eaqz3Wl1Q+Hk
zloZoPRtp22VJ1H6DJ3/6zvjovCzaHpvqyb7yxCfKEV5bZVSW3CZz0CacWqVVcUi
sfg3h/X9FBNRVT4gA65LCqp0EpCq5U27wQ69YZDAEaIeR0o7xinO+c7VHo1jVUj1
ua8=
-----END CERTIFICATE-----
`
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}
	websocketproxy.DefaultDialer.TLSClientConfig = &tls.Config{
		ServerName: "kubernetes",
		MinVersion: tls.VersionTLS12,
		RootCAs:    roots,
	}

	u2, err := url.Parse("wss://127.0.0.1:61586")
	if err != nil {
		log.Println(err)
	}

	wsp := websocketproxy.NewProxy(u2)
	wsp.Director = func(incoming *http.Request, out http.Header) {
		out.Set("Authorization", "Bearer "+token)
	}
	http.Handle("/proxies/k8s/api/v1/", http.StripPrefix("/proxies/k8s", wsp))
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
