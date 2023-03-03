//go:build tde
// +build tde

package knapsack

import "tde/pkg/tde"

func TDE_Knapsack(e *tde.E) {}

// func EvolveKnappsack(t *testing.Testing) {

// 	testParameters := map[*struct {
// 		prices  []float64
// 		weights []float64
// 	}]*struct {
// 		placement []int
// 	}{
// 		{
// 			prices:  []float64{0, 0, 0, 0, 0},
// 			weights: []float64{0, 0, 0, 0, 0},
// 		}: {
// 			placement: []int{0, 0, 0, 0, 1},
// 		},
// 		{
// 			prices:  []float64{0, 0, 0, 0, 0},
// 			weights: []float64{0, 0, 0, 0, 0},
// 		}: {
// 			placement: []int{0, 0, 0, 0, 1},
// 		},
// 		{
// 			prices:  []float64{0, 0, 0, 0, 0},
// 			weights: []float64{0, 0, 0, 0, 0},
// 		}: {
// 			placement: []int{0, 0, 0, 0, 1},
// 		},
// 		{
// 			prices:  []float64{0, 0, 0, 0, 0},
// 			weights: []float64{0, 0, 0, 0, 0},
// 		}: {
// 			placement: []int{0, 0, 0, 0, 1},
// 		},
// 	}

// 	t.Candidate(func(c *testing.Candidate) {

// 	})

// 	// for input, want := range testParameters {
// 	// 	err := assert.Equal(t, want.placement, Knappsack(input.weights, input.prices))
// 	// 	if err != nil {
// 	// 		t.Error(errors.Wrap(err, "Miscalculation"))
// 	// 	}
// 	// }

// 	// t.NewObjective("basic", func(o *testing.Objective) {

// 	// })
// }
