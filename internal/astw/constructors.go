package astw

import "go/ast"

func ConstructBasicNodes(nodeType NodeType) ast.Node {
	switch nodeType {
	case ArrayType:
		return &ast.ArrayType{}
	case AssignStmt:
		return &ast.AssignStmt{}
	case BadDecl:
		return &ast.BadDecl{}
	case BadExpr:
		return &ast.BadExpr{}
	case BadStmt:
		return &ast.BadStmt{}
	case BasicLit:
		return &ast.BasicLit{}
	case BinaryExpr:
		return &ast.BinaryExpr{}
	case BlockStmt:
		return &ast.BlockStmt{}
	case BranchStmt:
		return &ast.BranchStmt{}
	case CallExpr:
		return &ast.CallExpr{}
	case CaseClause:
		return &ast.CaseClause{}
	case ChanType:
		return &ast.ChanType{}
	case CommClause:
		return &ast.CommClause{}
	case Comment:
		return &ast.Comment{}
	case CommentGroup:
		return &ast.CommentGroup{}
	case CompositeLit:
		return &ast.CompositeLit{}
	case DeclStmt:
		return &ast.DeclStmt{}
	case DeferStmt:
		return &ast.DeferStmt{}
	case Ellipsis:
		return &ast.Ellipsis{}
	case EmptyStmt:
		return &ast.EmptyStmt{}
	case ExprStmt:
		return &ast.ExprStmt{}
	case Field:
		return &ast.Field{}
	case FieldList:
		return &ast.FieldList{}
	case File:
		return &ast.File{}
	case ForStmt:
		return &ast.ForStmt{}
	case FuncDecl:
		return &ast.FuncDecl{}
	case FuncLit:
		return &ast.FuncLit{}
	case FuncType:
		return &ast.FuncType{}
	case GenDecl:
		return &ast.GenDecl{}
	case GoStmt:
		return &ast.GoStmt{}
	case Ident:
		return &ast.Ident{}
	case IfStmt:
		return &ast.IfStmt{}
	case ImportSpec:
		return &ast.ImportSpec{}
	case IncDecStmt:
		return &ast.IncDecStmt{}
	case IndexExpr:
		return &ast.IndexExpr{}
	case IndexListExpr:
		return &ast.IndexListExpr{}
	case InterfaceType:
		return &ast.InterfaceType{}
	case KeyValueExpr:
		return &ast.KeyValueExpr{}
	case LabeledStmt:
		return &ast.LabeledStmt{}
	case MapType:
		return &ast.MapType{}
	case Package:
		return &ast.Package{}
	case ParenExpr:
		return &ast.ParenExpr{}
	case RangeStmt:
		return &ast.RangeStmt{}
	case ReturnStmt:
		return &ast.ReturnStmt{}
	case SelectorExpr:
		return &ast.SelectorExpr{}
	case SelectStmt:
		return &ast.SelectStmt{}
	case SendStmt:
		return &ast.SendStmt{}
	case SliceExpr:
		return &ast.SliceExpr{}
	case StarExpr:
		return &ast.StarExpr{}
	case StructType:
		return &ast.StructType{}
	case SwitchStmt:
		return &ast.SwitchStmt{}
	case TypeAssertExpr:
		return &ast.TypeAssertExpr{}
	case TypeSpec:
		return &ast.TypeSpec{}
	case TypeSwitchStmt:
		return &ast.TypeSwitchStmt{}
	case UnaryExpr:
		return &ast.UnaryExpr{}
	case ValueSpec:
		return &ast.ValueSpec{}
	}
	return nil
}

func ConstructNodeSlices(nodeType NodeType) any {
	switch nodeType {
	case DeclSlice:
		return []ast.Decl{}
	case ExprSlice:
		return []ast.Expr{}
	case FieldSlice:
		return []*ast.Field{}
	case IdentSlice:
		return []*ast.Ident{}
	case ImportSpecSlice:
		return []*ast.ImportSpec{}
	case SpecSlice:
		return []ast.Spec{}
	case StmtSlice:
		return []*ast.Stmt{}
	}
	return nil
}

// Important: Only use for test purposes
func ConstructMockInterfaceTypes(nodeType NodeType) ast.Node {
	switch nodeType {
	case Expr:
		return &ast.BadExpr{}
	case Stmt:
		return &ast.BadStmt{}
	case Decl:
		return &ast.BadDecl{}
	case Spec:
		return &ast.ValueSpec{}
	default:
		return nil
	}
}
