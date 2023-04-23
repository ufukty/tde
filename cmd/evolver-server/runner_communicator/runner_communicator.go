package runner_communicator

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	models "tde/models/transfer"

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

func (rc *RunnerCommunicator) sendToRunner(runner string, batch *Batch) error {
	fmt.Println(runner, batch)

	reqDTO := batch.GetRequestDTO()
	req, err := reqDTO.NewRequest("POST", "https://localhost")
	if err != nil {
		return errors.Wrap(err, "")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "")
	}

	resDTO := models.RunnerService_NewTest_Response{}
	err = resDTO.DeserializeResponse(res)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

func (rc *RunnerCommunicator) Send(batch *Batch) {
	batches := batch.Divide(len(rc.ip_addresses))
	for i, batch := range batches {
		rc.sendToRunner(rc.ip_addresses[i], batch)
	}
}
