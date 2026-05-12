package slicew

// Python's range function
//
//	Range(start = 0, stop, step = 1)
//	Range(3) => [0, 1, 2]
//	Range(1, 3) => [1, 2]
//	Range(0, 3, 2) => [0, 2]
func Range(args ...int) []int {
	var start, stop, step int
	switch len(args) {
	case 1:
		start = 0
		stop = args[0]
		step = 1
	case 2:
		start = args[0]
		stop = args[1]
		step = 1
	case 3:
		if args[2] == 0 {
			panic("0 is not accepted as step size")
		}
		start = args[0]
		stop = args[1]
		step = args[2]
	default:
		panic("Only 1, 2 or 3 parameter is accepted")
	}
	if start == stop {
		return []int{}
	}
	length := (stop - start + step - 1) / step
	seq := make([]int, length)
	for i := 0; i < length; i++ {
		seq[i] = start + (step * i)
	}
	return seq
}
