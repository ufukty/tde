package discovery

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
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

func findImportPath() (string, error) {
	cmd := exec.Command("go", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", errors.Wrap(err, "running 'go list' is failed on the working directory")
	}
	lines := strings.Split(out.String(), "\n")
	if len(lines) < 2 {
		return "", errors.New("'go list' don't print anything")
	} else if len(lines) > 2 {
		return "", errors.New("expected 1 lines of output from 'go list', got many")
	}
	return lines[0], nil
}

func InspectPackage(packagePath string) {
	return
}
