package paths

import (
	"fmt"
	"strings"
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

func (d Domain) String() string {
	return fmt.Sprintf("%s://%s:%s", d.Protocol, d.Domain, d.Port)
}

func (g Gateway) String() string {
	return fmt.Sprintf("%s%s", g.Root.String(), g.Listens)
}

func (s Service) String() string {
	return fmt.Sprintf("%s%s", s.Gateway.String(), s.Listens)
}

func (e Endpoint) String() string {
	return fmt.Sprintf("%s%s", e.Service.String(), e.Listens)
}

func checkPrefix(a Endpoint, b Endpoint) bool {
	return strings.HasPrefix(a.String(), b.String())
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

type Method string

const (
	GET    = Method("GET")
	POST   = Method("POST")
	PATCH  = Method("PATCH")
	DELETE = Method("DELETE")
	PUT    = Method("PUT")
)
