package stg

import (
	"fmt"
	"go/ast"
	"tde/internal/astw/traverse"
	"tde/internal/astw/types"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
	"tde/internal/evolution/genetics/nodes"
	"tde/internal/utilities/pick"
)

func listAppendableSpots(node ast.Node) (appandableNodes []*traverse.TraversableNode) {
	traverse.Traverse(traverse.GetTraversableNodeForASTNode(node), func(tNode *traverse.TraversableNode) bool {
		if tNode.PointsToNilSpot {
			appandableNodes = append(appandableNodes, tNode)
		}
		return true
	})
	return
}

func appendRandomly(dst *traverse.TraversableNode, ctx *context.Context, depthLimit int) (appended any, err error) {

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
			return nil, fmt.Errorf("Unhandled case for slice type node creation")
		}

	} else if dst.ExpectedType.IsInterfaceType() {
		switch dst.ExpectedType {
		case types.Expr:
			appended = nodes.Expr(ctx, depthLimit)
		case types.Stmt:
			appended = nodes.Stmt(ctx, depthLimit)
		case types.Decl:
			appended = nodes.Decl(ctx, depthLimit)
		case types.Spec:
			appended = nodes.Spec(ctx, depthLimit)
		case types.TypeExpr:
			appended = nodes.Type(ctx, depthLimit)
		default:
			return nil, fmt.Errorf("Unhandled case for interface type node creation")
		}

	} else if dst.ExpectedType.IsConcreteType() {

		switch dst.ExpectedType {
		// case types.BadDecl:
		// 	newNode = nc.BadDecl(ctx, depthLimit)
		// case types.BadExpr:
		// 	newNode = nc.BadExpr(ctx, depthLimit)
		// case types.BadStmt:
		// 	newNode = nc.BadStmt(ctx, depthLimit)
		// case types.Comment:
		// 	newNode = nc.Comment(ctx, depthLimit)
		// case types.CommentGroup:
		// 	newNode = nc.CommentGroup(ctx, depthLimit)
		// case types.File:
		// 	newNode = nc.File(ctx, depthLimit)
		// case types.Package:
		// 	newNode = nc.Package(ctx, depthLimit)

		case types.ArrayType:
			appended = nodes.ArrayType(ctx, depthLimit)
		case types.AssignStmt:
			appended = nodes.AssignStmt(ctx, depthLimit)
		case types.BasicLit:
			appended = nodes.BasicLit(ctx, depthLimit)
		case types.BinaryExpr:
			appended = nodes.BinaryExpr(ctx, depthLimit)
		case types.BlockStmt:
			appended = nodes.BlockStmt(ctx, depthLimit)
		case types.BranchStmt:
			appended = nodes.BranchStmt(ctx, depthLimit)
		case types.CallExpr:
			appended = nodes.CallExpr(ctx, depthLimit)
		case types.CaseClause:
			appended = nodes.CaseClause(ctx, depthLimit)
		case types.ChanType:
			appended = nodes.ChanType(ctx, depthLimit)
		case types.CommClause:
			appended = nodes.CommClause(ctx, depthLimit)
		case types.CompositeLit:
			appended = nodes.CompositeLit(ctx, depthLimit)
		case types.DeclStmt:
			appended = nodes.DeclStmt(ctx, depthLimit)
		case types.DeferStmt:
			appended = nodes.DeferStmt(ctx, depthLimit)
		case types.Ellipsis:
			appended = nodes.Ellipsis(ctx, depthLimit)
		case types.EmptyStmt:
			appended = nodes.EmptyStmt(ctx, depthLimit)
		case types.ExprStmt:
			appended = nodes.ExprStmt(ctx, depthLimit)
		case types.Field:
			appended = nodes.Field(ctx, depthLimit)
		case types.FieldList:
			appended = nodes.FieldList(ctx, depthLimit)
		case types.ForStmt:
			appended = nodes.ForStmt(ctx, depthLimit)
		case types.FuncDecl:
			appended = nodes.FuncDecl(ctx, depthLimit)
		case types.FuncLit:
			appended = nodes.FuncLit(ctx, depthLimit)
		case types.FuncType:
			appended = nodes.FuncType(ctx, depthLimit)
		case types.GenDecl:
			appended = nodes.GenDecl(ctx, depthLimit)
		case types.GoStmt:
			appended = nodes.GoStmt(ctx, depthLimit)
		case types.Ident:
			appended = nodes.Ident(ctx, depthLimit)
		case types.IfStmt:
			appended = nodes.IfStmt(ctx, depthLimit)
		case types.ImportSpec:
			appended = nodes.ImportSpec(ctx, depthLimit)
		case types.IncDecStmt:
			appended = nodes.IncDecStmt(ctx, depthLimit)
		case types.IndexExpr:
			appended = nodes.IndexExpr(ctx, depthLimit)
		case types.IndexListExpr:
			appended = nodes.IndexListExpr(ctx, depthLimit)
		case types.InterfaceType:
			appended = nodes.InterfaceType(ctx, depthLimit)
		case types.KeyValueExpr:
			appended = nodes.KeyValueExpr(ctx, depthLimit)
		case types.LabeledStmt:
			appended = nodes.LabeledStmt(ctx, depthLimit)
		case types.MapType:
			appended = nodes.MapType(ctx, depthLimit)
		case types.ParenExpr:
			appended = nodes.ParenExpr(ctx, depthLimit)
		case types.RangeStmt:
			appended = nodes.RangeStmt(ctx, depthLimit)
		case types.ReturnStmt:
			appended = nodes.ReturnStmt(ctx, depthLimit)
		case types.SelectorExpr:
			appended = nodes.SelectorExpr(ctx, depthLimit)
		case types.SelectStmt:
			appended = nodes.SelectStmt(ctx, depthLimit)
		case types.SendStmt:
			appended = nodes.SendStmt(ctx, depthLimit)
		case types.SliceExpr:
			appended = nodes.SliceExpr(ctx, depthLimit)
		case types.StarExpr:
			appended = nodes.StarExpr(ctx, depthLimit)
		case types.StructType:
			appended = nodes.StructType(ctx, depthLimit)
		case types.SwitchStmt:
			appended = nodes.SwitchStmt(ctx, depthLimit)
		case types.TypeAssertExpr:
			appended = nodes.TypeAssertExpr(ctx, depthLimit)
		case types.TypeSpec:
			appended = nodes.TypeSpec(ctx, depthLimit)
		case types.TypeSwitchStmt:
			appended = nodes.TypeSwitchStmt(ctx, depthLimit)
		case types.UnaryExpr:
			appended = nodes.UnaryExpr(ctx, depthLimit)
		case types.ValueSpec:
			appended = nodes.ValueSpec(ctx, depthLimit)
		default:
			return nil, fmt.Errorf("Unhandled case for concrete type node creation")
		}

	}
	if appended == nil {
		return nil, fmt.Errorf("Could not create an instance of ast.Node")
	}
	if ok := dst.Ref.Set(appended); !ok {

		fmt.Println("dst:", dst.ExpectedType, "appended:", appended)
		return appended, fmt.Errorf("Could not append created ast.Node instance to its place")
	}
	return appended, nil
}

// Picks an appandable spot randomly (either a nil field or end of a slice type field)
// Creates just one node and appends to choosen spot
func develop(astPkg *ast.Package, astFile *ast.File, astFuncDecl *ast.FuncDecl, depthLimit int) (appended any, err error) {
	availableSpots := listAppendableSpots(astFuncDecl.Body)
	if len(availableSpots) == 0 {
		return nil, fmt.Errorf("No available spots found in AST to place new node")
	}
	choosenSpot, err := pick.Pick(availableSpots)
	if err != nil {
		return nil, fmt.Errorf("picking one out of many available spots: %w", err)
	}
	ctx, err := ctxres.GetContextForSpot(
		astPkg,
		traverse.GetTraversableNodeForASTNode(astFuncDecl),
		choosenSpot,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to resolve context for choosen spot: %w", err)
	}
	newNode, err := appendRandomly(choosenSpot, ctx, depthLimit)
	if err != nil {
		return nil, fmt.Errorf("Failed to append a random node into choosen spot: %w", err)
	}
	return newNode, nil
}

func Develop(params *models.MutationParameters) error {
	for attempts := 0; attempts < 50; attempts++ {
		_, err := develop(params.Package, params.File, params.FuncDecl, 1)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("limit reached")
}
