package cfg

import (
	"reflect"
	"tde/internal/astw/clone"
	ast_utl "tde/internal/astw/utilities"
	"tde/internal/evaluation"

	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"testing"

	"github.com/kr/pretty"
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

func Test_Develop(t *testing.T) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}

	candidateFuncDecl := clone.FuncDecl(originalFuncDecl)
	newNode, err := Develop(astPkg, astFile, candidateFuncDecl, 1)
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on Develop"))
	}
	fmt.Println("typeOf: ", reflect.TypeOf(newNode))
	if ast_utl.CompareRecursivelyWithAddresses(candidateFuncDecl, originalFuncDecl) == true {
		pretty.Println(newNode)
		pretty.Println(candidateFuncDecl.Body)
		t.Error("Failed to see change on candidate")
	}
	if ok, _ := evaluation.SyntaxCheckSafe(candidateFuncDecl); ok {
		printer.Fprint(os.Stdout, token.NewFileSet(), candidateFuncDecl)
		fmt.Println("")
	}
}

func Benchmark_Develop(b *testing.B) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		b.Error(errors.Wrapf(err, "failed on prep"))
	}

	for i := 0; i < b.N; i++ {
		candidateFuncDecl := clone.FuncDecl(originalFuncDecl)
		newNode, err := Develop(astPkg, astFile, candidateFuncDecl, 1)
		if err != nil {
			b.Error(errors.Wrapf(err, "Failed on Develop"))
		}
		if ast_utl.CompareRecursivelyWithAddresses(candidateFuncDecl, originalFuncDecl) == true {
			if _, ok := newNode.(*ast.BranchStmt); ok { // empty branch statement always leads fail in ast->code convertion
				continue
			}
			b.Errorf("Failed to see change on candidate #%d\n", i)
		}
	}
}

func Test_DevelopProgressively(t *testing.T) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}
	var best = originalFuncDecl
	for i := 0; i < 2000; i++ {
		candidate := clone.FuncDecl(best)
		newNode, err := Develop(astPkg, astFile, candidate, 20)
		if err != nil {
			t.Error(errors.Wrapf(err, "Failed on Develop i = %d, typeOf = %v", i, reflect.TypeOf(newNode)))
		}
		if ok, _ := evaluation.SyntaxCheckSafe(best); ok {
			printer.Fprint(os.Stdout, token.NewFileSet(), best)
			fmt.Println("")
			fmt.Println("^", i, "---")
			best = candidate
		}
	}
}

func Test_DevelopFindUnbreakingChange(t *testing.T) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}

	nonBreakingChangeFound := false
	for i := 0; i < 200; i++ {
		candidate := clone.FuncDecl(originalFuncDecl)
		Develop(astPkg, astFile, candidate, 20)

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
