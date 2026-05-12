//go:build tde

package words

import "tde/pkg/testing"

func TDE_Reverse(t *testing.T) {
	testParameters := map[string]string{
		"Hello world": "dlrow olleH",
		"dlrow olleH": "Hello world",
	}

	for input, want := range testParameters {
		got := Reverse(input)
		if !t.Assert(got, want) {
			t.Fatalf("Reverse(%s) = %s (want: %s)", input, got, want)
		}
	}
}
