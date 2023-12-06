package pick

import (
	"fmt"
	"slices"
	"tde/internal/utilities/numerics"
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

func TestWeightedIndex(t *testing.T) {
	const totalRunPerCase = 10000

	testCases := [][]int64{
		{4, 5, 2, 7, 8, 9, 6, 3, 1, 2},
		{0, 1},
		{1},
		{1, 1, 1, 1, 1},
		{9, 1},
		{3, 2, 1, 0},
		{0, 1, 0},
	}

	for i, weights := range testCases {
		counters := make([]int64, len(weights))

		for j := 0; j < totalRunPerCase; j++ {
			i, _ := WeightedIndex(weights)
			counters[i]++
		}

		expectedFrequencies := numerics.DivideBySum(weights)
		gotFrequencies := numerics.DivideBySum(counters)

		fmt.Printf("i=%d, weights=%v, counters=%v, frequencies:\n\texpected : ", i, weights, counters)
		for j := 0; j < len(weights); j++ {
			fmt.Printf("%d%% ", int(expectedFrequencies[j]*100))
		}
		fmt.Printf("\n\tgot      : ")
		for j := 0; j < len(weights); j++ {
			fmt.Printf("%d%% ", int(gotFrequencies[j]*100))
		}
		fmt.Println()

		for j := 0; j < len(weights); j++ {
			expectedFrequency, gotFrequency := expectedFrequencies[j], gotFrequencies[j]
			if !(0.8*expectedFrequency <= gotFrequency && gotFrequency <= expectedFrequency*1.2) {
				t.Errorf("validation. frequency of %d. index", j)
			}
		}
	}
}

func TestWeightedIndexFloat64(t *testing.T) {
	const totalRunPerCase = 10000

	testCasesFloat64 := [][]float64{
		{0.4, 0.5, 0.2, 0.7, 0.8, 0.9, 0.6, 0.3, 0.1, 0.2},
		{0.0, 0.1},
		{0.1},
		{0.1, 0.1, 0.1, 0.1, 0.1},
		{0.9, 0.1},
		{0.3, 0.2, 0.1, 0.0},
		{0.0, 0.1, 0.0},
	}

	for i, weights := range testCasesFloat64 {
		counters := make([]int64, len(weights))

		for j := 0; j < totalRunPerCase; j++ {
			i, _ := WeightedIndex(weights)
			counters[i]++
		}

		expectedFrequencies := numerics.DivideBySum(weights)
		gotFrequencies := numerics.DivideBySum(counters)

		fmt.Printf("i=%d, weights=%v, counters=%v, frequencies:\n\texpected : ", i, weights, counters)
		for j := 0; j < len(weights); j++ {
			fmt.Printf("%d%% ", int(expectedFrequencies[j]*100))
		}
		fmt.Printf("\n\tgot      : ")
		for j := 0; j < len(weights); j++ {
			fmt.Printf("%d%% ", int(gotFrequencies[j]*100))
		}
		fmt.Println()

		for j := 0; j < len(weights); j++ {
			expectedFrequency, gotFrequency := expectedFrequencies[j], gotFrequencies[j]
			if !(0.5*expectedFrequency <= gotFrequency && gotFrequency <= expectedFrequency*1.5) {
				t.Errorf("validation. frequency of %d. index", j)
			}
		}
	}

}
