package remove_line

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/clone"
	ast_utl "tde/internal/astw/utilities"
	"testing"

	"github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
)

func loadTestPackage() (*ast.Package, *ast.File, *ast.FuncDecl, error) {
	_, astPkgs, err := ast_utl.LoadDir("../../../test-package")
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "could not load test package")
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["../../../test-package/walk.go"]
	funcDecl, err := ast_utl.FindFuncDecl(astPkg, "walkHelper")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "could not find test function")
	}
	return astPkg, astFile, funcDecl, nil
}

func Test_RemoveLine(t *testing.T) {
	_, _, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)
	ok := RemoveLine(modifiedFuncDecl.Body)
	if !ok {
		t.Error("return value")
	}

	codeForOriginal, err := ast_utl.String(originalFuncDecl)
	if err != nil {
		t.Error("validation prep")
	}
	codeForModified, err := ast_utl.String(modifiedFuncDecl)
	if err != nil {
		t.Error("validation prep")
	}

	fmt.Println("Differences in code:\n", diff.Diff(codeForOriginal, codeForModified))

	if ast_utl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
		t.Error("validation")
	}
}

func Test_RemoveLineMany(t *testing.T) {
	_, _, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	for i := 0; i < 1000; i++ {
		modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)

		ok := RemoveLine(modifiedFuncDecl.Body)
		if !ok {
			t.Error("return value")
		}

		if ast_utl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
			t.Error("validation", i)
		}
	}
}
