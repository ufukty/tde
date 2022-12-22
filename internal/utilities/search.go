package utilities

import "math"

func BinaryRangeSearch(values []float64, key float64) int {
	midIndex := func(lo, hi int) int {
		return int(math.Floor(float64(lo+hi) / 2))
	}

	values = append(values, math.MaxFloat64) // solve out of range access

	var (
		lo  = 0
		hi  = len(values) - 1
		mid int
	)

	for lo < hi {
		mid = midIndex(lo, hi)

		if values[mid] <= key && key < values[mid+1] {
			return mid
		}

		if values[mid] <= key {
			lo = mid
		} else {
			hi = mid
		}
	}

	return -1
}

func IsInSlice[T comparable](key T, slice []T) bool {
	for _, v := range slice {
		if key == v {
			return true
		}
	}
	return false
}
