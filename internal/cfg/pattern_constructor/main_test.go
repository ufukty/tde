package pattern_constructor

import (
	"go/ast"
	"testing"
)

func TestXxx(t *testing.T) {
	SliceOf(&ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{},
					Type:  nil,
					Tag:   &ast.BasicLit{},
				},
			},
		},
		Incomplete: false,
	})
}
