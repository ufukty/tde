package import_path

import (
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone"

	"fmt"
	"go/ast"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
)

func loadTestPackage() (*ast.Package, *ast.File, *ast.FuncDecl, error) {
	_, astPkgs, err := astwutl.LoadDir("testdata")
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "could not load test package")
	}
	astPkg := astPkgs["test_package"]
	astFile := astPkg.Files["testdata/walk.go"]
	funcDecl, err := astwutl.FindFuncDecl(astPkg, "walkHelper")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "could not find test function")
	}
	return astPkg, astFile, funcDecl, nil
}

func Test_ImportPackage(t *testing.T) {
	_, originalFile, _, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	packageNameToImport := "4e1c8b43-300e-549e-a7d8-2ddb6b803915"

	modifiedFile := clone.File(originalFile)
	ImportPackage(modifiedFile, packageNameToImport)

	codeForOriginal, err := astwutl.String(originalFile)
	if err != nil {
		t.Error("validation prep")
	}
	codeForModified, err := astwutl.String(modifiedFile)
	if err != nil {
		t.Error("validation prep")
	}

	fmt.Println("Differences in code:\n", diff.Diff(codeForOriginal, codeForModified))

	if astwutl.CompareRecursively(originalFile, modifiedFile) {
		t.Error("validation 1")
	}

	if !strings.Contains(codeForModified, packageNameToImport) {
		t.Error("validation 2")
	}
}

func Test_ImportPackageProgressively(t *testing.T) {
	_, originalFile, _, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	for i := 0; i < 100; i++ {
		packageNameToImport := uuid.New().String()

		modifiedFile := clone.File(originalFile)
		ImportPackage(modifiedFile, packageNameToImport)

		codeForModified, err := astwutl.String(modifiedFile)
		if err != nil {
			t.Error("validation prep")
		}

		if astwutl.CompareRecursively(originalFile, modifiedFile) {
			t.Error("validation 1")
		}

		if !strings.Contains(codeForModified, packageNameToImport) {
			t.Error("validation 2")
		}
	}
}
