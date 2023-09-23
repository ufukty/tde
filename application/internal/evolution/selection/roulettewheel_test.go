package selection

import (
	"fmt"
	models "tde/models/program"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_RouletteWheel(t *testing.T) {
	const (
		runs    = 100
		bullets = 10
	)
	var (
		inputFitnesses = []float64{1.0, 0.22, 0.20, 0.28, 0.18, 0.35, 0.93, 0.21, 0.12, 0.39, 0.33, 0.0, 0.34, 0.26, 0.28, 0.30, 0.34, 0.22, 0.0, 0.29, 0.21, 0.22, 0.39, 0.39, 0.32, 0.32, 0.15, 0.24, 0.92, 0.28, 0.19, 0.0, 0.74, 0.10, 0.22, 0.30, 0.16, 0.35, 1.0, 0.31}
		best10         = []float64{0.35, 0.35, 0.39, 0.39, 0.39, 0.74, 0.92, 0.93, 1.0, 1.0}
		worst10        = []float64{0.0, 0.0, 0.0, 0.10, 0.12, 0.15, 0.16, 0.18, 0.19, 0.20}
		candidates     = []*models.Candidate{}
	)
	for _, f := range inputFitnesses {
		candidates = append(candidates, &models.Candidate{Fitness: models.Fitness{AST: f, Code: 0.0, Program: 0.0, Solution: 0.0}})
	}

	imbalancedElimination := 0
	for i := 0; i < runs; i++ {
		survivors := RouletteWheel(candidates, bullets)
		if len(candidates)-bullets != len(survivors) {
			t.Fatal(fmt.Errorf("assert: len(candidates) = %d", len(survivors)))
		}
		var survivingWorst10 = 0
		var survivingBest10 = 0
		for _, cand := range survivors {
			if slices.Contains(best10, cand.Fitness.AST) {
				survivingBest10++
			}
			if slices.Contains(worst10, cand.Fitness.AST) {
				survivingWorst10++
			}
		}
		if survivingWorst10 > survivingBest10 {
			fmt.Printf("Run %d: Survival rate amongst worst and best tens are %d, %d !!!\n", i, survivingWorst10, survivingBest10)
			imbalancedElimination++
		} else {
			fmt.Printf("Run %d: Survival rate amongst worst and best tens are %d, %d\n", i, survivingWorst10, survivingBest10)
		}
	}
	fmt.Println()
	if imbalancedElimination > runs*0.33 {
		t.Fatal(fmt.Errorf("assert %d of the runs suffer imbalanced elimination", imbalancedElimination))
	}
	fmt.Printf("imbalanced eliminations = %d\n", imbalancedElimination)
}
