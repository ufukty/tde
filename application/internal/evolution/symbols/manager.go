package symbols

import (
	"fmt"
	"go/ast"
	"go/types"
	"tde/internal/utilities/strw"

	"golang.org/x/exp/maps"
	"golang.org/x/tools/go/packages"
)

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

// Multiple use
type Manager struct {
	main    *packages.Package
	imports []*packages.Package
}

func NewSymbolsManager(path string, allowedpackages []string) (*Manager, error) {
	cfg := &packages.Config{
		Mode: packages.NeedDeps |
			packages.NeedImports |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedFiles,
		BuildFlags: []string{"-tags", "tde"},
	}
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		return nil, fmt.Errorf("loading the package: %w", err)
	}
	if err = packageerrors(pkgs); err != nil {
		return nil, fmt.Errorf("checking package errors:\n%s", strw.IndentLines(err.Error(), 3))
	}

	var pkg *packages.Package
	for _, p := range pkgs {
		if p.ID == path {
			pkg = p
		}
	}
	if pkg == nil {
		return nil, fmt.Errorf("could not load the package %q", path)
	}

	return &Manager{main: pkg, imports: maps.Values(pkg.Imports)}, nil
}

func (sm *Manager) BinaryIdents() ([]*ast.Ident, map[string][]*ast.Ident) {
	inpackage := []*ast.Ident{}

	if sm.main.TypesInfo != nil && sm.main.TypesInfo.Defs != nil {
		fmt.Println(len(sm.main.TypesInfo.Defs))
		for i, o := range sm.main.TypesInfo.Defs {
			if o == nil {
				continue
			}
			t := o.Type()
			if t != nil && types.AssignableTo(t, types.Typ[types.Bool]) {
				inpackage = append(inpackage, i)
			}
		}
	}

	imported := map[string][]*ast.Ident{}
	for _, pkg := range sm.imports {
		fmt.Println(len(pkg.TypesInfo.Defs))
		if pkg.TypesInfo == nil || pkg.TypesInfo.Defs == nil {
			continue
		}
		for i, o := range pkg.TypesInfo.Defs {
			if o == nil || !o.Exported() {
				continue
			}
			t := o.Type()

			if t != nil && types.AssignableTo(t, types.Typ[types.Bool]) {
				if _, ok := imported[pkg.ID]; !ok {
					imported[pkg.ID] = []*ast.Ident{}
				}
				imported[pkg.ID] = append(imported[pkg.ID], i)
			}
		}
	}
	return inpackage, imported
}
