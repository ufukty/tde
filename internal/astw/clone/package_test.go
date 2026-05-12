package clone

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	"testing"

	"tde/internal/astw/astwutl"
)

func loadTestPackage() (*ast.Package, *ast.File, *ast.FuncDecl, error) {
	_, astPkgs, err := astwutl.LoadDir("testdata")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not load test package: %w", err)
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["testdata/walk.go"]
	funcDecl, err := astwutl.FindFuncDecl(astPkg, "WalkWithNils")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not find test function: %w", err)
	}
	return astPkg, astFile, funcDecl, nil
}

func Benchmark_CopyPackage(b *testing.B) {
	astPkg, _, _, err := loadTestPackage()
	if err != nil {
		b.Error(fmt.Errorf("failed on prep: %w", err))
	}
	for i := 0; i < b.N; i++ {
		Package(astPkg)
	}
}

func Test_StillPrintable(t *testing.T) {
	_, _, funcDecl, err := loadTestPackage()
	if err != nil {
		t.Error(fmt.Errorf("failed on prep: %w", err))
	}

	err = printer.Fprint(io.Discard, token.NewFileSet(), funcDecl)
	if err != nil {
		t.Error(fmt.Errorf("failed on printing original function declaration before even clone: %w", err))
	}

	err = printer.Fprint(io.Discard, token.NewFileSet(), FuncDecl(funcDecl))
	if err != nil {
		t.Error(fmt.Errorf("failed on printing original function declaration after clone: %w", err))
	}
}
