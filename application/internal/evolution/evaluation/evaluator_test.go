package evaluation

import (
	"fmt"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/evaluation/slotmgr"
	models "tde/models/program"
	"testing"
)

const ( // applies to all test cases in this package
	POPULATION = 3
)

func prepareEvaluator(ctx *models.Context) (*Evaluator, error) {
	pkgs, err := list.ListPackagesInDir("testdata")
	if err != nil {
		return nil, fmt.Errorf("listing packages in target dir: %w", err)
	}
	sample, err := inject.WithCreatingSample(pkgs.First().Module.Dir, pkgs.First(), "TDE_WordReverse")
	if err != nil {
		return nil, fmt.Errorf("creating sample: %w", err)
	}
	sm := slotmgr.New(sample, "internal/evolution/evaluation/testdata", "words.go")
	evaluator := NewEvaluator(sm, ctx)
	return evaluator, nil
}

func prepareSubjects(ctx *models.Context) models.Subjects {
	subjects := models.Subjects{}
	for i := 0; i < POPULATION; i++ { //
		subjects.Add(ctx.NewSubject())
	}
	return subjects
}

func prepare() (*Evaluator, models.Subjects, error) {
	ctx, err := models.LoadContext("../../..", "testdata", "WordReverse")
	if err != nil {
		return nil, nil, fmt.Errorf("finding the context for package: %w", err)
	}
	e, err := prepareEvaluator(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("creating evaluator: %w", err)
	}
	s := prepareSubjects(ctx)
	return e, s, nil
}

// FIXME: check fitness has populated after syntax errors
func Test_Pipeline(t *testing.T) {
	evaluator, subjects, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	if err := evaluator.Pipeline(subjects); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}
}
