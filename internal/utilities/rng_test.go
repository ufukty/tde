package utilities

import (
	"fmt"
	"log"
	"math"
	"testing"
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
