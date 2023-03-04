//go:build tde
// +build tde

package knapsack

import (
	"tde/pkg/tde"
)

func TDE_Knapsack(e *tde.E) {

	testParameters := map[*struct {
		prices  []float64
		weights []float64
	}]*struct {
		placement []int
	}{
		{
			prices:  []float64{0, 0, 0, 0, 0},
			weights: []float64{0, 0, 0, 0, 0},
		}: {
			placement: []int{0, 0, 0, 0, 1},
		},
		{
			prices:  []float64{0, 0, 0, 0, 0},
			weights: []float64{0, 0, 0, 0, 0},
		}: {
			placement: []int{0, 0, 0, 0, 1},
		},
		{
			prices:  []float64{0, 0, 0, 0, 0},
			weights: []float64{0, 0, 0, 0, 0},
		}: {
			placement: []int{0, 0, 0, 0, 1},
		},
		{
			prices:  []float64{0, 0, 0, 0, 0},
			weights: []float64{0, 0, 0, 0, 0},
		}: {
			placement: []int{0, 0, 0, 0, 1},
		},
	}

	for input, want := range testParameters {
		output := Knapsack(input.weights, input.prices)
		e.AssertEqual(output, want)
	}

	// t.NewObjective("basic", func(o *testing.Objective) {

	// })
}
