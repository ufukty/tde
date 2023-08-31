//go:build tde

package words

import "tde/pkg/testing"

func Test_WordReverse(t *testing.T) {
	testParameters := map[string]string{
		"Hello world": "dlrow olleH",
		"dlrow olleH": "Hello world",
	}

	for input, want := range testParameters {
		got := WordReverse(input)
		if !t.Assert(got, want) {
			t.Fatalf("WordReverse(%s) = %s (want: %s)", input, got, want)
		}
	}
}
