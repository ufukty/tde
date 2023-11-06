package testing

import (
	"math"
	"testing"
)

func areFloatsEqual(a, b float64) bool {
	return math.Floor(100000*a) == math.Floor(100000*b)
}

func Test_StringDistance(t *testing.T) {
	testCases := []struct {
		a    string
		b    string
		dist float64
	}{
		{"123", "123", 0.0},
		{"lll", "lll", 0.0},
		{"öçşiğü", "öçşiğü", 0.0},
		{"???", "???", 0.0},
		{"123", "123", 0.0},
		{"12", "12x", 1 - 2.0/3.0},
		{"123x", "123y", 1 - 3.0/4.0},
		{"123xy", "123yx", 1 - 3.0/5.0},
		{"123", "xyz", 1.0},
	}

	for _, testCase := range testCases {
		if _, d := distanceString(testCase.a, testCase.b); !areFloatsEqual(d, testCase.dist) {
			t.Errorf("StringDistance('%s', '%s') is expected to be '%f', got '%f'", testCase.a, testCase.b, testCase.dist, d)
		}
	}
}

func Test_DistanceRules(t *testing.T) {
	type AssertionPair struct{ lhs, rhs any }
	type TC struct{ Less, More AssertionPair }
	tcases := map[string]TC{
		"integers": {
			AssertionPair{0, 0},
			AssertionPair{0, 1},
		},
		"absolute values": {
			AssertionPair{0, 0},
			AssertionPair{0, -1},
		},
		"strings": {
			AssertionPair{"0", "0"},
			AssertionPair{"0", "1"},
		},
		"diff. length strings": {
			AssertionPair{"0", "0"},
			AssertionPair{"0", "00"},
		},
		"same length, diff. char. strings": {
			AssertionPair{"Hello world", "Hello world"},
			AssertionPair{"Hello world", "Hell0 w0rld"},
		},
		"different length arrays": {
			AssertionPair{[]int{}, []int{}},
			AssertionPair{[]int{}, []int{0}},
		},
		"different items, same length": {
			AssertionPair{[]int{0}, []int{0}},
			AssertionPair{[]int{0}, []int{1}},
		},
		"missing item at the end": {
			AssertionPair{
				[]int{0, 1, 2, 3, 4, 5},
				[]int{0, 1, 2, 3, 4, 5}},
			AssertionPair{
				[]int{0, 1, 2, 3, 4, 5},
				[]int{0, 1, 2, 3, 4}},
		},
		"missing item at beginning": {
			AssertionPair{
				[]int{0, 1, 2, 3, 4, 5},
				[]int{0, 1, 2, 3, 4, 5}},
			AssertionPair{
				[]int{0, 1, 2, 3, 4, 5},
				[]int{1, 2, 3, 4, 5}},
		},
		"arrays with a missing item & a different value": {
			AssertionPair{
				[]int{0, 1, 2, 3, 4, 5},
				[]int{0, 1, 2, 3, 4, 0}},
			AssertionPair{
				[]int{0, 1, 2, 3, 4, 5},
				[]int{0, 1, 2, 3, 4}},
		},
	}

	for tname, tcase := range tcases {
		t.Run(tname, func(t *testing.T) {
			tde := &T{}
			tde.Assert(tcase.Less.lhs, tcase.Less.rhs)
			tde.Assert(tcase.More.lhs, tcase.More.rhs)
			l := tde.AssertionErrorDistance[0]
			r := tde.AssertionErrorDistance[1]
			if l >= r {
				t.Errorf("Δ(%v, %v)=%.3f >= Δ(%v, %v)=%.3f",
					tcase.Less.lhs, tcase.Less.rhs, l, tcase.More.lhs, tcase.More.rhs, r)
			}
		})
	}
}
