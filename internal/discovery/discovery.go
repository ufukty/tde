package discovery

import (
	"path/filepath"
	"tde/internal/utilities"
	"tde/models/in_program_models"

	"github.com/pkg/errors"
)

var (
	ModuleNotFound = errors.New("this directory is not part of a Go module")
)

// Returns the import path for the package inside working directory
func FindImportPathOfThePackage() (string, error) {
	out, _, err := utilities.RunCommandForOutput("go", "list")
	if err != nil {
		return "", errors.Wrap(err, "running 'go list' is failed on the working directory")
	}
	return utilities.StripOnlyLineFromCommandOuput(out)
}

// Returns the absolute path of the module that working directory is in it
func FindModulePath() (string, error) {
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

func Discover() (*in_program_models.DiscoveryResponse, error) {
	modulePath, err := FindModulePath()
	if err != nil {
		return nil, errors.Wrap(err, "failed on finding module absoute path")
	}
	packageImportPath, err := FindImportPathOfThePackage()
	if err != nil {
		return nil, errors.Wrap(err, "failed on target package's import path")
	}
	return &in_program_models.DiscoveryResponse{
		ModuleAbsolutePath:      modulePath,
		TargetPackageImportPath: packageImportPath,
	}, nil
}
