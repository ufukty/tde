package utilities

import (
	"fmt"
	"go/ast"
	"testing"
)

//go:generate serdeser module-ast-pkg.go
type SeparateParamsTest struct {
	ArchiveId string       `url:"aid"`
	Folder    string       `url:"folder"`
	Package   *ast.Package `json:"package"`
}

func Test_separateParams(t *testing.T) {
	var (
		aid    = "54e36ac4-a965-56cc-8a0c-9446cf8daeba"
		folder = "1665/Kutbil/Plaza"
	)

	req := SeparateParamsTest{
		ArchiveId: aid,
		Folder:    folder,
	}

	url, body := separateParams(req)

	fmt.Println("url:")
	for key, value := range url {
		fmt.Printf("%s: %s\n", key, value)
	}

	fmt.Println("\nbody:")
	for key, value := range body {
		fmt.Printf("%s: %v\n", key, value)
	}

	if got, ok := url["aid"]; !ok || aid != got {
		t.Errorf("got %s want %s", got, aid)
	}
}
