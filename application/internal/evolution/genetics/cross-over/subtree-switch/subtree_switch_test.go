package subtree_switch

import (
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone"

	"fmt"
	"go/ast"
	"testing"

	"github.com/pkg/errors"
)

func loadTestPackage() (*ast.FuncDecl, *ast.FuncDecl, error) {
	_, astPkgs, err := astwutl.LoadDir("testdata")
	if err != nil {
		return nil, nil, errors.Wrapf(err, "could not load test package")
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["testdata/walk.go"]
	funcDeclA, err := astwutl.FindFuncDecl(astPkg, "walkHelper")
	if err != nil {
		return nil, nil, errors.Wrap(err, "walkHelper")
	}
	funcDeclB, err := astwutl.FindFuncDecl(astFile, "walkAstTypeFieldsIfSet")
	if err != nil {
		return nil, nil, errors.Wrap(err, "walkAstTypeFieldsIfSet")
	}
	return funcDeclA, funcDeclB, nil
}

func Test_SubtreeSwitch(t *testing.T) {
	funcDeclA, funcDeclB, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	modifiedFuncDeclA, modifiedFuncDeclB := clone.FuncDecl(funcDeclA), clone.FuncDecl(funcDeclB)
	ok := SubtreeSwitch(modifiedFuncDeclA, modifiedFuncDeclB)
	if !ok {
		t.Error("false return")
	}

	if astwutl.CompareRecursively(modifiedFuncDeclA, funcDeclA) {
		t.Error("Comparison")
	}
	if astwutl.CompareRecursively(modifiedFuncDeclB, funcDeclB) {
		t.Error("Comparison")
	}

	if diff, err := astwutl.Diff(funcDeclA, modifiedFuncDeclA); err != nil {
		t.Error("print")
	} else if diff != "" {
		fmt.Println(diff)
	}

	if diff, err := astwutl.Diff(funcDeclB, modifiedFuncDeclB); err != nil {
		t.Error("print")
	} else if diff != "" {
		fmt.Println(diff)
	}
}
