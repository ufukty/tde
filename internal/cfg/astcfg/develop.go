package astcfg

import (
	astw "tde/internal/ast_wrapper"
	"tde/internal/cfg/astcfg/context"
	nc "tde/internal/cfg/astcfg/node_constructor"
	utl "tde/internal/utilities"

	"fmt"
	"go/ast"
)

func listAppendableSpots(node ast.Node) (appandableNodes []*astw.TraversableNode) {
	astw.Traverse(astw.GetTraversableNodeForASTNode(node), func(tNode *astw.TraversableNode) bool {
		if tNode.IsNil || tNode.ExpectedType.IsSliceType() {
			appandableNodes = append(appandableNodes, tNode)
		}
		return true
	})
	return
}

// TODO: implement
func getContext(root ast.Node, tNode *astw.TraversableNode) context.Context { // FIXME: input list +funcDecl, global scope
	ctx := context.NewContext()
	return ctx
}

func appendRandomly(tNode *astw.TraversableNode, ctx context.Context) {

	if tNode.ExpectedType.IsSliceType() && !tNode.IsNil {
		switch slice := tNode.Value.(type) {
		case []*ast.ImportSpec:
			tNode.Value = append(slice, nc.ImportSpec(ctx, 1))
		case []*ast.Ident:
			tNode.Value = append(slice, nc.Ident(ctx, 1))
		case []*ast.Field:
			tNode.Value = append(slice, nc.Field(ctx, 1))
		case []ast.Stmt:
			tNode.Value = append(slice, nc.Stmt(ctx, 1))
		case []ast.Decl:
			tNode.Value = append(slice, nc.Decl(ctx, 1))
		case []ast.Spec:
			tNode.Value = append(slice, nc.Spec(ctx, 1))
		case []ast.Expr:
			tNode.Value = append(slice, nc.Expr(ctx, 1))
		}

	} else if tNode.IsNil {

		astw.Traverse(*tNode.Parent, func(tSubnode *astw.TraversableNode) bool {
			if tSubnode == tNode {
				return true
			}

			if tSubnode.Parent == tNode {
				fmt.Println(tSubnode)
			}

			return false
		})

	}
}

// Picks an appandable spot randomly (either a nil field or end of a slice type field)
// Creates just one node and appends to choosen spot
func Develop(astFile *ast.File, astFuncDecl *ast.FuncDecl) {
	appandableSpots := listAppendableSpots(astFuncDecl)
	if len(appandableSpots) == 0 {
		return
	}
	choosenSpot := *utl.Pick(appandableSpots)
	ctx := getContext(astFile, choosenSpot)
w	appendRandomly(choosenSpot, ctx)

}
