package context

import (
	"fmt"
	"tde/internal/astw"
	"testing"

	"github.com/pkg/errors"
)

func Test_GetContextForSpot(t *testing.T) {
	_, astPkgs, err := astw.LoadDir("../../test_package")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on loading package"))
	}
	astPkg := astPkgs["test_package"]
	funcDecl, err := astw.FindFuncDecl(astPkg, "WalkWithNils")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on preparation"))
	}

	tFuncDecl := astw.GetTraversableNodeForASTNode(funcDecl)
	funcBody := astw.GetTraversableNodeForASTNode(funcDecl.Body).GetTraversableSubnodes()
	choosenSpot := funcBody[len(funcBody)-1]

	ctx, err := GetContextForSpot(
		astPkg,
		tFuncDecl,
		choosenSpot,
	)
	if err != nil {
		t.Error(errors.Wrapf(err, ""))
	}

	fmt.Println("Context:\n", ctx)
}
