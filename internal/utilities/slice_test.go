package utilities

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Test_SliceRemoveItems(t *testing.T) {
	inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	inputExcept := []int{0, 2, 3, 4, 9}
	want := []int{1, 5, 6, 7, 8}

	got := SliceRemoveItems(inputSlice, inputExcept)
	if slices.Compare(want, got) != 0 {
		t.Error("validation")
	}
}
