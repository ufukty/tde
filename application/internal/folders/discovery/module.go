package discovery

import (
	"tde/internal/folders/list"
	"tde/internal/utilities"

	"fmt"
	"path/filepath"

	"github.com/pkg/errors"
)

var (
	ModuleNotFound = errors.New("this directory is not part of a Go module")
)

// Returns the absolute path of the module that working directory is in it
func ModuleRoot() (string, error) {
	path, _, err := utilities.RunCommandForOutput("go", "env", "GOMOD")
	if err != nil {
		return "", errors.Wrap(err, "failed to run 'go env GOMOD'")
	}
	path, err = utilities.StripOnlyLineFromCommandOuput(path)
	if err != nil {
		return "", errors.Wrap(err, "could not strip GOMOD path from the output of 'go env GOMOD'")
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
