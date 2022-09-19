package traits

import (
	"github.com/kiaedev/kiae/api/egress"
	"github.com/saltbo/gopkg/strutil"
)

type Sidecar struct {
	Egress []Egress `json:"egress"`
}

type Egress struct {
	Host     string `json:"host"`
	Port     uint32 `json:"port"`
	Protocol string `json:"protocol"`
	External bool   `json:"external"`
}

func NewSidecar(egresses []*egress.Egress) *Sidecar {
	results := make([]Egress, 0)
	for _, eg := range egresses {
		results = append(results, Egress{
			Host:     eg.Host,
			Port:     eg.Port,
			Protocol: eg.Protocol,
			External: eg.Type == egress.Egress_INTERNET,
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
