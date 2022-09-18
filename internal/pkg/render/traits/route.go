package traits

import (
	"net/url"

	"github.com/kiaedev/kiae/api/entry"
	"github.com/kiaedev/kiae/api/route"
)

type RouteTrait struct {
	Name     string    `json:"name"`
	Gateways []Gateway `json:"gateways,omitempty"`
	Routes   []Route   `json:"routes,omitempty"`
}

func NewRouteTrait(name string, entries []*entry.Entry, rs []*route.Route) *RouteTrait {
	gateways := make([]Gateway, 0, len(entries))
	for _, e := range entries {
		gateways = append(gateways, Gateway{e.Gateway, e.Host})
	}

	routes := make([]Route, 0, len(rs))
	for _, r := range rs {
		nr := Route{
			URI:        r.Path,
			Methods:    r.Methods,
			ExactMatch: false,
		}
		if r.Type == route.Route_REDIRECT {
			nr.Redirect = &Redirect{URL: NewURL(r.Redirect.Url), Code: r.Redirect.Code}
		} else if r.Type == route.Route_DIRECT_RESPONSE {
			nr.DirectResponse = r.GetMock()
		} else {
			nr.Forward = r.GetForward()
		}

		routes = append(routes, nr)
	}
	routes = append(routes, Route{Name: "default"})
	return &RouteTrait{
		Name:     name,
		Gateways: gateways,
		Routes:   routes,
	}
}

func (c *RouteTrait) GetName() string {
	return c.Name + "-routes"
}

func (c *RouteTrait) GetType() string {
	return "k-route"
}

type Gateway struct {
	Name string `json:"name"`
	Host string `json:"host"`
}

type Route struct {
	Name           string                `json:"name"`
	URI            string                `json:"uri,omitempty"`
	Methods        []string              `json:"methods,omitempty"`
	ExactMatch     bool                  `json:"exactMatch,omitempty"`
	DirectResponse *route.DirectResponse `json:"directResponse,omitempty"`
	Redirect       *Redirect             `json:"redirect,omitempty"`
	Forward        *route.Forward        `json:"forward,omitempty"`
}

type Redirect struct {
	URL  URL    `json:"url,omitempty"`
	Code uint32 `json:"code,omitempty"`
}

type URL struct {
	Schema string `json:"schema,omitempty"`
	Host   string `json:"host,omitempty"`
	Path   string `json:"path,omitempty"`
}

func NewURL(urlStr string) URL {
	u, _ := url.Parse(urlStr)
	return URL{
		Schema: u.Scheme,
		Host:   u.Host,
		Path:   u.Path,
	}
}
