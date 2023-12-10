package symbols

import (
	"fmt"
	"testing"
)

type tcase struct {
	name        string
	funcname    string
	path        string
	allowedpkgs []string
}

var tcs = []tcase{
	{"fmt", "writePadding",
		"fmt", []string{"math"}},
	{"reflect", "DeepEqual",
		"reflect", []string{}},
	{"./testdata/evolution", "WalkWithNils",
		"./testdata/evolution", []string{"fmt", "math"}},
	{"./testdata/words", "Reverse",
		"./testdata/words", []string{"fmt", "math"}},
}

func Test_SymbolsManager(t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			sm, err := NewSymbolsManager(tc.path, tc.allowedpkgs)
			if err != nil {
				t.Fatal(fmt.Errorf("prep: %w", err))
			}

			in, im := sm.BinaryIdents()
			for _, i := range in {
				fmt.Println(i)
			}
			for p, i := range im {
				fmt.Println(p, i)
			}
		})
	}
}
