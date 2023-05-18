package pattern_constructor

import "go/ast"

// Slice of struct, field, interface, or any type
func SliceOf(expr ast.Expr) *ast.ArrayType {
	return &ast.ArrayType{
		// Lbrack: 0,
		// Len:    nil,
		Elt: expr,
	}
}
