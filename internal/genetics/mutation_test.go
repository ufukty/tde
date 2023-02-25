package genetics

import (
	"fmt"
	"tde/internal/astw/clone"
	ast_utl "tde/internal/astw/utilities"
	"tde/internal/cfg"
	"tde/internal/evaluation"

	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func loadTestPackage() (*ast.Package, *ast.File, *ast.FuncDecl, error) {
	_, astPkgs, err := ast_utl.LoadDir("../test_package")
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "could not load test package")
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["../test_package/walk.go"]
	funcDecl, err := ast_utl.FindFuncDecl(astPkg, "WalkWithNils")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "could not find test function")
	}
	return astPkg, astFile, funcDecl, nil
}

func Test_DevelopWithoutSyntaxError(t *testing.T) {
	astPkg, astFile, funcDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}

	bufferOrg := bytes.NewBuffer([]byte{})
	ast.Fprint(bufferOrg, token.NewFileSet(), funcDecl, nil)

	err = cfg.Develop(astPkg, astFile, funcDecl, 2)
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on call to Develop"))
	}

	if !evaluation.SyntaxCheckUnsafe(funcDecl) {
		t.Error("fails on SyntaxCheck")
	} else {
		printer.Fprint(os.Stdout, token.NewFileSet(), funcDecl)
	}
}

func Test_DevelopFindUnbreakingChange(t *testing.T) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}

	nonBreakingChangeFound := false
	for i := 0; i < 20; i++ {
		candidate := clone.FuncDecl(originalFuncDecl)
		cfg.Develop(astPkg, astFile, candidate, 2)

		if ok, _ := evaluation.SyntaxCheckSafe(candidate); ok {
			printer.Fprint(os.Stdout, token.NewFileSet(), candidate)
			fmt.Println("\n---")
			nonBreakingChangeFound = true
		}
	}

	if !nonBreakingChangeFound {
		t.Error("No non-breaking candidates found.")
	}
}
