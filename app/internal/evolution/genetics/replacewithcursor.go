package genetics

import (
	"fmt"
	"go/ast"
	"reflect"
)

// (n)ext
func replaceOnParentWithCursor(c cursor, n any) error {

	switch p := c.parent.(type) {

	// case *ast.Package:
	// case *ast.Comment:
	// case *ast.CommentGroup:
	// case *ast.BadExpr:
	// case *ast.Ident:
	// case *ast.BasicLit:
	// case *ast.BadStmt:
	// case *ast.EmptyStmt:
	// case *ast.BadDecl:

	case *ast.Field:
		switch c.fi {
		case 0:
			if n, ok := n.([]*ast.Ident); ok {
				p.Names = n
				return nil
			}
			return fmt.Errorf("expected \"[]*ast.Ident\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 2:
			if n, ok := n.(*ast.BasicLit); ok {
				p.Tag = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.BasicLit\" got %T", n)
		}

	case *ast.FieldList:
		switch c.fi {
		case 0:
			if n, ok := n.([]*ast.Field); ok {
				p.List = n
				return nil
			}
			return fmt.Errorf("expected \"[]*ast.Field\" got %T", n)
		}

	// Expressions

	case *ast.Ellipsis:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Elt = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.FuncLit:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.FuncType); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.FuncType\" got %T", n)
		case 1:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.BlockStmt\" got %T", n)
		}

	case *ast.CompositeLit:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.([]ast.Expr); ok {
				p.Elts = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Expr\" got %T", n)
		}

	case *ast.ParenExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.SelectorExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.(*ast.Ident); ok {
				p.Sel = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.Ident\" got %T", n)
		}

	case *ast.IndexExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Index = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.IndexListExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.([]ast.Expr); ok {
				p.Indices = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Expr\" got %T", n)
		}

	case *ast.SliceExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Low = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 2:
			if n, ok := n.(ast.Expr); ok {
				p.High = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 3:
			if n, ok := n.(ast.Expr); ok {
				p.Max = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.TypeAssertExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.CallExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Fun = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.([]ast.Expr); ok {
				p.Args = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Expr\" got %T", n)
		}

	case *ast.StarExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.UnaryExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.BinaryExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Y = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.KeyValueExpr:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Key = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Value = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}
	// Types
	case *ast.ArrayType:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Len = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Elt = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.StructType:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.FieldList); ok {
				p.Fields = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.FieldList\" got %T", n)
		}

	case *ast.FuncType:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.FieldList); ok {
				p.TypeParams = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.FieldList\" got %T", n)
		case 1:
			if n, ok := n.(*ast.FieldList); ok {
				p.Params = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.FieldList\" got %T", n)
		case 2:
			if n, ok := n.(*ast.FieldList); ok {
				p.Results = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.FieldList\" got %T", n)
		}

	case *ast.InterfaceType:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.FieldList); ok {
				p.Methods = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.FieldList\" got %T", n)
		}

	case *ast.MapType:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Key = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Value = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.ChanType:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Value = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	// Statements

	case *ast.DeclStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Decl); ok {
				p.Decl = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Decl\" got %T", n)
		}

	case *ast.LabeledStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.Ident); ok {
				p.Label = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.Ident\" got %T", n)
		case 1:
			if n, ok := n.(ast.Stmt); ok {
				p.Stmt = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Stmt\" got %T", n)
		}

	case *ast.ExprStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.SendStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Chan = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Value = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.IncDecStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.AssignStmt:
		switch c.fi {
		case 0:
			if n, ok := n.([]ast.Expr); ok {
				p.Lhs = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.([]ast.Expr); ok {
				p.Rhs = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Expr\" got %T", n)
		}

	case *ast.GoStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.CallExpr); ok {
				p.Call = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.CallExpr\" got %T", n)
		}

	case *ast.DeferStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.CallExpr); ok {
				p.Call = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.CallExpr\" got %T", n)
		}

	case *ast.ReturnStmt:
		switch c.fi {
		case 0:
			if n, ok := n.([]ast.Expr); ok {
				p.Results = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Expr\" got %T", n)
		}

	case *ast.BranchStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.Ident); ok {
				p.Label = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.Ident\" got %T", n)
		}

	case *ast.BlockStmt:
		switch c.fi {
		case 0:
			if n, ok := n.([]ast.Stmt); ok {
				p.List = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Stmt\" got %T", n)
		}

	case *ast.IfStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Stmt); ok {
				p.Init = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Stmt\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Cond = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 2:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.BlockStmt\" got %T", n)
		case 3:
			if n, ok := n.(ast.Stmt); ok {
				p.Else = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Stmt\" got %T", n)
		}

	case *ast.CaseClause:
		switch c.fi {
		case 0:
			if n, ok := n.([]ast.Expr); ok {
				p.List = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.([]ast.Stmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Stmt\" got %T", n)
		}

	case *ast.SwitchStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Stmt); ok {
				p.Init = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Stmt\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Tag = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 2:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.BlockStmt\" got %T", n)
		}

	case *ast.TypeSwitchStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Stmt); ok {
				p.Init = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Stmt\" got %T", n)
		case 1:
			if n, ok := n.(ast.Stmt); ok {
				p.Assign = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Stmt\" got %T", n)
		case 2:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.BlockStmt\" got %T", n)
		}

	case *ast.CommClause:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Stmt); ok {
				p.Comm = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Stmt\" got %T", n)
		case 1:
			if n, ok := n.([]ast.Stmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Stmt\" got %T", n)
		}

	case *ast.SelectStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.BlockStmt\" got %T", n)
		}

	case *ast.ForStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Stmt); ok {
				p.Init = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Stmt\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Cond = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 2:
			if n, ok := n.(ast.Stmt); ok {
				p.Post = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Stmt\" got %T", n)
		case 3:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.BlockStmt\" got %T", n)
		}

	case *ast.RangeStmt:
		switch c.fi {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Key = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Value = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 2:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 3:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.BlockStmt\" got %T", n)
		}

	// Declarations

	case *ast.ImportSpec:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.Ident); ok {
				p.Name = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.Ident\" got %T", n)
		case 1:
			if n, ok := n.(*ast.BasicLit); ok {
				p.Path = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.BasicLit\" got %T", n)
		}

	case *ast.ValueSpec:
		switch c.fi {
		case 0:
			if n, ok := n.([]*ast.Ident); ok {
				p.Names = n
				return nil
			}
			return fmt.Errorf("expected \"[]*ast.Ident\" got %T", n)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		case 2:
			if n, ok := n.([]ast.Expr); ok {
				p.Values = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Expr\" got %T", n)
		}

	case *ast.TypeSpec:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.Ident); ok {
				p.Name = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.Ident\" got %T", n)
		case 1:
			if n, ok := n.(*ast.FieldList); ok {
				p.TypeParams = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.FieldList\" got %T", n)
		case 2:
			if n, ok := n.(ast.Expr); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("expected \"ast.Expr\" got %T", n)
		}

	case *ast.GenDecl:
		switch c.fi {
		case 0:
			if n, ok := n.([]ast.Spec); ok {
				p.Specs = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Spec\" got %T", n)
		}

	case *ast.FuncDecl:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.FieldList); ok {
				p.Recv = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.FieldList\" got %T", n)
		case 1:
			if n, ok := n.(*ast.Ident); ok {
				p.Name = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.Ident\" got %T", n)
		case 2:
			if n, ok := n.(*ast.FuncType); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.FuncType\" got %T", n)
		case 3:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.BlockStmt\" got %T", n)
		}

	// Files and packages

	case *ast.File:
		switch c.fi {
		case 0:
			if n, ok := n.(*ast.Ident); ok {
				p.Name = n
				return nil
			}
			return fmt.Errorf("expected \"*ast.Ident\" got %T", n)
		case 1:
			if n, ok := n.([]ast.Decl); ok {
				p.Decls = n
				return nil
			}
			return fmt.Errorf("expected \"[]ast.Decl\" got %T", n)
		}

	default:
		panic(fmt.Sprintf("Apply: unexpected node type %T", reflect.TypeOf(p)))
	}

	return fmt.Errorf("unhandled case for parent type: %T", c.parent)
}
