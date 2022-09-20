package traits

import (
	"fmt"

	"github.com/kiaedev/kiae/api/egress"
	"github.com/saltbo/gopkg/strutil"
)

type Egress struct {
	Hosts []string     `json:"hosts"`
	Port  *egress.Port `json:"port"`
}

type Sidecar struct {
	Egress []Egress `json:"egress"`
}

func NewSidecar(egresses []*egress.Egress) *Sidecar {
	portHosts := make(map[*egress.Port][]string)
	for _, eg := range egresses {
		for _, port := range eg.Ports {
			portHosts[port] = append(portHosts[port], fmt.Sprintf("./%s", eg.Host))
		}
	}

	results := make([]Egress, 0)
	for port, hosts := range portHosts {
		results = append(results, Egress{
			Hosts: hosts,
			Port:  port,
		})
	}
	return &Sidecar{results}
}

func (m *Sidecar) GetName() string {
	return "sidecar-" + strutil.RandomText(5)
}

func (m *Sidecar) GetType() string {
	return "k-sidecar"
}
