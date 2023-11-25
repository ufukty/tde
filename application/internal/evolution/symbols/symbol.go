package symbols

import (
	"fmt"
	"go/ast"
)

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

func (s Symbol) String() string {
	if s.Belongs == nil {
		return s.Ident.Name
	}
	return fmt.Sprintf("%s.%s", s.Belongs, s.Ident)
}
