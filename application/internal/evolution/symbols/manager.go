package symbols

import (
	"fmt"
	"go/ast"
	"go/types"
	"slices"
	"tde/internal/utilities/mapw"
	"tde/internal/utilities/strw"

	"golang.org/x/tools/go/packages"
)

// Multiple use
type Inspector struct {
	pkgs []*packages.Package
}

func packageerrors(pkgs []*packages.Package) error {
	err := ""
	packages.Visit(pkgs, func(p *packages.Package) bool {
		if len(p.Errors) > 0 {
			err += fmt.Sprintf("found %d errors in %q:\n", len(p.Errors), p.ID)
			for _, e := range p.Errors {
				err += fmt.Sprintf("   %s\n", e.Error())
			}
		}
		return true
	}, nil)
	if err == "" {
		return nil
	}
	return fmt.Errorf(err)
}

// pkgid is the import path of the current package
// allowedpkgs consists by canonicalized directory paths
func NewSymbolsInspector(allowedpkgs []string) (*Inspector, error) {
	cfg := &packages.Config{
		Mode: packages.NeedDeps |
			packages.NeedImports |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedFiles,
		BuildFlags: []string{"-tags", "tde"},
	}
	pkgs, err := packages.Load(cfg, allowedpkgs...)
	if err != nil {
		return nil, fmt.Errorf("loading the package: %w", err)
	}
	if err = packageerrors(pkgs); err != nil {
		return nil, fmt.Errorf("checking package errors:\n%s", strw.IndentLines(err.Error(), 3))
	}
	return &Inspector{pkgs: pkgs}, nil
}

func appenduniq[T comparable](s []T, v T) []T {
	if slices.Index(s, v) == -1 {
		return append(s, v)
	}
	return s
}

// helper of *Manager.AssignableTo. inspects only one scope
func symbolsFromScope(s *types.Scope, defs map[*ast.Ident]types.Object, t types.Type) []*ast.Ident {
	idents := mapw.Reverse(defs)
	symbols := []*ast.Ident{}
	for _, elem := range s.Names() {
		if o := s.Lookup(elem); o != nil && o.Exported() {
			if v := o.Type(); v != nil {
				if i, ok := idents[o]; ok {
					if containsSomethingAssignableTo(v, t) {
						symbols = appenduniq(symbols, i)
					}
				}
			}
		}
	}
	return symbols
}

// returns the list of all symbols defined in current/imported/importable packages that either
// itself or its a field, element or result can be assignable to a variable in type of "t".
func (si *Inspector) SymbolsAssignableTo(t types.Type) map[string][]*ast.Ident {
	symbols := map[string][]*ast.Ident{}
	for _, pkg := range si.pkgs {
		if pkg.TypesInfo != nil && pkg.TypesInfo.Defs != nil {
			if scopesymbols := symbolsFromScope(pkg.Types.Scope(), pkg.TypesInfo.Defs, t); len(scopesymbols) > 0 {
				symbols[pkg.ID] = scopesymbols
			}
		}
	}
	return symbols
}
