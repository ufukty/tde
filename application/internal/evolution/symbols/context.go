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
	c.ByType[typ] = append(c.ByType[typ])
}
