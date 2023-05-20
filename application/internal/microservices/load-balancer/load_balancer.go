package load_balancer

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"tde/internal/microservices/logger"
	service_discovery "tde/internal/microservices/service-discovery"
	"tde/internal/microservices/service-discovery/models/services"
	"tde/internal/utilities"
)

var (
	ErrNoHostAvailable = errors.New("No hosts are available right now. Please try again later.")
)

type LoadBalancer struct {
	sd            *service_discovery.ServiceDiscovery
	pool          map[string]*httputil.ReverseProxy
	index         int
	targetPort    string
	targetService services.ServiceName
	stripFromPath string
	addToPath     string
}

var log = logger.NewLogger("LoadBalancer")

// service: ip address and host
// hosts: ip addresses of available hosts
// port: port which will be used as forwarded target
func New(sd *service_discovery.ServiceDiscovery, targetService services.ServiceName, targetPort, stripFromPath, addToPath string) *LoadBalancer {
	return &LoadBalancer{
		sd:            sd,
		pool:          map[string]*httputil.ReverseProxy{},
		index:         0,
		targetPort:    targetPort,
		targetService: targetService,
		stripFromPath: stripFromPath,
		addToPath:     addToPath,
	}
}

func (lb *LoadBalancer) appendPort(host string) string {
	return fmt.Sprintf("%s%s", host, lb.targetPort)
}

func (lb *LoadBalancer) getProxyForHost(host string) *httputil.ReverseProxy {
	if proxy, ok := lb.pool[host]; ok {
		return proxy
	} else {
		proxy = httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   lb.appendPort(host),
			Path:   lb.addToPath,
		})
		lb.pool[host] = proxy
		return proxy
	}
}

func (lb *LoadBalancer) Next() (*httputil.ReverseProxy, error) {
	var hosts = lb.sd.ListPrivateIPs(lb.targetService)
	if len(hosts) == 0 {
		return nil, ErrNoHostAvailable
	} else if len(hosts) <= lb.index {
		lb.index = utilities.URandIntN(len(hosts))
	}
	var next = hosts[lb.index]
	lb.index = (lb.index + 1) % len(hosts)
	return lb.getProxyForHost(next), nil
}

func (lb *LoadBalancer) Forward(w http.ResponseWriter, r *http.Request) {
	var rp, err = lb.Next()
	if err != nil {
		http.Error(w, ErrNoHostAvailable.Error(), http.StatusInternalServerError)
		return
	}
	// r.Host = getTargetAddressAndIPForHost(host)
	r.URL.Path, _ = strings.CutPrefix(r.URL.Path, lb.stripFromPath)
	log.Println("Forwarding to", r.URL.Path)
	rp.ServeHTTP(w, r)
}
