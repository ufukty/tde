package testing

import "tde/internal/evolution/evaluation/distance"

func (t *T) Assert(a, b any) bool {
	eq, d := distance.Distance(a, b)
	t.AssertionResults = append(t.AssertionResults, eq)
	t.AssertionErrorDistance = append(t.AssertionErrorDistance, d)
	return eq
}

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
