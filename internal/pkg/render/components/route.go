package components

import (
	"github.com/kiaedev/kiae/api/entry"
	"github.com/kiaedev/kiae/api/route"
)

type RouteComponent struct {
	Name     string    `json:"name"`
	Gateways []Gateway `json:"gateways,omitempty"`
	Routes   []Route   `json:"routes,omitempty"`
}

func NewRouteComponent(name string, entries []*entry.Entry, rs []*route.Route) *RouteComponent {
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
		if r.GetForward() != nil {
			nr.Forward = r.GetForward()
		}
		if r.GetRedirect() != nil {
			nr.Redirect = r.GetRedirect()
		}
		if r.GetMock() != nil {
			nr.DirectResponse = r.GetMock()
		}

		routes = append(routes, nr)
	}
	routes = append(routes, Route{Name: "default"})
	return &RouteComponent{
		Name:     name,
		Gateways: gateways,
		Routes:   routes,
	}
}

func (c *RouteComponent) GetName() string {
	return c.Name + "-routes"
}

func (c *RouteComponent) GetType() string {
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
	Redirect       *route.Redirect       `json:"redirect,omitempty"`
	Forward        *route.Forward        `json:"forward,omitempty"`
}
