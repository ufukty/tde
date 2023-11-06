package distance

import (
	"log"
	"math"
	"os"
	"reflect"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

func distanceNumber[T constraints.Integer | constraints.Float](a, b T) (eq bool, d float64) {
	return a == b, math.Abs(float64(a) - float64(b))
}

//   - O(n) string comparison
//   - fast, but not best [like levenshtein which is O(n^2)]
//   - logic: similarity of two strings is based on the portion of characters
//     are in same position at two strings
func distanceString(a, b string) (eq bool, d float64) {
	var (
		lenA        = float64(len(a))
		lenB        = float64(len(b))
		lenInner    = math.Min(lenA, lenB)
		lenOuter    = math.Max(lenA, lenB)
		commonChars = 0
	)
	for i := 0; i < int(lenInner); i++ {
		if a[i] == b[i] {
			commonChars++
		}
	}
	var (
		similarity = float64(commonChars) / lenOuter
		distance   = 1 - similarity
	)
	return a == b, distance
}

func distanceByte(a, b byte) (eq bool, d float64) {
	if a == b {
		return true, 0.0
	}
	return false, 1.0
}

func distanceBool(a, b bool) (eq bool, d float64) {
	if a == b {
		return true, 0.0
	}
	return false, 1.0
}

func Distance(a, b any) (eq bool, d float64) {
	if ax, bx, ok := convert[[]byte](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]int](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]int8](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]int16](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]int32](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]int64](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]uint8](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]uint16](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]uint32](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]uint64](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]float32](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]float64](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]bool](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]string](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]error](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[[]any](a, b); ok {
		return distanceArrays(ax, bx)

	} else if ax, bx, ok := convert[byte](a, b); ok {
		return distanceByte(ax, bx)

	} else if ax, bx, ok := convert[int](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[int8](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[int16](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[int32](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[int64](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[uint8](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[uint16](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[uint32](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[uint64](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[float32](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[float64](a, b); ok {
		return distanceNumber(ax, bx)

	} else if ax, bx, ok := convert[bool](a, b); ok {
		return distanceBool(ax, bx)

	} else if ax, bx, ok := convert[string](a, b); ok {
		return distanceString(ax, bx)

	} else if ax, bx, ok := convert[error](a, b); ok {
		return distanceString(ax.Error(), bx.Error())

	} else {
		log.Fatalf("Can not calculate the distance for types %q and %q", reflect.TypeOf(a), reflect.TypeOf(b))
		os.Exit(1)
		return false, -1.0
	}
}

// FIXME: Multi-Objective GP TODO: how to incorporate execution path
// things to consider for measuring the distance:
// - diff. array lengths
// - common items and duplicates
// - item orders
// - item distances for exclusive items
// - item distances for exclusive items with unmatching indexes
func distanceArrays[T any](a, b []T) (eq bool, d float64) {
	d = 0.0
	for i := 0; i < len(a) && i < len(b); i++ {
		if _, d_i := Distance(a[i], b[i]); d_i != 0.0 {
			d += sigmoid(d_i)
		}
	}
	d += float64(abs(len(a)-len(b))) * 1.0
	return true, d / float64(min(len(a), len(b)))
}

func maxSizesForSubsets[T any](l, r []T) (union, intersects, differences int) {
	lmax := max(len(l), len(r))
	lmin := min(len(l), len(r))
	return lmax, lmin, lmax - lmin
}

// time: O(4*n) space: O(5*n)
func subsets[T comparable](l, r []T) (intersect, diffl, diffr []T) {
	mapl := make(map[T]bool, len(l))
	mapr := make(map[T]bool, len(r))
	_, isect, diff := maxSizesForSubsets(l, r)
	commons := make(map[T]bool, isect)
	uniquesl := make(map[T]bool, diff)
	uniquesr := make(map[T]bool, diff)

	for _, v := range l {
		mapl[v] = true
	}
	for _, v := range r {
		mapr[v] = true
	}
	for v := range mapl {
		if _, found := mapr[v]; found {
			commons[v] = true
		} else {
			uniquesl[v] = true
		}
	}
	for v := range mapr {
		if _, found := mapl[v]; !found {
			uniquesr[v] = true
		}
	}
	return maps.Keys(commons), maps.Keys(uniquesl), maps.Keys(uniquesr)
}

func distanceUnorderedArrays[T comparable](l, r []T) (eq bool, d float64) {
	isect, diffl, diffr := subsets(l, r)
	d = float64(len(isect)) / float64(len(isect)+len(diffl)+len(diffr))
	return d == 0.0, d
}

// common items are not necessarily to be in same indexes. cause it will skip any uncommon item between
func countOrderedCommonItems[T any](l, r []T, li, ri int) int {
	commons := 0
	for li+commons < len(l) && ri+commons < len(r) {
		if eq, _ := Distance(l[li+commons], r[ri+commons]); eq {
			commons++
		} else {
			cl := countOrderedCommonItems(l, r, li+commons+1, ri+commons)
			cr := countOrderedCommonItems(l, r, li+commons, ri+commons+1)
			if cl < cr {
				return commons + cl
			} else {
				return commons + cr
			}
		}
	}
	return commons
}

func distanceOrderedArrays[T any](l, r []T) (eq bool, d float64) {
	commons := countOrderedCommonItems(l, r, 0, 0)
	mx := max(len(l), len(r))
	return commons == mx, 1 - float64(commons)/float64(mx)
}
