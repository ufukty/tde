package slotmgr

import (
	"go/ast"
	"tde/internal/folders/discovery"
	"tde/internal/folders/inject"
	"tde/internal/folders/list"
	models "tde/models/program"

	"fmt"
	"testing"

	"github.com/google/uuid"
)

func Test_SlotManager(t *testing.T) {
	var (
		pkgs   list.Packages
		sample string
		err    error
	)
	if pkgs, err = list.ListPackagesInDir("../../../examples/word-reverse"); err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	var pkg = pkgs.First()
	if sample, err = inject.WithCreatingSample(pkg.Module.Dir, pkg, "TDE_WordReverse"); err != nil {
		t.Error(fmt.Errorf("prep: %w", err))
	}
	fmt.Println("sample module dir:", sample)

	var candidates = []*models.Candidate{}
	for i := 0; i < 10; i++ {
		candidates = append(candidates, &models.Candidate{
			UUID:    models.CandidateID(uuid.New().String()),
			BreedID: models.BreedID(uuid.New().String()),
			File:    []byte(`hello world`),
		})
	}

	var td = &discovery.CombinedDetails{
		Package: pkg,
		Target: &discovery.TargetFunction{
			Name: "",
			Path: "testdata/word-reverse/word_reverse.go",
			Line: 0,
		},
		Test: &discovery.TestFunction{
			Name:  "TDE_WordReverse",
			Path:  "testdata/word-reverse/word_reverse_tde.go",
			Line:  0,
			Calls: []*ast.CallExpr{},
		},
	}
	var sm = New(sample, td)
	if err := sm.PlaceCandidatesIntoSlots(candidates); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}

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
