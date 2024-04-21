package pick

import (
	"fmt"
	"tde/internal/utilities/numerics"
	"tde/internal/utilities/randoms"
	"tde/internal/utilities/setops"

	"golang.org/x/exp/constraints"
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
	cleaned := setops.Diff(s, e)
	if len(cleaned) == 0 {
		return *new(T), ErrEmptySlice
	}
	return Pick(cleaned)
}

func Coin() bool {
	p, _ := Pick([]bool{true, false})
	return p
}

func Weighted[T any, N constraints.Integer | constraints.Float](slice []T, weight []N) (T, error) {
	i, err := WeightedIndex(weight)
	if err != nil {
		return *new(T), err
	}
	return slice[i], nil
}

func WeightedIndex[N constraints.Integer | constraints.Float](weights []N) (int, error) {
	if len(weights) == 0 {
		return -1, ErrEmptySlice
	}
	weightsCumulative := numerics.Cumulate(weights)
	rnd := N(randoms.UniformCryptoFloat() * float64(weightsCumulative[len(weightsCumulative)-1]))
	index := numerics.BisectRight(weightsCumulative, rnd)
	return index, nil
}
