package slot_manager

import (
	"path/filepath"
	"tde/internal/folders/preparation"
	"tde/internal/folders/types"
	"tde/models/in_program_models"
	"testing"

	"github.com/pkg/errors"
)

func Test_SlotManager(t *testing.T) {
	absPath, err := filepath.Abs("../../")
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	clone, err := preparation.Prepare(
		types.AbsolutePath(absPath),
		types.InModulePath("examples/word_reverse"),
		"tde/examples/word_reverse",
		"TDE_WordReverse",
	)
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	candidates := []*in_program_models.Candidate{}
	for i := 0; i < 1000; i++ {
		candidates = append(candidates, in_program_models.NewCandidate())
	}

	config := &Config{
		OriginalModule:     types.AbsolutePath(clone),
		Package:            types.InModulePath("examples/word_reverse"),
		PackageImport:      "tde/examples/word_reverse",
		ImplementationFile: types.InModulePath("examples/word_reverse/word_reverse.go"),
		TestFile:           types.InModulePath("examples/word_reverse/word_reverse_tde.go"),
		TestName:           "TDE_WordReverse",
	}
	session := NewSession(config)
	session.PlaceCandidatesIntoSlots(candidates)
}
