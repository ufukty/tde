//go:build tdgp
// +build tdgp

package wordreverse

import (
	"GoGP/testing"
)

func TDGP_WordReverse(e *testing.Evolution) {
	testParameters := map[string]string{
		"Hello world":         "dlrow olleH",
		"dlrow olleH":         "Hello world",
		"The quick brown fox": "xof nworb kciuq ehT",
	}

	e.GetCandidate(func(c *testing.TestCandidate) {
		for input, want := range testParameters {
			output := c.CandidateFunction(input)
			c.AssertEqual(output, want)
		}

	})

}
