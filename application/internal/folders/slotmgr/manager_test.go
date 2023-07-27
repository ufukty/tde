package slotmgr

import (
	"tde/internal/folders/inject"
	"tde/internal/folders/list"
	"tde/internal/folders/types"
	models "tde/models/program"

	"fmt"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func Test_SlotManager(t *testing.T) {
	absPath, err := filepath.Abs("../../../")
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}
	var pkgInfo = &list.Package{
		ImportPath: "tde/examples/word-reverse",
	}
	clone, err := inject.WithCreatingSample(absPath, "examples/word-reverse", pkgInfo, "TDE_WordReverse")
	if err != nil {
		t.Error(errors.Wrapf(err, "prep"))
	}

	var candidates = []*models.Candidate{}
	for i := 0; i < 10; i++ {
		candidates = append(candidates, &models.Candidate{
			UUID:    models.CandidateID(uuid.New().String()),
			BreedID: models.BreedID(uuid.New().String()),
			File:    []byte(`hello world`),
		})
	}

	var (
		modulePath = clone
		config     = &types.TestDetails{
			PackagePath:  "testdata/word-reverse",
			Package:      pkgInfo,
			ImplFuncFile: "testdata/word-reverse/word_reverse.go",
			TestFuncFile: "testdata/word-reverse/word_reverse_tde.go",
			TestFuncName: "TDE_WordReverse",
		}
	)
	sm := New(modulePath, config)
	sm.PlaceCandidatesIntoSlots(candidates)

	fmt.Println(sm.tmp)
	sm.Print()
}
