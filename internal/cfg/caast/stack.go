package caast

import (
	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

type Stack struct {
	vars      [][]*ast.Ident
	constants [][]*ast.Ident
}

func (s *Stack) Recurse() {
	s.vars = append(s.vars, []*ast.Ident{})
	s.constants = append(s.constants, []*ast.Ident{})
}

func (s *Stack) Return() {
	s.vars = s.vars[:len(s.vars)-1]
	s.constants = s.constants[:len(s.vars)-1]
}

func (s *Stack) DeclareVariable(ident *ast.Ident) {
	i := len(s.vars) - 1
	s.vars[i] = append(s.vars[i], ident)
}

func (s *Stack) DeclareConstant(ident *ast.Ident) {
	i := len(s.vars) - 1
	s.constants[i] = append(s.constants[i], ident)
}

func (s *Stack) Fill(node ast.Node) {
	astutil.Apply(node, func(c *astutil.Cursor) bool {
		if c.Parent() == node {
			// recurse into variable/constant declarations
			switch c.Node().(type) {
			case *ast.DeclStmt:

			}

		}

		return c.Node() == node
	}, nil)
}

func (s *Stack) ListVars() {}

func IsInPath(ancestry []ast.Node, node ast.Node) bool {
	for _, a := range ancestry {
		if node == a {
			return true
		}
	}
	return false
}
