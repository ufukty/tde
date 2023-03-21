package subtree_switch

import (
	"fmt"
	"tde/internal/astw/clone"
	ast_utl "tde/internal/astw/utilities"

	"go/ast"
	"testing"

	"github.com/pkg/errors"
)

func loadTestPackage() (*ast.FuncDecl, *ast.FuncDecl, error) {
	_, astPkgs, err := ast_utl.LoadDir("../../../test_package")
	if err != nil {
		return nil, nil, errors.Wrapf(err, "could not load test package")
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["../../../test_package/walk.go"]
	funcDeclA, err := ast_utl.FindFuncDecl(astPkg, "walkHelper")
	if err != nil {
		return nil, nil, errors.Wrap(err, "walkHelper")
	}
	funcDeclB, err := ast_utl.FindFuncDecl(astFile, "walkAstTypeFieldsIfSet")
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

	if ast_utl.CompareRecursively(modifiedFuncDeclA, funcDeclA) {
		t.Error("Comparison")
	}
	if ast_utl.CompareRecursively(modifiedFuncDeclB, funcDeclB) {
		t.Error("Comparison")
	}

	if diff, err := ast_utl.Diff(funcDeclA, modifiedFuncDeclA); err != nil {
		t.Error("print")
	} else if diff != "" {
		fmt.Println(diff)
	}

	if diff, err := ast_utl.Diff(funcDeclB, modifiedFuncDeclB); err != nil {
		t.Error("print")
	} else if diff != "" {
		fmt.Println(diff)
	}
}
