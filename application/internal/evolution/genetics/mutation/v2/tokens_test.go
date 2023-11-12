package mutation

import (
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone"

	"fmt"
	"go/ast"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
)

func loadTestPackage() (*ast.Package, *ast.File, *ast.FuncDecl, error) {
	_, astPkgs, err := astwutl.LoadDir("testdata")
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "could not load test package")
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["testdata/walk.go"]
	funcDecl, err := astwutl.FindFuncDecl(astPkg, "walkHelper")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "could not find test function")
	}
	return astPkg, astFile, funcDecl, nil
}

func Test_Operator(t *testing.T) {
	_, _, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)
	choosenNode, newToken, ok := Perform(modifiedFuncDecl.Body)
	if !ok {
		t.Error("return value")
	}

	codeForOriginal, err := astwutl.String(originalFuncDecl)
	if err != nil {
		t.Error("validation prep")
	}
	codeForModified, err := astwutl.String(modifiedFuncDecl)
	if err != nil {
		t.Error("validation prep")
	}

	fmt.Println("Differences in code:\n", diff.Diff(codeForOriginal, codeForModified))

	if astwutl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
		t.Error("validation", choosenNode, newToken)
	}

}

func Test_Bulk(t *testing.T) {
	_, _, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	for i := 0; i < 1000; i++ {
		modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)
		choosenNode, newToken, ok := Perform(modifiedFuncDecl.Body)
		if !ok {
			t.Error("return value")
		}
		if astwutl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
			t.Errorf("validation i='%d' typeOf->choosenNode='%v' address->choosenNode='%p' newToken='%v'", i, reflect.TypeOf(choosenNode), choosenNode, newToken)
		}
	}

}
