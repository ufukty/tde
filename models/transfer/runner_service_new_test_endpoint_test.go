package models

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
)

func Test_RunnerService_NewTest_Request_Deserialize(t *testing.T) {
	body := bytes.NewReader([]byte(`{
		"Candidates": [
			{
				"CandidateID": "1",
				"FuncDeclSerialized": "func() { fmt.Println(\"Hello World\") }"
			}
		],
		"ArchiveID": "2",
		"FileTemplateSerialized": "3"
	}`))
	r := httptest.NewRequest("POST", "https://localhost", body)

	req := RunnerService_NewTest_Request{}
	err := req.Deserialize(r)
	if err != nil {
		t.Error(errors.Wrapf(err, "returned error"))
	}

	if len(req.Candidates) != 1 ||
		req.Candidates[0].CandidateID != "1" ||
		// req.Candidates[0].FuncDecl != "func() { fmt.Println(\"Hello World\") }" ||
		req.ArchiveID != "2" {
		t.Error(errors.Wrapf(err, "validation"))
	}
}

func Test_RunnerService_NewTest_Request_Serialize(t *testing.T) {

}
