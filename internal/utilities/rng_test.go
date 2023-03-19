package utilities

import (
	"fmt"
	"log"
	"math"
	"testing"

	"golang.org/x/exp/slices"
)

func IndexOfMax(values []int) int {
	indexOfMax := 0
	for i := 1; i < len(values); i++ {
		if values[indexOfMax] > values[i] {
			indexOfMax = i
		}
	}
	return indexOfMax
}

func Test_URandFloatForCrypto(t *testing.T) {
	totalNumberPerRun := 10000
	totalRun := 10
	mostFrequentRangePerRun := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for run := 0; run < totalRun; run++ {
		frequencies := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for i := 0; i < totalNumberPerRun; i++ {
			number := URandFloatForCrypto()
			if number < 0.0 || number > 1.0 {
				t.Error("Out of bounds random number")
			}
			frequencies[int(math.Floor(number*10))]++

		}
		log.Println(frequencies)
		mostFrequentRangePerRun[IndexOfMax(frequencies)]++
	}
	log.Println("mostFrequentRangePerRun:", mostFrequentRangePerRun)
}

func Test_Pick(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	freq := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 10000; i++ {
		freq[*Pick(array)]++
	}
	for i, fr := range freq {
		if fr == 0 {
			t.Errorf("TestPick didn't returned any number of %dth item.", i)
		}
	}
	fmt.Println(freq)
}

func Test_PickExcept(t *testing.T) {
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
			pick := *PickExcept(example.Slice, example.Except)
			if slices.Contains(example.Except, pick) {
				t.Errorf("validation i='%d', exception='%d'", i, pick)
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
			index := PickWeightedIndex(weights)
			counters[index]++
		}

		expectedFrequencies := ProportionItemsToTotal(weights)
		gotFrequencies := ProportionItemsToTotal(counters)

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
