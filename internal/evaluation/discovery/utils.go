package discovery

import (
	"go/ast"
	"go/token"
	"slices"
	"strings"
)

func ExpectedTargetFileAndFuncNameFor(testfile, testFunction string) (targetfile, targetFunction string) {
	targetfile = strings.TrimSuffix(testfile, "_tde.go") + ".go"
	targetFunction = strings.TrimPrefix(testFunction, "TDE_")
	return
}

func FunctionCalls(funcDecl *ast.FuncDecl) []*ast.CallExpr {
	list := []*ast.CallExpr{}
	ast.Inspect(funcDecl, func(node ast.Node) bool {
		if node, ok := node.(*ast.CallExpr); ok {
			if !slices.Contains(list, node) {
				list = append(list, node)
			}
		}
		return true
	})
	return list
}

func line(fset *token.FileSet, posToken token.Pos) int {
	return fset.Position(posToken).Line - 1
}
