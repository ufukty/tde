package test_package

import (
	"fmt"
	"go/ast"
)

func FindFuncDecl(toppest ast.Node, name string) (*ast.FuncDecl, error) {
	found := []*ast.FuncDecl{}

	ast.Inspect(toppest, func(n ast.Node) bool {
		if n != nil {
			if functionDeclaration, ok := n.(*ast.FuncDecl); ok {
				if functionDeclaration.Name.Name == name {
					found = append(found, functionDeclaration)
					return false
				}
			}
		}
		return true
	})

	if len(found) == 0 {
		return nil, fmt.Errorf("could not find function '%s'", name)
	} else if len(found) > 1 {
		return nil, fmt.Errorf("more than one function definition with the name  '%s'", name)
	} else {
		return found[0], nil
	}
}
