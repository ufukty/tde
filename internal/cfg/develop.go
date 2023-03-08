package cfg

import (
	"fmt"
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
		if tNode.PointsToNilSpot {
			appandableNodes = append(appandableNodes, tNode)
		}
		return true
	})
	return
}

func appendRandomly(dst *trav.TraversableNode, ctx *context.Context, depthLimit int) (appended any, err error) {

	if dst.ExpectedType.IsSliceType() {
		switch dst.Value.(type) {
		case []*ast.ImportSpec:
			appended = &[]*ast.ImportSpec{}
		case []*ast.Ident:
			appended = &[]*ast.Ident{}
		case []*ast.Field:
			appended = &[]*ast.Field{}
		case []ast.Stmt:
			appended = &[]ast.Stmt{}
		case []ast.Decl:
			appended = &[]ast.Decl{}
		case []ast.Spec:
			appended = &[]ast.Spec{}
		case []ast.Expr:
			appended = &[]ast.Expr{}
		default:
			return nil, errors.New("Unhandled case for slice type node creation")
		}

	} else if dst.ExpectedType.IsInterfaceType() {
		switch dst.ExpectedType {
		case ast_types.Expr:
			appended = nc.Expr(ctx, depthLimit)
		case ast_types.Stmt:
			appended = nc.Stmt(ctx, depthLimit)
		case ast_types.Decl:
			appended = nc.Decl(ctx, depthLimit)
		case ast_types.Spec:
			appended = nc.Spec(ctx, depthLimit)
		case ast_types.TypeExpr:
			appended = nc.Type(ctx, depthLimit)
		default:
			return nil, errors.New("Unhandled case for interface type node creation")
		}

	} else if dst.ExpectedType.IsConcreteType() {

		switch dst.ExpectedType {
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
			appended = nc.ArrayType(ctx, depthLimit)
		case ast_types.AssignStmt:
			appended = nc.AssignStmt(ctx, depthLimit)
		case ast_types.BasicLit:
			appended = nc.BasicLit(ctx, depthLimit)
		case ast_types.BinaryExpr:
			appended = nc.BinaryExpr(ctx, depthLimit)
		case ast_types.BlockStmt:
			appended = nc.BlockStmt(ctx, depthLimit)
		case ast_types.BranchStmt:
			appended = nc.BranchStmt(ctx, depthLimit)
		case ast_types.CallExpr:
			appended = nc.CallExpr(ctx, depthLimit)
		case ast_types.CaseClause:
			appended = nc.CaseClause(ctx, depthLimit)
		case ast_types.ChanType:
			appended = nc.ChanType(ctx, depthLimit)
		case ast_types.CommClause:
			appended = nc.CommClause(ctx, depthLimit)
		case ast_types.CompositeLit:
			appended = nc.CompositeLit(ctx, depthLimit)
		case ast_types.DeclStmt:
			appended = nc.DeclStmt(ctx, depthLimit)
		case ast_types.DeferStmt:
			appended = nc.DeferStmt(ctx, depthLimit)
		case ast_types.Ellipsis:
			appended = nc.Ellipsis(ctx, depthLimit)
		case ast_types.EmptyStmt:
			appended = nc.EmptyStmt(ctx, depthLimit)
		case ast_types.ExprStmt:
			appended = nc.ExprStmt(ctx, depthLimit)
		case ast_types.Field:
			appended = nc.Field(ctx, depthLimit)
		case ast_types.FieldList:
			appended = nc.FieldList(ctx, depthLimit)
		case ast_types.ForStmt:
			appended = nc.ForStmt(ctx, depthLimit)
		case ast_types.FuncDecl:
			appended = nc.FuncDecl(ctx, depthLimit)
		case ast_types.FuncLit:
			appended = nc.FuncLit(ctx, depthLimit)
		case ast_types.FuncType:
			appended = nc.FuncType(ctx, depthLimit)
		case ast_types.GenDecl:
			appended = nc.GenDecl(ctx, depthLimit)
		case ast_types.GoStmt:
			appended = nc.GoStmt(ctx, depthLimit)
		case ast_types.Ident:
			appended = nc.Ident(ctx, depthLimit)
		case ast_types.IfStmt:
			appended = nc.IfStmt(ctx, depthLimit)
		case ast_types.ImportSpec:
			appended = nc.ImportSpec(ctx, depthLimit)
		case ast_types.IncDecStmt:
			appended = nc.IncDecStmt(ctx, depthLimit)
		case ast_types.IndexExpr:
			appended = nc.IndexExpr(ctx, depthLimit)
		case ast_types.IndexListExpr:
			appended = nc.IndexListExpr(ctx, depthLimit)
		case ast_types.InterfaceType:
			appended = nc.InterfaceType(ctx, depthLimit)
		case ast_types.KeyValueExpr:
			appended = nc.KeyValueExpr(ctx, depthLimit)
		case ast_types.LabeledStmt:
			appended = nc.LabeledStmt(ctx, depthLimit)
		case ast_types.MapType:
			appended = nc.MapType(ctx, depthLimit)
		case ast_types.ParenExpr:
			appended = nc.ParenExpr(ctx, depthLimit)
		case ast_types.RangeStmt:
			appended = nc.RangeStmt(ctx, depthLimit)
		case ast_types.ReturnStmt:
			appended = nc.ReturnStmt(ctx, depthLimit)
		case ast_types.SelectorExpr:
			appended = nc.SelectorExpr(ctx, depthLimit)
		case ast_types.SelectStmt:
			appended = nc.SelectStmt(ctx, depthLimit)
		case ast_types.SendStmt:
			appended = nc.SendStmt(ctx, depthLimit)
		case ast_types.SliceExpr:
			appended = nc.SliceExpr(ctx, depthLimit)
		case ast_types.StarExpr:
			appended = nc.StarExpr(ctx, depthLimit)
		case ast_types.StructType:
			appended = nc.StructType(ctx, depthLimit)
		case ast_types.SwitchStmt:
			appended = nc.SwitchStmt(ctx, depthLimit)
		case ast_types.TypeAssertExpr:
			appended = nc.TypeAssertExpr(ctx, depthLimit)
		case ast_types.TypeSpec:
			appended = nc.TypeSpec(ctx, depthLimit)
		case ast_types.TypeSwitchStmt:
			appended = nc.TypeSwitchStmt(ctx, depthLimit)
		case ast_types.UnaryExpr:
			appended = nc.UnaryExpr(ctx, depthLimit)
		case ast_types.ValueSpec:
			appended = nc.ValueSpec(ctx, depthLimit)
		default:
			return nil, errors.New("Unhandled case for concrete type node creation")
		}

	}
	if appended == nil {
		return nil, errors.New("Could not create an instance of ast.Node")
	}
	if ok := dst.Ref.Set(appended); !ok {

		fmt.Println("dst:", dst.ExpectedType, "appended:", appended)
		return appended, errors.New("Could not append created ast.Node instance to its place")
	}
	return appended, nil
}

// Picks an appandable spot randomly (either a nil field or end of a slice type field)
// Creates just one node and appends to choosen spot
func Develop(astPkg *ast.Package, astFile *ast.File, astFuncDecl *ast.FuncDecl, depthLimit int) (appended any, err error) {
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
