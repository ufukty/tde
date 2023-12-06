package dict

import "slices"

// Ordered Map. Not performant as builtin map
type Ordered[K comparable, V any] struct {
	mapping  map[K]V
	ordering []K
}

func (od *Ordered[K, V]) Set(k K, v V) {
	od.mapping[k] = v
	od.ordering = append(od.ordering, k)
}

func (od *Ordered[K, V]) Get(k K) (value V, ok bool) {
	v, ok := od.mapping[k]
	return v, ok
}

func (od Ordered[K, V]) Delete(k K) {}

func (od Ordered[K, V]) InsertAt(k K, v V, index int) {}

func (od Ordered[K, V]) Keys() []K {
	return slices.Clone(od.ordering)
}
