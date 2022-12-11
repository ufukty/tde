package embedding

import (
	"testing"

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

func Test_EmbeddingConfig(t *testing.T) {
	ec := NewEmbeddingConfig(
		"../../examples/word-reverse/word_reverse",
		"../../examples/word-reverse/word_reverse/word_reverse.go",
		"../../examples/word-reverse/word_reverse/word_reverse_tde.go",
	)
	err := ec.Embed()
	if err != nil {
		t.Error(errors.Wrap(err, "Got error"))
	}
}
