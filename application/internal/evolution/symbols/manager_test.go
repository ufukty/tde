package symbols

import (
	"fmt"
	"go/ast"
	"go/types"
	"runtime"
	"strings"
	"tde/internal/evolution/evaluation/discovery"
	"tde/internal/evolution/evaluation/list"
	"tde/internal/evolution/symbols/canonicalize"
	"tde/internal/utilities/functional"
	"testing"
)

func Test_SymbolsInspector(t *testing.T) {
	type tcase struct {
		name        string
		allowedpkgs []string
	}

	var tcs = []tcase{
		{
			name:        "fmt",
			allowedpkgs: []string{"fmt"},
		},
		{
			name:        "reflect",
			allowedpkgs: []string{"reflect"},
		},
		{
			name:        "evolution",
			allowedpkgs: []string{"tde/internal/evolution/symbols/testdata/evolution", "fmt", "log", "slices", "maps"},
		},
		{
			name:        "words",
			allowedpkgs: []string{"tde/internal/evolution/symbols/testdata/words", "fmt", "math", "strconv"},
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
			si, err := NewSymbolsInspector(allowedpkgs)
			if err != nil {
				t.Fatal(fmt.Errorf("prep 2: %w", err))
			}

			symbols := si.SymbolsAssignableTo(types.Typ[types.String])
			for pkgid, symbols := range symbols {
				fmt.Printf("%s:\n\t[%s]\n", pkgid, strings.Join(functional.Map(symbols, func(i int, v *ast.Ident) string { return v.Name }), ", "))
			}
		})
	}
}
