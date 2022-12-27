package embedding

import (
	"testing"

	"github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
)

func Test_StampFunctionPrototype(t *testing.T) {
	originalHeader := "func (config *EmbeddingConfig) EmbedIntoFile(candidates []in_program_models.Candidate) {"
	suffixToAdd := "_Candidate00000000_0000_0000_0000_000000000000"
	want := "func (config *EmbeddingConfig) EmbedIntoFile_Candidate00000000_0000_0000_0000_000000000000(candidates []in_program_models.Candidate) {"
	output, err := StampFunctionPrototype(originalHeader, suffixToAdd)
	if err != nil {
		t.Error(errors.Wrap(err, "Function call has returned an error"))
	}
	if want != output {
		t.Errorf("Expected %s got %s", want, output)
	}
}

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
	got, err := PrepareTemplateForMainFile("tde/examples/word-reverse/word_reverse/word_reverse", "TDE_WordReverse", "00000000-0000-0000-0000-000000000001")
	if err != nil {
		t.Error(errors.Wrap(err, "call to prepare main_tde.go is failed"))
	}
	if got != want {
		t.Error("got != want, content got:\n\n", diff.Diff(got, want))
	}
}

func Test_EmbeddingConfig(t *testing.T) {
	var err error

	ec := NewEmbeddingConfig(
		"../../examples/word-reverse/word_reverse",
		"../../examples/word-reverse/word_reverse/word_reverse.go",
		"../../examples/word-reverse/word_reverse/word_reverse_tde.go",
		"tde/examples/word-reverse/word_reverse",
		"TDE_WordReverse",
	)
	err = ec.Embed()
	if err != nil {
		t.Error(errors.Wrap(err, "Got error"))
	}
}
