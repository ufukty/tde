//go:build tde
// +build tde

package word_reverse

import (
	"tde/pkg/tde"
)

func TDE_WordReverse(e *tde.E) {
	testParameters := map[string]string{
		"Hello world":         "dlrow olleH",
		"dlrow olleH":         "Hello world",
		"The quick brown fox": "xof nworb kciuq ehT",
	}

	e.TestCandidate(func(candidate *tde.C) {
		candidateFunction := candidate.Function.(func(string) string)
		for input, want := range testParameters {
			output := candidateFunction(input)
			candidate.AssertEqual(output, want)
		}
	})

}
