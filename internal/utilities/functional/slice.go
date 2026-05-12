package functional

func Map[T any, V any](source []T, c func(i int, value T) V) (dest []V) {
	for i, v := range source {
		dest = append(dest, c(i, v))
	}
	return
}

func Mapf[T any, V any](source []T, f func(i int, value T) (V, bool)) (dest []V) {
	for i, v := range source {
		if item, ok := f(i, v); ok {
			dest = append(dest, item)
		}
	}
	return
}
