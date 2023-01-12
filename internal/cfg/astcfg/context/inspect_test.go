package context

import (
	"go/ast"
	"tde/internal/utilities"
	"testing"
)

func Test_Inspect(t *testing.T) {
	Inspect(TEST_TREE, func(ctx Context, node ast.Node) {
		if _, ok := node.(*ast.ReturnStmt); ok {
			if utilities.FindSliceIndex(ctx.GetVariables(), ast.Ident{Name: "c"}) == -1 {
				// fmt.Println(ctx.GetVariables())
				t.Error("failed, variable c has not found")
			}
		}
	})
}
