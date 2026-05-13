package sets

func lookup[T comparable](s []T) map[T]bool {
	m := make(map[T]bool, len(s))
	for i := range s {
		m[s[i]] = true
	}
	return m
}

// returns L/R
func Diff[T comparable](l, r []T) []T {
	lu := lookup(r)
	d := make([]T, 0, len(l))
	for _, v := range l {
		if _, ok := lu[v]; !ok {
			d = append(d, v)
		}
	}
	return d
}
