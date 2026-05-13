package setops

func lookup[T comparable](s []T) map[T]bool {
	m := make(map[T]bool, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]] = true
	}
	return m
}

// returns L/R
func Diff[T comparable](l, r []T) []T {
	mr := lookup(r)
	d := make([]T, 0, len(l))
	for _, v := range l {
		if _, found := mr[v]; !found {
			d = append(d, v)
		}
	}
	return d
}
