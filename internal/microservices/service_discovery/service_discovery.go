package service_discovery

import (
	"tde/internal/microservices/service_discovery/models"

	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type ServiceDiscovery struct {
	configPath   string
	updateLock   sync.Mutex
	fileContent  models.File
	updatePeriod time.Duration
}

func NewServiceDiscovery(configPath string) *ServiceDiscovery {
	sd := ServiceDiscovery{
		configPath:   configPath,
		updatePeriod: time.Second * 5,
	}
	sd.readConfig()
	go sd.tick()
	return &sd
}

func (sd *ServiceDiscovery) readConfig() {
	if !sd.updateLock.TryLock() {
		return
	}
	log.Println("ServiceDiscovery: reading config file")
	fileHandler, err := os.Open(sd.configPath)
	if err != nil {
		log.Fatalln(errors.Wrapf(err, "Failed to open '%s'", sd.configPath))
	}
	defer fileHandler.Close()
	err = json.NewDecoder(fileHandler).Decode(&sd.fileContent)
	if err != nil {
		log.Fatalln(errors.Wrapf(err, "Failed to decode %s", sd.configPath))
	}
	log.Println("ServiceDiscovery: reading config file (done)")
	sd.updateLock.Unlock()
}

func (sd *ServiceDiscovery) tick() {
	for range time.Tick(sd.updatePeriod) {
		sd.readConfig()
	}
}

type Kind string

const (
	Runner  = Kind("runner")
	Evolver = Kind("evolver")
)

func (sd *ServiceDiscovery) LookupKind(kind Kind) (ip_addresses []string) {
	switch kind {
	case Runner:
		return sd.fileContent.Runner.GetIPs()

	case Evolver:
		return sd.fileContent.Evolver.GetIPs()
	}
	return nil
}
