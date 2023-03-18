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
