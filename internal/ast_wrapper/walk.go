package ast_wrapper

import (
	"go/ast"
	"reflect"

	"log"
)

func increaseLastChildIndex(childIndexTrace []int) {
	childIndexTrace[len(childIndexTrace)-1]++
}

// This mess is because simple comparison of n == nil doesn't work.
func isNodeNil(n ast.Node) bool {
	switch n := n.(type) {
	case *ast.ArrayType:
		return n == nil
	case *ast.AssignStmt:
		return n == nil
	case *ast.BadDecl:
		return n == nil
	case *ast.BadExpr:
		return n == nil
	case *ast.BadStmt:
		return n == nil
	case *ast.BasicLit:
		return n == nil
	case *ast.BinaryExpr:
		return n == nil
	case *ast.BlockStmt:
		return n == nil
	case *ast.BranchStmt:
		return n == nil
	case *ast.CallExpr:
		return n == nil
	case *ast.CaseClause:
		return n == nil
	case *ast.ChanType:
		return n == nil
	case *ast.CommClause:
		return n == nil
	case *ast.Comment:
		return n == nil
	case *ast.CommentGroup:
		return n == nil
	case *ast.CompositeLit:
		return n == nil
	case *ast.DeclStmt:
		return n == nil
	case *ast.DeferStmt:
		return n == nil
	case *ast.Ellipsis:
		return n == nil
	case *ast.EmptyStmt:
		return n == nil
	case *ast.ExprStmt:
		return n == nil
	case *ast.Field:
		return n == nil
	case *ast.FieldList:
		return n == nil
	case *ast.File:
		return n == nil
	case *ast.ForStmt:
		return n == nil
	case *ast.FuncDecl:
		return n == nil
	case *ast.FuncLit:
		return n == nil
	case *ast.FuncType:
		return n == nil
	case *ast.GenDecl:
		return n == nil
	case *ast.GoStmt:
		return n == nil
	case *ast.Ident:
		return n == nil
	case *ast.IfStmt:
		return n == nil
	case *ast.ImportSpec:
		return n == nil
	case *ast.IncDecStmt:
		return n == nil
	case *ast.IndexExpr:
		return n == nil
	case *ast.IndexListExpr:
		return n == nil
	case *ast.InterfaceType:
		return n == nil
	case *ast.KeyValueExpr:
		return n == nil
	case *ast.LabeledStmt:
		return n == nil
	case *ast.MapType:
		return n == nil
	case *ast.Package:
		return n == nil
	case *ast.ParenExpr:
		return n == nil
	case *ast.RangeStmt:
		return n == nil
	case *ast.ReturnStmt:
		return n == nil
	case *ast.SelectorExpr:
		return n == nil
	case *ast.SelectStmt:
		return n == nil
	case *ast.SendStmt:
		return n == nil
	case *ast.SliceExpr:
		return n == nil
	case *ast.StarExpr:
		return n == nil
	case *ast.StructType:
		return n == nil
	case *ast.SwitchStmt:
		return n == nil
	case *ast.TypeAssertExpr:
		return n == nil
	case *ast.TypeSpec:
		return n == nil
	case *ast.TypeSwitchStmt:
		return n == nil
	case *ast.UnaryExpr:
		return n == nil
	case *ast.ValueSpec:
		return n == nil
	}
	return true
}

type WalkCallbackFunction func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool

func walkAstTypeFieldsIfSet(parentTrace []ast.Node, childIndexTrace *[]int, callback WalkCallbackFunction, passUnassignedFields bool, vars ...any) {
	for _, field := range vars {
		switch field := field.(type) {
		case []*ast.Comment:
			for _, item := range field {
				walkHelper(item, passUnassignedFields, parentTrace, *childIndexTrace, callback)
				increaseLastChildIndex(*childIndexTrace)
			}
		case []*ast.CommentGroup:
			for _, item := range field {
				walkHelper(item, passUnassignedFields, parentTrace, *childIndexTrace, callback)
				increaseLastChildIndex(*childIndexTrace)
			}
		case []*ast.ImportSpec:
			for _, item := range field {
				walkHelper(item, passUnassignedFields, parentTrace, *childIndexTrace, callback)
				increaseLastChildIndex(*childIndexTrace)
			}
		case []*ast.Ident:
			for _, item := range field {
				walkHelper(item, passUnassignedFields, parentTrace, *childIndexTrace, callback)
				increaseLastChildIndex(*childIndexTrace)
			}
		case []*ast.Field:
			for _, item := range field {
				walkHelper(item, passUnassignedFields, parentTrace, *childIndexTrace, callback)
				increaseLastChildIndex(*childIndexTrace)
			}
		case []ast.Stmt:
			for _, item := range field {
				walkHelper(item, passUnassignedFields, parentTrace, *childIndexTrace, callback)
				increaseLastChildIndex(*childIndexTrace)
			}
		case []ast.Decl:
			for _, item := range field {
				walkHelper(item, passUnassignedFields, parentTrace, *childIndexTrace, callback)
				increaseLastChildIndex(*childIndexTrace)
			}
		case []ast.Spec:
			for _, item := range field {
				walkHelper(item, passUnassignedFields, parentTrace, *childIndexTrace, callback)
				increaseLastChildIndex(*childIndexTrace)
			}
		case []ast.Expr:
			for _, item := range field {
				walkHelper(item, passUnassignedFields, parentTrace, *childIndexTrace, callback)
				increaseLastChildIndex(*childIndexTrace)
			}
		case []ast.Node:
			for _, item := range field {
				walkHelper(item, passUnassignedFields, parentTrace, *childIndexTrace, callback)
				increaseLastChildIndex(*childIndexTrace)
			}
		case ast.Node:
			walkHelper(field, passUnassignedFields, parentTrace, *childIndexTrace, callback)
			increaseLastChildIndex(*childIndexTrace)
		default:
			log.Print("Field type is not covered by switch: ", reflect.TypeOf(field).String())
		}
	}
}

