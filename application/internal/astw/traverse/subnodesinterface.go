package traverse

import "go/ast"

func subnodesForIface(n *Node, c constraints) []*Node {
	if n.IsNil {
		return []*Node{}
	}

	switch n.ref.Get().(type) {
	case *ast.Expr:
		return subnodesForConcrete(n, c)
	case *ast.Stmt:
		return subnodesForConcrete(n, c)
	case *ast.Decl:
		return subnodesForConcrete(n, c)
	case *ast.Spec:
		return subnodesForConcrete(n, c)
	}
	return []*Node{}
}
