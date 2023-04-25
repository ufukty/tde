package service_discovery

import (
	"testing"
)

func Test_ServiceDiscovery(t *testing.T) {
	filename = "service_discovery_test.json"

	reg := NewServiceDiscovery()

	ip_addresses := reg.LookupKind(Runner)
	if len(ip_addresses) == 0 {
		t.Error("validation")
	}

	first := ip_addresses[0]
	if first != "10.150.0.3" {
		t.Error("validation 2")
	}
}
