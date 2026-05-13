package discovery

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"tde/internal/evolution/evaluation/list"
)

var ModuleNotFound = fmt.Errorf("this directory is not part of a Go module")

func gomod() (string, error) {
	cmd := exec.Command("go", "env", "GOMOD")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("running go command: shell: %w", err)
	}
	output := stdout.String()
	if lines := strings.Count(output, "\n"); lines != 2 {
		return "", fmt.Errorf("unexpected number of lines: %d", lines)
	}
	line := strings.SplitN(output, "\n", 2)[0]
	return strings.TrimSpace(line), nil
}

// Returns the absolute path of the module that working directory is in it
func ModuleRoot() (string, error) {
	path, err := gomod()
	if err != nil {
		return "", fmt.Errorf("finding GOMOD: %w", err)
	}
	if path == "/dev/null" {
		return "", ModuleNotFound
	}
	return filepath.Dir(path), nil
}

func WhereAmI() (module string, pkg *list.Package, err error) {
	module, err = ModuleRoot()
	if err != nil {
		return "", nil, fmt.Errorf("getting module root: %w", err)
	}
	var pkgs list.Packages
	pkgs, err = list.ListAllPackages(".")
	if err != nil {
		return "", nil, fmt.Errorf("listing packages in current dir: %w", err)
	}
	pkg = pkgs.First()
	return
}
