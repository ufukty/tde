package evolution

import "testing"

func TestBinarySearchSmallestOfGreaters(t *testing.T) {

	values := []float64{
		0.5, 1.2, 1.5, 1.9, 2.4, 3.5, 5.0, 5.1, 5.3,
	}

	testCases := []struct {
		key  float64
		want int
	}{
		{0.0, 0},
		{0.4, 0},
		{0.5, 1},
		{0.6, 1},
		{1.2, 2},
		{1.499999, 2},
		{1.5, 3},
		{1.8, 3},
		{1.9, 4},
		{2.3, 4},
		{2.4, 5},
		{3.4, 5},
		{3.5, 6},
		{4.9, 6},
		{5.0, 7},
		{5.000001, 7},
		{5.1, 8},
		{5.2, 8},
	}

	for _, testCase := range testCases {
		got := BinarySearchSmallestOfGreaters(values, testCase.key)
		if got != testCase.want {
			t.Errorf("For %f; expected %d, got %d", testCase.key, testCase.want, got)
		}
	}

}
