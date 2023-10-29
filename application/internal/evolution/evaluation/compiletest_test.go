package evaluation

import (
	"fmt"
	"tde/internal/utilities"
	models "tde/models/program"
	"testing"
)

func Test_Compile(t *testing.T) {
	evaluator, ctx, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	for _, layer := range []models.Layer{models.Code, models.Program, models.Candidate, models.Solution} {
		for i, ast := range examples[layer] {
			fmt.Println(">>> testing the example", layer, i)

			subjects := models.Subjects{}
			subject := ctx.NewSubject()
			subject.AST = ast
			subjects.Add(subject)

			if err := evaluator.print(subject); err != nil {
				t.Fatal(fmt.Errorf("prep, printing: %w", err))
			}
			if err = evaluator.mount(subject); err != nil {
				t.Fatal(fmt.Errorf("prep, mounting: %w", err))
			}
			if err := evaluator.compile(subject); err != nil {
				t.Fatal(fmt.Errorf("act: %w", err))
			}

			if subject.Fitness.Layer() < models.Code {
				t.Errorf("assert, got=%q expected=Code (%s/%d, fitness=%.3f)\n%s",
					subject.Fitness.Layer(), layer, i, subject.Fitness.Flat(), utilities.IndentLines(string(subject.Code), 4))
			}
		}
	}
}

func Test_Test(t *testing.T) {
	evaluator, ctx, err := prepare()
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	for _, layer := range []models.Layer{models.Program, models.Candidate, models.Solution} {
		for i, ast := range examples[layer] {
			fmt.Println(">>> testing the example", layer, i)

			subjects := models.Subjects{}
			subject := ctx.NewSubject()
			subject.AST = ast
			subjects.Add(subject)

			if err := evaluator.print(subject); err != nil {
				t.Fatal(fmt.Errorf("prep, printing: %w", err))
			}
			if err = evaluator.mount(subject); err != nil {
				t.Fatal(fmt.Errorf("prep, mounting: %w", err))
			}
			if err := evaluator.compile(subject); err != nil {
				t.Fatal(fmt.Errorf("act: %w", err))
			}
			if err := evaluator.test(subject); err != nil {
				t.Fatal(fmt.Errorf("act: %w", err))
			}

			if subject.Fitness.Layer() < models.Program {
				t.Errorf("assert, got=%q expected=Code (%s/%d, fitness=%.3f)\n%s",
					subject.Fitness.Layer(), layer, i, subject.Fitness.Flat(), utilities.IndentLines(string(subject.Code), 4))
			}
		}
	}
}
