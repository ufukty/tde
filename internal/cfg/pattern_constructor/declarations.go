package pattern_constructor

import (
	"go/ast"
	"go/token"
)

func ImportPackage(path string) *ast.GenDecl {
	return &ast.GenDecl{
		// Doc:    &ast.CommentGroup{},
		// TokPos: 0,
		Tok: token.IMPORT,
		// Lparen: 0,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				// Doc:     &ast.CommentGroup{},
				// Name:    &ast.Ident{},
				Path: &ast.BasicLit{
					// ValuePos: 0,
					Kind:  token.STRING,
					Value: path,
				},
				// Comment: &ast.CommentGroup{},
				// EndPos:  0,
			},
		},
		// Rparen: 0,
	}
}
