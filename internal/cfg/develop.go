package cfg

import (
	"fmt"
	"reflect"
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

func appendRandomly(tNode *trav.TraversableNode, ctx *context.Context) {

	if tNode.ExpectedType.IsSliceType() {
		switch tNode.Value.(type) {
		case []*ast.ImportSpec:
			createdNode := nc.ImportSpec(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case []*ast.Ident:
			createdNode := nc.Ident(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case []*ast.Field:
			createdNode := nc.Field(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case []ast.Stmt:
			createdNode := nc.Stmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case []ast.Decl:
			createdNode := nc.Decl(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case []ast.Spec:
			createdNode := nc.Spec(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case []ast.Expr:
			createdNode := nc.Expr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		}

	} else if tNode.ExpectedType.IsInterfaceType() {
		switch tNode.ExpectedType {
		case ast_types.Expr:
			createdNode := nc.Expr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.Stmt:
			createdNode := nc.Stmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.Decl:
			createdNode := nc.Decl(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.Spec:
			createdNode := nc.Spec(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.TypeExpr:
			createdNode := nc.Type(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		}

	} else if tNode.ExpectedType.IsConcreteType() {

		switch tNode.ExpectedType {
		// case ast_types.BadDecl:
		// createdNode:=nc.BadDecl(ctx, 1)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.BadExpr:
		// createdNode:=nc.BadExpr(ctx, 1)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.BadStmt:
		// createdNode:=nc.BadStmt(ctx, 1)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.Comment:
		// createdNode:=nc.Comment(ctx, 1)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.CommentGroup:
		// createdNode:=nc.CommentGroup(ctx, 1)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.File:
		// createdNode:=nc.File(ctx, 1)
		// 	tNode.Ref.Set(createdNode)
		// case ast_types.Package:
		// createdNode:=nc.Package(ctx, 1)
		// 	tNode.Ref.Set(createdNode)
		case ast_types.ArrayType:
			createdNode := nc.ArrayType(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.AssignStmt:
			createdNode := nc.AssignStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.BasicLit:
			createdNode := nc.BasicLit(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.BinaryExpr:
			createdNode := nc.BinaryExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.BlockStmt:
			createdNode := nc.BlockStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.BranchStmt:
			createdNode := nc.BranchStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.CallExpr:
			createdNode := nc.CallExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.CaseClause:
			createdNode := nc.CaseClause(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.ChanType:
			createdNode := nc.ChanType(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.CommClause:
			createdNode := nc.CommClause(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.CompositeLit:
			createdNode := nc.CompositeLit(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.DeclStmt:
			createdNode := nc.DeclStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.DeferStmt:
			createdNode := nc.DeferStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.Ellipsis:
			createdNode := nc.Ellipsis(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.EmptyStmt:
			createdNode := nc.EmptyStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.ExprStmt:
			createdNode := nc.ExprStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.Field:
			createdNode := nc.Field(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.FieldList:
			createdNode := nc.FieldList(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.ForStmt:
			createdNode := nc.ForStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.FuncDecl:
			createdNode := nc.FuncDecl(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.FuncLit:
			createdNode := nc.FuncLit(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.FuncType:
			createdNode := nc.FuncType(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.GenDecl:
			createdNode := nc.GenDecl(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.GoStmt:
			createdNode := nc.GoStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.Ident:
			createdNode := nc.Ident(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.IfStmt:
			createdNode := nc.IfStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.ImportSpec:
			createdNode := nc.ImportSpec(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.IncDecStmt:
			createdNode := nc.IncDecStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.IndexExpr:
			createdNode := nc.IndexExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.IndexListExpr:
			createdNode := nc.IndexListExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.InterfaceType:
			createdNode := nc.InterfaceType(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.KeyValueExpr:
			createdNode := nc.KeyValueExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.LabeledStmt:
			createdNode := nc.LabeledStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.MapType:
			createdNode := nc.MapType(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.ParenExpr:
			createdNode := nc.ParenExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.RangeStmt:
			createdNode := nc.RangeStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.ReturnStmt:
			createdNode := nc.ReturnStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.SelectorExpr:
			createdNode := nc.SelectorExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.SelectStmt:
			createdNode := nc.SelectStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.SendStmt:
			createdNode := nc.SendStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.SliceExpr:
			createdNode := nc.SliceExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.StarExpr:
			createdNode := nc.StarExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.StructType:
			createdNode := nc.StructType(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.SwitchStmt:
			createdNode := nc.SwitchStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.TypeAssertExpr:
			createdNode := nc.TypeAssertExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.TypeSpec:
			createdNode := nc.TypeSpec(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.TypeSwitchStmt:
			createdNode := nc.TypeSwitchStmt(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.UnaryExpr:
			createdNode := nc.UnaryExpr(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		case ast_types.ValueSpec:
			createdNode := nc.ValueSpec(ctx, 1)
			fmt.Println(reflect.ValueOf(createdNode))
			tNode.Ref.Set(createdNode)
		}

	}
}

// Picks an appandable spot randomly (either a nil field or end of a slice type field)
// Creates just one node and appends to choosen spot
func Develop(astPkg *ast.Package, astFile *ast.File, astFuncDecl *ast.FuncDecl) error {
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
	appendRandomly(choosenSpot, ctx)
	return nil
}
