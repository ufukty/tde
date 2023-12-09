package genetics

import (
	"fmt"
	"go/ast"
	"reflect"
	"tde/internal/evolution/genetics/mutation/v1/stg/ctxres/context"
	"tde/internal/evolution/genetics/mutation/v1/stg/nodes"
)

func sliced[T any](v T) []T {
	return []T{v}
}

// NOTE: use only for full replacement of field value with randomly generated values
func regenerateField(c cursor, ctx *context.Context) error {

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
			p.Names = sliced(nodes.Ident(ctx, 1))
		case 1:
			p.Type = nodes.Expr(ctx, 1)
		case 2:
			p.Tag = nodes.BasicLit(ctx, 1)
		}

	case *ast.FieldList:
		switch c.fi {
		case 0:
			p.List = sliced(nodes.Field(ctx, 1))
		}

	// Expressions

	case *ast.Ellipsis:
		switch c.fi {
		case 0:
			p.Elt = nodes.Expr(ctx, 1)
		}

	case *ast.FuncLit:
		switch c.fi {
		case 0:
			p.Type = nodes.FuncType(ctx, 1)
		case 1:
			p.Body = nodes.BlockStmt(ctx, 1)
		}

	case *ast.CompositeLit:
		switch c.fi {
		case 0:
			p.Type = nodes.Expr(ctx, 1)
		case 1:
			p.Elts = sliced(nodes.Expr(ctx, 1))
		}

	case *ast.ParenExpr:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		}

	case *ast.SelectorExpr:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		case 1:
			p.Sel = nodes.Ident(ctx, 1)
		}

	case *ast.IndexExpr:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		case 1:
			p.Index = nodes.Expr(ctx, 1)
		}

	case *ast.IndexListExpr:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		case 1:
			p.Indices = sliced(nodes.Expr(ctx, 1))
		}

	case *ast.SliceExpr:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		case 1:
			p.Low = nodes.Expr(ctx, 1)
		case 2:
			p.High = nodes.Expr(ctx, 1)
		case 3:
			p.Max = nodes.Expr(ctx, 1)
		}

	case *ast.TypeAssertExpr:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		case 1:
			p.Type = nodes.Expr(ctx, 1)
		}

	case *ast.CallExpr:
		switch c.fi {
		case 0:
			p.Fun = nodes.Expr(ctx, 1)
		case 1:
			p.Args = sliced(nodes.Expr(ctx, 1))
		}

	case *ast.StarExpr:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		}

	case *ast.UnaryExpr:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		}

	case *ast.BinaryExpr:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		case 1:
			p.Y = nodes.Expr(ctx, 1)
		}

	case *ast.KeyValueExpr:
		switch c.fi {
		case 0:
			p.Key = nodes.Expr(ctx, 1)
		case 1:
			p.Value = nodes.Expr(ctx, 1)
		}
	// Types
	case *ast.ArrayType:
		switch c.fi {
		case 0:
			p.Len = nodes.Expr(ctx, 1)
		case 1:
			p.Elt = nodes.Expr(ctx, 1)
		}

	case *ast.StructType:
		switch c.fi {
		case 0:
			p.Fields = nodes.FieldList(ctx, 1)
		}

	case *ast.FuncType:
		switch c.fi {
		case 0:
			p.TypeParams = nodes.FieldList(ctx, 1)
		case 1:
			p.Params = nodes.FieldList(ctx, 1)
		case 2:
			p.Results = nodes.FieldList(ctx, 1)
		}

	case *ast.InterfaceType:
		switch c.fi {
		case 0:
			p.Methods = nodes.FieldList(ctx, 1)
		}

	case *ast.MapType:
		switch c.fi {
		case 0:
			p.Key = nodes.Expr(ctx, 1)
		case 1:
			p.Value = nodes.Expr(ctx, 1)
		}

	case *ast.ChanType:
		switch c.fi {
		case 0:
			p.Value = nodes.Expr(ctx, 1)
		}

	// Statements

	case *ast.DeclStmt:
		switch c.fi {
		case 0:
			p.Decl = nodes.Decl(ctx, 1)
		}

	case *ast.LabeledStmt:
		switch c.fi {
		case 0:
			p.Label = nodes.Ident(ctx, 1)
		case 1:
			p.Stmt = nodes.Stmt(ctx, 1)
		}

	case *ast.ExprStmt:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		}

	case *ast.SendStmt:
		switch c.fi {
		case 0:
			p.Chan = nodes.Expr(ctx, 1)
		case 1:
			p.Value = nodes.Expr(ctx, 1)
		}

	case *ast.IncDecStmt:
		switch c.fi {
		case 0:
			p.X = nodes.Expr(ctx, 1)
		}

	case *ast.AssignStmt:
		switch c.fi {
		case 0:
			p.Lhs = sliced(nodes.Expr(ctx, 1))
		case 1:
			p.Rhs = sliced(nodes.Expr(ctx, 1))
		}

	case *ast.GoStmt:
		switch c.fi {
		case 0:
			p.Call = nodes.CallExpr(ctx, 1)
		}

	case *ast.DeferStmt:
		switch c.fi {
		case 0:
			p.Call = nodes.CallExpr(ctx, 1)
		}

	case *ast.ReturnStmt:
		switch c.fi {
		case 0:
			p.Results = sliced(nodes.Expr(ctx, 1))
		}

	case *ast.BranchStmt:
		switch c.fi {
		case 0:
			p.Label = nodes.Ident(ctx, 1)
		}

	case *ast.BlockStmt:
		switch c.fi {
		case 0:
			p.List = sliced(nodes.Stmt(ctx, 1))
		}

	case *ast.IfStmt:
		switch c.fi {
		case 0:
			p.Init = nodes.Stmt(ctx, 1)
		case 1:
			p.Cond = nodes.Expr(ctx, 1)
		case 2:
			p.Body = nodes.BlockStmt(ctx, 1)
		case 3:
			p.Else = nodes.Stmt(ctx, 1)
		}

	case *ast.CaseClause:
		switch c.fi {
		case 0:
			p.List = sliced(nodes.Expr(ctx, 1))
		case 1:
			p.Body = sliced(nodes.Stmt(ctx, 1))
		}

	case *ast.SwitchStmt:
		switch c.fi {
		case 0:
			p.Init = nodes.Stmt(ctx, 1)
		case 1:
			p.Tag = nodes.Expr(ctx, 1)
		case 2:
			p.Body = nodes.BlockStmt(ctx, 1)
		}

	case *ast.TypeSwitchStmt:
		switch c.fi {
		case 0:
			p.Init = nodes.Stmt(ctx, 1)
		case 1:
			p.Assign = nodes.Stmt(ctx, 1)
		case 2:
			p.Body = nodes.BlockStmt(ctx, 1)
		}

	case *ast.CommClause:
		switch c.fi {
		case 0:
			p.Comm = nodes.Stmt(ctx, 1)
		case 1:
			p.Body = sliced(nodes.Stmt(ctx, 1))
		}

	case *ast.SelectStmt:
		switch c.fi {
		case 0:
			p.Body = nodes.BlockStmt(ctx, 1)
		}

	case *ast.ForStmt:
		switch c.fi {
		case 0:
			p.Init = nodes.Stmt(ctx, 1)
		case 1:
			p.Cond = nodes.Expr(ctx, 1)
		case 2:
			p.Post = nodes.Stmt(ctx, 1)
		case 3:
			p.Body = nodes.BlockStmt(ctx, 1)
		}

	case *ast.RangeStmt:
		switch c.fi {
		case 0:
			p.Key = nodes.Expr(ctx, 1)
		case 1:
			p.Value = nodes.Expr(ctx, 1)
		case 2:
			p.X = nodes.Expr(ctx, 1)
		case 3:
			p.Body = nodes.BlockStmt(ctx, 1)
		}

	// Declarations

	case *ast.ImportSpec:
		switch c.fi {
		case 0:
			p.Name = nodes.Ident(ctx, 1)
		case 1:
			p.Path = nodes.BasicLit(ctx, 1)
		}

	case *ast.ValueSpec:
		switch c.fi {
		case 0:
			p.Names = sliced(nodes.Ident(ctx, 1))
		case 1:
			p.Type = nodes.Expr(ctx, 1)
		case 2:
			p.Values = sliced(nodes.Expr(ctx, 1))
		}

	case *ast.TypeSpec:
		switch c.fi {
		case 0:
			p.Name = nodes.Ident(ctx, 1)
		case 1:
			p.TypeParams = nodes.FieldList(ctx, 1)
		case 2:
			p.Type = nodes.Expr(ctx, 1)
		}

	case *ast.GenDecl:
		switch c.fi {
		case 0:
			p.Specs = sliced(nodes.Spec(ctx, 1))
		}

	case *ast.FuncDecl:
		switch c.fi {
		case 0:
			p.Recv = nodes.FieldList(ctx, 1)
		case 1:
			p.Name = nodes.Ident(ctx, 1)
		case 2:
			p.Type = nodes.FuncType(ctx, 1)
		case 3:
			p.Body = nodes.BlockStmt(ctx, 1)
		}

	// Files and packages

	case *ast.File:
		switch c.fi {
		case 0:
			p.Name = nodes.Ident(ctx, 1)
		case 1:
			p.Decls = sliced(nodes.Decl(ctx, 1))
		}

	default:
		panic(fmt.Sprintf("Apply: unexpected node type %T", reflect.TypeOf(p)))
	}

	return ErrTypeNotFound
}
