package embedding

import (
	"testing"

	"github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
)

func Test_PrepareTemplateForMainFile(t *testing.T) {
	want := `//go:build tde
// +build tde

package main

import (
	targetPackage "tde/examples/word-reverse/word_reverse/word_reverse"
	"tde/models/in_program_models"
	"tde/pkg/tde"

	"flag"
)

var candidateUUID string

func init() {
	flag.StringVar(
		&candidateUUID,
		"candidate-uuid",
		"",
		"Tested candidate's identifier. The program will use this only for output and it is not essential for program to run.",
	)
}

func main() {
	flag.Parse()
	var (
		testFunction  = targetPackage.TDE_WordReverse
		candidateUUID = in_program_models.CandidateID(candidateUUID)
		e			  = tde.NewE(candidateUUID)
	)
	testFunction(e)
	e.Export()
}`
	got, err := prepareTemplateForMainFile("tde/examples/word-reverse/word_reverse/word_reverse", "TDE_WordReverse", "00000000-0000-0000-0000-000000000001")
	if err != nil {
		t.Error(errors.Wrap(err, "call to prepare main_tde.go is failed"))
	}
	if got != want {
		t.Error("got != want, content got:\n\n", diff.Diff(got, want))
	}
}

func Test_injectMainPackage(t *testing.T) {

}
