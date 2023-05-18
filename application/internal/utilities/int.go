package utilities

import "math"

func Min[N Number](values ...N) N {
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}

func Max[N Number](values ...N) N {
	max := values[0]
	for i := 1; i < len(values); i++ {
		if max < values[i] {
			max = values[i]
		}
	}
	return max
}

func Floor(i float64) int {
	return int(math.Floor(i))
}

func Ceil(i float64) int {
	return int(math.Ceil(i))
}

func GetPrimeFactors(number int) (factors []int) {
	var (
		total       = 1
		remaining   = number
		last_factor = 2
	)

	for last_factor <= number {
		if remaining%last_factor == 0 {
			factors = append(factors, last_factor)
			total *= last_factor
			remaining /= last_factor
			if remaining == 1 {
				return
			}
		} else {
			last_factor++
		}
	}
	return
}
