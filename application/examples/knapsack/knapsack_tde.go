//go:build tde
// +build tde

package knapsack

import "tde/pkg/testing"

func TDE_Knapsack(e *testing.T) {

	examples := map[*struct {
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
		// add more example
	}

	for input, want := range examples {
		output := Knapsack(input.weights, input.prices)
		e.Assert(output, want)
	}
}