func walkHelper(n ast.Node, passUnassignedFields bool, parentTrace []ast.Node, childIndexTrace []int, callback WalkCallbackFunction) {

	if isNodeNil(n) {
		if !passUnassignedFields {
			callback(n, parentTrace, childIndexTrace)
		}
		return
	} else if !callback(n, parentTrace, childIndexTrace) {
		return
	}

	parentTrace = append(parentTrace, n)
	childIndexTrace = append(childIndexTrace, 0)

	switch n := n.(type) {

	// Comments and fields

	case *ast.Comment:
		// nothing to do
	case *ast.CommentGroup:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.List)
	case *ast.Field:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Names, n.Type, n.Tag)
	case *ast.FieldList:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.List)

	// Expressions

	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
		// nothing to do
	case *ast.Ellipsis:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Elt)
	case *ast.FuncLit:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Type, n.Body)
	case *ast.CompositeLit:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Type, n.Elts)
	case *ast.ParenExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X)
	case *ast.SelectorExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X, n.Sel)
	case *ast.IndexExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X, n.Index)
	case *ast.IndexListExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X, n.Indices)
	case *ast.SliceExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X, n.Low, n.High, n.Max)
	case *ast.TypeAssertExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X, n.Type)
	case *ast.CallExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Fun, n.Args)
	case *ast.StarExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X)
	case *ast.UnaryExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X)
	case *ast.BinaryExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X, n.Y)
	case *ast.KeyValueExpr:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Key, n.Value)

	// Types

	case *ast.ArrayType:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Len, n.Elt)
	case *ast.StructType:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Fields)
	case *ast.FuncType:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.TypeParams, n.Params, n.Results)
	case *ast.InterfaceType:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Methods)
	case *ast.MapType:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Key, n.Value)
	case *ast.ChanType:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Value)

	// Statements

	case *ast.BadStmt:
		// nothing to do
	case *ast.DeclStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Decl)
	case *ast.EmptyStmt:
		// nothing to do
	case *ast.LabeledStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Label, n.Stmt)
	case *ast.ExprStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X)
	case *ast.SendStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Chan, n.Value)
	case *ast.IncDecStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.X)
	case *ast.AssignStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Lhs, n.Rhs)
	case *ast.GoStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Call)
	case *ast.DeferStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Call)
	case *ast.ReturnStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Results)
	case *ast.BranchStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Label)
	case *ast.BlockStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.List)
	case *ast.IfStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Init, n.Cond, n.Body, n.Else)
	case *ast.CaseClause:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.List, n.Body)
	case *ast.SwitchStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Init, n.Tag, n.Body)
	case *ast.TypeSwitchStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Init, n.Assign, n.Body)
	case *ast.CommClause:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Comm, n.Body)
	case *ast.SelectStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Body)
	case *ast.ForStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Init, n.Cond, n.Post, n.Body)
	case *ast.RangeStmt:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Key, n.Value, n.X, n.Body)

	// Specifications

	case *ast.ImportSpec:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Name, n.Path)
	case *ast.ValueSpec:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Names, n.Type, n.Values)
	case *ast.TypeSpec:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Name, n.TypeParams, n.Type)
	case *ast.BadDecl:
		// nothing to do

	// Declarations

	case *ast.GenDecl:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Specs)
	case *ast.FuncDecl:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Recv, n.Name, n.Type, n.Body)

	// File & Package

	case *ast.File:
		walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, n.Name, n.Decls)
	case *ast.Package:
		for _, f := range n.Files {
			walkAstTypeFieldsIfSet(parentTrace, &childIndexTrace, callback, passUnassignedFields, f.Name, f.Decls)
		}
	}
}

// Walk recursively visits every ast.Node compliant type nodes of the AST, calls
// the callback function once per node, optionally can pass
// unassigned (nil) fields of nodes when lists children of them
func Walk(root ast.Node, passUnassignedFields bool, callback WalkCallbackFunction) {
	walkHelper(root, passUnassignedFields, []ast.Node{}, []int{}, callback)
}
