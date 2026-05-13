package subtreeswitch

import (
	"fmt"
	"go/ast"
	"testing"

	"tde/internal/ast/astwutl"
	"tde/internal/ast/clone"
	"tde/internal/ast/compare"
	"tde/internal/ast/find"
	"tde/internal/ast/parse"
)

func loadTestPackage() (*ast.FuncDecl, *ast.FuncDecl, error) {
	_, astPkgs, err := parse.Dir("testdata")
	if err != nil {
		return nil, nil, fmt.Errorf("could not load test package: %w", err)
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["testdata/walk.go"]
	funcDeclA, err := find.Function(astPkg, "walkHelper")
	if err != nil {
		return nil, nil, fmt.Errorf("walkHelper: %w", err)
	}
	funcDeclB, err := find.Function(astFile, "walkAstTypeFieldsIfSet")
	if err != nil {
		return nil, nil, fmt.Errorf("walkAstTypeFieldsIfSet: %w", err)
	}
	return funcDeclA, funcDeclB, nil
}

func Test_SubtreeSwitch(t *testing.T) {
	funcDeclA, funcDeclB, err := loadTestPackage()
	if err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}

	modifiedFuncDeclA, modifiedFuncDeclB := clone.FuncDecl(funcDeclA), clone.FuncDecl(funcDeclB)
	ok := SubtreeSwitch(modifiedFuncDeclA, modifiedFuncDeclB)
	if !ok {
		t.Error("false return")
	}

	if compare.Recursively(modifiedFuncDeclA, funcDeclA) {
		t.Error("Comparison")
	}
	if compare.Recursively(modifiedFuncDeclB, funcDeclB) {
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
