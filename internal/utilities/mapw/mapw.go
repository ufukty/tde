package mapw

// returns m / c, with O(n)
func Diff[M map[K]V, K comparable, V any](m, c M) M {
	d := make(M, len(m))
	for k := range m {
		if _, ok := c[k]; !ok {
			d[k] = m[k]
		}
	}
	return d
}

func FindKey[K, V comparable](m map[K]V, v V) (K, bool) {
	for k, v1 := range m {
		if v1 == v {
			return k, true
		}
	}
	return *new(K), false
}

