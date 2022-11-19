package components

import (
	"fmt"

	"github.com/kiaedev/kiae/api/gateway"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
)

type IstioGateway struct {
	Name          string        `json:"name"`
	Hosts         []string      `json:"hosts,omitempty"`
	HttpsEnabled  bool          `json:"https_enabled,omitempty"`
	HttpsRedirect bool          `json:"https_redirect,omitempty"`
	CertIssuer    string        `json:"cert_issuer,omitempty"`
	CustomPorts   *gateway.Port `json:"custom_ports,omitempty"`
}

func NewIstioGateway(gw *gateway.Gateway) *IstioGateway {
	wrapHosts := func(hosts []string) []string {
		newHosts := make([]string, 0, len(hosts))
		for _, host := range hosts {
			newHosts = append(newHosts, fmt.Sprintf("*.%s", host))
		}

		return newHosts
	}
	return &IstioGateway{
		Name:          gw.Name,
		Hosts:         wrapHosts(gw.Hosts),
		HttpsEnabled:  gw.HttpsEnabled,
		HttpsRedirect: gw.HttpsRedirect,
		CertIssuer:    gw.CertIssuer,
		CustomPorts:   gw.CustomPorts,
	}
}

func (m *IstioGateway) GetName() string {
	return m.Name
}

func (m *IstioGateway) GetType() string {
	return "k-istio-gateway"
}

func (m *IstioGateway) GetTraits() []common.ApplicationTrait {
	return nil
}
