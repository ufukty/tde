package discovery

import (
	"os"
	"path/filepath"
	"tde/internal/utilities"

	"github.com/pkg/errors"
)

var (
	ModuleNotFound = errors.New("this directory is not part of a Go module")
)

// Returns the import path for the package inside working directory
func GetImportPathOfPackage() (string, error) {
	out, _, err := utilities.RunCommandForOutput("go", "list")
	if err != nil {
		return "", errors.Wrap(err, "running 'go list' is failed on the working directory")
	}
	return utilities.StripOnlyLineFromCommandOuput(out)
}

func GetPackagePathInModule(modulePath string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "get working directory")
	}
	wdRelToMod, err := filepath.Rel(modulePath, wd)
	if err != nil {
		return "", errors.Wrap(err, "get relative path of working dir to module root")
	}
	return wdRelToMod, nil
}

// Returns the absolute path of the module that working directory is in it
func GetModulePath() (string, error) {
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

// func GetModulePathOfLocation(location string) (string, error) {
// 	cmd := exec.Command("go", "env", "GOMOD")
// 	cmd.Dir = location
// 	modPath, err := cmd.Output()
// 	if err != nil {
// 		return "", errors.Wrap(err, "go env GOMOD")
// 	}
// 	return string(modPath), nil
// }
