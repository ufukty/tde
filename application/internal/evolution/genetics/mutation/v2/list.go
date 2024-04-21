package mutation

import (
	"go/ast"
)

func ListSubnodesInType[T any](n ast.Node) (subnodes []T) {
	ast.Inspect(n, func(m ast.Node) bool {
		if n == m {
			return true
		}
		if m, ok := m.(T); ok {
			subnodes = append(subnodes, m)
		}
		return false
	})
	return
}

func Demo() {
	fd := &ast.FuncDecl{}
	for _, n := range fd.Body.List {
		if subnodes := ListSubnodesInType[ast.AssignStmt](n); len(subnodes) > 0 {

		} else if subnodes := ListSubnodesInType[ast.ForStmt](n); len(subnodes) > 0 {

		} else if subnodes := ListSubnodesInType[ast.CompositeLit](n); len(subnodes) > 0 {

		} else if subnodes := ListSubnodesInType[ast.AssignStmt](n); len(subnodes) > 0 {

		} else if subnodes := ListSubnodesInType[ast.AssignStmt](n); len(subnodes) > 0 {

		} else if subnodes := ListSubnodesInType[ast.AssignStmt](n); len(subnodes) > 0 {

		} else if subnodes := ListSubnodesInType[ast.AssignStmt](n); len(subnodes) > 0 {

		} else if subnodes := ListSubnodesInType[ast.AssignStmt](n); len(subnodes) > 0 {

		} else if subnodes := ListSubnodesInType[ast.AssignStmt](n); len(subnodes) > 0 {

		} else {
			
		}

	}
}
