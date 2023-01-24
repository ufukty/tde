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

type Number64 interface {
	~float64 | ~int64
}

func GetCumulative[N Number64](input []N) []N {
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

