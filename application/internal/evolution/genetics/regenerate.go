package genetics

import (
	"fmt"
	"go/ast"
	"reflect"
	"tde/internal/evolution/genetics/nodes"
)

func sliced[T any](v T) []T {
	return []T{v}
}

// NOTE: use only for full replacement of field value with randomly generated values
func regenerateField(nc *nodes.Creator, c cursor) error {

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
			Names, err := nc.Ident(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Names = sliced(Names)
		case 1:
			Type, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Type = Type
		case 2:
			Tag, err := nc.BasicLit(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Tag = Tag
		}

	case *ast.FieldList:
		switch c.fi {
		case 0:
			List, err := nc.Field(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.List = sliced(List)
		}

	// Expressions

	case *ast.Ellipsis:
		switch c.fi {
		case 0:
			Elt, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Elt = Elt
		}

	case *ast.FuncLit:
		switch c.fi {
		case 0:
			Type, err := nc.FuncType(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Type = Type
		case 1:
			Body, err := nc.BlockStmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Body = Body
		}

	case *ast.CompositeLit:
		switch c.fi {
		case 0:
			Type, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Type = Type
		case 1:
			Elts, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Elts = sliced(Elts)
		}

	case *ast.ParenExpr:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		}

	case *ast.SelectorExpr:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		case 1:
			Sel, err := nc.Ident(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Sel = Sel
		}

	case *ast.IndexExpr:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		case 1:
			Index, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Index = Index
		}

	case *ast.IndexListExpr:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		case 1:
			Indices, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Indices = sliced(Indices)
		}

	case *ast.SliceExpr:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		case 1:
			Low, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Low = Low
		case 2:
			High, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.High = High
		case 3:
			Max, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Max = Max
		}

	case *ast.TypeAssertExpr:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		case 1:
			Type, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Type = Type
		}

	case *ast.CallExpr:
		switch c.fi {
		case 0:
			Fun, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Fun = Fun
		case 1:
			Args, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Args = sliced(Args)
		}

	case *ast.StarExpr:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		}

	case *ast.UnaryExpr:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		}

	case *ast.BinaryExpr:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		case 1:
			Y, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Y = Y
		}

	case *ast.KeyValueExpr:
		switch c.fi {
		case 0:
			Key, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Key = Key
		case 1:
			Value, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Value = Value
		}
	// Types
	case *ast.ArrayType:
		switch c.fi {
		case 0:
			Len, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Len = Len
		case 1:
			Elt, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Elt = Elt
		}

	case *ast.StructType:
		switch c.fi {
		case 0:
			Fields, err := nc.FieldList(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Fields = Fields
		}

	case *ast.FuncType:
		switch c.fi {
		case 0:
			TypeParams, err := nc.FieldList(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.TypeParams = TypeParams
		case 1:
			Params, err := nc.FieldList(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Params = Params
		case 2:
			Results, err := nc.FieldList(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Results = Results
		}

	case *ast.InterfaceType:
		switch c.fi {
		case 0:
			Methods, err := nc.FieldList(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Methods = Methods
		}

	case *ast.MapType:
		switch c.fi {
		case 0:
			Key, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Key = Key
		case 1:
			Value, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Value = Value
		}

	case *ast.ChanType:
		switch c.fi {
		case 0:
			Value, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Value = Value
		}

	// Statements

	case *ast.DeclStmt:
		switch c.fi {
		case 0:
			Decl, err := nc.Decl(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Decl = Decl
		}

	case *ast.LabeledStmt:
		switch c.fi {
		case 0:
			Label, err := nc.Ident(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Label = Label
		case 1:
			Stmt, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Stmt = Stmt
		}

	case *ast.ExprStmt:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		}

	case *ast.SendStmt:
		switch c.fi {
		case 0:
			Chan, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Chan = Chan
		case 1:
			Value, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Value = Value
		}

	case *ast.IncDecStmt:
		switch c.fi {
		case 0:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		}

	case *ast.AssignStmt:
		switch c.fi {
		case 0:
			Lhs, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Lhs = sliced(Lhs)
		case 1:
			Rhs, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Rhs = sliced(Rhs)
		}

	case *ast.GoStmt:
		switch c.fi {
		case 0:
			Call, err := nc.CallExpr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Call = Call
		}

	case *ast.DeferStmt:
		switch c.fi {
		case 0:
			Call, err := nc.CallExpr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Call = Call
		}

	case *ast.ReturnStmt:
		switch c.fi {
		case 0:
			Results, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Results = sliced(Results)
		}

	case *ast.BranchStmt:
		switch c.fi {
		case 0:
			Label, err := nc.Ident(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Label = Label
		}

	case *ast.BlockStmt:
		switch c.fi {
		case 0:
			List, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.List = sliced(List)
		}

	case *ast.IfStmt:
		switch c.fi {
		case 0:
			Init, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Init = Init
		case 1:
			Cond, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Cond = Cond
		case 2:
			Body, err := nc.BlockStmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Body = Body
		case 3:
			Else, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Else = Else
		}

	case *ast.CaseClause:
		switch c.fi {
		case 0:
			List, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.List = sliced(List)
		case 1:
			Body, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Body = sliced(Body)
		}

	case *ast.SwitchStmt:
		switch c.fi {
		case 0:
			Init, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Init = Init
		case 1:
			Tag, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Tag = Tag
		case 2:
			Body, err := nc.BlockStmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Body = Body
		}

	case *ast.TypeSwitchStmt:
		switch c.fi {
		case 0:
			Init, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Init = Init
		case 1:
			Assign, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Assign = Assign
		case 2:
			Body, err := nc.BlockStmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Body = Body
		}

	case *ast.CommClause:
		switch c.fi {
		case 0:
			Comm, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Comm = Comm
		case 1:
			Body, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Body = sliced(Body)
		}

	case *ast.SelectStmt:
		switch c.fi {
		case 0:
			Body, err := nc.BlockStmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Body = Body
		}

	case *ast.ForStmt:
		switch c.fi {
		case 0:
			Init, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Init = Init
		case 1:
			Cond, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Cond = Cond
		case 2:
			Post, err := nc.Stmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Post = Post
		case 3:
			Body, err := nc.BlockStmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Body = Body
		}

	case *ast.RangeStmt:
		switch c.fi {
		case 0:
			Key, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Key = Key
		case 1:
			Value, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Value = Value
		case 2:
			X, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.X = X
		case 3:
			Body, err := nc.BlockStmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Body = Body
		}

	// Declarations

	case *ast.ImportSpec:
		switch c.fi {
		case 0:
			Name, err := nc.Ident(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Name = Name
		case 1:
			Path, err := nc.BasicLit(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Path = Path
		}

	case *ast.ValueSpec:
		switch c.fi {
		case 0:
			Names, err := nc.Ident(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Names = sliced(Names)
		case 1:
			Type, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Type = Type
		case 2:
			Values, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Values = sliced(Values)
		}

	case *ast.TypeSpec:
		switch c.fi {
		case 0:
			Name, err := nc.Ident(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Name = Name
		case 1:
			TypeParams, err := nc.FieldList(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.TypeParams = TypeParams
		case 2:
			Type, err := nc.Expr(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Type = Type
		}

	case *ast.GenDecl:
		switch c.fi {
		case 0:
			Specs, err := nc.Spec(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Specs = sliced(Specs)
		}

	case *ast.FuncDecl:
		switch c.fi {
		case 0:
			Recv, err := nc.FieldList(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Recv = Recv
		case 1:
			Name, err := nc.Ident(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Name = Name
		case 2:
			Type, err := nc.FuncType(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Type = Type
		case 3:
			Body, err := nc.BlockStmt(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Body = Body
		}

	// Files and packages

	case *ast.File:
		switch c.fi {
		case 0:
			Name, err := nc.Ident(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Name = Name
		case 1:
			Decls, err := nc.Decl(1)
			if err != nil {
				return fmt.Errorf("generating replacement node: %w", err)
			}
			p.Decls = sliced(Decls)
		}

	default:
		panic(fmt.Sprintf("Apply: unexpected node type %T", reflect.TypeOf(p)))
	}

	return ErrTypeNotFound
}
