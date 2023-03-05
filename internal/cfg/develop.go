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

func appendRandomly(tNode *trav.TraversableNode, ctx *context.Context, depthLimit int) (ast.Node, error) {
	var newNode ast.Node

	if tNode.ExpectedType.IsSliceType() {
		switch tNode.Value.(type) {
		case []*ast.ImportSpec:
			newNode = nc.ImportSpec(ctx, depthLimit)
		case []*ast.Ident:
			newNode = nc.Ident(ctx, depthLimit)
		case []*ast.Field:
			newNode = nc.Field(ctx, depthLimit)
		case []ast.Stmt:
			newNode = nc.Stmt(ctx, depthLimit)
		case []ast.Decl:
			newNode = nc.Decl(ctx, depthLimit)
		case []ast.Spec:
			newNode = nc.Spec(ctx, depthLimit)
		case []ast.Expr:
			newNode = nc.Expr(ctx, depthLimit)
		default:
			return nil, errors.New("Unhandled case for slice type node creation")
		}

	} else if tNode.ExpectedType.IsInterfaceType() {
		switch tNode.ExpectedType {
		case ast_types.Expr:
			newNode = nc.Expr(ctx, depthLimit)
		case ast_types.Stmt:
			newNode = nc.Stmt(ctx, depthLimit)
		case ast_types.Decl:
			newNode = nc.Decl(ctx, depthLimit)
		case ast_types.Spec:
			newNode = nc.Spec(ctx, depthLimit)
		case ast_types.TypeExpr:
			newNode = nc.Type(ctx, depthLimit)
		default:
			return nil, errors.New("Unhandled case for interface type node creation")
		}

	} else if tNode.ExpectedType.IsConcreteType() {

		switch tNode.ExpectedType {
		// case ast_types.BadDecl:
		// 	newNode = nc.BadDecl(ctx, depthLimit)
		// case ast_types.BadExpr:
		// 	newNode = nc.BadExpr(ctx, depthLimit)
		// case ast_types.BadStmt:
		// 	newNode = nc.BadStmt(ctx, depthLimit)
		// case ast_types.Comment:
		// 	newNode = nc.Comment(ctx, depthLimit)
		// case ast_types.CommentGroup:
		// 	newNode = nc.CommentGroup(ctx, depthLimit)
		// case ast_types.File:
		// 	newNode = nc.File(ctx, depthLimit)
		// case ast_types.Package:
		// 	newNode = nc.Package(ctx, depthLimit)

		case ast_types.ArrayType:
			newNode = nc.ArrayType(ctx, depthLimit)
		case ast_types.AssignStmt:
			newNode = nc.AssignStmt(ctx, depthLimit)
		case ast_types.BasicLit:
			newNode = nc.BasicLit(ctx, depthLimit)
		case ast_types.BinaryExpr:
			newNode = nc.BinaryExpr(ctx, depthLimit)
		case ast_types.BlockStmt:
			newNode = nc.BlockStmt(ctx, depthLimit)
		case ast_types.BranchStmt:
			newNode = nc.BranchStmt(ctx, depthLimit)
		case ast_types.CallExpr:
			newNode = nc.CallExpr(ctx, depthLimit)
		case ast_types.CaseClause:
			newNode = nc.CaseClause(ctx, depthLimit)
		case ast_types.ChanType:
			newNode = nc.ChanType(ctx, depthLimit)
		case ast_types.CommClause:
			newNode = nc.CommClause(ctx, depthLimit)
		case ast_types.CompositeLit:
			newNode = nc.CompositeLit(ctx, depthLimit)
		case ast_types.DeclStmt:
			newNode = nc.DeclStmt(ctx, depthLimit)
		case ast_types.DeferStmt:
			newNode = nc.DeferStmt(ctx, depthLimit)
		case ast_types.Ellipsis:
			newNode = nc.Ellipsis(ctx, depthLimit)
		case ast_types.EmptyStmt:
			newNode = nc.EmptyStmt(ctx, depthLimit)
		case ast_types.ExprStmt:
			newNode = nc.ExprStmt(ctx, depthLimit)
		case ast_types.Field:
			newNode = nc.Field(ctx, depthLimit)
		case ast_types.FieldList:
			newNode = nc.FieldList(ctx, depthLimit)
		case ast_types.ForStmt:
			newNode = nc.ForStmt(ctx, depthLimit)
		case ast_types.FuncDecl:
			newNode = nc.FuncDecl(ctx, depthLimit)
		case ast_types.FuncLit:
			newNode = nc.FuncLit(ctx, depthLimit)
		case ast_types.FuncType:
			newNode = nc.FuncType(ctx, depthLimit)
		case ast_types.GenDecl:
			newNode = nc.GenDecl(ctx, depthLimit)
		case ast_types.GoStmt:
			newNode = nc.GoStmt(ctx, depthLimit)
		case ast_types.Ident:
			newNode = nc.Ident(ctx, depthLimit)
		case ast_types.IfStmt:
			newNode = nc.IfStmt(ctx, depthLimit)
		case ast_types.ImportSpec:
			newNode = nc.ImportSpec(ctx, depthLimit)
		case ast_types.IncDecStmt:
			newNode = nc.IncDecStmt(ctx, depthLimit)
		case ast_types.IndexExpr:
			newNode = nc.IndexExpr(ctx, depthLimit)
		case ast_types.IndexListExpr:
			newNode = nc.IndexListExpr(ctx, depthLimit)
		case ast_types.InterfaceType:
			newNode = nc.InterfaceType(ctx, depthLimit)
		case ast_types.KeyValueExpr:
			newNode = nc.KeyValueExpr(ctx, depthLimit)
		case ast_types.LabeledStmt:
			newNode = nc.LabeledStmt(ctx, depthLimit)
		case ast_types.MapType:
			newNode = nc.MapType(ctx, depthLimit)
		case ast_types.ParenExpr:
			newNode = nc.ParenExpr(ctx, depthLimit)
		case ast_types.RangeStmt:
			newNode = nc.RangeStmt(ctx, depthLimit)
		case ast_types.ReturnStmt:
			newNode = nc.ReturnStmt(ctx, depthLimit)
		case ast_types.SelectorExpr:
			newNode = nc.SelectorExpr(ctx, depthLimit)
		case ast_types.SelectStmt:
			newNode = nc.SelectStmt(ctx, depthLimit)
		case ast_types.SendStmt:
			newNode = nc.SendStmt(ctx, depthLimit)
		case ast_types.SliceExpr:
			newNode = nc.SliceExpr(ctx, depthLimit)
		case ast_types.StarExpr:
			newNode = nc.StarExpr(ctx, depthLimit)
		case ast_types.StructType:
			newNode = nc.StructType(ctx, depthLimit)
		case ast_types.SwitchStmt:
			newNode = nc.SwitchStmt(ctx, depthLimit)
		case ast_types.TypeAssertExpr:
			newNode = nc.TypeAssertExpr(ctx, depthLimit)
		case ast_types.TypeSpec:
			newNode = nc.TypeSpec(ctx, depthLimit)
		case ast_types.TypeSwitchStmt:
			newNode = nc.TypeSwitchStmt(ctx, depthLimit)
		case ast_types.UnaryExpr:
			newNode = nc.UnaryExpr(ctx, depthLimit)
		case ast_types.ValueSpec:
			newNode = nc.ValueSpec(ctx, depthLimit)
		default:
			return nil, errors.New("Unhandled case for concrete type node creation")
		}

	}
	if newNode == nil {
		return nil, errors.New("Could not create an instance of ast.Node")
	}
	if ok := tNode.Ref.Set(newNode); !ok {
		return newNode, errors.New("Could not append created ast.Node instance to its place")
	}
	return newNode, nil
}

// Picks an appandable spot randomly (either a nil field or end of a slice type field)
// Creates just one node and appends to choosen spot
func Develop(astPkg *ast.Package, astFile *ast.File, astFuncDecl *ast.FuncDecl, depthLimit int) (ast.Node, error) {
	availableSpots := listAppendableSpots(astFuncDecl.Body)
	if len(availableSpots) == 0 {
		return nil, errors.New("No available spots found in AST to place new node")
	}
	choosenSpot := *utl.Pick(availableSpots)
	ctx, err := context_resolution.GetContextForSpot(
		astPkg,
		trav.GetTraversableNodeForASTNode(astFuncDecl),
		choosenSpot,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to resolve context for choosen spot.")
	}
	newNode, err := appendRandomly(choosenSpot, ctx, depthLimit)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to append a random node into choosen spot")
	}
	return newNode, nil
}
