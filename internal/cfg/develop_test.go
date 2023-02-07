package cfg

import (
	"tde/internal/astw/clone"
	ast_utl "tde/internal/astw/utilities"

	"testing"

	"github.com/pkg/errors"
)

func Test_Develop(t *testing.T) {
	_, astPkgs, err := ast_utl.LoadDir("../test_package")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed test prep"))
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["../test_package/walk.go"]
	originalFuncDecl, _ := ast_utl.FindFuncDecl(astPkg, "WalkWithNils")

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
	_, astPkgs, err := ast_utl.LoadDir("../test_package")
	if err != nil {
		b.Error(errors.Wrapf(err, "Failed test prep"))
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["../test_package/walk.go"]
	originalFuncDecl, _ := ast_utl.FindFuncDecl(astPkg, "WalkWithNils")

	for i := 0; i < b.N; i++ {
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
	_, astPkgs, err := ast_utl.LoadDir("../test_package")
	if err != nil {
		t.Error(errors.Wrapf(err, "Failed test prep"))
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["../test_package/walk.go"]
	funcDecl, _ := ast_utl.FindFuncDecl(astPkg, "WalkWithNils")

	for i := 0; i < 2000; i++ {
		err := Develop(astPkg, astFile, funcDecl, 1)
		if err != nil {
			t.Error(errors.Wrapf(err, "Failed on Develop"))
		}
	}
}
