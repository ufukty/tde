//go:build tde
// +build tde

package wordreverse

import (
	"GoGP/testing/evolution"
)

func TDE_WordReverse(e *evolution.E) {
	testParameters := map[string]string{
		"Hello world":         "dlrow olleH",
		"dlrow olleH":         "Hello world",
		"The quick brown fox": "xof nworb kciuq ehT",
	}

	e.TestCandidate(func(candidate *evolution.C) {
		candidateFunction := candidate.Function.(func (string) string)
		for input, want := range testParameters {
			output := candidateFunction(input)
			candidate.AssertEqual(output, want)
		}
	})

}
