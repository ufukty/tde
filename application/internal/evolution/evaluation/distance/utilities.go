package distance

import "math"

func convert[T any](a, b any) (T, T, bool) {
	if a, ok := a.(T); ok {
		if b, ok := b.(T); ok {
			return a, b, true
		}
	}
	return *new(T), *new(T), false
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Pow(math.E, -1*x))
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
