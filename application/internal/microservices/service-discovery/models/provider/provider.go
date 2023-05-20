package provider

import "tde/internal/microservices/service-discovery/models/services"

type Provider interface {
	ListPrivateIPs(services.ServiceName) []string
}
