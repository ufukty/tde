package evaluation

import (
	"fmt"
	"sort"
	"tde/internal/evolution/evaluation/inject"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/evaluation/slotmgr"
	"tde/internal/evolution/models"
	"tde/internal/utilities"
	"testing"
)

func prepareEvaluator(ctx *models.Context) (*Evaluator, error) {
	pkgs, err := list.ListPackagesInDir("testdata/context")
	if err != nil {
		return nil, fmt.Errorf("listing packages in target dir: %w", err)
	}
	sample, err := inject.WithCreatingSample(pkgs.First().Module.Dir, pkgs.First(), "TDE_WordReverse")
	if err != nil {
		return nil, fmt.Errorf("creating sample: %w", err)
	}
	sm := slotmgr.New(sample, "internal/evolution/evaluation/testdata/context", "words.go")
	evaluator := NewEvaluator(sm, ctx)
	return evaluator, nil
}

func prepare() (*Evaluator, *models.Context, error) {
	ctx, err := models.LoadContext("../../..", "testdata/context", "WordReverse")
	if err != nil {
		return nil, nil, fmt.Errorf("finding the context for package: %w", err)
	}
	e, err := prepareEvaluator(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("creating evaluator: %w", err)
	}
	return e, ctx, nil
}

func Test_Pipeline(t *testing.T) {
	evaluator, ctx, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	sorted := []*models.Subject{}
	for _, layer := range []models.Layer{models.AST, models.Code, models.Program, models.Candidate, models.Solution} {
		for i, ast := range examples[layer] {
			t.Run(fmt.Sprintf("%s-%d", layer, i), func(t *testing.T) {
				subjects := models.Subjects{}
				subject := ctx.NewSubject()
				subject.AST = ast
				subjects.Add(subject)

				if err := evaluator.Pipeline(subjects); err != nil {
					t.Fatal(fmt.Errorf("act: %w", err))
				}

				if subject.Fitness.Layer() != layer {
					t.Errorf("assert, layer mistmatch: got=%q (%s/%d, fitness=%.3f)\n%s",
						subject.Fitness.Layer(), layer, i, subject.Fitness.Flat(), utilities.IndentLines(string(subject.Code), 4))
				}

				sorted = append(sorted, subject)
			})
		}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Fitness.Flat() < sorted[j].Fitness.Flat()
	})
	fmt.Println("sorted by fitnesses:")
	for _, subj := range sorted {
		fmt.Printf("%s %10s %.10f\n", subj.Sid, subj.Fitness.Layer(), subj.Fitness.Flat())
	}
}
