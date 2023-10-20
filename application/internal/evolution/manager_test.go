package evolution

import (
	"fmt"
	"tde/internal/evolution/evaluation"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/evaluation/slotmgr"
	models "tde/models/program"
	"testing"
)

func prepare(parameters *models.Parameters) (*SolutionSearch, error) {
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
	em := NewSolutionSearch(ev, parameters, ctx)
	return em, nil
}

func Test_SolutionSearch(t *testing.T) {
	parameters := &models.Parameters{
		Population:  0,
		Generations: 0,
		Size:        0,
		Packages:    []string{},
		Code: models.SearchParameters{
			Cap:         20,
			Generations: 3,
			Depth:       3,
			Evaluations: 2,
		},
		Program: models.SearchParameters{
			Cap:         4,
			Generations: 2,
			Depth:       2,
			Evaluations: 2,
		},
		Candidate: models.SearchParameters{
			Cap:         0,
			Generations: 2,
			Depth:       0,
			Evaluations: 0,
		},
	}
	em, err := prepare(parameters)
	if err != nil {
		t.Fatalf("prep: %w", err)
	}
	em.Loop()
}
