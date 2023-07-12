package serviced

import (
	"tde/internal/microservices/serviced/models/services"
	"testing"
	"time"
)

func Test_ServiceDiscovery(t *testing.T) {
	var filename = "test/service_discovery_stage.json"
	sd := NewServiceDiscovery(filename, time.Second*5)

	runners := sd.ListPrivateIPs(services.Runner)
	if len(runners) != 3 {
		t.Error("validation")
	}
}

func Test_LocalConfig(t *testing.T) {
	var filename = "test/service_discovery_local.json"
	sd := NewServiceDiscovery(filename, time.Second*5)

	runners := sd.ListPrivateIPs(services.Runner)
	if len(runners) != 4 {
		t.Error("validation")
	}
}
