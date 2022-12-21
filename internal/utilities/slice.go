package utilities

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

type ChainableSlice[T any] []T

func (s *ChainableSlice[T]) ForEach(f func(i int, v T)) *ChainableSlice[T] {
	for i, v := range *s {
		f(i, v)
	}
	return s
}

func (s *ChainableSlice[T]) Filter(f func(i int, v T) bool) *ChainableSlice[T] {
	filtered := ChainableSlice[T]{}
	for i, v := range *s {
		if f(i, v) {
			filtered = append(filtered, (*s)[i])
		}
	}
	return &filtered
}
