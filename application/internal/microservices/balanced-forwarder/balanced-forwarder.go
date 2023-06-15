package balanced_forwarder

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
	ErrNoHostAvailable = errors.New("no hosts are available right now")
)

type BalancedForwarder struct {
	sd            *service_discovery.ServiceDiscovery
	pool          map[string]*httputil.ReverseProxy
	index         int
	targetPort    string
	targetService services.ServiceName
	stripFromPath string
	addToPath     string
}

var log = logger.NewLogger("BalancedForwarder")

// service: ip address and host
// hosts: ip addresses of available hosts
// port: port which will be used as forwarded target
func New(sd *service_discovery.ServiceDiscovery, targetService services.ServiceName, targetPort, stripFromPath, addToPath string) *BalancedForwarder {
	return &BalancedForwarder{
		sd:            sd,
		pool:          map[string]*httputil.ReverseProxy{},
		index:         0,
		targetPort:    targetPort,
		targetService: targetService,
		stripFromPath: stripFromPath,
		addToPath:     addToPath,
	}
}

func (bf *BalancedForwarder) appendPort(host string) string {
	return fmt.Sprintf("%s%s", host, bf.targetPort)
}

func (bf *BalancedForwarder) getProxyForHost(host string) *httputil.ReverseProxy {
	if proxy, ok := bf.pool[host]; ok {
		return proxy
	} else {
		proxy = httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   bf.appendPort(host),
			Path:   bf.addToPath,
		})
		bf.pool[host] = proxy
		return proxy
	}
}

func (bf *BalancedForwarder) Next() (*httputil.ReverseProxy, error) {
	var hosts = bf.sd.ListPrivateIPs(bf.targetService)
	if len(hosts) == 0 {
		return nil, ErrNoHostAvailable
	} else if len(hosts) <= bf.index {
		bf.index = utilities.URandIntN(len(hosts))
	}
	var next = hosts[bf.index]
	bf.index = (bf.index + 1) % len(hosts)
	return bf.getProxyForHost(next), nil
}

func (bf *BalancedForwarder) Forward(w http.ResponseWriter, r *http.Request) {
	var rp, err = bf.Next()
	if err != nil {
		http.Error(w, ErrNoHostAvailable.Error(), http.StatusInternalServerError)
		return
	}
	// r.Host = getTargetAddressAndIPForHost(host)
	r.URL.Path, _ = strings.CutPrefix(r.URL.Path, bf.stripFromPath)
	log.Println("Forwarding to", r.URL.Path)
	rp.ServeHTTP(w, r)
}
