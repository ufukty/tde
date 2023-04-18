package utilities

import "golang.org/x/exp/slices"

func CompareSlices[T comparable](l, r []T) bool {
	if len(l) != len(r) {
		return false
	}
	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			return false
		}
	}

	return true
}

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func GetCumulative[N Number](input []N) []N {
	var output = []N{}
	var total = N(0)
	for _, v := range input {
		total += v
		output = append(output, total)
	}
	return output
}

func ForEach[T any](slice []T, callback func(index int, value T)) {
	for i, v := range slice {
		callback(i, v)
	}
}

func SliceRemoveLast[T any](slice []T) []T {
	return slice[:len(slice)-1]
}

func SliceLast[T any](slice []T) T {
	return slice[len(slice)-1]
}

func SliceZipToMap[K comparable, V any](a []K, b []V) map[K]V {
	pairs := map[K]V{}
	l := Min(len(a), len(b))
	for i := 0; i < l; i++ {
		pairs[a[i]] = b[i]
	}
	return pairs
}

func SliceZipToSlice[T any](a, b []T) []*[2]T {
	pairs := []*[2]T{}
	l := Min(len(a), len(b))
	for i := 0; i < l; i++ {
		pairs = append(pairs, &[2]T{a[i], b[i]})
	}
	return pairs
}

func Map[T any, V any](slice []T, callback func(i int, value T) V) []V {
	list := []V{}
	for i, v := range slice {
		list = append(list, callback(i, v))
	}
	return list
}

func FilteredMap[T any, V any](slice []T, callback func(i int, value T) (V, bool)) []V {
	list := []V{}
	for i, v := range slice {
		if item, ok := callback(i, v); ok {
			list = append(list, item)
		}
	}
	return list
}

func SliceExceptItem[T comparable](s []T, v T) []T {
	n := slices.Clone(s)
	if i := slices.Index(n, v); i != -1 {
		return append(n[:i], n[i+1:]...)
	}
	return n
}

func SliceExceptItems[T comparable](s []T, e []T) []T {
	for _, it := range e {
		s = SliceExceptItem(s, it)
	}
	return s
}

// example input -> output:
//
//	[0, 4, 6, 10] -> [0, 0.2, 0.3, 0.5]
func ProportionItemsToTotal[N Number](slice []N) (proportions []float64) {
	total := N(0)
	for _, item := range slice {
		total += item
	}
	for _, item := range slice {
		proportions = append(proportions, float64(item)/float64(total))
	}
	return
}

func SlicePop[T any](slice []T) ([]T, T) {
	lastItem := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return slice, lastItem
}

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
