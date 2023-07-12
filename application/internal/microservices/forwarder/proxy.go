package forwarder

import (
	"tde/internal/microservices/loadbalancer"
	"tde/internal/microservices/serviced"
	"tde/internal/microservices/serviced/models/services"

	"net/http"
	"net/http/httputil"
	"strings"
)

type reverseProxyPool struct {
	lb        *loadbalancer.LoadBalancer
	pool      map[string]*httputil.ReverseProxy // host:handler
	generator func(string) *httputil.ReverseProxy
}

func newPool(lb *loadbalancer.LoadBalancer, generator func(host string) *httputil.ReverseProxy) *reverseProxyPool {
	return &reverseProxyPool{
		pool:      map[string]*httputil.ReverseProxy{},
		lb:        lb,
		generator: generator,
	}
}

func (rpp *reverseProxyPool) Get() (*httputil.ReverseProxy, error) {
	var next, err = rpp.lb.Next()
	if err == loadbalancer.ErrNoHostAvailable {
		return nil, err
	}
	if _, ok := rpp.pool[next]; !ok {
		rpp.pool[next] = rpp.generator(next)
	}
	return rpp.pool[next], nil
}

func newReverseProxyGenerator(pathByGateway, port string) func(host string) *httputil.ReverseProxy {
	return func(host string) *httputil.ReverseProxy {
		var target = host + ":" + port
		// see link to understand usage of rewrite
		// https://www.ory.sh/hop-by-hop-header-vulnerability-go-standard-library-reverse-proxy/
		return &httputil.ReverseProxy{
			Rewrite: func(pr *httputil.ProxyRequest) {
				pr.SetXForwarded()

				pr.Out.URL.Scheme = "http"
				pr.Out.URL.Host = target
				pr.Out.URL.Path = strings.TrimPrefix(pr.In.URL.Path, pathByGateway)
				pr.Out.URL.RawPath = strings.TrimPrefix(pr.In.URL.RawPath, pathByGateway)
			},
		}
	}
}

func generateLoadBalancedProxyHandler(pool *reverseProxyPool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var forwarder, err = pool.Get()
		if err != nil {
			http.Error(w, "Service you want to access is not available at the moment. Please try again later.", http.StatusBadGateway)
			return
		}
		forwarder.ServeHTTP(w, r)
	}
}

func NewLoadBalancedProxy(sd *serviced.ServiceDiscovery, service services.ServiceName, port string, pathByGateway string) (http.HandlerFunc, error) {
	var lb = loadbalancer.New(sd, service)
	if _, err := lb.Next(); err == loadbalancer.ErrNoHostAvailable {
		return nil, err
	}
	var pool = newPool(lb, newReverseProxyGenerator(pathByGateway, port))
	return generateLoadBalancedProxyHandler(pool), nil
}
