package kubeproxy

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/koding/websocketproxy"
	"github.com/saltbo/gopkg/httputil"
)

type Proxy struct {
}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func test() {
	// todo implement the cluster management
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
	_ = roots.AppendCertsFromPEM([]byte(rootPEM))
	tlsConfig := &tls.Config{
		ServerName: "kubernetes",
		MinVersion: tls.VersionTLS12,
		RootCAs:    roots,
	}

	// http proxy
	ku, _ := url.Parse("https://127.0.0.1:61586")
	var authnHeaders = make(http.Header)
	authnHeaders.Set("Authorization", "Bearer "+token)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = tlsConfig
	http.Handle("/proxies/k8s/", http.StripPrefix("/proxies/k8s", httputil.NewReverseProxy(ku, authnHeaders)))

	// websocket proxy
	kuw, _ := url.Parse(ku.String())
	kuw.Scheme = "wss"
	websocketproxy.DefaultDialer.TLSClientConfig = tlsConfig
	wsp := websocketproxy.NewProxy(kuw)
	wsp.Director = func(incoming *http.Request, out http.Header) {
		// TODO: fetch the auth token from db
		out.Set("Authorization", "Bearer "+token)
	}
	router := mux.NewRouter()
	router.Handle("/k8s/api/v1/namespaces/{ns}/pods/{pod}/exec", http.StripPrefix("/k8s", wsp))
}
