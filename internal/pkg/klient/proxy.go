package klient

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/koding/websocketproxy"
	"k8s.io/client-go/rest"
)

type Proxy struct {
	localHpProxy *httputil.ReverseProxy
	localWsProxy *websocketproxy.WebsocketProxy
}

func NewProxy(cfg *rest.Config) *Proxy {
	target, _ := url.Parse(cfg.Host)
	tlsConfig, _ := rest.TLSConfigFor(cfg)

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = tlsConfig.Clone()
	wTransport, _ := rest.HTTPWrappersForConfig(cfg, transport)
	localHpProxy := httputil.NewSingleHostReverseProxy(target)
	localHpProxy.Transport = wTransport

	wssScheme := func(u *url.URL) *url.URL { ur := *u; ur.Scheme = "wss"; return &ur }
	localWsProxy := websocketproxy.NewProxy(wssScheme(target))
	localWsProxy.Dialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig:  tlsConfig.Clone(),
	}
	return &Proxy{
		localHpProxy: localHpProxy,
		localWsProxy: localWsProxy,
	}
}

func (p *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	cgw := req.Header.Get("X-Cluster-Gateway")
	if cgw != "" {
		clusterProxyPathFormat := "/cluster.core.oam.dev/v1alpha1/clustergateways/%s/%s"
		req.URL.Path = fmt.Sprintf(clusterProxyPathFormat, cgw, req.URL.Path)
	}

	if req.Header.Get("Upgrade") == "websocket" {
		p.localWsProxy.ServeHTTP(rw, req)
		return
	}

	// fixme: debug why not support websocket, can be replace the localWsProxy?
	p.localHpProxy.ServeHTTP(rw, req)
}
