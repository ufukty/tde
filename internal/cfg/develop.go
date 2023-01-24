package cfg

import (
	astw "tde/internal/ast_wrapper"
	"tde/internal/cfg/context"
	nc "tde/internal/cfg/node_constructor"
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

	if tNode.ExpectedType.IsSliceType() {

		var added ast.Node
		switch tNode.Value.(type) {
		case []*ast.ImportSpec:
			added = nc.ImportSpec(ctx, 1)
			tNode.Ref.Set(added)
		case []*ast.Ident:
			added = nc.Ident(ctx, 1)
			tNode.Ref.Set(added)
		case []*ast.Field:
			added = nc.Field(ctx, 1)
			tNode.Ref.Set(added)
		case []ast.Stmt:
			added = nc.Stmt(ctx, 1)
			tNode.Ref.Set(added)
		case []ast.Decl:
			added = nc.Decl(ctx, 1)
			tNode.Ref.Set(added)
		case []ast.Spec:
			added = nc.Spec(ctx, 1)
			tNode.Ref.Set(added)
		case []ast.Expr:
			added = nc.Expr(ctx, 1)
			tNode.Ref.Set(added)
		}

		fmt.Printf("Added into slice: %+v\n", added)

	} else if tNode.ExpectedType.IsInterfaceType() {

		var added ast.Node
		switch tNode.Value.(type) {
		case ast.Expr:
			added = nc.Expr(ctx, 1)
			tNode.Ref.Set(added)
		case ast.Stmt:
			added = nc.Stmt(ctx, 1)
			tNode.Ref.Set(added)
		case ast.Decl:
			added = nc.Decl(ctx, 1)
			tNode.Ref.Set(added)
		case ast.Spec:
			added = nc.Spec(ctx, 1)
			tNode.Ref.Set(added)
		}

		fmt.Printf("Interface added: %+v\n", added)

	} else if tNode.ExpectedType.IsConcreteType() {

		var add ast.Node
		switch tNode.ExpectedType {
		case astw.ArrayType:
			add = nc.ArrayType(ctx, 1)
			tNode.Ref.Set(add)
		case astw.AssignStmt:
			add = nc.AssignStmt(ctx, 1)
			tNode.Ref.Set(add)
		// case astw.BadDecl:
		// add = nc.BadDecl(ctx, 1)
		// 	tNode.Ref.Set(add)
		// case astw.BadExpr:
		// add = nc.BadExpr(ctx, 1)
		// 	tNode.Ref.Set(add)
		// case astw.BadStmt:
		// add = nc.BadStmt(ctx, 1)
		// 	tNode.Ref.Set(add)
		case astw.BasicLit:
			add = nc.BasicLit(ctx, 1)
			tNode.Ref.Set(add)
		case astw.BinaryExpr:
			add = nc.BinaryExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.BlockStmt:
			add = nc.BlockStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.BranchStmt:
			add = nc.BranchStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.CallExpr:
			add = nc.CallExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.CaseClause:
			add = nc.CaseClause(ctx, 1)
			tNode.Ref.Set(add)
		case astw.ChanType:
			add = nc.ChanType(ctx, 1)
			tNode.Ref.Set(add)
		case astw.CommClause:
			add = nc.CommClause(ctx, 1)
			tNode.Ref.Set(add)
		// case astw.Comment:
		// add = nc.Comment(ctx, 1)
		// 	tNode.Ref.Set(add)
		// case astw.CommentGroup:
		// add = nc.CommentGroup(ctx, 1)
		// 	tNode.Ref.Set(add)
		case astw.CompositeLit:
			add = nc.CompositeLit(ctx, 1)
			tNode.Ref.Set(add)
		case astw.DeclStmt:
			add = nc.DeclStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.DeferStmt:
			add = nc.DeferStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.Ellipsis:
			add = nc.Ellipsis(ctx, 1)
			tNode.Ref.Set(add)
		case astw.EmptyStmt:
			add = nc.EmptyStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.ExprStmt:
			add = nc.ExprStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.Field:
			add = nc.Field(ctx, 1)
			tNode.Ref.Set(add)
		case astw.FieldList:
			add = nc.FieldList(ctx, 1)
			tNode.Ref.Set(add)
		// case astw.File:
		// add = nc.File(ctx, 1)
		// 	tNode.Ref.Set(add)
		case astw.ForStmt:
			add = nc.ForStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.FuncDecl:
			add = nc.FuncDecl(ctx, 1)
			tNode.Ref.Set(add)
		case astw.FuncLit:
			add = nc.FuncLit(ctx, 1)
			tNode.Ref.Set(add)
		case astw.FuncType:
			add = nc.FuncType(ctx, 1)
			tNode.Ref.Set(add)
		case astw.GenDecl:
			add = nc.GenDecl(ctx, 1)
			tNode.Ref.Set(add)
		case astw.GoStmt:
			add = nc.GoStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.Ident:
			add = nc.Ident(ctx, 1)
			tNode.Ref.Set(add)
		case astw.IfStmt:
			add = nc.IfStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.ImportSpec:
			add = nc.ImportSpec(ctx, 1)
			tNode.Ref.Set(add)
		case astw.IncDecStmt:
			add = nc.IncDecStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.IndexExpr:
			add = nc.IndexExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.IndexListExpr:
			add = nc.IndexListExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.InterfaceType:
			add = nc.InterfaceType(ctx, 1)
			tNode.Ref.Set(add)
		case astw.KeyValueExpr:
			add = nc.KeyValueExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.LabeledStmt:
			add = nc.LabeledStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.MapType:
			add = nc.MapType(ctx, 1)
			tNode.Ref.Set(add)
		// case astw.Package:
		// add = nc.Package(ctx, 1)
		// 	tNode.Ref.Set(add)
		case astw.ParenExpr:
			add = nc.ParenExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.RangeStmt:
			add = nc.RangeStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.ReturnStmt:
			add = nc.ReturnStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.SelectorExpr:
			add = nc.SelectorExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.SelectStmt:
			add = nc.SelectStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.SendStmt:
			add = nc.SendStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.SliceExpr:
			add = nc.SliceExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.StarExpr:
			add = nc.StarExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.StructType:
			add = nc.StructType(ctx, 1)
			tNode.Ref.Set(add)
		case astw.SwitchStmt:
			add = nc.SwitchStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.TypeAssertExpr:
			add = nc.TypeAssertExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.TypeSpec:
			add = nc.TypeSpec(ctx, 1)
			tNode.Ref.Set(add)
		case astw.TypeSwitchStmt:
			add = nc.TypeSwitchStmt(ctx, 1)
			tNode.Ref.Set(add)
		case astw.UnaryExpr:
			add = nc.UnaryExpr(ctx, 1)
			tNode.Ref.Set(add)
		case astw.ValueSpec:
			add = nc.ValueSpec(ctx, 1)
			tNode.Ref.Set(add)
		}

		fmt.Printf("Concrete instance added: %+v\n", add)

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
	appendRandomly(choosenSpot, ctx)

}
