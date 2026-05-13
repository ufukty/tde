package numerics

import (
	"math"
)

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func Cumulate[N number](input []N) []N {
	var output []N
	total := N(0)
	for _, v := range input {
		total += v
		output = append(output, total)
	}
	return output
}

func midIndex(lo, hi int) int {
	return int(math.Floor(float64(lo+hi) / 2))
}

// Returns the index of leftmost (smallest) element greater than (or equal to) the key. range: [0, len]
func BisectLeft[N number](values []N, key N) int {
	var (
		low  = 0
		high = len(values)
		mid  int
	)
	for low < high {
		mid = midIndex(low, high)
		if values[mid] < key {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

// Returns the index of leftmost (smallest) element greater than the key. range: [0, len]
func BisectRight[N number](values []N, key N) int {
	var (
		low  = 0
		high = len(values)
		mid  int
	)
	for low < high {
		mid = midIndex(low, high)
		if key < values[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// example input -> output:
//
//	[0, 4, 6, 10] -> [0, 0.2, 0.3, 0.5]
func DivideBySum[N number](slice []N) (proportions []float64) {
	total := Sum(slice)
	for _, item := range slice {
		proportions = append(proportions, float64(item)/float64(total))
	}
	return
}

func Sum[N number](s []N) N {
	total := N(0)
	for _, item := range s {
		total += item
	}
	return total
}
