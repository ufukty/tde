package batch_post

import (
	"bytes"
	"go/ast"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
)

func Test_RunnerService_NewTest_Request_Deserialize(t *testing.T) {
	body := bytes.NewReader([]byte(`{
		"Subjects": [
			{
				"Sid": "1",
				"FuncDeclSerialized": "func() { fmt.Println(\"Hello World\") }"
			}
		],
		"ArchiveID": "2",
		"FileTemplateSerialized": "3"
	}`))
	r := httptest.NewRequest("POST", "https://localhost", body)

	req := Request{}
	err := req.ParseRequest(r)
	if err != nil {
		t.Error(errors.Wrapf(err, "returned error"))
	}

	if len(req.Subjects) != 1 ||
		req.Subjects[0].Sid != "1" ||
		// req.Subjects[0].FuncDecl != "func() { fmt.Println(\"Hello World\") }" ||
		req.ArchiveID != "2" {
		t.Error(errors.Wrapf(err, "validation"))
	}
}

func Test_RunnerService_NewTest_Request_Serialize(t *testing.T) {
	content1 := Request{
		Subjects: []Subject{{
			Sid: "1",
			FuncDecl: &ast.FuncDecl{
				Name: &ast.Ident{Name: "blabla"},
			},
		}},
		ArchiveID: "2",
		FileTemplate: &ast.File{
			Name: &ast.Ident{Name: "blibli"},
		},
	}

	req, err := content1.NewRequest("POST", "https://localhost")
	if err != nil {
		t.Error(errors.Wrapf(err, "process"))
	}

	content2 := Request{}
	content2.ParseRequest(req)

	areSame := content1.ArchiveID == content2.ArchiveID &&
		content1.FileTemplate.Name.Name == content2.FileTemplate.Name.Name
	if !areSame {
		t.Error("validation")
	}
}
