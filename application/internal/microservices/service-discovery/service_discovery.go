package service_discovery

import (
	"tde/internal/microservices/service-discovery/models"

	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type ServiceDiscovery struct {
	models.ServiceDiscoveryFile `yaml:",inline"`
	configPath                  string
	updateLock                  sync.Mutex
	updatePeriod                time.Duration
}

func NewServiceDiscovery(configPath string, updatePeriod time.Duration) *ServiceDiscovery {
	sd := ServiceDiscovery{
		configPath:   configPath,
		updatePeriod: updatePeriod,
	}
	sd.readConfig()
	go sd.tick()
	return &sd
}

func (sd *ServiceDiscovery) readConfig() {
	if !sd.updateLock.TryLock() {
		return
	}
	fileHandler, err := os.Open(sd.configPath)
	if err != nil {
		log.Fatalln(errors.Wrapf(err, "Failed to open '%s'", sd.configPath))
	}
	defer fileHandler.Close()
	err = json.NewDecoder(fileHandler).Decode(&sd)
	if err != nil {
		log.Fatalln(errors.Wrapf(err, "Failed to decode %s", sd.configPath))
	}
	sd.updateLock.Unlock()
}

func (sd *ServiceDiscovery) tick() {
	for range time.Tick(sd.updatePeriod) {
		sd.readConfig()
	}
}
