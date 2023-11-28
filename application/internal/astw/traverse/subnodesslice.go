package traverse

import (
	"go/ast"
	"tde/internal/astw/types"
)

// adds placeholders for appendable spots
func subnodesForSliceField(n *Node, c constraints) (subnodes []*Node) {
	if n.IsNil {
		return
	}

	switch sl := n.ref.Get().(type) {
	case []*ast.Comment:

		for i, item := range sl {
			subnodes = append(subnodes, subnode(n, newSliceItemRef(&sl, i), item == nil, types.Comment, 0))
		}

	case []*ast.CommentGroup:
		for i, item := range sl {
			subnodes = append(subnodes, subnode(n, newSliceItemRef(&sl, i), item == nil, types.CommentGroup, 0))
		}

	case []*ast.ImportSpec:
		for i, item := range sl {
			subnodes = append(subnodes, subnode(n, newSliceItemRef(&sl, i), item == nil, types.ImportSpec, 0))
		}

	case []*ast.Ident:
		for i, item := range sl {
			subnodes = append(subnodes, subnode(n, newSliceItemRef(&sl, i), item == nil, types.Ident, 0))
		}

	case []*ast.Field:
		for i, item := range sl {
			subnodes = append(subnodes, subnode(n, newSliceItemRef(&sl, i), item == nil, types.Field, 0))
		}

	case []ast.Stmt:
		for i, item := range sl {
			subnodes = append(subnodes, subnode(n, newSliceItemRef(&sl, i), item == nil, types.TypeFor(item), 0))
		}

	case []ast.Decl:
		for i, item := range sl {
			subnodes = append(subnodes, subnode(n, newSliceItemRef(&sl, i), item == nil, types.TypeFor(item), 0))
		}

	case []ast.Spec:
		for i, item := range sl {
			subnodes = append(subnodes, subnode(n, newSliceItemRef(&sl, i), item == nil, types.TypeFor(item), 0))
		}

	case []ast.Expr:
		for i, item := range sl {
			subnodes = append(subnodes, subnode(n, newSliceItemRef(&sl, i), item == nil, types.TypeFor(item), 0))
		}
	}

	return
}
