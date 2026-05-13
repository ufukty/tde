package pick

import (
	"fmt"
	"slices"
	"testing"
)

func Test_Pick(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	freq := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 10000; i++ {
		p, _ := Pick(array)
		freq[p]++
	}
	for i, fr := range freq {
		if fr == 0 {
			t.Errorf("TestPick didn't returned any number of %dth item.", i)
		}
	}
	fmt.Println(freq)
}

func Test_PickExceptInt(t *testing.T) {
	examples := []struct {
		Slice  []int
		Except []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, []int{0, 2, 3, 4, 8}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 2, 3, 4, 9}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 2, 3, 4, 5, 9}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for i, example := range examples {
		for j := 0; j < 200; j++ {
			p, err := Except(example.Slice, example.Except)
			if err != nil {
				t.Fatal(fmt.Errorf("act %d/%d: %w", i, j, err))
			}
			if slices.Contains(example.Except, p) {
				t.Errorf("validation i='%d', exception='%d'", i, p)
			}

		}
	}
}
