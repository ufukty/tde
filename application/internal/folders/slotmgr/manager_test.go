package slotmgr

import (
	"tde/internal/folders/discovery"
	"tde/internal/folders/inject"
	"tde/internal/folders/list"
	models "tde/models/program"

	"fmt"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
)

func Test_SlotManager(t *testing.T) {
	absPath, err := filepath.Abs("../../../")
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	var pkgInfo = &list.Package{
		ImportPath: "tde/examples/word-reverse",
	}
	sample, err := inject.WithCreatingSample(absPath, "examples/word-reverse", pkgInfo, "TDE_WordReverse")
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	var candidates = []*models.Candidate{}
	for i := 0; i < 10; i++ {
		candidates = append(candidates, &models.Candidate{
			UUID:    models.CandidateID(uuid.New().String()),
			BreedID: models.BreedID(uuid.New().String()),
			File:    []byte(`hello world`),
		})
	}

	var td = &discovery.TestDetails{
		PackagePath:  "testdata/word-reverse",
		Package:      pkgInfo,
		ImplFuncFile: "testdata/word-reverse/word_reverse.go",
		TestFuncFile: "testdata/word-reverse/word_reverse_tde.go",
		TestFuncName: "TDE_WordReverse",
	}
	var sm = New(sample, td)
	if err := sm.PlaceCandidatesIntoSlots(candidates); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}

	fmt.Println("sample module dir:", sample)
	fmt.Println("slot manager tmp:", sm.tmp)
	sm.Print()

	if len(sm.slots.assigned) < 10 {
		t.Fatal("assert. less then 10 slots assigned")
	}
	if len(sm.slots.free) != 0 {
		t.Fatal("assert. there should be no free slots")
	}

	sm.FreeAllSlots()
	sm.Print()

	if len(sm.slots.assigned) != 0 {
		t.Fatal("assert. there should be no assigned slots")
	}
	if len(sm.slots.free) < 10 {
		t.Fatal("assert. less then 10 slots freed")
	}
}
