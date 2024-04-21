package astwutl

import (
	"tde/internal/utilities/slicew"

	"go/ast"
)

// Returns true if they are same.
// Returns true for those types that don't have non-node type field.
// Returns false for changes in addresses also. In addition to only value changes just like CompareNonNodeFields
func CompareNonNodeFieldsWithAddresses(a, b ast.Node) bool {
	switch a := a.(type) {

	case *ast.Comment:
		if b, ok := b.(*ast.Comment); ok {
			return a.Slash == b.Slash && // token.Pos
				&a.Slash == &b.Slash && // token.Pos
				a.Text == b.Text && // string
				&a.Text == &b.Text // string
		} else {
			return false
		}

	case *ast.FieldList:
		if b, ok := b.(*ast.FieldList); ok {
			return a.Opening == b.Opening && // token.Pos
				&a.Opening == &b.Opening && // token.Pos
				a.Closing == b.Closing && // token.Pos
				&a.Closing == &b.Closing // token.Pos
		} else {
			return false
		}

	case *ast.BadExpr:
		if b, ok := b.(*ast.BadExpr); ok {
			return a.From == b.From && // token.Pos
				&a.From == &b.From && // token.Pos
				a.To == b.To && // token.Pos
				&a.To == &b.To // token.Pos
		} else {
			return false
		}

	case *ast.Ident:
		if b, ok := b.(*ast.Ident); ok {
			return a.NamePos == b.NamePos && // token.Pos
				&a.NamePos == &b.NamePos && // token.Pos
				a.Name == b.Name && // string
				&a.Name == &b.Name // string
		} else {
			return false
		}

	case *ast.Ellipsis:
		if b, ok := b.(*ast.Ellipsis); ok {
			return a.Ellipsis == b.Ellipsis && // token.Pos
				&a.Ellipsis == &b.Ellipsis // token.Pos
		} else {
			return false
		}

	case *ast.BasicLit:
		if b, ok := b.(*ast.BasicLit); ok {
			return a.ValuePos == b.ValuePos && // token.Pos
				&a.ValuePos == &b.ValuePos && // token.Pos
				a.Kind == b.Kind && // token.Token
				&a.Kind == &b.Kind && // token.Token
				a.Value == b.Value && // string
				&a.Value == &b.Value // string
		} else {
			return false
		}

	case *ast.CompositeLit:
		if b, ok := b.(*ast.CompositeLit); ok {
			return a.Lbrace == b.Lbrace && // token.Pos
				&a.Lbrace == &b.Lbrace && // token.Pos
				a.Rbrace == b.Rbrace && // token.Pos
				&a.Rbrace == &b.Rbrace && // token.Pos
				a.Incomplete == b.Incomplete && // bool
				&a.Incomplete == &b.Incomplete // bool
		} else {
			return false
		}

	case *ast.ParenExpr:
		if b, ok := b.(*ast.ParenExpr); ok {
			return a.Lparen == b.Lparen && // token.Pos
				&a.Lparen == &b.Lparen && // token.Pos
				a.Rparen == b.Rparen && // token.Pos
				&a.Rparen == &b.Rparen // token.Pos
		} else {
			return false
		}

	case *ast.IndexExpr:
		if b, ok := b.(*ast.IndexExpr); ok {
			return a.Lbrack == b.Lbrack && // token.Pos
				&a.Lbrack == &b.Lbrack && // token.Pos
				a.Rbrack == b.Rbrack && // token.Pos
				&a.Rbrack == &b.Rbrack // token.Pos
		} else {
			return false
		}

	case *ast.IndexListExpr:
		if b, ok := b.(*ast.IndexListExpr); ok {
			return a.Lbrack == b.Lbrack && // token.Pos
				&a.Lbrack == &b.Lbrack && // token.Pos
				a.Rbrack == b.Rbrack && // token.Pos
				&a.Rbrack == &b.Rbrack // token.Pos
		} else {
			return false
		}

	case *ast.SliceExpr:
		if b, ok := b.(*ast.SliceExpr); ok {
			return a.Lbrack == b.Lbrack && // token.Pos
				&a.Lbrack == &b.Lbrack && // token.Pos
				a.Slice3 == b.Slice3 && // bool
				&a.Slice3 == &b.Slice3 && // bool
				a.Rbrack == b.Rbrack && // token.Pos
				&a.Rbrack == &b.Rbrack // token.Pos
		} else {
			return false
		}

	case *ast.TypeAssertExpr:
		if b, ok := b.(*ast.TypeAssertExpr); ok {
			return a.Lparen == b.Lparen && // token.Pos
				&a.Lparen == &b.Lparen && // token.Pos
				a.Rparen == b.Rparen && // token.Pos
				&a.Rparen == &b.Rparen // token.Pos
		} else {
			return false
		}

	case *ast.CallExpr:
		if b, ok := b.(*ast.CallExpr); ok {
			return a.Lparen == b.Lparen && // token.Pos
				&a.Lparen == &b.Lparen && // token.Pos
				a.Ellipsis == b.Ellipsis && // token.Pos
				&a.Ellipsis == &b.Ellipsis && // token.Pos
				a.Rparen == b.Rparen && //  token.Pos
				&a.Rparen == &b.Rparen //  token.Pos
		} else {
			return false
		}

	case *ast.StarExpr:
		if b, ok := b.(*ast.StarExpr); ok {
			return a.Star == b.Star && // token.Pos
				&a.Star == &b.Star // token.Pos
		} else {
			return false
		}

	case *ast.UnaryExpr:
		if b, ok := b.(*ast.UnaryExpr); ok {
			return a.OpPos == b.OpPos && // token.Pos
				&a.OpPos == &b.OpPos && // token.Pos
				a.Op == b.Op && // token.Token
				&a.Op == &b.Op // token.Token
		} else {
			return false
		}

	case *ast.BinaryExpr:
		if b, ok := b.(*ast.BinaryExpr); ok {
			return a.OpPos == b.OpPos && // token.Pos
				&a.OpPos == &b.OpPos && // token.Pos
				a.Op == b.Op && // token.Token
				&a.Op == &b.Op // token.Token
		} else {
			return false
		}

	case *ast.KeyValueExpr:
		if b, ok := b.(*ast.KeyValueExpr); ok {
			return a.Colon == b.Colon && // token.Pos
				&a.Colon == &b.Colon // token.Pos
		} else {
			return false
		}

	case *ast.ArrayType:
		if b, ok := b.(*ast.ArrayType); ok {
			return a.Lbrack == b.Lbrack && // token.Pos
				&a.Lbrack == &b.Lbrack // token.Pos
		} else {
			return false
		}

	case *ast.StructType:
		if b, ok := b.(*ast.StructType); ok {
			return a.Struct == b.Struct && // token.Pos
				&a.Struct == &b.Struct && // token.Pos
				a.Incomplete == b.Incomplete && // bool
				&a.Incomplete == &b.Incomplete // bool
		} else {
			return false
		}

	case *ast.FuncType:
		if b, ok := b.(*ast.FuncType); ok {
			return a.Func == b.Func && // token.Pos
				&a.Func == &b.Func // token.Pos
		} else {
			return false
		}

	case *ast.InterfaceType:
		if b, ok := b.(*ast.InterfaceType); ok {
			return a.Interface == b.Interface && // token.Pos
				&a.Interface == &b.Interface && // token.Pos
				a.Incomplete == b.Incomplete && // bool
				&a.Incomplete == &b.Incomplete // bool
		} else {
			return false
		}

	case *ast.MapType:
		if b, ok := b.(*ast.MapType); ok {
			return a.Map == b.Map && // token.Pos
				&a.Map == &b.Map // token.Pos
		} else {
			return false
		}

	case *ast.ChanType:
		if b, ok := b.(*ast.ChanType); ok {
			return a.Begin == b.Begin && // token.Pos
				&a.Begin == &b.Begin && // token.Pos
				a.Arrow == b.Arrow && // token.Pos
				&a.Arrow == &b.Arrow // token.Pos
		} else {
			return false
		}

	case *ast.BadStmt:
		if b, ok := b.(*ast.BadStmt); ok {
			return a.From == b.From && // token.Pos
				&a.From == &b.From && // token.Pos
				a.To == b.To && // token.Pos
				&a.To == &b.To // token.Pos
		} else {
			return false
		}

	case *ast.EmptyStmt:
		if b, ok := b.(*ast.EmptyStmt); ok {
			return a.Semicolon == b.Semicolon && // token.Pos
				&a.Semicolon == &b.Semicolon && // token.Pos
				a.Implicit == b.Implicit && // bool
				&a.Implicit == &b.Implicit // bool
		} else {
			return false
		}

	case *ast.LabeledStmt:
		if b, ok := b.(*ast.LabeledStmt); ok {
			return a.Colon == b.Colon && // token.Pos
				&a.Colon == &b.Colon // token.Pos
		} else {
			return false
		}

	case *ast.SendStmt:
		if b, ok := b.(*ast.SendStmt); ok {
			return a.Arrow == b.Arrow && // token.Pos
				&a.Arrow == &b.Arrow // token.Pos
		} else {
			return false
		}

	case *ast.IncDecStmt:
		if b, ok := b.(*ast.IncDecStmt); ok {
			return a.TokPos == b.TokPos && // token.Pos
				&a.TokPos == &b.TokPos && // token.Pos
				a.Tok == b.Tok && // token.Token
				&a.Tok == &b.Tok // token.Token
		} else {
			return false
		}

	case *ast.AssignStmt:
		if b, ok := b.(*ast.AssignStmt); ok {
			return a.TokPos == b.TokPos && // token.Pos
				&a.TokPos == &b.TokPos && // token.Pos
				a.Tok == b.Tok && // token.Token
				&a.Tok == &b.Tok // token.Token
		} else {
			return false
		}

	case *ast.GoStmt:
		if b, ok := b.(*ast.GoStmt); ok {
			return a.Go == b.Go && // token.Pos
				&a.Go == &b.Go // token.Pos
		} else {
			return false
		}

	case *ast.DeferStmt:
		if b, ok := b.(*ast.DeferStmt); ok {
			return a.Defer == b.Defer && // token.Pos
				&a.Defer == &b.Defer // token.Pos
		} else {
			return false
		}

	case *ast.ReturnStmt:
		if b, ok := b.(*ast.ReturnStmt); ok {
			return a.Return == b.Return && // token.Pos
				&a.Return == &b.Return // token.Pos
		} else {
			return false
		}

	case *ast.BranchStmt:
		if b, ok := b.(*ast.BranchStmt); ok {
			return a.TokPos == b.TokPos && // token.Pos
				&a.TokPos == &b.TokPos && // token.Pos
				a.Tok == b.Tok && // token.Token
				&a.Tok == &b.Tok // token.Token
		} else {
			return false
		}

	case *ast.BlockStmt:
		if b, ok := b.(*ast.BlockStmt); ok {
			return a.Lbrace == b.Lbrace && // token.Pos
				&a.Lbrace == &b.Lbrace && // token.Pos
				a.Rbrace == b.Rbrace && // token.Pos
				&a.Rbrace == &b.Rbrace // token.Pos
		} else {
			return false
		}

	case *ast.IfStmt:
		if b, ok := b.(*ast.IfStmt); ok {
			return a.If == b.If && // token.Pos
				&a.If == &b.If // token.Pos
		} else {
			return false
		}

	case *ast.CaseClause:
		if b, ok := b.(*ast.CaseClause); ok {
			return a.Case == b.Case && // token.Pos
				&a.Case == &b.Case && // token.Pos
				a.Colon == b.Colon && // token.Pos
				&a.Colon == &b.Colon // token.Pos
		} else {
			return false
		}

	case *ast.SwitchStmt:
		if b, ok := b.(*ast.SwitchStmt); ok {
			return a.Switch == b.Switch && // token.Pos
				&a.Switch == &b.Switch // token.Pos
		} else {
			return false
		}

	case *ast.TypeSwitchStmt:
		if b, ok := b.(*ast.TypeSwitchStmt); ok {
			return a.Switch == b.Switch && // token.Pos
				&a.Switch == &b.Switch // token.Pos
		} else {
			return false
		}

	case *ast.CommClause:
		if b, ok := b.(*ast.CommClause); ok {
			return a.Case == b.Case && // token.Pos
				&a.Case == &b.Case && // token.Pos
				a.Colon == b.Colon && // token.Pos
				&a.Colon == &b.Colon // token.Pos
		} else {
			return false
		}

	case *ast.SelectStmt:
		if b, ok := b.(*ast.SelectStmt); ok {
			return a.Select == b.Select && // token.Pos
				&a.Select == &b.Select // token.Pos
		} else {
			return false
		}

	case *ast.ForStmt:
		if b, ok := b.(*ast.ForStmt); ok {
			return a.For == b.For && // token.Pos
				&a.For == &b.For // token.Pos
		} else {
			return false
		}

	case *ast.RangeStmt:
		if b, ok := b.(*ast.RangeStmt); ok {
			return a.For == b.For && // token.Pos
				&a.For == &b.For && // token.Pos
				a.TokPos == b.TokPos && // token.Pos
				&a.TokPos == &b.TokPos && // token.Pos
				a.Tok == b.Tok && // token.Token
				&a.Tok == &b.Tok // token.Token
		} else {
			return false
		}

	case *ast.ImportSpec:
		if b, ok := b.(*ast.ImportSpec); ok {
			return a.EndPos == b.EndPos && // token.Pos
				&a.EndPos == &b.EndPos // token.Pos
		} else {
			return false
		}

	case *ast.TypeSpec:
		if b, ok := b.(*ast.TypeSpec); ok {
			return a.Assign == b.Assign && // token.Pos
				&a.Assign == &b.Assign // token.Pos
		} else {
			return false
		}

	case *ast.BadDecl:
		if b, ok := b.(*ast.BadDecl); ok {
			return a.From == b.From && // token.Pos
				&a.From == &b.From && // token.Pos
				a.To == b.To && // token.Pos
				&a.To == &b.To // token.Pos
		} else {
			return false
		}

	case *ast.GenDecl:
		if b, ok := b.(*ast.GenDecl); ok {
			return a.TokPos == b.TokPos && // token.Pos
				&a.TokPos == &b.TokPos && // token.Pos
				a.Tok == b.Tok && // token.Token
				&a.Tok == &b.Tok && // token.Token
				a.Lparen == b.Lparen && // token.Pos
				&a.Lparen == &b.Lparen && // token.Pos
				a.Rparen == b.Rparen && // token.Pos
				&a.Rparen == &b.Rparen // token.Pos
		} else {
			return false
		}

	case *ast.File:
		if b, ok := b.(*ast.File); ok {
			return a.Package == b.Package && // token.Pos
				&a.Package == &b.Package // token.Pos
		} else {
			return false
		}

	case *ast.Package:
		if b, ok := b.(*ast.Package); ok {
			return a.Name == b.Name && // string
				&a.Name == &b.Name // string
		}

	}
	return true
}

