package ast_wrapper

import (
	"go/ast"
)

var NodeSlice ast.Node // Used by "walkEachNd" in parentTrace to imply traversed struct field is a slice type eg: []Expr, []Stmt

func walkEachNode[T ast.Node](
	nodeSlice []T,
	parentTrace []ast.Node,
	childIndexTrace []int,
	callback func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int),
) {
	parentTrace = append(parentTrace, NodeSlice)
	for i, v := range nodeSlice {
		walkHelper(v, parentTrace, append(childIndexTrace, i), callback)
	}
}

// 	// *ast.Comment:
// 	// *ast.BadExpr:
// 	// *ast.Ident:
// 	// *ast.BasicLit:
// 	// *ast.BadStmt
// 	// *ast.EmptyStmt
// 	// *ast.BadDecl:
// 	// *ast.Package
// 	// *ast.File
// 	*ast.CommentGroup: []any{n.List},
// 	*ast.Field:        []any{n.Names, n.Type, n.Tag},
// 	*ast.FieldList:    []any{n.List},
// 	// Expressions
// 	*ast.Ellipsis:       []any{n.Elt},
// 	*ast.FuncLit:        []any{n.Type, n.Body},
// 	*ast.CompositeLit:   []any{n.Type, n.Elts},
// 	*ast.ParenExpr:      []any{n.X},
// 	*ast.SelectorExpr:   []any{n.X, n.Sel},
// 	*ast.IndexExpr:      []any{n.X, n.Index},
// 	*ast.IndexListExpr:  []any{n.X, n.Indices},
// 	*ast.SliceExpr:      []any{n.X, n.Low, n.High, n.Max},
// 	*ast.TypeAssertExpr: []any{n.X, n.Type},
// 	*ast.CallExpr:       []any{n.Fun, n.Args},
// 	*ast.StarExpr:       []any{n.X},
// 	*ast.UnaryExpr:      []any{n.X},
// 	*ast.BinaryExpr:     []any{n.X, n.Y},
// 	*ast.KeyValueExpr:   []any{n.Key, n.Value},
// 	// Types
// 	*ast.ArrayType:     []any{n.Len, n.Elt},
// 	*ast.StructType:    []any{n.Fields},
// 	*ast.FuncType:      []any{n.TypeParams, n.Params, n.Results},
// 	*ast.InterfaceType: []any{n.Methods},
// 	*ast.MapType:       []any{n.Key, n.Value},
// 	*ast.ChanType:      []any{n.Value},
// 	// Statements
// 	*ast.DeclStmt:       []any{n.Decl},
// 	*ast.LabeledStmt:    []any{n.Label, n.Stmt},
// 	*ast.ExprStmt:       []any{n.X},
// 	*ast.SendStmt:       []any{n.Chan, n.Value},
// 	*ast.IncDecStmt:     []any{n.X},
// 	*ast.AssignStmt:     []any{n.Lhs, n.Rhs},
// 	*ast.GoStmt:         []any{n.Call},
// 	*ast.DeferStmt:      []any{n.Call},
// 	*ast.ReturnStmt:     []any{n.Results},
// 	*ast.BranchStmt:     []any{n.Label},
// 	*ast.BlockStmt:      []any{n.List},
// 	*ast.IfStmt:         []any{n.Init, n.Cond, n.Body, n.Else},
// 	*ast.CaseClause:     []any{n.List, n.Body},
// 	*ast.SwitchStmt:     []any{n.Init, n.Tag, n.Body},
// 	*ast.TypeSwitchStmt: []any{n.Init, n.Assign, n.Body},
// 	*ast.CommClause:     []any{n.Comm, n.Body},
// 	*ast.SelectStmt:     []any{n.Body},
// 	*ast.ForStmt:        []any{n.Init, n.Cond, n.Post, n.Body},
// 	*ast.RangeStmt:      []any{n.Key, n.Value, n.X, n.Body},
// 	// Declarations
// 	*ast.ImportSpec: []any{n.Name, n.Path},
// 	*ast.ValueSpec:  []any{n.Names, n.Type, n.Values},
// 	*ast.TypeSpec:   []any{n.Name, n.TypeParams, n.Type},
// 	*ast.GenDecl:    []any{n.Specs},
// 	*ast.FuncDecl:   []any{n.Recv, n.Name, n.Type, n.Body},


