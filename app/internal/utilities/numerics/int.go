package numerics

func PrimeFactors(number int) (factors []int) {
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
