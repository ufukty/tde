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
	store  map[models.Sid]*models.Subject
}

func New(config *reader.Config, sd *serviced.ServiceDiscovery) *Cache {
	return &Cache{
		// lb:    load_balancer.New(sd, services.Customs, config.APIGateway.RouterPrivate, "/customs/"),
		sd:     sd,
		config: config,
		store:  *new(map[models.Sid]*models.Subject),
	}
}

func (c *Cache) Get(id models.Sid) (*models.Subject, error) {
	if subject, ok := c.store[id]; ok {
		return subject, nil
	}

	// cache miss
	var customs = c.sd.ListPrivateIPs(services.Customs)
	if len(customs) == 0 {
		return nil, ErrNoHosts
	}

	endpoints.AstPackageSend(&endpoints.AstPackageRequest{
		ArchiveId: "",
	})

	return nil, nil
}

func (c *Cache) GetModuleAST(archiveId string) (*models.Subject, error) {
	// if subject, ok := c.store[id]; ok {
	// 	return subject, nil
	// }

	// cache miss
	var customs = c.sd.ListPrivateIPs(services.Customs)
	if len(customs) == 0 {
		return nil, ErrNoHosts
	}

	var res, err = endpoints.AstPackageSend(&endpoints.AstPackageRequest{
		ArchiveId: archiveId,
	})
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	fmt.Println(res.Package)

	return nil, nil
}

func (c *Cache) Set(id models.Sid, subject *models.Subject) error {
	c.store[id] = subject

	// update db

	return nil
}
