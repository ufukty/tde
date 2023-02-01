package test_package

import "go/ast"

type NodeField struct {
	Value any // could be instance of ast.Node or []ast.Node, call NodeField.Type.IsSlice() to distinguish
	Type  NodeType
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
			{n.Names, IdentSlice},
			{n.Type, Expr},
			{n.Tag, BasicLit},
		}

	case *ast.FieldList:
		return []NodeField{
			{n.List, FieldSlice},
		}

	case *ast.BadExpr:
		return []NodeField{}

	case *ast.Ident:
		return []NodeField{}

	case *ast.Ellipsis:
		return []NodeField{
			{n.Elt, Expr},
		}

	case *ast.BasicLit:
		return []NodeField{}

	case *ast.FuncLit:
		return []NodeField{
			{n.Type, FuncType},
			{n.Body, BlockStmt},
		}

	case *ast.CompositeLit:
		return []NodeField{
			{n.Type, Expr},
			{n.Elts, ExprSlice},
		}

	case *ast.ParenExpr:
		return []NodeField{
			{n.X, Expr},
		}

	case *ast.SelectorExpr:
		return []NodeField{
			{n.X, Expr},
			{n.Sel, Ident},
		}

	case *ast.IndexExpr:
		return []NodeField{
			{n.X, Expr},
			{n.Index, Expr},
		}

	case *ast.IndexListExpr:
		return []NodeField{
			{n.X, Expr},
			{n.Indices, ExprSlice},
		}

	case *ast.SliceExpr:
		return []NodeField{
			{n.X, Expr},
			{n.Low, Expr},
			{n.High, Expr},
			{n.Max, Expr},
		}

	case *ast.TypeAssertExpr:
		return []NodeField{
			{n.X, Expr},
			{n.Type, Expr},
		}

	case *ast.CallExpr:
		return []NodeField{
			{n.Fun, Expr},
			{n.Args, ExprSlice},
		}

	case *ast.StarExpr:
		return []NodeField{
			{n.X, Expr},
		}

	case *ast.UnaryExpr:
		return []NodeField{
			{n.X, Expr},
		}

	case *ast.BinaryExpr:
		return []NodeField{
			{n.X, Expr},
			{n.Y, Expr},
		}

	case *ast.KeyValueExpr:
		return []NodeField{
			{n.Key, Expr},
			{n.Value, Expr},
		}

	case *ast.ArrayType:
		return []NodeField{
			{n.Len, Expr},
			{n.Elt, Expr},
		}

	case *ast.StructType:
		return []NodeField{
			{n.Fields, FieldList},
		}

	case *ast.FuncType:
		return []NodeField{
			{n.TypeParams, FieldList},
			{n.Params, FieldList},
			{n.Results, FieldList},
		}

	case *ast.InterfaceType:
		return []NodeField{
			{n.Methods, FieldList},
		}

	case *ast.MapType:
		return []NodeField{
			{n.Key, Expr},
			{n.Value, Expr},
		}

	case *ast.ChanType:
		return []NodeField{
			{n.Value, Expr},
		}

	case *ast.BadStmt:
		return []NodeField{}

	case *ast.DeclStmt:
		return []NodeField{
			{n.Decl, Decl},
		}

	case *ast.EmptyStmt:
		return []NodeField{}

	case *ast.LabeledStmt:
		return []NodeField{
			{n.Label, Ident},
			{n.Stmt, Stmt},
		}

	case *ast.ExprStmt:
		return []NodeField{
			{n.X, Expr},
		}

	case *ast.SendStmt:
		return []NodeField{
			{n.Chan, Expr},
			{n.Value, Expr},
		}

	case *ast.IncDecStmt:
		return []NodeField{
			{n.X, Expr},
		}

	case *ast.AssignStmt:
		return []NodeField{
			{n.Lhs, ExprSlice},
			{n.Rhs, ExprSlice},
		}

	case *ast.GoStmt:
		return []NodeField{
			{n.Call, CallExpr},
		}

	case *ast.DeferStmt:
		return []NodeField{
			{n.Call, CallExpr},
		}

	case *ast.ReturnStmt:
		return []NodeField{
			{n.Results, ExprSlice},
		}

	case *ast.BranchStmt:
		return []NodeField{
			{n.Label, Ident},
		}

	case *ast.BlockStmt:
		return []NodeField{
			{n.List, StmtSlice},
		}

	case *ast.IfStmt:
		return []NodeField{
			{n.Init, Stmt},
			{n.Cond, Expr},
			{n.Body, BlockStmt},
			{n.Else, Stmt},
		}

	case *ast.CaseClause:
		return []NodeField{
			{n.List, ExprSlice},
			{n.Body, StmtSlice},
		}

	case *ast.SwitchStmt:
		return []NodeField{
			{n.Init, Stmt},
			{n.Tag, Expr},
			{n.Body, BlockStmt},
		}

	case *ast.TypeSwitchStmt:
		return []NodeField{
			{n.Init, Stmt},
			{n.Assign, Stmt},
			{n.Body, BlockStmt},
		}

	case *ast.CommClause:
		return []NodeField{
			{n.Comm, Stmt},
			{n.Body, StmtSlice},
		}

	case *ast.SelectStmt:
		return []NodeField{
			{n.Body, BlockStmt},
		}

	case *ast.ForStmt:
		return []NodeField{
			{n.Init, Stmt},
			{n.Cond, Expr},
			{n.Post, Stmt},
			{n.Body, BlockStmt},
		}

	case *ast.RangeStmt:
		return []NodeField{
			{n.Key, Expr},
			{n.Value, Expr},
			{n.X, Expr},
			{n.Body, BlockStmt},
		}

	case *ast.ImportSpec:
		return []NodeField{
			{n.Name, Ident},
			{n.Path, BasicLit},
		}

	case *ast.ValueSpec:
		return []NodeField{
			{n.Names, IdentSlice},
			{n.Type, Expr},
			{n.Values, ExprSlice},
		}

	case *ast.TypeSpec:
		return []NodeField{
			{n.Name, Ident},
			{n.TypeParams, FieldList},
			{n.Type, Expr},
		}

	case *ast.BadDecl:
		return []NodeField{}

	case *ast.GenDecl:
		return []NodeField{
			{n.Specs, SpecSlice},
		}

	case *ast.FuncDecl:
		return []NodeField{
			{n.Recv, FieldList},
			{n.Name, Ident},
			{n.Type, FuncType},
			{n.Body, BlockStmt},
		}

	case *ast.File:
		return []NodeField{
			{n.Name, Ident},
			{n.Decls, DeclSlice},
			{n.Imports, ImportSpecSlice},
			{n.Unresolved, IdentSlice},
			// {n.Comments, CommentGroupSlice},
		}

		// case *ast.Package:
		// 	return []NodeField{
		// 	{n.Files, map[string]*File},
		// }
	}
	return nil
}
