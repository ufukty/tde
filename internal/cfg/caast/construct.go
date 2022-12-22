package caast

import "go/ast"

func Construct(node ast.Node, nodeType NodeType) ast.Node {
	var (
		stack       = Stack{}
		isPerformed = false
	)

	ancestry := []ast.Node{}

	ast.Inspect(node, func(n ast.Node) bool {

		if !isPerformed {
			return false
		}

		if n == nil {
			stack.Return()
		}

		if n == node {

		} else if IsInPath(ancestry, n) {
			stack.Fill(n)
			stack.Recurse()
			return true
		} else {
			return false
		}

		// if utilities.Coin() {

		// }
		return false
	})

	switch nodeType {

	// MARK: Expressions and types

	// case Field: // struct field
	// 	return &ast.Field{}
	// case FieldList:
	// 	return &ast.FieldList{}
	// case Ellipsis:
	// 	return &ast.Ellipsis{} // ... variadic argument spreading
	case CompositeLit:
		return &ast.CompositeLit{
			Type:       nil,
			Lbrace:     0,
			Elts:       []ast.Expr{},
			Rbrace:     0,
			Incomplete: false,
		}
	case BasicLit: // basic literal
		kind, value := GenerateRandomLiteral()
		return &ast.BasicLit{
			Kind:  kind,
			Value: value,
		}
	// case FuncLit:
	// 	return &ast.FuncLit{}
	case ParenExpr:
		return &ast.ParenExpr{}
	case SelectorExpr:
		return &ast.SelectorExpr{Sel: &ast.Ident{}}
	case IndexExpr:
		return &ast.IndexExpr{}
	case IndexListExpr:
		return &ast.IndexListExpr{}
	case SliceExpr:
		return &ast.SliceExpr{}
	case TypeAssertExpr:
		return &ast.TypeAssertExpr{}
	case CallExpr:
		return &ast.CallExpr{}
	case StarExpr:
		return &ast.StarExpr{}
	case UnaryExpr:
		return &ast.UnaryExpr{}
	case BinaryExpr:
		return &ast.BinaryExpr{}
	case KeyValueExpr:
		return &ast.KeyValueExpr{}

	// MARK: Types

	case ArrayType:
		return &ast.ArrayType{}
	case StructType:
		return &ast.StructType{}
	// case FuncType:
	// 	return &ast.FuncType{}
	// case InterfaceType:
	// 	return &ast.InterfaceType{}
	case MapType:
		return &ast.MapType{}
	// case ChanType:
	// 	return &ast.ChanType{}

	// MARK: Statements

	case BadStmt:
		return &ast.BadStmt{}
	case DeclStmt:
		return &ast.DeclStmt{}
	case EmptyStmt:
		return &ast.EmptyStmt{}
	case LabeledStmt:
		return &ast.LabeledStmt{}
	case ExprStmt:
		return &ast.ExprStmt{}
	case SendStmt:
		return &ast.SendStmt{}
	case IncDecStmt:
		return &ast.IncDecStmt{}
	case AssignStmt:
		return &ast.AssignStmt{}
	// case GoStmt:
	// 	return &ast.GoStmt{}
	// case DeferStmt:
	// 	return &ast.DeferStmt{}
	case ReturnStmt:
		return &ast.ReturnStmt{Results: GenerateReturnStatement()}
	// case BranchStmt:
	// 	return &ast.BranchStmt{}
	case BlockStmt:
		return &ast.BlockStmt{List: GenerateRandomNumberOfStatements()}
	case IfStmt:
		return &ast.IfStmt{
			Cond: GenerateBooleanExpression(),
			Body: GenerateBlockStatement(),
			Else: GenerateBlockStatement(),
		}
	// case CaseClause:
	// 	return &ast.CaseClause{}
	// case SwitchStmt:
	// 	return &ast.SwitchStmt{}
	// case TypeSwitchStmt:
	// 	return &ast.TypeSwitchStmt{}
	// case CommClause:
	// 	return &ast.CommClause{}
	// case SelectStmt:
	// 	return &ast.SelectStmt{}
	case ForStmt:
		return &ast.ForStmt{}
	case RangeStmt:
		return &ast.RangeStmt{}

	// MARK: Declarations/Spec

	case ImportSpec:
		return &ast.ImportSpec{}
	case ValueSpec:
		return &ast.ValueSpec{}
	case TypeSpec:
		return &ast.TypeSpec{}

		// MARK: Declarations/Other

		// case BadDecl:
		// 	return &ast.BadDecl{}
		// case GenDecl:
		// 	return &ast.GenDecl{}
		// case FuncDecl:
		// 	return &ast.FuncDecl{}

		// MARK: Files & packages

		// case File:
		// 	return &ast.File{}
		// case Package:
		// 	return &ast.Package{}
	}

	return nil
}
