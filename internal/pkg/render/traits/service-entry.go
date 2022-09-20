package traits

import (
	"github.com/kiaedev/kiae/api/egress"
	"github.com/saltbo/gopkg/strutil"
)

type ServiceEntry struct {
	Hosts []string       `json:"hosts"`
	Ports []*egress.Port `json:"ports"`
}

func NewServiceEntry(egresses []*egress.Egress) *ServiceEntry {
	hosts := make([]string, 0)
	ports := make([]*egress.Port, 0)
	for _, eg := range egresses {
		if eg.Type == egress.Egress_INTERNET {
			hosts = append(hosts, eg.Host)
			ports = append(ports, eg.Ports...)
		}
	}

	return &ServiceEntry{hosts, ports}
}

func (m *ServiceEntry) GetName() string {
	return "se-" + strutil.RandomText(5)
}

func (m *ServiceEntry) GetType() string {
	return "k-service-entry"
}
