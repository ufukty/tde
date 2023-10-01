package runner_communicator

import (
	"go/ast"
	models "tde/models/program"

	"testing"
)

func Test_Send(t *testing.T) {
	var (
		rc = RunnerCommunicator{
			ip_addresses: []string{"127.0.0.1:8081", "127.0.0.1:8081", "127.0.0.1:8081"},
		}
		batch = Batch{
			File: &ast.File{Name: &ast.Ident{Name: "blabla"}},
			Subjects: []*models.Subject{
				{Sid: "MR"}, {Sid: "45"}, {Sid: "LJ"}, {Sid: "1A"}, {Sid: "Nb"}, {Sid: "nP"}, {Sid: "gJ"},
				{Sid: "WI"}, {Sid: "GS"}, {Sid: "KV"}, {Sid: "Q0"}, {Sid: "et"}, {Sid: "XA"}, {Sid: "Nx"},
				{Sid: "6n"}, {Sid: "4p"}, {Sid: "N9"}, {Sid: "zH"}, {Sid: "Jy"}, {Sid: "BP"}, {Sid: "Rq"},
				{Sid: "i8"}, {Sid: "DI"}, {Sid: "Ct"}, {Sid: "tw"}, {Sid: "m7"}, {Sid: "Lw"}, {Sid: "YI"},
				{Sid: "lm"},
			},
		}
	)

	rc.Send(&batch)
}
