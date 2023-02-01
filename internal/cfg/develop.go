package cfg

import (
	"tde/internal/astw"
	"tde/internal/cfg/context"
	nc "tde/internal/cfg/node_constructor"
	utl "tde/internal/utilities"

	"go/ast"

	"github.com/pkg/errors"
)

func listAppendableSpots(node ast.Node) (appandableNodes []*astw.TraversableNode) {
	astw.Traverse(astw.GetTraversableNodeForASTNode(node), func(tNode *astw.TraversableNode) bool {
		if tNode.PointsToNilSpot || tNode.ExpectedType.IsSliceType() {
			appandableNodes = append(appandableNodes, tNode)
		}
		return true
	})
	return
}

func appendRandomly(tNode *astw.TraversableNode, ctx *context.Context) {

	if tNode.ExpectedType.IsSliceType() {
		switch tNode.Value.(type) {
		case []*ast.ImportSpec:
			tNode.Ref.Set(nc.ImportSpec(ctx, 1))
		case []*ast.Ident:
			tNode.Ref.Set(nc.Ident(ctx, 1))
		case []*ast.Field:
			tNode.Ref.Set(nc.Field(ctx, 1))
		case []ast.Stmt:
			tNode.Ref.Set(nc.Stmt(ctx, 1))
		case []ast.Decl:
			tNode.Ref.Set(nc.Decl(ctx, 1))
		case []ast.Spec:
			tNode.Ref.Set(nc.Spec(ctx, 1))
		case []ast.Expr:
			tNode.Ref.Set(nc.Expr(ctx, 1))
		}

	} else if tNode.ExpectedType.IsInterfaceType() {
		switch tNode.Value.(type) {
		case ast.Expr:
			tNode.Ref.Set(nc.Expr(ctx, 1))
		case ast.Stmt:
			tNode.Ref.Set(nc.Stmt(ctx, 1))
		case ast.Decl:
			tNode.Ref.Set(nc.Decl(ctx, 1))
		case ast.Spec:
			tNode.Ref.Set(nc.Spec(ctx, 1))
		}

	} else if tNode.ExpectedType.IsConcreteType() {

		switch tNode.ExpectedType {
		// case astw.BadDecl:
		// 	tNode.Ref.Set(nc.BadDecl(ctx, 1))
		// case astw.BadExpr:
		// 	tNode.Ref.Set(nc.BadExpr(ctx, 1))
		// case astw.BadStmt:
		// 	tNode.Ref.Set(nc.BadStmt(ctx, 1))
		// case astw.Comment:
		// 	tNode.Ref.Set(nc.Comment(ctx, 1))
		// case astw.CommentGroup:
		// 	tNode.Ref.Set(nc.CommentGroup(ctx, 1))
		// case astw.File:
		// 	tNode.Ref.Set(nc.File(ctx, 1))
		// case astw.Package:
		// 	tNode.Ref.Set(nc.Package(ctx, 1))
		case astw.ArrayType:
			tNode.Ref.Set(nc.ArrayType(ctx, 1))
		case astw.AssignStmt:
			tNode.Ref.Set(nc.AssignStmt(ctx, 1))
		case astw.BasicLit:
			tNode.Ref.Set(nc.BasicLit(ctx, 1))
		case astw.BinaryExpr:
			tNode.Ref.Set(nc.BinaryExpr(ctx, 1))
		case astw.BlockStmt:
			tNode.Ref.Set(nc.BlockStmt(ctx, 1))
		case astw.BranchStmt:
			tNode.Ref.Set(nc.BranchStmt(ctx, 1))
		case astw.CallExpr:
			tNode.Ref.Set(nc.CallExpr(ctx, 1))
		case astw.CaseClause:
			tNode.Ref.Set(nc.CaseClause(ctx, 1))
		case astw.ChanType:
			tNode.Ref.Set(nc.ChanType(ctx, 1))
		case astw.CommClause:
			tNode.Ref.Set(nc.CommClause(ctx, 1))
		case astw.CompositeLit:
			tNode.Ref.Set(nc.CompositeLit(ctx, 1))
		case astw.DeclStmt:
			tNode.Ref.Set(nc.DeclStmt(ctx, 1))
		case astw.DeferStmt:
			tNode.Ref.Set(nc.DeferStmt(ctx, 1))
		case astw.Ellipsis:
			tNode.Ref.Set(nc.Ellipsis(ctx, 1))
		case astw.EmptyStmt:
			tNode.Ref.Set(nc.EmptyStmt(ctx, 1))
		case astw.ExprStmt:
			tNode.Ref.Set(nc.ExprStmt(ctx, 1))
		case astw.Field:
			tNode.Ref.Set(nc.Field(ctx, 1))
		case astw.FieldList:
			tNode.Ref.Set(nc.FieldList(ctx, 1))
		case astw.ForStmt:
			tNode.Ref.Set(nc.ForStmt(ctx, 1))
		case astw.FuncDecl:
			tNode.Ref.Set(nc.FuncDecl(ctx, 1))
		case astw.FuncLit:
			tNode.Ref.Set(nc.FuncLit(ctx, 1))
		case astw.FuncType:
			tNode.Ref.Set(nc.FuncType(ctx, 1))
		case astw.GenDecl:
			tNode.Ref.Set(nc.GenDecl(ctx, 1))
		case astw.GoStmt:
			tNode.Ref.Set(nc.GoStmt(ctx, 1))
		case astw.Ident:
			tNode.Ref.Set(nc.Ident(ctx, 1))
		case astw.IfStmt:
			tNode.Ref.Set(nc.IfStmt(ctx, 1))
		case astw.ImportSpec:
			tNode.Ref.Set(nc.ImportSpec(ctx, 1))
		case astw.IncDecStmt:
			tNode.Ref.Set(nc.IncDecStmt(ctx, 1))
		case astw.IndexExpr:
			tNode.Ref.Set(nc.IndexExpr(ctx, 1))
		case astw.IndexListExpr:
			tNode.Ref.Set(nc.IndexListExpr(ctx, 1))
		case astw.InterfaceType:
			tNode.Ref.Set(nc.InterfaceType(ctx, 1))
		case astw.KeyValueExpr:
			tNode.Ref.Set(nc.KeyValueExpr(ctx, 1))
		case astw.LabeledStmt:
			tNode.Ref.Set(nc.LabeledStmt(ctx, 1))
		case astw.MapType:
			tNode.Ref.Set(nc.MapType(ctx, 1))
		case astw.ParenExpr:
			tNode.Ref.Set(nc.ParenExpr(ctx, 1))
		case astw.RangeStmt:
			tNode.Ref.Set(nc.RangeStmt(ctx, 1))
		case astw.ReturnStmt:
			tNode.Ref.Set(nc.ReturnStmt(ctx, 1))
		case astw.SelectorExpr:
			tNode.Ref.Set(nc.SelectorExpr(ctx, 1))
		case astw.SelectStmt:
			tNode.Ref.Set(nc.SelectStmt(ctx, 1))
		case astw.SendStmt:
			tNode.Ref.Set(nc.SendStmt(ctx, 1))
		case astw.SliceExpr:
			tNode.Ref.Set(nc.SliceExpr(ctx, 1))
		case astw.StarExpr:
			tNode.Ref.Set(nc.StarExpr(ctx, 1))
		case astw.StructType:
			tNode.Ref.Set(nc.StructType(ctx, 1))
		case astw.SwitchStmt:
			tNode.Ref.Set(nc.SwitchStmt(ctx, 1))
		case astw.TypeAssertExpr:
			tNode.Ref.Set(nc.TypeAssertExpr(ctx, 1))
		case astw.TypeSpec:
			tNode.Ref.Set(nc.TypeSpec(ctx, 1))
		case astw.TypeSwitchStmt:
			tNode.Ref.Set(nc.TypeSwitchStmt(ctx, 1))
		case astw.UnaryExpr:
			tNode.Ref.Set(nc.UnaryExpr(ctx, 1))
		case astw.ValueSpec:
			tNode.Ref.Set(nc.ValueSpec(ctx, 1))
		}

	}
}

// Picks an appandable spot randomly (either a nil field or end of a slice type field)
// Creates just one node and appends to choosen spot
func Develop(astPkg *ast.Package, astFile *ast.File, astFuncDecl *ast.FuncDecl) error {
	availableSpots := listAppendableSpots(astFuncDecl)
	if len(availableSpots) == 0 {
		return errors.New("No available spots found in AST to place new node")
	}
	choosenSpot := *utl.Pick(availableSpots)
	ctx, err := context.GetContextForSpot(
		astPkg,
		astw.GetTraversableNodeForASTNode(astFuncDecl),
		choosenSpot,
	)
	if err != nil {
		return errors.Wrap(err, "Failed")
	}
	appendRandomly(choosenSpot, ctx)
	return nil
}
