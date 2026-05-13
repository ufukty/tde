package pick

import (
	"fmt"

	"tde/internal/utilities/randoms"
	"tde/internal/utilities/sets"
)

var ErrEmptySlice = fmt.Errorf("empty slice")

func Pick[T any](s []T) (T, error) {
	if len(s) == 0 {
		return *new(T), ErrEmptySlice
	}
	return s[randoms.UniformIntN(len(s))], nil
}

func Except[T comparable](s []T, e []T) (T, error) {
	if len(s) == 0 {
		return *new(T), ErrEmptySlice
	}
	cleaned := sets.Diff(s, e)
	if len(cleaned) == 0 {
		return *new(T), ErrEmptySlice
	}
	return Pick(cleaned)
}

func Coin() bool {
	p, _ := Pick([]bool{true, false})
	return p
}

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}
