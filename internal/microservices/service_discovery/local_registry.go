package service_discovery

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

var (
	filename     = "~/service_discovery.json"
	updatePeriod = time.Second * 5
)

type ServiceDiscovery struct {
	updateLock  sync.Mutex
	fileContent serviceDiscoveryFile
}

func NewServiceDiscovery() *ServiceDiscovery {
	sd := ServiceDiscovery{}
	sd.readConfig()
	go sd.tick()
	return &sd
}

func (sd *ServiceDiscovery) readConfig() {
	if !sd.updateLock.TryLock() {
		return
	}
	log.Println("ServiceDiscovery: reading config file")
	fileHandler, err := os.Open(filename)
	if err != nil {
		panic(errors.Wrapf(err, "Failed to open %s", filename))
	}
	defer fileHandler.Close()
	err = json.NewDecoder(fileHandler).Decode(&sd.fileContent)
	if err != nil {
		panic(errors.Wrapf(err, "Failed to decode %s", filename))
	}
	log.Println("ServiceDiscovery: reading config file (done)")
	sd.updateLock.Unlock()
}

func (sd *ServiceDiscovery) tick() {
	for range time.Tick(updatePeriod) {
		sd.readConfig()
	}
}

type Kind string

const (
	Runner  = Kind("runner")
	Evolver = Kind("evolver")
)

func (sd *ServiceDiscovery) LookupKind(kind Kind) (ip_addresses []string) {
	spew.Println(sd.fileContent.Runner.Digitalocean)
	switch kind {
	case Runner:
		return sd.fileContent.Runner.GetIPs()

	case Evolver:
		return sd.fileContent.Evolver.GetIPs()
	}
	return nil
}
