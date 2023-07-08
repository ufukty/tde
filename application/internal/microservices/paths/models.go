package paths

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/exp/maps"
)

type Domain struct {
	Protocol string // eg. http, https
	Domain   string // eg. deepthinker.app
	Port     string // eg. 8080
}

type Gateway struct {
	Root    Domain
	Listens string // eg. /api/v1.0.0
}

type Service struct {
	Gateway Gateway
	Listens string
}

type Endpoint struct {
	Service Service
	Listens string
	Method  Method
}

func (d Domain) Url() string {
	return fmt.Sprintf("%s://%s:%s", d.Protocol, d.Domain, d.Port)
}

func (g Gateway) Url() string {
	return fmt.Sprintf("%s%s", g.Root.Url(), g.Listens)
}

func (s Service) Url() string {
	return fmt.Sprintf("%s%s", s.Gateway.Url(), s.Listens)
}

func (e Endpoint) Url() string {
	return fmt.Sprintf("%s%s", e.Service.Url(), e.Listens)
}

func checkPrefix(a Endpoint, b Endpoint) bool {
	return strings.HasPrefix(a.Url(), b.Url())
}

func Sort(eps []Endpoint) []Endpoint {
	for i := 0; i < len(eps); i++ {
		for j := 1; j < len(eps); j++ {
			if checkPrefix(eps[j], eps[j-1]) {
				eps[j-1], eps[j] = eps[j], eps[j-1]
			}
		}
	}
	return eps
}

//go:generate stringer -type=Method
type Method int

const (
	GET = Method(iota)
	POST
	PATCH
	DELETE
	PUT
)

func RouteRegisterer(handlers map[Endpoint]http.HandlerFunc) func(*mux.Router) {
	return func(r *mux.Router) {
		r = r.UseEncodedPath()
		for _, ep := range Sort(maps.Keys(handlers)) {
			var handler = handlers[ep]
			log.Printf("Registering route: %-6s %s\n", ep.Method, ep.Listens)
			r.PathPrefix(ep.Listens).Methods(ep.Method.String()).HandlerFunc(handler)
		}
	}
}
