package provider

import "tde/internal/microservices/serviced/models/services"

type Provider interface {
	ListPrivateIPs(services.ServiceName) []string
}
