package mapw


func FindKey[K, V comparable](m map[K]V, v V) (K, bool) {
	for k, v1 := range m {
		if v1 == v {
			return k, true
		}
	}
	return *new(K), false
}

