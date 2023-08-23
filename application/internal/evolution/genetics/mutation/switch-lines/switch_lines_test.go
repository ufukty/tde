package switch_lines

import (
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone"

	"fmt"
	"go/ast"
	"testing"

	"github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
)

func loadTestPackage() (*ast.Package, *ast.File, *ast.FuncDecl, error) {
	_, astPkgs, err := astwutl.LoadDir("../../../test-package")
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "could not load test package")
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["../../../test-package/walk.go"]
	funcDecl, err := astwutl.FindFuncDecl(astPkg, "walkHelper")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "could not find test function")
	}
	return astPkg, astFile, funcDecl, nil
}

func Test_SiblingSwap(t *testing.T) {
	_, _, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)
	ok := SiblingSwap(modifiedFuncDecl.Body)
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
		t.Error("validation")
	}
}

func Test_SiblingSwapMany(t *testing.T) {
	_, _, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	for i := 0; i < 1000; i++ {
		modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)

		ok := SiblingSwap(modifiedFuncDecl.Body)
		if !ok {
			t.Error("return value")
		}

		if astwutl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
			t.Error("validation", i)
		}
	}
}
