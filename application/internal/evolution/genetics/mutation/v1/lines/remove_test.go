package lines

import (
	"tde/internal/astw/astwutl"
	"tde/internal/astw/clone"
	"tde/internal/evolution/genetics/mutation/v1/models"

	"fmt"
	"testing"

	"github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
)

func Test_RemoveLine(t *testing.T) {
	_, _, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)
	ctx := &models.MutationParameters{
		FuncDecl: modifiedFuncDecl,
	}
	if err := RemoveLine(ctx); err != nil {
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

func Test_RemoveLineMany(t *testing.T) {
	_, _, originalFuncDecl, err := loadTestPackage()
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	for i := 0; i < 1000; i++ {
		modifiedFuncDecl := clone.FuncDecl(originalFuncDecl)
		ctx := &models.MutationParameters{
			FuncDecl: modifiedFuncDecl,
		}
		if err := RemoveLine(ctx); err != nil {
			t.Fatal(fmt.Errorf("act: %w", err))
		}

		if astwutl.CompareRecursively(originalFuncDecl, modifiedFuncDecl) {
			t.Error("validation", i)
		}
	}
}
