package discovery

import (
	"fmt"
	"path/filepath"

	"tde/internal/evolution/evaluation/list"
	"tde/internal/utilities/osw"
)

var ModuleNotFound = fmt.Errorf("this directory is not part of a Go module")

// Returns the absolute path of the module that working directory is in it
func ModuleRoot() (string, error) {
	path, _, err := osw.RunCommandForOutput("go", "env", "GOMOD")
	if err != nil {
		return "", fmt.Errorf("failed to run 'go env GOMOD': %w", err)
	}
	path, err = osw.StripOnlyLineFromCommandOuput(path)
	if err != nil {
		return "", fmt.Errorf("could not strip GOMOD path from the output of 'go env GOMOD': %w", err)
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
