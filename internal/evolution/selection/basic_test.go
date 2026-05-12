package selection

import (
	"fmt"
	"testing"
)

func Test_Random(t *testing.T) {
	var datasets = [][]float64{
		{1.0},
		{1.0, 1.0},
		{1.0, 1.0, 1.0},
	}

	for _, dataset := range datasets {
		subjects := subjectsForDataset(dataset)
		for pick := 0; pick <= len(dataset); pick++ {
			t.Run(fmt.Sprintf("%v>%d", dataset, pick), func(t *testing.T) {
				fmt.Printf("%v>%d\n", dataset, pick)
				selection := Random(subjects, pick)
				if len(selection) != pick {
					t.Fatal(fmt.Errorf("assert, selection length: expected %d, got %d items", pick, len(selection)))
				}
			})
		}
	}
}
