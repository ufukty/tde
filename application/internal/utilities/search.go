package utilities

func midIndex(lo, hi int) int {
	return Floor(float64(lo+hi) / 2)
}

// Returns the index of leftmost (smallest) element greater than (or equal to) the key. range: [0, len]
func BisectLeft[N Number](values []N, key N) int {
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
func BisectRight[N Number](values []N, key N) int {
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

func IsInSlice[T comparable](key T, slice []T) bool {
	for _, v := range slice {
		if key == v {
			return true
		}
	}
	return false
}
