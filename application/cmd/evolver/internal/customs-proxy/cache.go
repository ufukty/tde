package customs_proxy

import (
	"tde/cmd/customs/endpoints"
	"tde/config/reader"
	"tde/internal/microservices/serviced"
	"tde/internal/microservices/serviced/models/services"
	models "tde/models/program"

	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrNoHosts = errors.New("No instances of customs are available. Check service discovery file.")
)

type Cache struct {
	config *reader.Config
	sd     *serviced.ServiceDiscovery
	store  map[models.CandidateID]*models.Candidate
}

func New(config *reader.Config, sd *serviced.ServiceDiscovery) *Cache {
	return &Cache{
		// lb:    load_balancer.New(sd, services.Customs, config.APIGateway.RouterPrivate, "/customs/"),
		sd:     sd,
		config: config,
		store:  *new(map[models.CandidateID]*models.Candidate),
	}
}

func (c *Cache) Get(id models.CandidateID) (*models.Candidate, error) {
	if candidate, ok := c.store[id]; ok {
		return candidate, nil
	}

	// cache miss
	var customs = c.sd.ListPrivateIPs(services.Customs)
	if len(customs) == 0 {
		return nil, ErrNoHosts
	}

	var req = endpoints.AstPackageRequest{
		ArchiveId: "",
	}
	req.NewRequest("GET", fmt.Sprintf("%s:%s/ast", customs[0], c.config.Customs.RouterPrivate))

	return nil, nil
}

func (c *Cache) GetModuleAST(archiveId string) (*models.Candidate, error) {
	// if candidate, ok := c.store[id]; ok {
	// 	return candidate, nil
	// }

	// cache miss
	var customs = c.sd.ListPrivateIPs(services.Customs)
	if len(customs) == 0 {
		return nil, ErrNoHosts
	}

	var req = endpoints.AstPackageRequest{
		ArchiveId: archiveId,
	}

	var res, err = req.Send("GET", fmt.Sprintf("%s:%s/ast", customs[0], c.config.Customs.RouterPrivate))
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	fmt.Println(res.Package)

	return nil, nil
}

func (c *Cache) Set(id models.CandidateID, candidate *models.Candidate) error {
	c.store[id] = candidate

	// update db

	return nil
}
