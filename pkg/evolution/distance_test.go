package evolution

import (
	"math"
	"testing"
)

func areFloatsEqual(a, b float64) bool {
	return math.Floor(100000*a) == math.Floor(100000*b)
}

func Test_StringDistance(t *testing.T) {

	testCases := []struct {
		a    string
		b    string
		dist float64
	}{
		{"123", "123", 0.0},
		{"lll", "lll", 0.0},
		{"öçşiğü", "öçşiğü", 0.0},
		{"???", "???", 0.0},
		{"123", "123", 0.0},
		{"12", "12x", 1 - 2.0/3.0},
		{"123x", "123y", 1 - 3.0/4.0},
		{"123xy", "123yx", 1 - 3.0/5.0},
		{"123", "xyz", 1.0},
	}

	for _, testCase := range testCases {
		if output := StringDistance(testCase.a, testCase.b); !areFloatsEqual(output, testCase.dist) {
			t.Errorf("StringDistance('%s', '%s') is expected to be '%f', got '%f'", testCase.a, testCase.b, testCase.dist, output)
		}
	}
}
