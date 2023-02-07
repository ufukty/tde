package cfg

import (
	trav "tde/internal/astw/traverse"
	ast_types "tde/internal/astw/types"
	"tde/internal/cfg/context_resolution"
	"tde/internal/cfg/context_resolution/context"
	nc "tde/internal/cfg/node_constructor"
	utl "tde/internal/utilities"

	"go/ast"

	"github.com/pkg/errors"
)

func listAppendableSpots(node ast.Node) (appandableNodes []*trav.TraversableNode) {
	trav.Traverse(trav.GetTraversableNodeForASTNode(node), func(tNode *trav.TraversableNode) bool {
		if tNode.PointsToNilSpot || tNode.ExpectedType.IsSliceType() {
			appandableNodes = append(appandableNodes, tNode)
		}
		return true
	})
	return
}

func appendRandomly(tNode *trav.TraversableNode, ctx *context.Context, depthLimit int) {

	if tNode.ExpectedType.IsSliceType() {
		switch tNode.Value.(type) {
		case []*ast.ImportSpec:
			tNode.Ref.Set(nc.ImportSpec(ctx, depthLimit))
		case []*ast.Ident:
			tNode.Ref.Set(nc.Ident(ctx, depthLimit))
		case []*ast.Field:
			tNode.Ref.Set(nc.Field(ctx, depthLimit))
		case []ast.Stmt:
			tNode.Ref.Set(nc.Stmt(ctx, depthLimit))
		case []ast.Decl:
			tNode.Ref.Set(nc.Decl(ctx, depthLimit))
		case []ast.Spec:
			tNode.Ref.Set(nc.Spec(ctx, depthLimit))
		case []ast.Expr:
			tNode.Ref.Set(nc.Expr(ctx, depthLimit))
		}

	} else if tNode.ExpectedType.IsInterfaceType() {
		switch tNode.ExpectedType {
		case ast_types.Expr:
			tNode.Ref.Set(nc.Expr(ctx, depthLimit))
		case ast_types.Stmt:
			tNode.Ref.Set(nc.Stmt(ctx, depthLimit))
		case ast_types.Decl:
			tNode.Ref.Set(nc.Decl(ctx, depthLimit))
		case ast_types.Spec:
			tNode.Ref.Set(nc.Spec(ctx, depthLimit))
		case ast_types.TypeExpr:
			tNode.Ref.Set(nc.Type(ctx, depthLimit))
		}

	} else if tNode.ExpectedType.IsConcreteType() {

		switch tNode.ExpectedType {
		// case ast_types.BadDecl:
		// createdNode:=nc.BadDecl(ctx, depthLimit)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.BadExpr:
		// createdNode:=nc.BadExpr(ctx, depthLimit)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.BadStmt:
		// createdNode:=nc.BadStmt(ctx, depthLimit)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.Comment:
		// createdNode:=nc.Comment(ctx, depthLimit)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.CommentGroup:
		// createdNode:=nc.CommentGroup(ctx, depthLimit)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.File:
		// createdNode:=nc.File(ctx, depthLimit)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.Package:
		// createdNode:=nc.Package(ctx, depthLimit)
		// 	tNode.Ref.Set(createdNode)
		case ast_types.ArrayType:
			tNode.Ref.Set(nc.ArrayType(ctx, depthLimit))
		case ast_types.AssignStmt:
			tNode.Ref.Set(nc.AssignStmt(ctx, depthLimit))
		case ast_types.BasicLit:
			tNode.Ref.Set(nc.BasicLit(ctx, depthLimit))
		case ast_types.BinaryExpr:
			tNode.Ref.Set(nc.BinaryExpr(ctx, depthLimit))
		case ast_types.BlockStmt:
			tNode.Ref.Set(nc.BlockStmt(ctx, depthLimit))
		case ast_types.BranchStmt:
			tNode.Ref.Set(nc.BranchStmt(ctx, depthLimit))
		case ast_types.CallExpr:
			tNode.Ref.Set(nc.CallExpr(ctx, depthLimit))
		case ast_types.CaseClause:
			tNode.Ref.Set(nc.CaseClause(ctx, depthLimit))
		case ast_types.ChanType:
			tNode.Ref.Set(nc.ChanType(ctx, depthLimit))
		case ast_types.CommClause:
			tNode.Ref.Set(nc.CommClause(ctx, depthLimit))
		case ast_types.CompositeLit:
			tNode.Ref.Set(nc.CompositeLit(ctx, depthLimit))
		case ast_types.DeclStmt:
			tNode.Ref.Set(nc.DeclStmt(ctx, depthLimit))
		case ast_types.DeferStmt:
			tNode.Ref.Set(nc.DeferStmt(ctx, depthLimit))
		case ast_types.Ellipsis:
			tNode.Ref.Set(nc.Ellipsis(ctx, depthLimit))
		case ast_types.EmptyStmt:
			tNode.Ref.Set(nc.EmptyStmt(ctx, depthLimit))
		case ast_types.ExprStmt:
			tNode.Ref.Set(nc.ExprStmt(ctx, depthLimit))
		case ast_types.Field:
			tNode.Ref.Set(nc.Field(ctx, depthLimit))
		case ast_types.FieldList:
			tNode.Ref.Set(nc.FieldList(ctx, depthLimit))
		case ast_types.ForStmt:
			tNode.Ref.Set(nc.ForStmt(ctx, depthLimit))
		case ast_types.FuncDecl:
			tNode.Ref.Set(nc.FuncDecl(ctx, depthLimit))
		case ast_types.FuncLit:
			tNode.Ref.Set(nc.FuncLit(ctx, depthLimit))
		case ast_types.FuncType:
			tNode.Ref.Set(nc.FuncType(ctx, depthLimit))
		case ast_types.GenDecl:
			tNode.Ref.Set(nc.GenDecl(ctx, depthLimit))
		case ast_types.GoStmt:
			tNode.Ref.Set(nc.GoStmt(ctx, depthLimit))
		case ast_types.Ident:
			tNode.Ref.Set(nc.Ident(ctx, depthLimit))
		case ast_types.IfStmt:
			tNode.Ref.Set(nc.IfStmt(ctx, depthLimit))
		case ast_types.ImportSpec:
			tNode.Ref.Set(nc.ImportSpec(ctx, depthLimit))
		case ast_types.IncDecStmt:
			tNode.Ref.Set(nc.IncDecStmt(ctx, depthLimit))
		case ast_types.IndexExpr:
			tNode.Ref.Set(nc.IndexExpr(ctx, depthLimit))
		case ast_types.IndexListExpr:
			tNode.Ref.Set(nc.IndexListExpr(ctx, depthLimit))
		case ast_types.InterfaceType:
			tNode.Ref.Set(nc.InterfaceType(ctx, depthLimit))
		case ast_types.KeyValueExpr:
			tNode.Ref.Set(nc.KeyValueExpr(ctx, depthLimit))
		case ast_types.LabeledStmt:
			tNode.Ref.Set(nc.LabeledStmt(ctx, depthLimit))
		case ast_types.MapType:
			tNode.Ref.Set(nc.MapType(ctx, depthLimit))
		case ast_types.ParenExpr:
			tNode.Ref.Set(nc.ParenExpr(ctx, depthLimit))
		case ast_types.RangeStmt:
			tNode.Ref.Set(nc.RangeStmt(ctx, depthLimit))
		case ast_types.ReturnStmt:
			tNode.Ref.Set(nc.ReturnStmt(ctx, depthLimit))
		case ast_types.SelectorExpr:
			tNode.Ref.Set(nc.SelectorExpr(ctx, depthLimit))
		case ast_types.SelectStmt:
			tNode.Ref.Set(nc.SelectStmt(ctx, depthLimit))
		case ast_types.SendStmt:
			tNode.Ref.Set(nc.SendStmt(ctx, depthLimit))
		case ast_types.SliceExpr:
			tNode.Ref.Set(nc.SliceExpr(ctx, depthLimit))
		case ast_types.StarExpr:
			tNode.Ref.Set(nc.StarExpr(ctx, depthLimit))
		case ast_types.StructType:
			tNode.Ref.Set(nc.StructType(ctx, depthLimit))
		case ast_types.SwitchStmt:
			tNode.Ref.Set(nc.SwitchStmt(ctx, depthLimit))
		case ast_types.TypeAssertExpr:
			tNode.Ref.Set(nc.TypeAssertExpr(ctx, depthLimit))
		case ast_types.TypeSpec:
			tNode.Ref.Set(nc.TypeSpec(ctx, depthLimit))
		case ast_types.TypeSwitchStmt:
			tNode.Ref.Set(nc.TypeSwitchStmt(ctx, depthLimit))
		case ast_types.UnaryExpr:
			tNode.Ref.Set(nc.UnaryExpr(ctx, depthLimit))
		case ast_types.ValueSpec:
			tNode.Ref.Set(nc.ValueSpec(ctx, depthLimit))
		}

	}
}

// Picks an appandable spot randomly (either a nil field or end of a slice type field)
// Creates just one node and appends to choosen spot
func Develop(astPkg *ast.Package, astFile *ast.File, astFuncDecl *ast.FuncDecl, depthLimit int) error {
	availableSpots := listAppendableSpots(astFuncDecl.Body)
	if len(availableSpots) == 0 {
		return errors.New("No available spots found in AST to place new node")
	}
	choosenSpot := *utl.Pick(availableSpots)
	ctx, err := context_resolution.GetContextForSpot(
		astPkg,
		trav.GetTraversableNodeForASTNode(astFuncDecl),
		choosenSpot,
	)
	if err != nil {
		return errors.Wrap(err, "Failed")
	}
	appendRandomly(choosenSpot, ctx, depthLimit)
	return nil
}
