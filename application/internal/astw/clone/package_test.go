package clone

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	"tde/internal/astw/astwutl"

	"testing"

	"github.com/pkg/errors"
)

func loadTestPackage() (*ast.Package, *ast.File, *ast.FuncDecl, error) {
	_, astPkgs, err := astwutl.LoadDir("../../test-package")
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "could not load test package")
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["../../test-package/walk.go"]
	funcDecl, err := astwutl.FindFuncDecl(astPkg, "WalkWithNils")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "could not find test function")
	}
	return astPkg, astFile, funcDecl, nil
}

func Benchmark_CopyPackage(b *testing.B) {
	astPkg, _, _, err := loadTestPackage()
	if err != nil {
		b.Error(errors.Wrapf(err, "failed on prep"))
	}
	for i := 0; i < b.N; i++ {
		Package(astPkg)
	}
}

func Test_StillPrintable(t *testing.T) {
	_, _, funcDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed on prep"))
	}

	err = printer.Fprint(io.Discard, token.NewFileSet(), funcDecl)
	if err != nil {
		t.Error(errors.Wrap(err, "failed on printing original function declaration before even clone"))
	}

	err = printer.Fprint(io.Discard, token.NewFileSet(), FuncDecl(funcDecl))
	if err != nil {
		t.Error(errors.Wrap(err, "failed on printing original function declaration after clone"))
	}

}
