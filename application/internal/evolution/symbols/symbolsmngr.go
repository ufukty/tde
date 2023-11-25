package symbols

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"slices"

	"golang.org/x/exp/maps"
)

// Use it to get list of symbols defined at the package and in imports.
// Excludes the symbols defined inside a function (because the info is not available at initialization)
type SymbolsMngr struct {
	pkg  *types.Package
	ast  *ast.Package
	info *types.Info
	fset *token.FileSet
	// scopes  map[ast.Node]*types.Scope
	Context *Context
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

// func (sm *SymbolsMngr) prepareScopes() {
// 	for idt, scp := range sm.info.Scopes {
// 		sm.scopes[idt] = scp
// 	}
// }

func (sm *SymbolsMngr) prepareContext() error {
	// the "universe"
	sm.Context.ReviewScopeContent(NewScopeContent(types.Universe), nil)

	// the package
	sm.Context.ReviewScopeContent(NewScopeContent(sm.pkg.Scope()), nil)

	// imports
	for _, pkg := range sm.pkg.Imports() {
		sm.Context.ReviewScopeContent(NewScopeContent(pkg.Scope()), pkg)
	}

	return nil
}

func NewSymbolsManager(path string) (*SymbolsMngr, error) {
	sm := &SymbolsMngr{
		fset: token.NewFileSet(),
		// scopes: map[ast.Node]*types.Scope{},
		Context: &Context{
			Symbols: []*Symbol{},
			ByType:  map[types.Type][]*Symbol{},
		},
	}
	if err := sm.analyze(path); err != nil {
		return nil, fmt.Errorf("analyze: %w", err)
	}
	// sm.prepareScopes()
	if err := sm.prepareContext(); err != nil {
		return nil, fmt.Errorf("prepareContext: %w", err)
	}
	return sm, nil
}
