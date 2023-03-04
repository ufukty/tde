package cfg

import (
	"tde/internal/astw/clone"
	ast_utl "tde/internal/astw/utilities"
	"tde/internal/evaluation"

	"fmt"
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

func Test_Develop(t *testing.T) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}

	candidateFuncDecl := clone.FuncDecl(originalFuncDecl)
	err = Develop(astPkg, astFile, candidateFuncDecl, 1)
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed on Develop"))
	}
	if ast_utl.CompareRecursively(candidateFuncDecl, originalFuncDecl) == true {
		t.Error("Failed to see change on candidate")
	}
}

func Benchmark_Develop(b *testing.B) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		b.Error(errors.Wrapf(err, "failed on prep"))
	}

	for i := 0; i < b.N; i++ {
		if i%100 == 0 {
			fmt.Println(i)
		}
		candidateFuncDecl := clone.FuncDecl(originalFuncDecl)
		err := Develop(astPkg, astFile, candidateFuncDecl, 1)
		if err != nil {
			b.Error(errors.Wrapf(err, "Failed on Develop"))
		}
		if ast_utl.CompareRecursively(candidateFuncDecl, originalFuncDecl) == true {
			b.Errorf("Failed to see change on candidate #%d\n", i)
		}
	}
}

func Test_SequentialDevelop(t *testing.T) {
	astPkg, astFile, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}

	for i := 0; i < 2000; i++ {
		err := Develop(astPkg, astFile, originalFuncDecl, 1)
		if err != nil {
			t.Error(errors.Wrapf(err, "Failed on Develop"))
		}
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
		Develop(astPkg, astFile, candidate, 2)

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
