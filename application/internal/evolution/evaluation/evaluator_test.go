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

func prepare() (*Evaluator, models.Subjects, error) {
	ctx, err := models.LoadContext("../../..", "testdata", "TDE_WordReverse")
	if err != nil {
		return nil, nil, fmt.Errorf("finding the context for package: %w", err)
	}
	subjects := models.Subjects{}
	for i := 0; i < POPULATION; i++ { //
		subjects.Add(ctx.NewSubject())
	}
	pkgs, err := list.ListPackagesInDir("testdata")
	if err != nil {
		return nil, nil, fmt.Errorf("listing packages in target dir: %w", err)
	}
	sample, err := inject.WithCreatingSample("../../..", pkgs.First(), "TDE_WordReverse")
	if err != nil {
		return nil, nil, fmt.Errorf("creating sample: %w", err)
	}
	sm := slotmgr.New(sample, "internal/evolution/evaluation/testdata", "words.go")
	if err := sm.PlaceSubjectsIntoSlots(subjects); err != nil {
		return nil, nil, fmt.Errorf("passing subjects into slot manager: %w", err)
	}
	syntaxCheckAndProduceCode(ctx, subjects)
	evaluator := NewEvaluator(sm, ctx)
	return evaluator, subjects, nil
}

func Test_PrintAndCompile(t *testing.T) {
	evaluator, subjects, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	if err := evaluator.runSubjects(subjects); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}
}

// FIXME: check fitness has populated after syntax errors
func Test_Pipeline(t *testing.T) {
	evaluator, subjects, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	evaluator.sm.Print()

	if err := evaluator.Pipeline(subjects); err != nil {
		t.Fatal(fmt.Errorf("act: %w", err))
	}
}
