package upload

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

func Test_ArchiveDirectory(t *testing.T) {
	path, err := ArchiveDirectory("../../../", true, []string{".git", "build", "docs", ".vscode"})
	if err != nil {
		t.Error(errors.Wrapf(err, "failed to archive directory"))
	}
	fmt.Println("Output zip:", path)
}

func Test_findModuleRoot(t *testing.T) {
	root, err := findModuleRoot()
	if err != nil {
		t.Error(errors.Wrapf(err, "failed"))
	}

	cmdWorkingDir := exec.Command("pwd", "-P")
	bytes, err := cmdWorkingDir.Output()
	if err != nil {
		t.Error(errors.Wrap(err, "failed on running \"pwd -P\" command"))
	}
	expected := filepath.Clean(string(bytes) + "/../../..")

	if root != expected {
		t.Error("failed: got wrong output:", root, "expected:", expected)
	}
}
