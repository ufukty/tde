package archive

import (
	"tde/internal/folders/discovery"
	"tde/internal/utilities"

	"fmt"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

func Test_ArchiveDirectory(t *testing.T) {
	path, err := Directory("../../../../", true, []string{".git", "build", "docs", ".vscode"})
	if err != nil {
		t.Error(errors.Wrapf(err, "failed to archive directory"))
	}
	fmt.Println("Output zip:", path)
}

func Test_findModuleRoot(t *testing.T) {
	root, err := discovery.FindModulePath()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed"))
	}

	workingDir, err := utilities.WorkingDir()
	if err != nil {
		t.Error(errors.Wrap(err, "failed on running \"pwd -P\" command"))
	}
	expected := filepath.Clean(workingDir + "/../../../..")

	if root != expected {
		t.Error("failed: got wrong output:", root, "expected:", expected)
	}
}
