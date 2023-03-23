package tester_package

import (
	"os"
	"tde/internal/embedding/models"
	"testing"

	"github.com/pkg/errors"
)

func Test_Inject(t *testing.T) {
	ec := models.NewEmbeddingConfig(
		"../../../examples/word_reverse",
		"../../../examples/word_reverse/word_reverse.go",
		"../../../examples/word_reverse/word_reverse_tde.go",
		"tde/examples/word_reverse",
		"TDE_WordReverse",
	)

	if err := Inject(ec); err != nil {
		t.Error(errors.Wrap(err, "returned error"))
	}

	if _, err := os.OpenFile("../../../examples/word_reverse/tde/main_tde.go", os.O_RDONLY, os.ModeAppend); err != nil {
		t.Error(errors.Wrap(err, "validation"))
	}
}
