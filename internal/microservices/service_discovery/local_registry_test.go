package service_discovery

import (
	"testing"
	"time"
)

func Test_ServiceDiscovery(t *testing.T) {
	filename = "service_discovery_test.json"

	reg := NewServiceDiscovery()

	ip_addresses := reg.LookupKind(Runner)
	if len(ip_addresses) == 0 {
		t.Error("validation")
	}

	time.Sleep(10 * time.Second)

	// first := ip_addresses[0]
}
