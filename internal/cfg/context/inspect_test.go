package context

import (
	"go/ast"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_InspectWithContext(t *testing.T) {

	identA := TEST_TREE.Decls[0].(*ast.FuncDecl).Type.Params.List[0].Names[0]
	identB := TEST_TREE.Decls[0].(*ast.FuncDecl).Type.Params.List[0].Names[1]
	identC := TEST_TREE.Decls[0].(*ast.FuncDecl).Body.List[0].(*ast.AssignStmt).Lhs[0].(*ast.Ident)
	returnStmt := TEST_TREE.Decls[0].(*ast.FuncDecl).Body.List[1].(*ast.ReturnStmt)

	InspectWithContext(TEST_TREE, func(ctx Context, node ast.Node) {
		// fmt.Println(reflect.TypeOf(node), ctx.Scopes)
		if node == returnStmt {
			if slices.Index(ctx.GetVariables(), identA) == -1 {
				t.Error("failed, input parameter 'a' is not detected")
			}
			if slices.Index(ctx.GetVariables(), identB) == -1 {
				t.Error("failed, input parameter 'b' is not detected")
			}
			if slices.Index(ctx.GetVariables(), identC) == -1 {
				t.Error("failed, variable 'c' is not detected")
			}
		}
	})
}