// returns true if they are same, uses BFS and value comparison instead of ponter comparison
// Returns false for changes in addresses also. In addition to only value changes just like CompareRecursively
func CompareRecursivelyWithAddresses(a, b ast.Node) bool {
	pairs := []*[2]ast.Node{}
	pairs = append(pairs, &[2]ast.Node{a, b})

	for len(pairs) > 0 {
		a, b := pairs[0][0], pairs[0][1]
		pairs = pairs[1:]

		if !CompareNonNodeFieldsWithAddresses(a, b) {
			return false
		}

		childrenA := ChildNodes(a)
		childrenB := ChildNodes(b)

		if len(childrenA) != len(childrenB) {
			return false
		}

		pairs = append(pairs, slicew.Zip(childrenA, childrenB)...)
	}

	return true
}

// Returns true if they are same.
// Returns true for those types that don't have non-node type field.
// Doesn't care about changes in addresses, in contrast to CompareNonNodeFieldWithAddresses.
func CompareNonNodeFields(a, b ast.Node) bool {
	switch a := a.(type) {

	case *ast.Comment:
		if b, ok := b.(*ast.Comment); ok {
			return a.Slash == b.Slash && // token.Pos
				a.Text == b.Text // string
		} else {
			return false
		}

	case *ast.FieldList:
		if b, ok := b.(*ast.FieldList); ok {
			return a.Opening == b.Opening && // token.Pos
				a.Closing == b.Closing // token.Pos
		} else {
			return false
		}

	case *ast.BadExpr:
		if b, ok := b.(*ast.BadExpr); ok {
			return a.From == b.From && // token.Pos
				a.To == b.To // token.Pos
		} else {
			return false
		}

	case *ast.Ident:
		if b, ok := b.(*ast.Ident); ok {
			return a.NamePos == b.NamePos && // token.Pos
				a.Name == b.Name // string
		} else {
			return false
		}

	case *ast.Ellipsis:
		if b, ok := b.(*ast.Ellipsis); ok {
			return a.Ellipsis == b.Ellipsis // token.Pos
		} else {
			return false
		}

	case *ast.BasicLit:
		if b, ok := b.(*ast.BasicLit); ok {
			return a.ValuePos == b.ValuePos && // token.Pos
				a.Kind == b.Kind && // token.Token
				a.Value == b.Value // string
		} else {
			return false
		}

	case *ast.CompositeLit:
		if b, ok := b.(*ast.CompositeLit); ok {
			return a.Lbrace == b.Lbrace && // token.Pos
				a.Rbrace == b.Rbrace && // token.Pos
				a.Incomplete == b.Incomplete // bool
		} else {
			return false
		}

	case *ast.ParenExpr:
		if b, ok := b.(*ast.ParenExpr); ok {
			return a.Lparen == b.Lparen && // token.Pos
				a.Rparen == b.Rparen // token.Pos
		} else {
			return false
		}

	case *ast.IndexExpr:
		if b, ok := b.(*ast.IndexExpr); ok {
			return a.Lbrack == b.Lbrack && // token.Pos
				a.Rbrack == b.Rbrack // token.Pos
		} else {
			return false
		}

	case *ast.IndexListExpr:
		if b, ok := b.(*ast.IndexListExpr); ok {
			return a.Lbrack == b.Lbrack && // token.Pos
				a.Rbrack == b.Rbrack // token.Pos
		} else {
			return false
		}

	case *ast.SliceExpr:
		if b, ok := b.(*ast.SliceExpr); ok {
			return a.Lbrack == b.Lbrack && // token.Pos
				a.Slice3 == b.Slice3 && // bool
				a.Rbrack == b.Rbrack // token.Pos
		} else {
			return false
		}

	case *ast.TypeAssertExpr:
		if b, ok := b.(*ast.TypeAssertExpr); ok {
			return a.Lparen == b.Lparen && // token.Pos
				a.Rparen == b.Rparen // token.Pos
		} else {
			return false
		}

	case *ast.CallExpr:
		if b, ok := b.(*ast.CallExpr); ok {
			return a.Lparen == b.Lparen && // token.Pos
				a.Ellipsis == b.Ellipsis && // token.Pos
				a.Rparen == b.Rparen //  token.Pos
		} else {
			return false
		}

	case *ast.StarExpr:
		if b, ok := b.(*ast.StarExpr); ok {
			return a.Star == b.Star // token.Pos
		} else {
			return false
		}

	case *ast.UnaryExpr:
		if b, ok := b.(*ast.UnaryExpr); ok {
			return a.OpPos == b.OpPos && // token.Pos
				a.Op == b.Op // token.Token
		} else {
			return false
		}

	case *ast.BinaryExpr:
		if b, ok := b.(*ast.BinaryExpr); ok {
			return a.OpPos == b.OpPos && // token.Pos
				a.Op == b.Op // token.Token
		} else {
			return false
		}

	case *ast.KeyValueExpr:
		if b, ok := b.(*ast.KeyValueExpr); ok {
			return a.Colon == b.Colon // token.Pos
		} else {
			return false
		}

	case *ast.ArrayType:
		if b, ok := b.(*ast.ArrayType); ok {
			return a.Lbrack == b.Lbrack // token.Pos
		} else {
			return false
		}

	case *ast.StructType:
		if b, ok := b.(*ast.StructType); ok {
			return a.Struct == b.Struct && // token.Pos
				a.Incomplete == b.Incomplete // bool
		} else {
			return false
		}

	case *ast.FuncType:
		if b, ok := b.(*ast.FuncType); ok {
			return a.Func == b.Func // token.Pos
		} else {
			return false
		}

	case *ast.InterfaceType:
		if b, ok := b.(*ast.InterfaceType); ok {
			return a.Interface == b.Interface && // token.Pos
				a.Incomplete == b.Incomplete // bool
		} else {
			return false
		}

	case *ast.MapType:
		if b, ok := b.(*ast.MapType); ok {
			return a.Map == b.Map // token.Pos
		} else {
			return false
		}

	case *ast.ChanType:
		if b, ok := b.(*ast.ChanType); ok {
			return a.Begin == b.Begin && // token.Pos
				a.Arrow == b.Arrow // token.Pos
		} else {
			return false
		}

	case *ast.BadStmt:
		if b, ok := b.(*ast.BadStmt); ok {
			return a.From == b.From && // token.Pos
				a.To == b.To // token.Pos
		} else {
			return false
		}

	case *ast.EmptyStmt:
		if b, ok := b.(*ast.EmptyStmt); ok {
			return a.Semicolon == b.Semicolon && // token.Pos
				a.Implicit == b.Implicit // bool
		} else {
			return false
		}

	case *ast.LabeledStmt:
		if b, ok := b.(*ast.LabeledStmt); ok {
			return a.Colon == b.Colon // token.Pos
		} else {
			return false
		}

	case *ast.SendStmt:
		if b, ok := b.(*ast.SendStmt); ok {
			return a.Arrow == b.Arrow // token.Pos
		} else {
			return false
		}

	case *ast.IncDecStmt:
		if b, ok := b.(*ast.IncDecStmt); ok {
			return a.TokPos == b.TokPos && // token.Pos
				a.Tok == b.Tok // token.Token
		} else {
			return false
		}

	case *ast.AssignStmt:
		if b, ok := b.(*ast.AssignStmt); ok {
			return a.TokPos == b.TokPos && // token.Pos
				a.Tok == b.Tok // token.Token
		} else {
			return false
		}

	case *ast.GoStmt:
		if b, ok := b.(*ast.GoStmt); ok {
			return a.Go == b.Go // token.Pos
		} else {
			return false
		}

	case *ast.DeferStmt:
		if b, ok := b.(*ast.DeferStmt); ok {
			return a.Defer == b.Defer // token.Pos
		} else {
			return false
		}

	case *ast.ReturnStmt:
		if b, ok := b.(*ast.ReturnStmt); ok {
			return a.Return == b.Return // token.Pos
		} else {
			return false
		}

	case *ast.BranchStmt:
		if b, ok := b.(*ast.BranchStmt); ok {
			return a.TokPos == b.TokPos && // token.Pos
				a.Tok == b.Tok // token.Token
		} else {
			return false
		}

	case *ast.BlockStmt:
		if b, ok := b.(*ast.BlockStmt); ok {
			return a.Lbrace == b.Lbrace && // token.Pos
				a.Rbrace == b.Rbrace // token.Pos
		} else {
			return false
		}

	case *ast.IfStmt:
		if b, ok := b.(*ast.IfStmt); ok {
			return a.If == b.If // token.Pos
		} else {
			return false
		}

	case *ast.CaseClause:
		if b, ok := b.(*ast.CaseClause); ok {
			return a.Case == b.Case && // token.Pos
				a.Colon == b.Colon // token.Pos
		} else {
			return false
		}

	case *ast.SwitchStmt:
		if b, ok := b.(*ast.SwitchStmt); ok {
			return a.Switch == b.Switch // token.Pos
		} else {
			return false
		}

	case *ast.TypeSwitchStmt:
		if b, ok := b.(*ast.TypeSwitchStmt); ok {
			return a.Switch == b.Switch // token.Pos
		} else {
			return false
		}

	case *ast.CommClause:
		if b, ok := b.(*ast.CommClause); ok {
			return a.Case == b.Case && // token.Pos
				a.Colon == b.Colon // token.Pos
		} else {
			return false
		}

	case *ast.SelectStmt:
		if b, ok := b.(*ast.SelectStmt); ok {
			return a.Select == b.Select // token.Pos
		} else {
			return false
		}

	case *ast.ForStmt:
		if b, ok := b.(*ast.ForStmt); ok {
			return a.For == b.For // token.Pos
		} else {
			return false
		}

	case *ast.RangeStmt:
		if b, ok := b.(*ast.RangeStmt); ok {
			return a.For == b.For && // token.Pos
				a.TokPos == b.TokPos && // token.Pos
				a.Tok == b.Tok // token.Token
		} else {
			return false
		}

	case *ast.ImportSpec:
		if b, ok := b.(*ast.ImportSpec); ok {
			return a.EndPos == b.EndPos // token.Pos
		} else {
			return false
		}

	case *ast.TypeSpec:
		if b, ok := b.(*ast.TypeSpec); ok {
			return a.Assign == b.Assign // token.Pos
		} else {
			return false
		}

	case *ast.BadDecl:
		if b, ok := b.(*ast.BadDecl); ok {
			return a.From == b.From && // token.Pos
				a.To == b.To // token.Pos
		} else {
			return false
		}

	case *ast.GenDecl:
		if b, ok := b.(*ast.GenDecl); ok {
			return a.TokPos == b.TokPos && // token.Pos
				a.Tok == b.Tok && // token.Token
				a.Lparen == b.Lparen && // token.Pos
				a.Rparen == b.Rparen // token.Pos
		} else {
			return false
		}

	case *ast.File:
		if b, ok := b.(*ast.File); ok {
			return a.Package == b.Package // token.Pos
		} else {
			return false
		}

	case *ast.Package:
		if b, ok := b.(*ast.Package); ok {
			return a.Name == b.Name // string
		}

	}
	return true
}

// returns true if they are same, uses BFS and value comparison instead of ponter comparison
// Doesn't care about changes in addresses, in contrast to CompareRecursivelWithAddresses.
func CompareRecursively(a, b ast.Node) bool {
	pairs := []*[2]ast.Node{}
	pairs = append(pairs, &[2]ast.Node{a, b})

	for len(pairs) > 0 {
		a, b := pairs[0][0], pairs[0][1]
		pairs = pairs[1:]

		if !CompareNonNodeFields(a, b) {
			return false
		}

		childrenA := ChildNodes(a)
		childrenB := ChildNodes(b)

		if len(childrenA) != len(childrenB) {
			return false
		}

		pairs = append(pairs, slicew.Zip(childrenA, childrenB)...)
	}

	return true
}
