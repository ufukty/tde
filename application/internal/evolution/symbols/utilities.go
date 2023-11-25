package symbols

import (
	"fmt"
	"go/ast"
)

func findFuncTypeParent(r ast.Node, ft *ast.FuncType) *ast.FuncDecl {
	var fd, cfd *ast.FuncDecl
	ast.Inspect(r, func(n ast.Node) bool {
		if n == r || n == nil || fd != nil {
			return true
		}

		switch n := n.(type) {
		case *ast.Package, *ast.File:
			fmt.Println("case *ast.Package, *ast.File:")
			return true

		case *ast.GenDecl:
			fmt.Println("case *ast.GenDecl:")
			return false

		case *ast.FuncDecl:
			fmt.Println("case *ast.FuncDecl:", n.Name)
			cfd = n
			return true

		case *ast.FuncType:
			fmt.Println("case *ast.FuncType:")
			if ft == n {
				fd = cfd
			}
			return false
		}

		return false
	})

	return fd
}
