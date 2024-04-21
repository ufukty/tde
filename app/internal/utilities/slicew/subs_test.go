package slicew

import (
	"fmt"
	"slices"
	"testing"

	"github.com/pkg/errors"
)

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
		buckets := Subs(items, numberOfBuckets)

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
