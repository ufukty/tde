package knapsack

import "testing"

func Test_Knapsack(t *testing.T) {

	cases := map[*struct {
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

	for input, output := range cases {
		AssertArrays(t, output.placement, Knapsack(input.weights, input.prices))
	}

}

func AssertArrays[T int | float64](t *testing.T, left, right []T) {
	if len(left) != len(right) {
		t.Fatalf("Array lengths are different")
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			t.Fatalf("Arrays have different items.")
		}
	}
}
