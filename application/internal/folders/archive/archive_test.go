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
	excludeDir := []string{".git", "build", "docs", ".vscode", "artifacts"}
	includeExt := []string{"go"}
	path, err := Directory("../../../../", true, excludeDir, excludeDir, includeExt, false)
	if err != nil {
		t.Error(errors.Wrapf(err, "failed to archive directory"))
	}
	fmt.Println("Output zip:", path)
}

func Test_findModuleRoot(t *testing.T) {
	root, err := discovery.ModuleRoot()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed"))
	}

	workingDir, err := utilities.WorkingDir()
	if err != nil {
		t.Error(errors.Wrap(err, "failed on running \"pwd -P\" command"))
	}
	expected := filepath.Clean(workingDir + "/../../..")

	if root != expected {
		t.Error("failed: got wrong output:", root, "expected:", expected)
	}
}
