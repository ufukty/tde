package runner_communicator

import (
	"tde/internal/microservices/service-discovery"

	"github.com/pkg/errors"
)

const IP_ADDRESSES_FILE_NAME = "runner_ip_addresses.txt"

var (
	ErrNoAvailableRunners  = errors.New("ErrNoAvailableRunners")
	ErrInternalServerError = errors.New("ErrInternalServerError")
)

type RunnerCommunicator struct {
	sd           *service_discovery.ServiceDiscovery
	ip_addresses []string
}

func NewRunnerCommunicator(sd *service_discovery.ServiceDiscovery) (*RunnerCommunicator, error) {
	rc := &RunnerCommunicator{
		sd: sd,
	}
	if err := rc.discover(); err != nil {
		return nil, err
	}
	return rc, nil
}

func (rc *RunnerCommunicator) discover() error {
	rc.ip_addresses = rc.sd.Runner.GetIPs()
	if !(len(rc.ip_addresses) > 0) {
		return ErrNoAvailableRunners
	}
	return nil
}

func (rc *RunnerCommunicator) sendToRunner(runner string, batch *Batch) error {
	req := batch.GetRequestDTO()
	_, err := req.Send("POST", "http://"+runner)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func (rc *RunnerCommunicator) Send(batch *Batch) error {
	if err := rc.discover(); err != nil {
		return ErrNoAvailableRunners
	}
	batches := batch.Divide(len(rc.ip_addresses))
	for i, batch := range batches {
		rc.sendToRunner(rc.ip_addresses[i], batch)
	}
	return nil
}
