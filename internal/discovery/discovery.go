package discovery

import (
	"fmt"
	"log"
	"os"

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
