package utilities

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

func (s *ChainableSlice[T]) Map(f func(i int, v T) T) *ChainableSlice[T] {
	list := ChainableSlice[T]{}
	for i, v := range *s {
		list = append(list, f(i, v))
	}
	return &list
}
