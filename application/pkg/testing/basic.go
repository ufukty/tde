package testing

import (
	"log"
	"math"
	"os"
	"reflect"
)

func convert[T any](a, b any) (T, T, bool) {
	if a, ok := a.(T); ok {
		if b, ok := b.(T); ok {
			return a, b, true
		}
	}
	return *new(T), *new(T), false
}

func distance(a, b any) float64 {
	if a, b, ok := convert[int](a, b); ok {
		return distanceInt(a, b)

	} else if a, b, ok := convert[int8](a, b); ok {
		return distanceInt(a, b)

	} else if a, b, ok := convert[int16](a, b); ok {
		return distanceInt(a, b)

	} else if a, b, ok := convert[int32](a, b); ok {
		return distanceInt(a, b)

	} else if a, b, ok := convert[int64](a, b); ok {
		return distanceInt(a, b)

	} else if a, b, ok := convert[uint8](a, b); ok {
		return distanceInt(a, b)

	} else if a, b, ok := convert[uint16](a, b); ok {
		return distanceInt(a, b)

	} else if a, b, ok := convert[uint32](a, b); ok {
		return distanceInt(a, b)

	} else if a, b, ok := convert[uint64](a, b); ok {
		return distanceInt(a, b)

	} else if a, b, ok := convert[float32](a, b); ok {
		return math.Abs(float64(a) - float64(b))

	} else if a, b, ok := convert[float64](a, b); ok {
		return math.Abs(a - b)

	} else if a, b, ok := convert[bool](a, b); ok {
		return distanceBool(a, b)

	} else if a, b, ok := convert[string](a, b); ok {
		return distanceString(a, b)

	} else if a, b, ok := convert[error](a, b); ok {
		return distanceString(a.Error(), b.Error())

	} else {
		log.Fatalf("Can not calculate the distance for types %q and %q", reflect.TypeOf(a), reflect.TypeOf(b))
		os.Exit(1)
		return -1.0
	}

}

// FIXME:
func distanceArrays[T any](a, b []T) float64 {
	d_m := 0.0
	if len(a) != len(b) {
		return 0.0 // FIXME:
	}
	for i := 0; i < len(a); i++ {
		if d_i := distance(a[i], b[i]); d_i != 0.0 {
			d_m += d_i
		}
	}
	return 0
}

func (t *T) Assert(a, b any) bool {
	comparison := a == b
	t.AssertionResults = append(t.AssertionResults, comparison)

	var d float64
	if a, b, ok := convert[[]int](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]int8](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]int16](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]int32](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]int64](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]uint8](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]uint16](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]uint32](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]uint64](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]float32](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]float64](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]bool](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]string](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]error](a, b); ok {
		d = distanceArrays(a, b)

	} else if a, b, ok := convert[[]any](a, b); ok {
		d = distanceArrays(a, b)

	} else {
		d = distance(a, b)
	}

	t.AssertionErrorDistance = append(t.AssertionErrorDistance, d)
	return comparison
}

// TODO:
// func assertArrays[C comparable](a, b []C) int {
// 	errorRate := 0.0
// 	if len(a) != len(b) {
// 		return
// 	}
// 	for i := 0; i < len(a); i++ {
// 		if a[i] != b[i] {
// 			return 1
// 		}
// 	}
// 	return 0
// }

// func (t *T) AssertArrays(left, right []any) {
// 	abs(len(left) - len(right))
// 	for i := 0; i < len(left); i++ {

// 	}

// 	}
// 	assertArrays(left, right)
// }

// func (t *T) AssertStructs(a, b any) {

// }

// func (tc *C) AssertNotEqual(output, notWant any) {
// 	result := output != notWant
// 	tc.AssertionResults = append(tc.AssertionResults, result)

// 	errorDistance := 0.0
// 	switch wantCasted := notWant.(type) {
// 	case string:
// 		if outputCasted, ok := output.(string); ok {
// 			errorDistance = StringDistance(wantCasted, outputCasted)
// 		} else {
// 			panic("<output> and <want> should be same type and one of string, int, float64")
// 		}
// 	case int:
// 		if outputCasted, ok := output.(int); ok {
// 			errorDistance = IntegerDistance(wantCasted, outputCasted)
// 		} else {
// 			panic("<output> and <want> should be same type and one of string, int, float64")
// 		}
// 	case float64:
// 		if outputCasted, ok := output.(float64); ok {
// 			errorDistance = FloatDistance(wantCasted, outputCasted)
// 		} else {
// 			panic("<output> and <want> should be same type and one of string, int, float64")
// 		}
// 	}
// 	tc.AssertionErrorDistance = append(tc.AssertionErrorDistance, errorDistance)
// }
