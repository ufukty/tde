package discovery

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

var (
	ModuleNotFound = errors.New("this directory is not part of a Go module")
)

type Request struct {
	functionName string
	dir          string
	testfile     string
	// evolution
	lineStart int
	lineEnd   int
}

func NewRequest(evolvingFilePath string, targetBlockStartingLine int, targetBlockEndingLine int) *Request {
	return &Request{
		// file:      evolvingFilePath,
		lineStart: targetBlockStartingLine,
		lineEnd:   targetBlockEndingLine,
	}
}

func (r *Request) Discover() {
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Panicln(errors.Wrap(err, "Couldn't list the entries of this directory."))
	}
	fmt.Println(entries)
}

// Returns the import path for the package inside working directory
func FindImportPathOfThePackage() (string, error) {
	out, err := RunCommandForOutput("go", "list")
	if err != nil {
		return "", errors.Wrap(err, "running 'go list' is failed on the working directory")
	}
	return StripOnlyLineFromCommandOuput(out)
}

// Returns the absolute path of the module that working directory is in it
func FindModulePath() (string, error) {
	path, err := RunCommandForOutput("go", "env", "GOMOD")
	if err != nil {
		return "", errors.Wrap(err, "failed to run 'go env GOMOD'")
	}
	path, err = StripOnlyLineFromCommandOuput(path)
	if err != nil {
		return "", errors.Wrap(err, "could not strip GOMOD path from the output of 'go env GOMOD'")
	}
	if path == "/dev/null" {
		return "", ModuleNotFound
	}
	return filepath.Dir(path), nil
}
