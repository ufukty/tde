package embedding

import (
	"fmt"
	"go/ast"
)

func SearchFunctionDeclaration(astFile *ast.File, funcName string) (*ast.FuncDecl, error) {
	found := []*ast.FuncDecl{}

	ast.Inspect(astFile, func(n ast.Node) bool {
		if n != nil {
			if functionDeclaration, ok := n.(*ast.FuncDecl); ok {
				if functionDeclaration.Name.Name == funcName {
					found = append(found, functionDeclaration)
					return false
				}
			}
		}
		return true
	})

	if len(found) == 0 {
		return nil, fmt.Errorf("could not find function '%s'", funcName)
	} else if len(found) > 1 {
		return nil, fmt.Errorf("more than one function definition with the name  '%s'", funcName)
	} else {
		return found[0], nil
	}
}

func ReplaceFunctionName(n *ast.FuncDecl, newName string) {
	n.Name.Name = newName
}
