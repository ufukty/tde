package node_relations

import (
	"testing"
)

func Test_CompareIndicesTraces(t *testing.T) {
	insertionPoint := []int{0, 3, 5, 1, 0, 0, 4}

	testCases := []struct {
		CurrentTrace []int
		Expected     NodeRelationToInsertionPoint
	}{
		{[]int{}, ANCESTOR},
		{[]int{0}, ANCESTOR},
		{[]int{0, 0}, ON_PATH},
		{[]int{0, 1}, ON_PATH},
		{[]int{0, 1, 0}, WONT_ACCESS},
		{[]int{0, 2}, ON_PATH},
		{[]int{0, 3}, ANCESTOR},
		{[]int{0, 3, 0}, ON_PATH},
		{[]int{0, 3, 5}, ANCESTOR},
		{[]int{0, 3, 5, 1, 0, 0, 3, 0}, WONT_ACCESS},
		{[]int{0, 3, 5, 1, 0, 0, 4}, ITSELF},
		{[]int{0, 3, 5, 1, 0, 0, 4, 0}, WONT_ACCESS},
		{[]int{0, 3, 5, 1, 0, 0, 5}, WONT_ACCESS},
		{[]int{1}, WONT_ACCESS},
		{[]int{1, 0}, WONT_ACCESS},
	}

	for testNumber, testCase := range testCases {
		got := CompareIndicesTraces(testCase.CurrentTrace, insertionPoint)
		if got != testCase.Expected {
			t.Errorf("Failed; for test number: %d, test subject: %v, expected: %v, got: %v\n",
				testNumber, testCase.CurrentTrace, testCase.Expected, got)
		}
	}

}
