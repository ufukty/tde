package numerics

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBisect(t *testing.T) {

	values := []float64{
		0.0, 0.5, 1.2, 1.5, 1.9, 2.4, 3.5, 5.0, 5.1,
	}

	testCases := []struct {
		key             float64
		wantBisectLeft  int
		wantBisectRight int
	}{
		{-0.4, 0, 0},
		{0.0, 0, 1},
		{0.003275058479104, 1, 1},
		{0.051653628821504, 1, 1},
		{0.073821599203328, 1, 1},
		{0.077568278691840, 1, 1},
		{0.141161587802112, 1, 1},
		{0.4, 1, 1},
		{0.499220254720000, 1, 1},
		{0.5, 1, 2},
		{0.523633630642176, 2, 2},
		{0.6, 2, 2},
		{0.6, 2, 2},
		{0.6, 2, 2},
		{0.606093078953984, 2, 2},
		{0.7, 2, 2},
		{0.758019668475904, 2, 2},
		{0.8, 2, 2},
		{0.9, 2, 2},
		{0.90702324039680, 2, 2},
		{1.119519152013312, 2, 2},
		{1.180454810025984, 2, 2},
		{1.187266319843328, 2, 2},
		{1.2, 2, 3},
		{1.252245156888576, 3, 3},
		{1.322561264680960, 3, 3},
		{1.365033785196544, 3, 3},
		{1.499999, 3, 3},
		{1.5, 3, 4},
		{1.72455274479616, 4, 4},
		{1.724739844276224, 4, 4},
		{1.755768258789376, 4, 4},
		{1.758445954891776, 4, 4},
		{1.785894388334592, 4, 4},
		{1.8, 4, 4},
		{1.9, 4, 5},
		{1.908036088102912, 5, 5},
		{1.944757871804416, 5, 5},
		{2.042095179137024, 5, 5},
		{2.122710612180992, 5, 5},
		{2.3, 5, 5},
		{2.4, 5, 6},
		{2.418839786618880, 6, 6},
		{3.098357419474944, 6, 6},
		{3.4, 6, 6},
		{3.5, 6, 7},
		{3.531016306163712, 7, 7},
		{3.915128013586432, 7, 7},
		{4.061231172288512, 7, 7},
		{4.127605047427072, 7, 7},
		{4.348672397541376, 7, 7},
		{4.9, 7, 7},
		{5.0, 7, 8},
		{5.000001, 8, 8},
		{5.035247104327680, 8, 8},
		{5.06287226019840, 8, 8},
		{5.077430578233344, 8, 8},
		{5.1, 8, 9},
		{5.2, 9, 9},
		{5.209656402116608, 9, 9},
		{5.24883175931904, 9, 9},
		{5.441704914059264, 9, 9},
		{5.446366853595136, 9, 9},
		{25.472660618215424, 9, 9},
	}

	for _, testCase := range testCases {
		gotRight := BisectRight(values, testCase.key)
		gotLeft := BisectLeft(values, testCase.key)
		fmt.Printf("For %f; Left: (expected: %d, got: %d), Right: (expected: %d, got: %d)\n", testCase.key, testCase.wantBisectLeft, gotLeft, testCase.wantBisectRight, gotRight)
		if gotRight != testCase.wantBisectRight || gotLeft != testCase.wantBisectLeft {
			t.Errorf("validation")
		}
	}

}

func TestBinaryRangeSearchProbabilities(t *testing.T) {
	weights := []int64{4, 5, 2, 7, 8, 9, 6, 3, 1, 2}
	cumulativeWeights := Cumulate(weights) // [4,9,11,18,26,35,41,44,45,47]

	var (
		totalCumulation = cumulativeWeights[len(cumulativeWeights)-1]
		freq            = [10]int{}
		totalRun        = 5000
	)

	for i := 0; i < totalRun; i++ {
		key := int64(rand.Intn(int(totalCumulation)))
		index := BisectRight(cumulativeWeights, key)
		freq[index]++
	}

	for i := 0; i < 10; i++ {
		expectedProb := float64(weights[i]) / float64(totalCumulation)
		got := float64(freq[i]) / float64(totalRun)
		fmt.Printf("i=%d, weight=%d, expected prob=%f, got=%f\n", i, weights[i], expectedProb, got)
		if !(0.8*expectedProb < got && got < expectedProb*1.2) {
			t.Error("validation")
		}
	}

}
