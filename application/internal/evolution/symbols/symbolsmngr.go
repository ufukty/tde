package symbols

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"slices"
	"tde/internal/utilities"

	"golang.org/x/exp/maps"
)

type SymbolsMngr struct {
	pkg     *types.Package
	info    *types.Info
	fset    *token.FileSet
	Context *Context
	scopes  map[ast.Node]*types.Scope
}

func (sm *SymbolsMngr) analyze(path string) error {
	pkgs, err := parser.ParseDir(sm.fset, path, nil, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("parseDir: %w", err)
	}
	if len(pkgs) == 0 {
		return fmt.Errorf("there is no package in the dir")
	}
	if slices.Contains(maps.Keys(pkgs), "test") && len(pkgs) != 2 {
		return fmt.Errorf("too many packages (%d) in the dir", len(pkgs))
	}
	pkg := maps.Values(pkgs)[0]
	files := maps.Values(pkg.Files)
	conf := types.Config{Importer: importer.Default()}
	sm.info = &types.Info{
		Defs:       map[*ast.Ident]types.Object{},
		Implicits:  map[ast.Node]types.Object{},
		InitOrder:  []*types.Initializer{},
		Instances:  map[*ast.Ident]types.Instance{},
		Scopes:     map[ast.Node]*types.Scope{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
		Types:      map[ast.Expr]types.TypeAndValue{},
		Uses:       map[*ast.Ident]types.Object{},
	}
	sm.pkg, err = conf.Check("main", sm.fset, files, sm.info)
	if err != nil {
		return fmt.Errorf("types.Config.Check: %w", err)
	}
	return nil
}

func (sm *SymbolsMngr) prepareContext() error {
	pkgscope := sm.pkg.Scope()
	for _, name := range pkgscope.Names() {
		obj := pkgscope.Lookup(name)
		idt, ok := utilities.MapSearchKey(sm.info.Defs, obj)
		if !ok {
			return fmt.Errorf("info.Uses doesn't have %q", name)
		}
		sm.Context.Package.Append(idt, obj.Type())
	}

	return nil
}

func NewSymbolsManager(path string) (*SymbolsMngr, error) {
	sm := &SymbolsMngr{
		fset:   token.NewFileSet(),
		scopes: map[ast.Node]*types.Scope{},
		Context: &Context{
			Imported:   map[*ast.Ident]IdentsAndTypes{},
			Importable: map[*ast.Ident]IdentsAndTypes{},
			Package:    map[*ast.Ident]types.Type{},
			File:       map[*ast.Ident]types.Type{},
			Function:   map[*ast.Ident]types.Type{},
			InFunction: map[*ast.Ident]types.Type{},
		},
	}
	if err := sm.analyze(path); err != nil {
		return nil, fmt.Errorf("analyze: %w", err)
	}
	if err := sm.prepareContext(); err != nil {
		return nil, fmt.Errorf("prepareContext: %w", err)
	}
	for idt, scp := range sm.info.Scopes {
		sm.scopes[idt] = scp
	}
	return sm, nil
}
