package symbols

import (
	"go/ast"
	"go/types"
)

type Context struct {
	Symbols []*Symbol
	ByType  map[types.Type][]*Symbol
}

func (c *Context) CreateSymbol(ident, belongs *ast.Ident, typ types.Type) {
	s := &Symbol{
		Belongs: belongs,
		Ident:   ident,
	}
	c.Symbols = append(c.Symbols, s)
	c.ByType[typ] = append(c.ByType[typ], s)
}

func (c *Context) ReviewScopeContent(sc *ScopeContent, of *types.Package) {
	for _, typ := range sc.TypeNamesBasic {
		c.CreateSymbol(ast.NewIdent(typ.Name()), ast.NewIdent(of.Name()), typ.Type())
	}
	for _, typ := range sc.TypeNamesNamed {
		c.CreateSymbol(ast.NewIdent(typ.Name()), ast.NewIdent(of.Name()), typ.Type())
	}
	for _, typ := range sc.TypeNamesInterface {
		c.CreateSymbol(ast.NewIdent(typ.Name()), ast.NewIdent(of.Name()), typ.Type())
	}
	for _, typ := range sc.Consts {
		c.CreateSymbol(ast.NewIdent(typ.Name()), ast.NewIdent(of.Name()), typ.Type())
	}
	for _, typ := range sc.Funcs {
		c.CreateSymbol(ast.NewIdent(typ.Name()), ast.NewIdent(of.Name()), typ.Type())
	}
	for _, typ := range sc.PkgNames {
		c.CreateSymbol(ast.NewIdent(typ.Name()), ast.NewIdent(of.Name()), typ.Type())
	}
	for _, typ := range sc.Vars {
		c.CreateSymbol(ast.NewIdent(typ.Name()), ast.NewIdent(of.Name()), typ.Type())
	}
	for _, typ := range sc.Labels {
		c.CreateSymbol(ast.NewIdent(typ.Name()), ast.NewIdent(of.Name()), typ.Type())
	}
	for _, typ := range sc.Builtins {
		c.CreateSymbol(ast.NewIdent(typ.Name()), ast.NewIdent(of.Name()), typ.Type())
	}
	for _, typ := range sc.Nils {
		c.CreateSymbol(ast.NewIdent(typ.Name()), ast.NewIdent(of.Name()), typ.Type())
	}
}
