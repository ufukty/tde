package testing

import (
	"math"

	"golang.org/x/exp/constraints"
)

func distanceInt[T constraints.Signed | constraints.Unsigned](a, b T) float64 {
	return math.Abs(float64(a) - float64(b))
}

func stringFloat(a, b float64) float64 {
	return math.Abs(a - b)
}

//   - O(n) string comparison
//   - fast, but not best [like levenshtein which is O(n^2)]
//   - logic: similarity of two strings is based on the portion of characters
//     are in same position at two strings
func distanceString(a, b string) float64 {
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
	return distance
}

func distanceBool(a, b bool) float64 {
	if a == b {
		return 0.0
	}
	return 1.0
}
