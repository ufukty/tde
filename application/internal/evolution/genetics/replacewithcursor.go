package genetics

import (
	"fmt"
	"go/ast"
	"reflect"
)

var ErrTypeMismatch = fmt.Errorf("type mismatch between next value and parent.field")
var ErrTypeNotFound = fmt.Errorf("type not found")

// (n)ext
func replaceOnParentWithCursor(c cursor, n any) error {

	switch p := c.parent.(type) {

	// case
	//  *ast.Package,
	// 	*ast.Comment,
	// 	*ast.BadExpr,
	// 	*ast.Ident,
	// 	*ast.BasicLit,
	// 	*ast.BadStmt,
	// 	*ast.EmptyStmt,
	// 	*ast.BadDecl:

	case *ast.CommentGroup:
		switch c.field {
		case 0:
			if n, ok := n.([]*ast.Comment); ok {
				p.List = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]*ast.Comment", n, ErrTypeMismatch)
		}

	case *ast.Field:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Doc = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.([]*ast.Ident); ok {
				p.Names = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]*ast.Ident", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(ast.Expr); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 3:
			if n, ok := n.(*ast.BasicLit); ok {
				p.Tag = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.BasicLit", n, ErrTypeMismatch)
		case 4:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Comment = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		}

	case *ast.FieldList:
		switch c.field {
		case 0:
			if n, ok := n.([]*ast.Field); ok {
				p.List = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]*ast.Field", n, ErrTypeMismatch)
		}

	// Expressions

	case *ast.Ellipsis:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Elt = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.FuncLit:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.FuncType); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.FuncType", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.BlockStmt", n, ErrTypeMismatch)
		}

	case *ast.CompositeLit:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.([]ast.Expr); ok {
				p.Elts = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.ParenExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.SelectorExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(*ast.Ident); ok {
				p.Sel = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.Ident", n, ErrTypeMismatch)
		}

	case *ast.IndexExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Index = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.IndexListExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.([]ast.Expr); ok {
				p.Indices = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.SliceExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Low = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(ast.Expr); ok {
				p.High = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 3:
			if n, ok := n.(ast.Expr); ok {
				p.Max = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.TypeAssertExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.CallExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Fun = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.([]ast.Expr); ok {
				p.Args = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.StarExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.UnaryExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.BinaryExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Y = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.KeyValueExpr:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Key = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Value = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}
	// Types
	case *ast.ArrayType:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Len = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Elt = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.StructType:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.FieldList); ok {
				p.Fields = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.FieldList", n, ErrTypeMismatch)
		}

	case *ast.FuncType:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.FieldList); ok {
				p.TypeParams = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.FieldList", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(*ast.FieldList); ok {
				p.Params = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.FieldList", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(*ast.FieldList); ok {
				p.Results = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.FieldList", n, ErrTypeMismatch)
		}

	case *ast.InterfaceType:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.FieldList); ok {
				p.Methods = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.FieldList", n, ErrTypeMismatch)
		}

	case *ast.MapType:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Key = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Value = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.ChanType:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Value = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	// Statements

	case *ast.DeclStmt:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Decl); ok {
				p.Decl = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Decl", n, ErrTypeMismatch)
		}

	case *ast.LabeledStmt:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.Ident); ok {
				p.Label = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.Ident", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Stmt); ok {
				p.Stmt = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Stmt", n, ErrTypeMismatch)
		}

	case *ast.ExprStmt:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.SendStmt:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Chan = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Value = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.IncDecStmt:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.AssignStmt:
		switch c.field {
		case 0:
			if n, ok := n.([]ast.Expr); ok {
				p.Lhs = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.([]ast.Expr); ok {
				p.Rhs = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.GoStmt:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.CallExpr); ok {
				p.Call = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CallExpr", n, ErrTypeMismatch)
		}

	case *ast.DeferStmt:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.CallExpr); ok {
				p.Call = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CallExpr", n, ErrTypeMismatch)
		}

	case *ast.ReturnStmt:
		switch c.field {
		case 0:
			if n, ok := n.([]ast.Expr); ok {
				p.Results = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Expr", n, ErrTypeMismatch)
		}

	case *ast.BranchStmt:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.Ident); ok {
				p.Label = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.Ident", n, ErrTypeMismatch)
		}

	case *ast.BlockStmt:
		switch c.field {
		case 0:
			if n, ok := n.([]ast.Stmt); ok {
				p.List = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Stmt", n, ErrTypeMismatch)
		}

	case *ast.IfStmt:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Stmt); ok {
				p.Init = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Stmt", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Cond = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.BlockStmt", n, ErrTypeMismatch)
		case 3:
			if n, ok := n.(ast.Stmt); ok {
				p.Else = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Stmt", n, ErrTypeMismatch)
		}

	case *ast.CaseClause:
		switch c.field {
		case 0:
			if n, ok := n.([]ast.Expr); ok {
				p.List = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.([]ast.Stmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Stmt", n, ErrTypeMismatch)
		}

	case *ast.SwitchStmt:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Stmt); ok {
				p.Init = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Stmt", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Tag = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.BlockStmt", n, ErrTypeMismatch)
		}

	case *ast.TypeSwitchStmt:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Stmt); ok {
				p.Init = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Stmt", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Stmt); ok {
				p.Assign = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Stmt", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.BlockStmt", n, ErrTypeMismatch)
		}

	case *ast.CommClause:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Stmt); ok {
				p.Comm = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Stmt", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.([]ast.Stmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Stmt", n, ErrTypeMismatch)
		}

	case *ast.SelectStmt:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.BlockStmt", n, ErrTypeMismatch)
		}

	case *ast.ForStmt:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Stmt); ok {
				p.Init = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Stmt", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Cond = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(ast.Stmt); ok {
				p.Post = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Stmt", n, ErrTypeMismatch)
		case 3:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.BlockStmt", n, ErrTypeMismatch)
		}

	case *ast.RangeStmt:
		switch c.field {
		case 0:
			if n, ok := n.(ast.Expr); ok {
				p.Key = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(ast.Expr); ok {
				p.Value = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(ast.Expr); ok {
				p.X = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 3:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.BlockStmt", n, ErrTypeMismatch)
		}

	// Declarations

	case *ast.ImportSpec:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Doc = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(*ast.Ident); ok {
				p.Name = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.Ident", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(*ast.BasicLit); ok {
				p.Path = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.BasicLit", n, ErrTypeMismatch)
		case 3:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Comment = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		}

	case *ast.ValueSpec:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Doc = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.([]*ast.Ident); ok {
				p.Names = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]*ast.Ident", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(ast.Expr); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 3:
			if n, ok := n.([]ast.Expr); ok {
				p.Values = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Expr", n, ErrTypeMismatch)
		case 4:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Comment = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		}

	case *ast.TypeSpec:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Doc = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(*ast.Ident); ok {
				p.Name = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.Ident", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(*ast.FieldList); ok {
				p.TypeParams = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.FieldList", n, ErrTypeMismatch)
		case 3:
			if n, ok := n.(ast.Expr); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "ast.Expr", n, ErrTypeMismatch)
		case 4:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Comment = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		}

	case *ast.GenDecl:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Doc = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.([]ast.Spec); ok {
				p.Specs = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Spec", n, ErrTypeMismatch)
		}

	case *ast.FuncDecl:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Doc = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(*ast.FieldList); ok {
				p.Recv = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.FieldList", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.(*ast.Ident); ok {
				p.Name = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.Ident", n, ErrTypeMismatch)
		case 3:
			if n, ok := n.(*ast.FuncType); ok {
				p.Type = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.FuncType", n, ErrTypeMismatch)
		case 4:
			if n, ok := n.(*ast.BlockStmt); ok {
				p.Body = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.BlockStmt", n, ErrTypeMismatch)
		}

	// Files and packages

	case *ast.File:
		switch c.field {
		case 0:
			if n, ok := n.(*ast.CommentGroup); ok {
				p.Doc = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.CommentGroup", n, ErrTypeMismatch)
		case 1:
			if n, ok := n.(*ast.Ident); ok {
				p.Name = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "*ast.Ident", n, ErrTypeMismatch)
		case 2:
			if n, ok := n.([]ast.Decl); ok {
				p.Decls = n
				return nil
			}
			return fmt.Errorf("%w: expected %q got %T", "[]ast.Decl", n, ErrTypeMismatch)
		}

	default:
		panic(fmt.Sprintf("Apply: unexpected node type %T", reflect.TypeOf(p)))
	}

	return ErrTypeNotFound
}
