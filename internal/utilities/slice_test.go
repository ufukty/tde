package utilities

import (
	"go/token"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_SliceRemoveItem(t *testing.T) {
	examples := map[*struct {
		Slice  []int
		Remove int
	}][]int{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0}: {1, 2, 3, 4, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 1}: {0, 2, 3, 4, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 2}: {0, 1, 3, 4, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3}: {0, 1, 2, 4, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 4}: {0, 1, 2, 3, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5}: {0, 1, 2, 3, 4, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 6}: {0, 1, 2, 3, 4, 5, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 7}: {0, 1, 2, 3, 4, 5, 6, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 8}: {0, 1, 2, 3, 4, 5, 6, 7, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9}: {0, 1, 2, 3, 4, 5, 6, 7, 8},
	}

	for input, want := range examples {
		got := SliceExceptItem(input.Slice, input.Remove)
		if slices.Compare(want, got) != 0 {
			t.Error("validation")
		}
	}

	got := SliceExceptItems([]token.Token{token.ASSIGN, token.DEFINE}, []token.Token{token.DEFINE})
	if slices.Compare(got, []token.Token{token.ASSIGN}) != 0 {
		t.Error("validation")
	}

}
func Test_SliceRemoveItems(t *testing.T) {
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
		got := SliceExceptItems(input.Slice, input.Except)
		if slices.Compare(want, got) != 0 {
			t.Error("validation")
		}
	}

	got := SliceExceptItems([]token.Token{token.ASSIGN, token.DEFINE}, []token.Token{token.DEFINE})
	if slices.Compare(got, []token.Token{token.ASSIGN}) != 0 {
		t.Error("validation")
	}

}
