package traverse

import (
	ast_types "tde/internal/astw/types"

	"go/ast"
)

type NodeField struct {
	Value any // could be instance of ast.Node or []ast.Node, call NodeField.Type.IsSlice() to distinguish
	Type  ast_types.NodeType
}

// NodeFields don't cover every "NodeType", but only structs that comply ast.Node
func NodeFields(n ast.Node) []NodeField {
	switch n := n.(type) {

	case *ast.Comment:
		return []NodeField{}

	case *ast.CommentGroup:
		return []NodeField{}

	case *ast.Field:
		return []NodeField{
			{n.Names, ast_types.IdentSlice},
			{n.Type, ast_types.Expr},
			{n.Tag, ast_types.BasicLit},
		}

	case *ast.FieldList:
		return []NodeField{
			{n.List, ast_types.FieldSlice},
		}

	case *ast.BadExpr:
		return []NodeField{}

	case *ast.Ident:
		return []NodeField{}

	case *ast.Ellipsis:
		return []NodeField{
			{n.Elt, ast_types.Expr},
		}

	case *ast.BasicLit:
		return []NodeField{}

	case *ast.FuncLit:
		return []NodeField{
			{n.Type, ast_types.FuncType},
			{n.Body, ast_types.BlockStmt},
		}

	case *ast.CompositeLit:
		return []NodeField{
			{n.Type, ast_types.Expr},
			{n.Elts, ast_types.ExprSlice},
		}

	case *ast.ParenExpr:
		return []NodeField{
			{n.X, ast_types.Expr},
		}

	case *ast.SelectorExpr:
		return []NodeField{
			{n.X, ast_types.Expr},
			{n.Sel, ast_types.Ident},
		}

	case *ast.IndexExpr:
		return []NodeField{
			{n.X, ast_types.Expr},
			{n.Index, ast_types.Expr},
		}

	case *ast.IndexListExpr:
		return []NodeField{
			{n.X, ast_types.Expr},
			{n.Indices, ast_types.ExprSlice},
		}

	case *ast.SliceExpr:
		return []NodeField{
			{n.X, ast_types.Expr},
			{n.Low, ast_types.Expr},
			{n.High, ast_types.Expr},
			{n.Max, ast_types.Expr},
		}

	case *ast.TypeAssertExpr:
		return []NodeField{
			{n.X, ast_types.Expr},
			{n.Type, ast_types.Expr},
		}

	case *ast.CallExpr:
		return []NodeField{
			{n.Fun, ast_types.Expr},
			{n.Args, ast_types.ExprSlice},
		}

	case *ast.StarExpr:
		return []NodeField{
			{n.X, ast_types.Expr},
		}

	case *ast.UnaryExpr:
		return []NodeField{
			{n.X, ast_types.Expr},
		}

	case *ast.BinaryExpr:
		return []NodeField{
			{n.X, ast_types.Expr},
			{n.Y, ast_types.Expr},
		}

	case *ast.KeyValueExpr:
		return []NodeField{
			{n.Key, ast_types.Expr},
			{n.Value, ast_types.Expr},
		}

	case *ast.ArrayType:
		return []NodeField{
			{n.Len, ast_types.Expr},
			{n.Elt, ast_types.Expr},
		}

	case *ast.StructType:
		return []NodeField{
			{n.Fields, ast_types.FieldList},
		}

	case *ast.FuncType:
		return []NodeField{
			{n.TypeParams, ast_types.FieldList},
			{n.Params, ast_types.FieldList},
			{n.Results, ast_types.FieldList},
		}

	case *ast.InterfaceType:
		return []NodeField{
			{n.Methods, ast_types.FieldList},
		}

	case *ast.MapType:
		return []NodeField{
			{n.Key, ast_types.Expr},
			{n.Value, ast_types.Expr},
		}

	case *ast.ChanType:
		return []NodeField{
			{n.Value, ast_types.Expr},
		}

	case *ast.BadStmt:
		return []NodeField{}

	case *ast.DeclStmt:
		return []NodeField{
			{n.Decl, ast_types.Decl},
		}

	case *ast.EmptyStmt:
		return []NodeField{}

	case *ast.LabeledStmt:
		return []NodeField{
			{n.Label, ast_types.Ident},
			{n.Stmt, ast_types.Stmt},
		}

	case *ast.ExprStmt:
		return []NodeField{
			{n.X, ast_types.Expr},
		}

	case *ast.SendStmt:
		return []NodeField{
			{n.Chan, ast_types.Expr},
			{n.Value, ast_types.Expr},
		}

	case *ast.IncDecStmt:
		return []NodeField{
			{n.X, ast_types.Expr},
		}

	case *ast.AssignStmt:
		return []NodeField{
			{n.Lhs, ast_types.ExprSlice},
			{n.Rhs, ast_types.ExprSlice},
		}

	case *ast.GoStmt:
		return []NodeField{
			{n.Call, ast_types.CallExpr},
		}

	case *ast.DeferStmt:
		return []NodeField{
			{n.Call, ast_types.CallExpr},
		}

	case *ast.ReturnStmt:
		return []NodeField{
			{n.Results, ast_types.ExprSlice},
		}

	case *ast.BranchStmt:
		return []NodeField{
			{n.Label, ast_types.Ident},
		}

	case *ast.BlockStmt:
		return []NodeField{
			{n.List, ast_types.StmtSlice},
		}

	case *ast.IfStmt:
		return []NodeField{
			{n.Init, ast_types.Stmt},
			{n.Cond, ast_types.Expr},
			{n.Body, ast_types.BlockStmt},
			{n.Else, ast_types.Stmt},
		}

	case *ast.CaseClause:
		return []NodeField{
			{n.List, ast_types.ExprSlice},
			{n.Body, ast_types.StmtSlice},
		}

	case *ast.SwitchStmt:
		return []NodeField{
			{n.Init, ast_types.Stmt},
			{n.Tag, ast_types.Expr},
			{n.Body, ast_types.BlockStmt},
		}

	case *ast.TypeSwitchStmt:
		return []NodeField{
			{n.Init, ast_types.Stmt},
			{n.Assign, ast_types.Stmt},
			{n.Body, ast_types.BlockStmt},
		}

	case *ast.CommClause:
		return []NodeField{
			{n.Comm, ast_types.Stmt},
			{n.Body, ast_types.StmtSlice},
		}

	case *ast.SelectStmt:
		return []NodeField{
			{n.Body, ast_types.BlockStmt},
		}

	case *ast.ForStmt:
		return []NodeField{
			{n.Init, ast_types.Stmt},
			{n.Cond, ast_types.Expr},
			{n.Post, ast_types.Stmt},
			{n.Body, ast_types.BlockStmt},
		}

	case *ast.RangeStmt:
		return []NodeField{
			{n.Key, ast_types.Expr},
			{n.Value, ast_types.Expr},
			{n.X, ast_types.Expr},
			{n.Body, ast_types.BlockStmt},
		}

	case *ast.ImportSpec:
		return []NodeField{
			{n.Name, ast_types.Ident},
			{n.Path, ast_types.BasicLit},
		}

	case *ast.ValueSpec:
		return []NodeField{
			{n.Names, ast_types.IdentSlice},
			{n.Type, ast_types.Expr},
			{n.Values, ast_types.ExprSlice},
		}

	case *ast.TypeSpec:
		return []NodeField{
			{n.Name, ast_types.Ident},
			{n.TypeParams, ast_types.FieldList},
			{n.Type, ast_types.Expr},
		}

	case *ast.BadDecl:
		return []NodeField{}

	case *ast.GenDecl:
		return []NodeField{
			{n.Specs, ast_types.SpecSlice},
		}

	case *ast.FuncDecl:
		return []NodeField{
			{n.Recv, ast_types.FieldList},
			{n.Name, ast_types.Ident},
			{n.Type, ast_types.FuncType},
			{n.Body, ast_types.BlockStmt},
		}

	case *ast.File:
		return []NodeField{
			{n.Name, ast_types.Ident},
			{n.Decls, ast_types.DeclSlice},
			{n.Imports, ast_types.ImportSpecSlice},
			{n.Unresolved, ast_types.IdentSlice},
			// {n.Comments, types.CommentGroupSlice},
		}

		// case *ast.Package:
		// 	return []NodeField{
		// 	{n.Files, map[string]*types.File},
		// }
	}
	return nil
}
