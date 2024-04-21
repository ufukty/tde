package symbols

import (
	"fmt"
	"go/ast"
	"go/types"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"tde/internal/evolution/evaluation/discovery"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/symbols/canonicalize"
	"tde/internal/utilities/functional"
	"testing"
)

func Test_SymbolsInspector(t *testing.T) {
	type tcase struct {
		name            string
		allowedpkgs     []string
		expectedSymbols map[string][]string
	}

	var tcs = []tcase{
		{
			name:        "fmt",
			allowedpkgs: []string{"fmt"},
			expectedSymbols: map[string][]string{
				"fmt": {"Sprint", "badIndexString"},
			},
		},
		{
			name:        "reflect",
			allowedpkgs: []string{"reflect"},
			expectedSymbols: map[string][]string{
				"reflect": {"funcName", "funcStr"},
			},
		},
		{
			name:        "evolution",
			allowedpkgs: []string{"tde/internal/evolution/symbols/testdata/evolution", "fmt", "log", "slices", "maps"},
			expectedSymbols: map[string][]string{
				"fmt": {"FormatString", "Sprint", "Sprintf", "Sprintln"},
				"log": {"Prefix"},
			},
		},
		{
			name:        "words",
			allowedpkgs: []string{"tde/internal/evolution/symbols/testdata/words", "fmt", "math", "strconv"},
			expectedSymbols: map[string][]string{
				"tde/internal/evolution/symbols/testdata/words": {"Indent", "Reverse"},
			},
		},
	}

	goroot := runtime.GOROOT()
	mod, err := discovery.ModuleRoot()
	if err != nil {
		t.Fatal(fmt.Errorf("prep moduleroot: %w", err))
	}

	pkgs, err := list.ListAllPackages(mod)
	if err != nil {
		t.Fatal(fmt.Errorf("prep listallpackages: %w", err))
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			allowedpkgs := canonicalize.CanonicalizePaths(goroot, mod, pkgs, tc.allowedpkgs)
			si, err := NewSymbolsInspector(filepath.Base(allowedpkgs[0]), allowedpkgs)
			if err != nil {
				t.Fatal(fmt.Errorf("prep 2: %w", err))
			}

			got := si.SymbolsAssignableTo(types.Typ[types.String])
			for pkgid, symbols := range got {
				fmt.Printf("%s:\n  [%s]\n", pkgid, strings.Join(functional.Map(symbols, func(i int, v *ast.Ident) string { return v.Name }), ", "))
			}
			if len(got) > len(allowedpkgs)+1 {
				t.Fatal(fmt.Errorf("assert 1: len(got)=%d > len(allowedpkgs)=%d + 1 (universe)", len(got), len(allowedpkgs)))
			}
			for pkgid, expectedPkgSymbols := range tc.expectedSymbols {
				if gotPkgSymbols, ok := got[pkgid]; ok {
					gotPkgSymbolsStrings := functional.Map(gotPkgSymbols, func(i int, v *ast.Ident) string { return v.Name })
					for _, expectedSymbol := range expectedPkgSymbols {
						if slices.Index(gotPkgSymbolsStrings, expectedSymbol) == -1 {
							t.Fatalf("assert 3: could not find %q in symbols for pkgid %q\n", expectedSymbol, pkgid)
						}
					}
				} else {
					t.Fatalf("assert 2: could not find the list of symbols for %q\n", pkgid)
				}
			}

		})
	}
}
