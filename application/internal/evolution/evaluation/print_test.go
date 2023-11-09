package evaluation

import (
	"fmt"
	"tde/internal/evolution/models"
	"tde/internal/utilities"
	"testing"
)

func Test_Print(t *testing.T) {
	ctx, err := models.LoadContext("../../..", "testdata/context", "TDE_WordReverse")
	if err != nil {
		t.Fatal(fmt.Errorf("prep, finding the context for package: %w", err))
	}
	e := Evaluator{ctx: ctx}

	for _, expectedLayer := range []models.Layer{models.AST, models.Code, models.Program, models.Candidate, models.Solution} {
		for i, example := range examples[expectedLayer] {
			t.Run(fmt.Sprintf("%s-%d", expectedLayer, i), func(t *testing.T) {

				subj := ctx.NewSubject()
				subj.AST = example
				e.print(subj)

				if expectedLayer == models.AST {
					if subj.Fitness.Layer() != models.AST {
						t.Errorf("assert, layer mismatch (expected layer=%s, got=%s, fitness=%.3f)\n%s",
							expectedLayer, subj.Fitness.Layer(), subj.Fitness.Flat(), utilities.IndentLines(string(subj.Code), 4))
					}
					if len(subj.Code) != 0 {
						t.Errorf("assert, unexpectedly populated (expected layer=%s, got=%s, fitness=%.3f)\n%s",
							expectedLayer, subj.Fitness.Layer(), subj.Fitness.Flat(), utilities.IndentLines(string(subj.Code), 4))
					}
				} else {
					if subj.Fitness.Layer() != models.Code {
						t.Errorf("assert, layer mismatch (expected layer=%s, got=%s, fitness=%.3f)",
							expectedLayer, subj.Fitness.Layer(), subj.Fitness.Flat())
					}
					if len(subj.Code) == 0 {
						t.Errorf("assert, unexpectedly empty (expected layer=%s, got=%s, fitness=%.3f)",
							expectedLayer, subj.Fitness.Layer(), subj.Fitness.Flat())
					}
				}
			})
		}
	}

}
