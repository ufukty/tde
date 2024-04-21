package evolution

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/evaluation/slotmgr"
	"tde/internal/evolution/models"
	"testing"
)

func prepareSolutionSearch() (*SolutionSearch, error) {
	pkgs, err := list.ListPackagesInDir("testdata/words")
	if err != nil {
		return nil, fmt.Errorf("listing packages in the testdata package: %w", err)
	}
	sample, err := inject.WithCreatingSample(pkgs.First().Module.Dir, pkgs.First(), "TDE_WordReverse")
	if err != nil {
		return nil, fmt.Errorf("creating the sample (injected one): %w", err)
	}
	sm := slotmgr.New(sample, pkgs.First().PathInModule(), "words.go")
	ctx, err := models.LoadContext(pkgs.First().Module.Dir, "testdata/words", "WordReverse")
	if err != nil {
		return nil, fmt.Errorf("loading the context: %w", err)
	}
	ev := evaluation.NewEvaluator(sm, ctx)
	em := NewSolutionSearch(ev, defaults, ctx)
	return em, nil
}

func Test_SolutionSearch(t *testing.T) {
	em, err := prepareSolutionSearch()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	if err := em.Loop(); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}
}
