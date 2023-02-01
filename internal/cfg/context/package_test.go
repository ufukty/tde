package context

import (
	trav "tde/internal/astw/traverse"
	utl "tde/internal/astw/utilities"

	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func Test_GetContextForSpot(t *testing.T) {
	_, astPkgs, err := utl.LoadDir("../../test_package")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on loading package"))
	}
	astPkg := astPkgs["test_package"]
	funcDecl, err := utl.FindFuncDecl(astPkg, "WalkWithNils")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on preparation"))
	}

	tFuncDecl := trav.GetTraversableNodeForASTNode(funcDecl)
	funcBody := trav.GetTraversableNodeForASTNode(funcDecl.Body).GetTraversableSubnodes()
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
