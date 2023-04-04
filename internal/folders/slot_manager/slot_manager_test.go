package slot_manager

import (
	"fmt"
	"path/filepath"
	"tde/internal/folders/preparation"
	"tde/internal/folders/types"
	"tde/models/in_program_models"
	"testing"

	"github.com/pkg/errors"
)

func Test_SlotManager(t *testing.T) {
	absPath, err := filepath.Abs("../../../")
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
	for i := 0; i < 10; i++ {
		cand := &in_program_models.Candidate{}
		cand.File = []byte(`hello world`)
		candidates = append(candidates, cand)
	}

	var (
		modulePath = types.AbsolutePath(clone)
		config     = &types.TestDetails{
			PackagePath:   types.InModulePath("examples/word_reverse"),
			PackageImport: "tde/examples/word_reverse",
			ImplFuncFile:  types.InModulePath("examples/word_reverse/word_reverse.go"),
			TestFuncFile:  types.InModulePath("examples/word_reverse/word_reverse_tde.go"),
			TestFuncName:  "TDE_WordReverse",
		}
	)
	session := NewSession(modulePath, config)
	session.PlaceCandidatesIntoSlots(candidates)

	fmt.Println(session.tmp)
}
