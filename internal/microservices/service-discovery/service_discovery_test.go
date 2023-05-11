package service_discovery

import (
	"testing"
	"time"
)

func Test_ServiceDiscovery(t *testing.T) {
	var filename = "test/service_discovery_stage.json"
	sd := NewServiceDiscovery(filename, time.Second*5)

	runners := sd.Runner.ListPrivateIPs()
	if len(runners) != 10 {
		t.Error("validation")
	}
}
