package setops

import (
	"go/token"
	"slices"
	"testing"
)

func Test_Diff(t *testing.T) {
	examples := map[*struct {
		Slice  []int
		Except []int
	}][]int{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{}}:                             {0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 2, 3, 4, 9}}:                {1, 5, 6, 7, 8},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}: {},
		{[]int{}, []int{}}: {},
	}

	for input, want := range examples {
		got := Diff(input.Slice, input.Except)
		if slices.Compare(want, got) != 0 {
			t.Error("validation")
		}
	}

	got := Diff([]token.Token{token.ASSIGN, token.DEFINE}, []token.Token{token.DEFINE})
	if slices.Compare(got, []token.Token{token.ASSIGN}) != 0 {
		t.Error("validation")
	}

}
