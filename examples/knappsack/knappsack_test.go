package knappsack

import (
	"testing"

	"github.com/pkg/errors"
)

func AssertArrays[T int | float64](left, right []T) error {
	if len(left) != len(right) {
		return errors.New("Array lengths are different")
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return errors.New("Arrays have different items.")
		}
	}
	return nil
}

func Test_Knappsack(t *testing.T) {

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
		err := AssertArrays(output.placement, Knappsack(input.weights, input.prices))
		if err != nil {
			t.Error(errors.Wrap(err, "Miscalculation"))
		}
	}

}
