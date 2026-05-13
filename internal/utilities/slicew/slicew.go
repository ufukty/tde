package slicew

func Pop[T any](slice []T) ([]T, T) {
	return slice[:len(slice)-1], slice[len(slice)-1]
}
