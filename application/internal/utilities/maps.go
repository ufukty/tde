package utilities

import (
	"golang.org/x/exp/slices"
)

// A map implementation that doesn't lose ordering. Not performant as builtin map

type OrderedDict[K comparable, V any] struct {
	mapping  map[K]V
	ordering []K
}

func (od *OrderedDict[K, V]) Set(k K, v V) {
	od.mapping[k] = v
	od.ordering = append(od.ordering, k)
}

func (od *OrderedDict[K, V]) Get(k K) (value V, ok bool) {
	v, ok := od.mapping[k]
	return v, ok
}

func (od OrderedDict[K, V]) Delete(k K) {}

func (od OrderedDict[K, V]) InsertAt(k K, v V, index int) {}

func (od OrderedDict[K, V]) Keys() []K {
	return slices.Clone(od.ordering)
}

// separates the keys and values of a map into two array
func MapItems[K comparable, V any](in map[K]V) ([]K, []V) {
	keys, values := make([]K, len(in)), make([]V, len(in))
	for k, v := range in {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

// returns m / c, with O(n)
func MapDiff[M map[K]V, K comparable, V any](m, c M) M {
	d := make(M, len(m))
	for k := range m {
		if _, ok := c[k]; !ok {
			d[k] = m[k]
		}
	}
	return d
}

func MapSearchKey[K, V comparable](m map[K]V, v V) (K, bool) {
	for k, v1 := range m {
		if v1 == v {
			return k, true
		}
	}
	return *new(K), false
}