func walkHelper(n ast.Node, parentTrace []ast.Node, childIndexTrace []int, callback func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int)) {
	callback(n, parentTrace, childIndexTrace)
	parentTrace = append(parentTrace, n)
	switch n := n.(type) {

	// MARK: Comments and fields

	case *ast.Comment:
		// nothing to do

	case *ast.CommentGroup:
		if n.List != nil {
			walkEachNode(n.List, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.Field:
		if n.Names != nil {
			walkEachNode(n.Names, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Type != nil {
			walkHelper(n.Type, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.Tag != nil {
			walkHelper(n.Tag, parentTrace, append(childIndexTrace, 2), callback)
		}

	case *ast.FieldList:
		if n.List != nil {
			walkEachNode(n.List, parentTrace, append(childIndexTrace, 0), callback)
		}

	// MARK: Expressions

	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
		// nothing to do

	case *ast.Ellipsis:
		if n.Elt != nil {
			walkHelper(n.Elt, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.FuncLit:
		if n.Type != nil {
			walkHelper(n.Type, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Body != nil {
			walkHelper(n.Body, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.CompositeLit:
		if n.Type != nil {
			walkHelper(n.Type, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Elts != nil {
			walkEachNode(n.Elts, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.ParenExpr:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.SelectorExpr:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Sel != nil {
			walkHelper(n.Sel, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.IndexExpr:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Index != nil {
			walkHelper(n.Index, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.IndexListExpr:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Indices != nil {
			walkEachNode(n.Indices, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.SliceExpr:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Low != nil {
			walkHelper(n.Low, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.High != nil {
			walkHelper(n.High, parentTrace, append(childIndexTrace, 2), callback)
		}
		if n.Max != nil {
			walkHelper(n.Max, parentTrace, append(childIndexTrace, 3), callback)
		}

	case *ast.TypeAssertExpr:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Type != nil {
			walkHelper(n.Type, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.CallExpr:
		if n.Fun != nil {
			walkHelper(n.Fun, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Args != nil {
			walkEachNode(n.Args, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.StarExpr:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.UnaryExpr:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.BinaryExpr:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Y != nil {
			walkHelper(n.Y, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.KeyValueExpr:
		if n.Key != nil {
			walkHelper(n.Key, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Value != nil {
			walkHelper(n.Value, parentTrace, append(childIndexTrace, 1), callback)
		}

	// MARK: Types

	case *ast.ArrayType:
		if n.Len != nil {
			walkHelper(n.Len, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Elt != nil {
			walkHelper(n.Elt, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.StructType:
		if n.Fields != nil {
			walkHelper(n.Fields, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.FuncType:
		if n.TypeParams != nil {
			walkHelper(n.TypeParams, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Params != nil {
			walkHelper(n.Params, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.Results != nil {
			walkHelper(n.Results, parentTrace, append(childIndexTrace, 2), callback)
		}

	case *ast.InterfaceType:
		if n.Methods != nil {
			walkHelper(n.Methods, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.MapType:
		if n.Key != nil {
			walkHelper(n.Key, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Value != nil {
			walkHelper(n.Value, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.ChanType:
		if n.Value != nil {
			walkHelper(n.Value, parentTrace, append(childIndexTrace, 0), callback)
		}

	// MARK: Statements

	case *ast.BadStmt:
		// nothing to do

	case *ast.DeclStmt:
		if n.Decl != nil {
			walkHelper(n.Decl, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.EmptyStmt:
		// nothing to do

	case *ast.LabeledStmt:
		if n.Label != nil {
			walkHelper(n.Label, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Stmt != nil {
			walkHelper(n.Stmt, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.ExprStmt:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.SendStmt:
		if n.Chan != nil {
			walkHelper(n.Chan, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Value != nil {
			walkHelper(n.Value, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.IncDecStmt:
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.AssignStmt:
		if n.Lhs != nil {
			walkEachNode(n.Lhs, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Rhs != nil {
			walkEachNode(n.Rhs, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.GoStmt:
		if n.Call != nil {
			walkHelper(n.Call, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.DeferStmt:
		if n.Call != nil {
			walkHelper(n.Call, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.ReturnStmt:
		if n.Results != nil {
			walkEachNode(n.Results, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.BranchStmt:
		if n.Label != nil {
			walkHelper(n.Label, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.BlockStmt:
		if n.List != nil {
			walkEachNode(n.List, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.IfStmt:
		if n.Init != nil {
			walkHelper(n.Init, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Cond != nil {
			walkHelper(n.Cond, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.Body != nil {
			walkHelper(n.Body, parentTrace, append(childIndexTrace, 2), callback)
		}
		if n.Else != nil {
			walkHelper(n.Else, parentTrace, append(childIndexTrace, 3), callback)
		}

	case *ast.CaseClause:
		if n.List != nil {
			walkEachNode(n.List, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Body != nil {
			walkEachNode(n.Body, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.SwitchStmt:
		if n.Init != nil {
			walkHelper(n.Init, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Tag != nil {
			walkHelper(n.Tag, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.Body != nil {
			walkHelper(n.Body, parentTrace, append(childIndexTrace, 2), callback)
		}

	case *ast.TypeSwitchStmt:
		if n.Init != nil {
			walkHelper(n.Init, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Assign != nil {
			walkHelper(n.Assign, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.Body != nil {
			walkHelper(n.Body, parentTrace, append(childIndexTrace, 2), callback)
		}

	case *ast.CommClause:
		if n.Comm != nil {
			walkHelper(n.Comm, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Body != nil {
			walkEachNode(n.Body, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.SelectStmt:
		if n.Body != nil {
			walkHelper(n.Body, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.ForStmt:
		if n.Init != nil {
			walkHelper(n.Init, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Cond != nil {
			walkHelper(n.Cond, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.Post != nil {
			walkHelper(n.Post, parentTrace, append(childIndexTrace, 2), callback)
		}
		if n.Body != nil {
			walkHelper(n.Body, parentTrace, append(childIndexTrace, 3), callback)
		}

	case *ast.RangeStmt:
		if n.Key != nil {
			walkHelper(n.Key, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Value != nil {
			walkHelper(n.Value, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.X != nil {
			walkHelper(n.X, parentTrace, append(childIndexTrace, 2), callback)
		}
		if n.Body != nil {
			walkHelper(n.Body, parentTrace, append(childIndexTrace, 3), callback)
		}

	// MARK: Declarations

	case *ast.ImportSpec:
		if n.Name != nil {
			walkHelper(n.Name, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Path != nil {
			walkHelper(n.Path, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.ValueSpec:
		if n.Names != nil {
			walkEachNode(n.Names, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Type != nil {
			walkHelper(n.Type, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.Values != nil {
			walkEachNode(n.Values, parentTrace, append(childIndexTrace, 2), callback)
		}

	case *ast.TypeSpec:
		if n.Name != nil {
			walkHelper(n.Name, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.TypeParams != nil {
			walkHelper(n.TypeParams, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.Type != nil {
			walkHelper(n.Type, parentTrace, append(childIndexTrace, 2), callback)
		}

	case *ast.BadDecl:
		// nothing to do

	case *ast.GenDecl:
		if n.Specs != nil {
			walkEachNode(n.Specs, parentTrace, append(childIndexTrace, 0), callback)
		}

	case *ast.FuncDecl:
		if n.Recv != nil {
			walkHelper(n.Recv, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Name != nil {
			walkHelper(n.Name, parentTrace, append(childIndexTrace, 1), callback)
		}
		if n.Type != nil {
			walkHelper(n.Type, parentTrace, append(childIndexTrace, 2), callback)
		}
		if n.Body != nil {
			walkHelper(n.Body, parentTrace, append(childIndexTrace, 3), callback)
		}

	case *ast.File:
		if n.Name != nil {
			walkHelper(n.Name, parentTrace, append(childIndexTrace, 0), callback)
		}
		if n.Decls != nil {
			walkEachNode(n.Decls, parentTrace, append(childIndexTrace, 1), callback)
		}

	case *ast.Package:
		for _, f := range n.Files {
			if f != nil {
				walkHelper(f, parentTrace, append(childIndexTrace, 0), callback)
			}
		}

	}
}

// walk calls the callback function for every joint (eg. []Stmt..., []Expr..., Stmt, Expr, Token)
func Walk(root ast.Node, callback func(n ast.Node, parentTrace []ast.Node, childIndexTrace []int)) {
	walkHelper(root, []ast.Node{}, []int{}, callback)
}

// Inspects (Walks) the partial-AST of root; chooses one insertion point and make an appropriate type Node append
func Enrich(root ast.Node) {

}
