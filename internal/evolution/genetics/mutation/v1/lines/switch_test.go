package lines

import (
	"fmt"
	"go/ast"
	"testing"

	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone"
	"tde/internal/evolution/genetics/mutation/v1/models"

	"github.com/kylelemons/godebug/diff"
)

func loadTestPackage() (*ast.Package, *ast.File, *ast.FuncDecl, error) {
	_, astPkgs, err := astwutl.LoadDir("testdata")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not load test package: %w", err)
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["testdata/walk.go"]
	funcDecl, err := astwutl.FindFuncDecl(astPkg, "walkHelper")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not find test function: %w", err)
	}
	return astPkg, astFile, funcDecl, nil
}

func Test_SiblingSwap(t *testing.T) {
	_, _, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}

	modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)
	ctx := &models.MutationParameters{
		FuncDecl: modifiedFuncDecl,
	}
	if err := SwapLines(ctx); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
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
		t.Error(fmt.Errorf("prep: %w", err))
	}

	for i := 0; i < 1000; i++ {
		modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)
		ctx := &models.MutationParameters{
			FuncDecl: modifiedFuncDecl,
		}

		if err := SwapLines(ctx); err != nil {
			t.Fatal(fmt.Errorf("act: %w", err))
		}

		if astwutl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
			t.Error("validation", i)
		}
	}
}
