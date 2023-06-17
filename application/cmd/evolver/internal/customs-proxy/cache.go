package customs_proxy

import (
	config_reader "tde/internal/microservices/config-reader"
	load_balancer "tde/internal/microservices/load-balancer"
	service_discovery "tde/internal/microservices/service-discovery"
	"tde/internal/microservices/service-discovery/models/services"
	models "tde/models/program"

	"github.com/pkg/errors"
)

type Cache struct {
	lb    *load_balancer.LoadBalancer
	store map[models.CandidateID]*models.Candidate
}

func New(config *config_reader.Config, sd *service_discovery.ServiceDiscovery) *Cache {
	return &Cache{
		lb:    load_balancer.New(sd, services.Customs, config.APIGateway.RouterPrivate, "/customs/"),
		store: *new(map[models.CandidateID]*models.Candidate),
	}
}

func (c *Cache) Get(id models.CandidateID) (*models.Candidate, error) {
	if candidate, ok := c.store[id]; ok {
		return candidate, nil
	}

	// cache miss
	var customs, err = c.lb.Next()
	if err != nil {
		return nil, errors.Wrap(err, "No instances of Customs are accessible")
	}

	return nil, nil
}

func (c *Cache) Set(id models.CandidateID, candidate *models.Candidate) error {
	c.store[id] = candidate

	// update db

	return nil
}
