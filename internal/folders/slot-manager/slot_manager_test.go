package slot_manager

import (
	"tde/internal/folders/preparation"
	"tde/internal/folders/types"
	models "tde/models/program"

	"fmt"
	"path/filepath"
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
		types.InModulePath("examples/word-reverse"),
		"tde/examples/word-reverse",
		"TDE_WordReverse",
	)
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	candidates := []*models.Candidate{}
	for i := 0; i < 10; i++ {
		cand := &models.Candidate{}
		cand.File = []byte(`hello world`)
		candidates = append(candidates, cand)
	}

	var (
		modulePath = types.AbsolutePath(clone)
		config     = &types.TestDetails{
			PackagePath:   types.InModulePath("examples/word-reverse"),
			PackageImport: "tde/examples/word-reverse",
			ImplFuncFile:  types.InModulePath("examples/word-reverse/word_reverse.go"),
			TestFuncFile:  types.InModulePath("examples/word-reverse/word_reverse_tde.go"),
			TestFuncName:  "TDE_WordReverse",
		}
	)
	session := NewSession(modulePath, config)
	session.PlaceCandidatesIntoSlots(candidates)

	fmt.Println(session.tmp)
}
