package testing

func convert[T any](a, b any) (T, T, bool) {
	if a, ok := a.(T); ok {
		if b, ok := b.(T); ok {
			return a, b, true
		}
	}
	return *new(T), *new(T), false
}

func (t *T) Assert(a, b any) bool {
	eq, d := distance(a, b)
	t.AssertionResults = append(t.AssertionResults, eq)
	t.AssertionErrorDistance = append(t.AssertionErrorDistance, d)
	return eq
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
