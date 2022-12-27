//go:build tde
// +build tde

package word_reverse

import (
	"tde/pkg/tde"
)

func TDE_WordReverse(e *tde.E) {
	e.SetConfig(tde.Config{
		MaxCompute:           100,
		MaxMemory:            1000,
		MaxSize:              1000,
		MaxTime:              10,
		ComputeToMemoryRatio: 3 / 2,
	})
	
	testParameters := map[string]string{
		"Hello world":         "dlrow olleH",
		"dlrow olleH":         "Hello world",
		"The quick brown fox": "xof nworb kciuq ehT",
	}

	s := St{}
	for input, want := range testParameters {
		output := s.WordReverse(input)
		e.AssertEqual(output, want)
	}
}
