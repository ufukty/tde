package evolution

import (
	"tde/internal/microservices/service_discovery"
	models "tde/models/program"
)

type RemoteEvaluator struct {
	sd *service_discovery.ServiceDiscovery
}

func (re *RemoteEvaluator) Pipeline(candidates []*models.Candidate) {
	ip_addresses := re.sd.LookupKind(service_discovery.Runner)
	return ip_addresses
}
