package symbols

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"
	"strings"
	"tde/internal/astw/traced"
	"tde/internal/utilities/mapw"
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

func crumb(in string) string {
	return fmt.Sprintf("`%s`", in)
}

func findMeaningfulPathToScope(info *types.Info, r ast.Node, s *types.Scope) string {
	sn, ok := mapw.FindKey(info.Scopes, s)
	if !ok {
		return ""
	}
	path := ""
	for i, n := range traced.Parents(r, sn) {
		if i != 0 {
			path += "/"
		}
		switch a := n.(type) {
		case *ast.File:
			path += a.Name.Name
		case *ast.TypeSpec:
			path += a.Name.Name
		case *ast.FuncDecl:
			path += a.Name.Name
		default:
			path += crumb(strings.Split(reflect.TypeOf(a).String(), ".")[1])
		}
	}
	return path
}
