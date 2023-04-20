package runner_communicator

import (
	"tde/models"
	"testing"
)

func Test_Send(t *testing.T) {
	var (
		rc = RunnerCommunicator{
			ip_addresses: []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"},
		}
		batch = Batch{
			FileTemplate: "",
			Candidates: []*models.Candidate{
				{UUID: "MR"}, {UUID: "45"}, {UUID: "LJ"}, {UUID: "1A"}, {UUID: "Nb"}, {UUID: "nP"}, {UUID: "gJ"},
				{UUID: "WI"}, {UUID: "GS"}, {UUID: "KV"}, {UUID: "Q0"}, {UUID: "et"}, {UUID: "XA"}, {UUID: "Nx"},
				{UUID: "6n"}, {UUID: "4p"}, {UUID: "N9"}, {UUID: "zH"}, {UUID: "Jy"}, {UUID: "BP"}, {UUID: "Rq"},
				{UUID: "i8"}, {UUID: "DI"}, {UUID: "Ct"}, {UUID: "tw"}, {UUID: "m7"}, {UUID: "Lw"}, {UUID: "YI"},
				{UUID: "lm"},
			},
		}
	)

	rc.Send(&batch)
}
