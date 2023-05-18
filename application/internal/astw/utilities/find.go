package utilities

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

func FindFuncDeclInPkg(pkg *ast.Package, name string) (*ast.File, *ast.FuncDecl, error) {
	funcDecl := []*ast.FuncDecl{}
	file := []*ast.File{}

	var fileVisited *ast.File

	ast.Inspect(pkg, func(n ast.Node) bool {
		if n != nil {
			if file, ok := n.(*ast.File); ok {
				fileVisited = file
			}
			if functionDeclaration, ok := n.(*ast.FuncDecl); ok {
				if functionDeclaration.Name.Name == name {
					funcDecl = append(funcDecl, functionDeclaration)
					file = append(file, fileVisited)
					return false
				}
			}
		}
		return true
	})

	if len(funcDecl) == 0 {
		return nil, nil, fmt.Errorf("could not find function '%s'", name)
	} else if len(funcDecl) > 1 {
		return nil, nil, fmt.Errorf("more than one function definition with the name  '%s'", name)
	} else {
		return file[0], funcDecl[0], nil
	}
}
