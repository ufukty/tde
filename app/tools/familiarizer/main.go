package familiarizer

import "go/ast"

func FindTargetFuncCall(funcDecl *ast.FuncDecl, name string) (forStmt *ast.ForStmt, assignStmt *ast.AssignStmt) {
	var currentForStmt *ast.ForStmt

	ast.Inspect(funcDecl, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		inForSearch := forStmt == nil && currentForStmt != nil
		searchForStmt := forStmt == nil && currentForStmt == nil

		switch {
		case inForSearch:

		case searchForStmt:
		}
		if forStmt == nil && currentForStmt != nil {
			// if callExpr, ok := n.(*ast.CallExpr); ok {

			// }
		}
		return assignStmt == nil
	})

	return nil, nil
}

func main() {

}
