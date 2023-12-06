package slicew

func WithoutLast[T any](slice []T) []T {
	return slice[:len(slice)-1]
}

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

func Zip[T any](a, b []T) []*[2]T {
	pairs := []*[2]T{}
	l := min(len(a), len(b))
	for i := 0; i < l; i++ {
		pairs = append(pairs, &[2]T{a[i], b[i]})
	}
	return pairs
}

func Pop[T any](slice []T) ([]T, T) {
	lastItem := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return slice, lastItem
}
