package archive

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"tde/internal/evaluation/discovery"
)

func Test_ArchiveDirectory(t *testing.T) {
	excludeDir := []string{".git", "build", "docs", ".vscode", "artifacts"}
	includeExt := []string{"go"}
	path, err := Directory("../../../../", true, excludeDir, excludeDir, includeExt, false)
	if err != nil {
		t.Error(fmt.Errorf("failed to archive directory: %w", err))
	}
	fmt.Println("Output zip:", path)
}

func Test_findModuleRoot(t *testing.T) {
	root, err := discovery.ModuleRoot()
	if err != nil {
		t.Error(fmt.Errorf("failed: %w", err))
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Error(fmt.Errorf("getting the working directory: %w", err))
	}
	expected := filepath.Clean(wd + "/../../../..")

	if root != expected {
		t.Error("failed: got wrong output:", root, "expected:", expected)
	}
}
