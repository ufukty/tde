package genetics

import (
	"fmt"
	"go/ast"
	"reflect"
	"tde/internal/astw/types"
)

type field struct {
	ptr      any            // eg. ast.Stmt, *ast.BlockStmt, *[]ast.Stmt
	expected types.NodeType // this is the field's type, not the assigned value's type
}

// n can be a `ast.Node`, or a slice that is connected to a `ast.Node`
func children(n any) []field {

	if n == nil || isNil(n) {
		return []field{}
	}

	switch m := n.(type) {

	// case *ast.Package:
	// case *ast.Comment:
	// case *ast.CommentGroup:
	// case *ast.BadExpr:
	// case *ast.Ident:
	// case *ast.BasicLit:
	// case *ast.BadStmt:
	// case *ast.EmptyStmt:
	// case *ast.BadDecl:

	// MARK: Concrete array

	case *[]*ast.Field:
		s := make([]field, len(*m))
		for _, mv := range *m {
			s = append(s, field{mv, types.Field})
		}
		return s

	case *[]*ast.Ident:
		s := make([]field, len(*m))
		for _, mv := range *m {
			s = append(s, field{mv, types.Ident})
		}
		return s

	// MARK: Interface array

	case *[]ast.Decl:
		s := make([]field, len(*m))
		for _, mv := range *m {
			s = append(s, field{mv, types.Decl})
		}
		return s

	case *[]ast.Expr:
		s := make([]field, len(*m))
		for _, mv := range *m {
			s = append(s, field{mv, types.Expr})
		}
		return s

	case *[]ast.Spec:
		s := make([]field, len(*m))
		for _, mv := range *m {
			s = append(s, field{mv, types.Spec})
		}
		return s

	case *[]ast.Stmt:
		s := make([]field, len(*m))
		for _, mv := range *m {
			s = append(s, field{mv, types.Stmt})
		}
		return s

	// MARK: Concrete types

	case *ast.Field:
		return []field{
			{&m.Names, types.IdentSlice},
			{m.Type, types.Expr},
			{m.Tag, types.BasicLit},
		}

	case *ast.FieldList:
		return []field{
			{&m.List, types.FieldSlice},
		}

	// Expressions

	case *ast.Ellipsis:
		return []field{
			{m.Elt, types.Expr},
		}

	case *ast.FuncLit:
		return []field{
			{m.Type, types.FuncType},
			{m.Body, types.BlockStmt},
		}

	case *ast.CompositeLit:
		return []field{
			{m.Type, types.Expr},
			{&m.Elts, types.ExprSlice},
		}

	case *ast.ParenExpr:
		return []field{
			{m.X, types.Expr},
		}

	case *ast.SelectorExpr:
		return []field{
			{m.X, types.Expr},
			{m.Sel, types.Ident},
		}

	case *ast.IndexExpr:
		return []field{
			{m.X, types.Expr},
			{m.Index, types.Expr},
		}

	case *ast.IndexListExpr:
		return []field{
			{m.X, types.Expr},
			{&m.Indices, types.ExprSlice},
		}

	case *ast.SliceExpr:
		return []field{
			{m.X, types.Expr},
			{m.Low, types.Expr},
			{m.High, types.Expr},
			{m.Max, types.Expr},
		}

	case *ast.TypeAssertExpr:
		return []field{
			{m.X, types.Expr},
			{m.Type, types.Expr},
		}

	case *ast.CallExpr:
		return []field{
			{m.Fun, types.Expr},
			{&m.Args, types.ExprSlice},
		}

	case *ast.StarExpr:
		return []field{
			{m.X, types.Expr},
		}

	case *ast.UnaryExpr:
		return []field{
			{m.X, types.Expr},
		}

	case *ast.BinaryExpr:
		return []field{
			{m.X, types.Expr},
			{m.Y, types.Expr},
		}

	case *ast.KeyValueExpr:
		return []field{
			{m.Key, types.Expr},
			{m.Value, types.Expr},
		}
	// Types
	case *ast.ArrayType:
		return []field{
			{m.Len, types.Expr},
			{m.Elt, types.Expr},
		}

	case *ast.StructType:
		return []field{
			{m.Fields, types.FieldList},
		}

	case *ast.FuncType:
		return []field{
			{m.TypeParams, types.FieldList},
			{m.Params, types.FieldList},
			{m.Results, types.FieldList},
		}

	case *ast.InterfaceType:
		return []field{
			{m.Methods, types.FieldList},
		}

	case *ast.MapType:
		return []field{
			{m.Key, types.Expr},
			{m.Value, types.Expr},
		}

	case *ast.ChanType:
		return []field{
			{m.Value, types.Expr},
		}

	// Statements

	case *ast.DeclStmt:
		return []field{
			{m.Decl, types.Decl},
		}

	case *ast.LabeledStmt:
		return []field{
			{m.Label, types.Ident},
			{m.Stmt, types.Stmt},
		}

	case *ast.ExprStmt:
		return []field{
			{m.X, types.Expr},
		}

	case *ast.SendStmt:
		return []field{
			{m.Chan, types.Expr},
			{m.Value, types.Expr},
		}

	case *ast.IncDecStmt:
		return []field{
			{m.X, types.Expr},
		}

	case *ast.AssignStmt:
		return []field{
			{&m.Lhs, types.ExprSlice},
			{&m.Rhs, types.ExprSlice},
		}

	case *ast.GoStmt:
		return []field{
			{m.Call, types.CallExpr},
		}

	case *ast.DeferStmt:
		return []field{
			{m.Call, types.CallExpr},
		}

	case *ast.ReturnStmt:
		return []field{
			{&m.Results, types.ExprSlice},
		}

	case *ast.BranchStmt:
		return []field{
			{m.Label, types.Ident},
		}

	case *ast.BlockStmt:
		return []field{
			{&m.List, types.StmtSlice},
		}

	case *ast.IfStmt:
		return []field{
			{m.Init, types.Stmt},
			{m.Cond, types.Expr},
			{m.Body, types.BlockStmt},
			{m.Else, types.Stmt},
		}

	case *ast.CaseClause:
		return []field{
			{&m.List, types.ExprSlice},
			{&m.Body, types.StmtSlice},
		}

	case *ast.SwitchStmt:
		return []field{
			{m.Init, types.Stmt},
			{m.Tag, types.Expr},
			{m.Body, types.BlockStmt},
		}

	case *ast.TypeSwitchStmt:
		return []field{
			{m.Init, types.Stmt},
			{m.Assign, types.Stmt},
			{m.Body, types.BlockStmt},
		}

	case *ast.CommClause:
		return []field{
			{m.Comm, types.Stmt},
			{&m.Body, types.StmtSlice},
		}

	case *ast.SelectStmt:
		return []field{
			{m.Body, types.BlockStmt},
		}

	case *ast.ForStmt:
		return []field{
			{m.Init, types.Stmt},
			{m.Cond, types.Expr},
			{m.Post, types.Stmt},
			{m.Body, types.BlockStmt},
		}

	case *ast.RangeStmt:
		return []field{
			{m.Key, types.Expr},
			{m.Value, types.Expr},
			{m.X, types.Expr},
			{m.Body, types.BlockStmt},
		}

	// Declarations

	case *ast.ImportSpec:
		return []field{
			{m.Name, types.Ident},
			{m.Path, types.BasicLit},
		}

	case *ast.ValueSpec:
		return []field{
			{&m.Names, types.IdentSlice},
			{m.Type, types.Expr},
			{&m.Values, types.ExprSlice},
		}

	case *ast.TypeSpec:
		return []field{
			{m.Name, types.Ident},
			{m.TypeParams, types.FieldList},
			{m.Type, types.Expr},
		}

	case *ast.GenDecl:
		return []field{
			{&m.Specs, types.SpecSlice},
		}

	case *ast.FuncDecl:
		return []field{
			{m.Recv, types.FieldList},
			{m.Name, types.Ident},
			{m.Type, types.FuncType},
			{m.Body, types.BlockStmt},
		}

	// Files and packages

	case *ast.File:
		return []field{
			{m.Name, types.Ident},
			{&m.Decls, types.DeclSlice},
		}

	default:
		panic(fmt.Sprintf("Apply: unexpected node type %T", reflect.TypeOf(m)))
	}
}
