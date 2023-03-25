package slot_manager

func slicePop[T any](slice []T) ([]T, T) {
	lastItem := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return slice, lastItem
}
