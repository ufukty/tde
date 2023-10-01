package utilities

import (
	"fmt"
	"go/token"
	"testing"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

func Test_SliceRemoveItem(t *testing.T) {
	examples := map[*struct {
		Slice  []int
		Remove int
	}][]int{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0}: {1, 2, 3, 4, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 1}: {0, 2, 3, 4, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 2}: {0, 1, 3, 4, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3}: {0, 1, 2, 4, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 4}: {0, 1, 2, 3, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5}: {0, 1, 2, 3, 4, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 6}: {0, 1, 2, 3, 4, 5, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 7}: {0, 1, 2, 3, 4, 5, 6, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 8}: {0, 1, 2, 3, 4, 5, 6, 7, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9}: {0, 1, 2, 3, 4, 5, 6, 7, 8},
	}

	for input, want := range examples {
		got := SliceExceptItem(input.Slice, input.Remove)
		if slices.Compare(want, got) != 0 {
			t.Error("validation")
		}
	}

	got := SliceExceptItems([]token.Token{token.ASSIGN, token.DEFINE}, []token.Token{token.DEFINE})
	if slices.Compare(got, []token.Token{token.ASSIGN}) != 0 {
		t.Error("validation")
	}

}
func Test_SliceRemoveItems(t *testing.T) {
	examples := map[*struct {
		Slice  []int
		Except []int
	}][]int{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{}}:                             {0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 2, 3, 4, 9}}:                {1, 5, 6, 7, 8},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}: {},
		{[]int{}, []int{}}: {},
	}

	for input, want := range examples {
		got := SliceExceptItems(input.Slice, input.Except)
		if slices.Compare(want, got) != 0 {
			t.Error("validation")
		}
	}

	got := SliceExceptItems([]token.Token{token.ASSIGN, token.DEFINE}, []token.Token{token.DEFINE})
	if slices.Compare(got, []token.Token{token.ASSIGN}) != 0 {
		t.Error("validation")
	}

}

func TestRange(t *testing.T) {
	type TestCase struct {
		InputStart int
		InputStop  int
		InputStep  int
		Want       []int
	}
	var testCases = []TestCase{
		// 1 argument
		{-1, 0, -1, []int{}},
		{-1, 2, -1, []int{0, 1}},
		{-1, 4, -1, []int{0, 1, 2, 3}},
		// 2 argument
		{0, 0, -1, []int{}},
		{2, 2, -1, []int{}},
		{0, 2, -1, []int{0, 1}},
		{0, 20, -1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}},
		// 3 argument
		{0, 0, 1, []int{}},
		{0, 0, 2, []int{}},
		{0, 0, 3, []int{}},
		{2, 2, 3, []int{}},
		{0, 20, 1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}},
		{0, 20, 2, []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}},
		{1, 20, 3, []int{1, 4, 7, 10, 13, 16, 19}},
	}

	var got []int
	for i, testCase := range testCases {
		if i < 3 {
			got = Range(testCase.InputStop)
		} else if i < 7 {
			got = Range(testCase.InputStart, testCase.InputStop)
		} else {
			got = Range(testCase.InputStart, testCase.InputStop, testCase.InputStep)
		}

		if slices.Compare(testCase.Want, got) != 0 {
			t.Error("validation. test index:", i, "case:", testCase, "got:", got)
		}
	}
}

func Test_DivideIntoBuckets(t *testing.T) {
	checkBucketSize := func(items []int, bucket []int, numberOfBuckets int) error {
		maxExpected := len(items)/numberOfBuckets + 1
		minExpected := len(items) / numberOfBuckets

		if len(bucket) > maxExpected || len(bucket) < minExpected {
			return fmt.Errorf("bucket size out of range, expected %d <= size <= %d, got %d", minExpected, maxExpected, len(bucket))
		}
		return nil
	}

	checkAllItemsPlaced := func(items []int, buckets [][]int) error {
		merged := []int{}
		for _, bucket := range buckets {
			merged = append(merged, bucket...)
		}

		for _, item := range items {
			if slices.Index(merged, item) == -1 {
				return fmt.Errorf("None of the buckets has the item: %v", item)
			}
		}

		return nil
	}

	checkItems := func(items []int, bucket []int) error {
		for _, c := range bucket {
			found := false
			for _, cc := range items {
				if c == cc {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("subject %v not found in original bucket", c)
			}
		}
		return nil
	}

	items := Range(10001)

	for numberOfBuckets := 2; numberOfBuckets < 10; numberOfBuckets++ {
		buckets := DivideIntoBuckets(items, numberOfBuckets)

		if len(buckets) != numberOfBuckets {
			t.Errorf("expected %d buckets, got %d", numberOfBuckets, len(buckets))
		}

		if err := checkAllItemsPlaced(items, buckets); err != nil {
			t.Error(errors.Wrap(err, fmt.Sprintf("numberOfBuckets: %d", numberOfBuckets)))
		}

		for _, subBucket := range buckets {
			if err := checkBucketSize(items, subBucket, numberOfBuckets); err != nil {
				t.Error(errors.Wrap(err, fmt.Sprintf("numberOfBuckets: %d", numberOfBuckets)))
			}
			if err := checkItems(items, subBucket); err != nil {
				t.Error(errors.Wrap(err, fmt.Sprintf("numberOfBuckets: %d", numberOfBuckets)))
			}
		}

	}
}
