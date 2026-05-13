package find

import (
	"fmt"
	"go/ast"
)

func FunctionInFile(f *ast.File, name string) (*ast.FuncDecl, error) {
	if f != nil && f.Decls != nil {
		for _, d := range f.Decls {
			if fd, ok := d.(*ast.FuncDecl); ok && fd.Name != nil && fd.Name.Name == name {
				return fd, nil
			}
		}
	}
	return nil, fmt.Errorf("not found")
}

func FunctionInPackage(p *ast.Package, name string) (*ast.FuncDecl, error) {
	if p != nil && p.Files != nil {
		for _, f := range p.Files {
			if fd, err := FunctionInFile(f, name); err == nil {
				return fd, nil
			}
		}
	}
	return nil, fmt.Errorf("not found")
}
