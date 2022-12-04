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

	e.GetCandidate(func(c *evolution.TestCandidate) {
		for input, want := range testParameters {
			output := c.CandidateFunction(input)
			c.AssertEqual(output, want)
		}

	})

}
