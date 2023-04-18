package runner_communicator

import (
	"os"
	"strings"
	"tde/internal/utilities"
	"tde/models"

	"github.com/pkg/errors"
)

const IP_ADDRESSES_FILE_NAME = "runner_ip_addresses.txt"

type RunnerCommunicator struct {
	ip_addresses []string
}

func NewRunnerCommunicator() (*RunnerCommunicator, error) {
	rm := &RunnerCommunicator{}

	ips, err := os.ReadFile(IP_ADDRESSES_FILE_NAME)
	if err != nil {
		return nil, errors.Wrap(err, "could not learn the ip addresses of runners")
	}

	rm.ip_addresses = strings.Split(string(ips), "\n")

	if !(len(rm.ip_addresses) > 0) {
		return nil, errors.New("no runners")
	}

	return rm, nil
}

func (rc *RunnerCommunicator) bundleCandidates(candidates []*models.Candidate) [][]*models.Candidate {
	var (
		noCandidates = len(candidates)
		noRunners    = len(rc.ip_addresses)
		batchSize    = utilities.Ceil(float64(noCandidates) / float64(noRunners))
	)

	var bundles = make([][]*models.Candidate, noRunners)
	for range utilities.Range(noRunners) {
		bundles = append(bundles, make([]*models.Candidate, batchSize))
	}

	for i, candidate := range candidates {
		bundles[i%noRunners] = append(bundles[i%noRunners], candidate)
	}

	return bundles
}

func (rc *RunnerCommunicator) Send(candidates []*models.Candidate) {
	bundles := bundleCandidates(candidates)

}
