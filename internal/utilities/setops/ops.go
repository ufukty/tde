package setops

import (
	"golang.org/x/exp/maps"
)

func prepLookupMap[T comparable](s []T) map[T]bool {
	m := make(map[T]bool, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]] = true
	}
	return m
}

func maxSizesForSubsets[T any](l, r []T) (union, intersects, differences int) {
	lmax := max(len(l), len(r))
	lmin := min(len(l), len(r))
	return lmax, lmin, lmax - lmin
}

// time: O(4*n) space: O(5*n)
func Combined[T comparable](l, r []T) (isect, diffl, diffr []T) {
	ml, mr := prepLookupMap(l), prepLookupMap(r)

	_, isect_s, diff_s := maxSizesForSubsets(l, r)
	isect_m := make(map[T]bool, isect_s)
	diffl_m := make(map[T]bool, diff_s)
	diffr_m := make(map[T]bool, diff_s)

	for v := range ml {
		if _, found := mr[v]; found {
			isect_m[v] = true
		} else {
			diffl_m[v] = true
		}
	}
	for v := range mr {
		if _, found := ml[v]; !found {
			diffr_m[v] = true
		}
	}
	return maps.Keys(isect_m), maps.Keys(diffl_m), maps.Keys(diffr_m)
}

func Intersect[T comparable](l, r []T) []T {
	ml := prepLookupMap(l)
	i := make([]T, 0, min(len(l), len(r)))
	for _, v := range r {
		if _, found := ml[v]; found {
			i = append(i, v)
		}
	}
	return i
}

// returns L/R (L - R)
func Diff[T comparable](l, r []T) []T {
	mr := prepLookupMap(r)
	d := make([]T, 0, len(l))
	for _, v := range l {
		if _, found := mr[v]; !found {
			d = append(d, v)
		}
	}
	return d
}

func Union[T comparable](l, r []T) []T {
	u := make([]T, 0, max(len(l), len(r)))
	for _, v := range l {
		u[len(u)] = v
	}
	m := prepLookupMap(u)
	for _, v := range r {
		if _, found := m[v]; !found {
			u = append(u, v)
		}
	}
	return u
}
