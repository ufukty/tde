package service_discovery

import (
	"testing"
)

func Test_ServiceDiscovery(t *testing.T) {
	var filename = "test/service_discovery_stage.json"
	sd := NewServiceDiscovery(filename)

	runners := sd.Runner.GetIPs()
	if len(runners) != 10 {
		t.Error("validation")
	}
}
