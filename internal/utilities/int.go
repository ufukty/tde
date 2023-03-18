package utilities

import "math"

func Min[N Number](a, b N) N {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max[N Number](a, b N) N {
	if a < b {
		return b
	} else {
		return a
	}
}

func Floor(i float64) int {
	return int(math.Floor(i))
}
