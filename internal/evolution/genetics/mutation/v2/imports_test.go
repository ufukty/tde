package mutation

import (
	"fmt"
	"strings"
	"testing"

	"tde/internal/ast/astwutl"
	"tde/internal/ast/clone"
	"tde/internal/ast/compare"

	"github.com/google/uuid"
	"github.com/kylelemons/godebug/diff"
)

func Test_ImportPackage(t *testing.T) {
	_, originalFile, _, err := loadTestPackage()
	if err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
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

	if compare.Recursively(originalFile, modifiedFile) {
		t.Error("validation 1")
	}

	if !strings.Contains(codeForModified, packageNameToImport) {
		t.Error("validation 2")
	}
}

func Test_ImportPackageProgressively(t *testing.T) {
	_, originalFile, _, err := loadTestPackage()
	if err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}

	for i := 0; i < 100; i++ {
		packageNameToImport := uuid.New().String()

		modifiedFile := clone.File(originalFile)
		ImportPackage(modifiedFile, packageNameToImport)

		codeForModified, err := astwutl.String(modifiedFile)
		if err != nil {
			t.Error("validation prep")
		}

		if compare.Recursively(originalFile, modifiedFile) {
			t.Error("validation 1")
		}

		if !strings.Contains(codeForModified, packageNameToImport) {
			t.Error("validation 2")
		}
	}
}
