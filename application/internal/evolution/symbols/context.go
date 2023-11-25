package symbols

import (
	"go/ast"
	"go/types"
)

//go:generate stringer -type Type

type Type int

const (
	TypeDefiningExpr = Type(iota)
	Boolean          // or boolean returning expression
	Int
	String
	Char
)

type IdentsAndTypes map[*ast.Ident]types.Type

func (mp IdentsAndTypes) Append(idt *ast.Ident, typ types.Type) {
	mp[idt] = typ
}

type Symbol struct {
	Belongs, Ident *ast.Ident
}

func (s Symbol) Node() ast.Node {
	if s.Belongs == nil {
		return s.Ident
	}
	return &ast.SelectorExpr{
		Sel: s.Belongs,
		X:   s.Ident,
	}
}

type Context struct {
	Imported   map[*ast.Ident]IdentsAndTypes
	Importable map[*ast.Ident]IdentsAndTypes
	Package    IdentsAndTypes
	File       IdentsAndTypes
	Function   IdentsAndTypes
	InFunction IdentsAndTypes // ?

	ByType map[Type][]Symbol
}
