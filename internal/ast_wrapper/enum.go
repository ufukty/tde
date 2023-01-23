package ast_wrapper

import "go/ast"

type NodeType int

const (
	ArrayType = NodeType(iota)
	AssignStmt
	BadDecl
	BadExpr
	BadStmt
	BasicLit
	BinaryExpr
	BlockStmt
	BranchStmt
	CallExpr
	CaseClause
	ChanType
	CommClause
	Comment
	CommentGroup
	CompositeLit
	DeclStmt
	DeferStmt
	Ellipsis
	EmptyStmt
	ExprStmt
	Field
	FieldList
	File
	ForStmt
	FuncDecl
	FuncLit
	FuncType
	GenDecl
	GoStmt
	Ident
	IfStmt
	ImportSpec
	IncDecStmt
	IndexExpr
	IndexListExpr
	InterfaceType
	KeyValueExpr
	LabeledStmt
	MapType
	Package
	ParenExpr
	RangeStmt
	ReturnStmt
	SelectorExpr
	SelectStmt
	SendStmt
	SliceExpr
	StarExpr
	StructType
	SwitchStmt
	TypeAssertExpr
	TypeSpec
	TypeSwitchStmt
	UnaryExpr
	ValueSpec

	// Interfaces
	Expr
	Stmt
	Decl
	Spec

	// Slices
	CommentGroupSlice
	CommentSlice
	DeclSlice
	ExprSlice
	FieldSlice
	IdentSlice
	ImportSpecSlice
	SpecSlice
	StmtSlice
)

func (nt NodeType) String() string {
	switch nt {

	case ArrayType:
		return "ArrayType"
	case AssignStmt:
		return "AssignStmt"
	case BadDecl:
		return "BadDecl"
	case BadExpr:
		return "BadExpr"
	case BadStmt:
		return "BadStmt"
	case BasicLit:
		return "BasicLit"
	case BinaryExpr:
		return "BinaryExpr"
	case BlockStmt:
		return "BlockStmt"
	case BranchStmt:
		return "BranchStmt"
	case CallExpr:
		return "CallExpr"
	case CaseClause:
		return "CaseClause"
	case ChanType:
		return "ChanType"
	case CommClause:
		return "CommClause"
	case Comment:
		return "Comment"
	case CommentGroup:
		return "CommentGroup"
	case CompositeLit:
		return "CompositeLit"
	case DeclStmt:
		return "DeclStmt"
	case DeferStmt:
		return "DeferStmt"
	case Ellipsis:
		return "Ellipsis"
	case EmptyStmt:
		return "EmptyStmt"
	case ExprStmt:
		return "ExprStmt"
	case Field:
		return "Field"
	case FieldList:
		return "FieldList"
	case File:
		return "File"
	case ForStmt:
		return "ForStmt"
	case FuncDecl:
		return "FuncDecl"
	case FuncLit:
		return "FuncLit"
	case FuncType:
		return "FuncType"
	case GenDecl:
		return "GenDecl"
	case GoStmt:
		return "GoStmt"
	case Ident:
		return "Ident"
	case IfStmt:
		return "IfStmt"
	case ImportSpec:
		return "ImportSpec"
	case IncDecStmt:
		return "IncDecStmt"
	case IndexExpr:
		return "IndexExpr"
	case IndexListExpr:
		return "IndexListExpr"
	case InterfaceType:
		return "InterfaceType"
	case KeyValueExpr:
		return "KeyValueExpr"
	case LabeledStmt:
		return "LabeledStmt"
	case MapType:
		return "MapType"
	case Package:
		return "Package"
	case ParenExpr:
		return "ParenExpr"
	case RangeStmt:
		return "RangeStmt"
	case ReturnStmt:
		return "ReturnStmt"
	case SelectorExpr:
		return "SelectorExpr"
	case SelectStmt:
		return "SelectStmt"
	case SendStmt:
		return "SendStmt"
	case SliceExpr:
		return "SliceExpr"
	case StarExpr:
		return "StarExpr"
	case StructType:
		return "StructType"
	case SwitchStmt:
		return "SwitchStmt"
	case TypeAssertExpr:
		return "TypeAssertExpr"
	case TypeSpec:
		return "TypeSpec"
	case TypeSwitchStmt:
		return "TypeSwitchStmt"
	case UnaryExpr:
		return "UnaryExpr"
	case ValueSpec:
		return "ValueSpec"
	case Expr:
		return "Expr"
	case Stmt:
		return "Stmt"
	case Decl:
		return "Decl"
	case Spec:
		return "Spec"
	case CommentGroupSlice:
		return "CommentGroupSlice"
	case CommentSlice:
		return "CommentSlice"
	case DeclSlice:
		return "DeclSlice"
	case ExprSlice:
		return "ExprSlice"
	case FieldSlice:
		return "FieldSlice"
	case IdentSlice:
		return "IdentSlice"
	case ImportSpecSlice:
		return "ImportSpecSlice"
	case SpecSlice:
		return "SpecSlice"
	case StmtSlice:
		return "StmtSlice"
	}
	return ""
}

func (nt NodeType) IsSliceType() bool {
	switch nt {
	case CommentGroupSlice, CommentSlice, DeclSlice, ExprSlice, FieldSlice,
		IdentSlice, ImportSpecSlice, SpecSlice, StmtSlice:
		return true
	default:
		return false
	}
}

func (nt NodeType) IsInterfaceType() bool {
	switch nt {
	case Expr, Stmt, Decl, Spec:
		return true
	default:
		return false
	}
}

func (nt NodeType) IsConcreteType() bool {
	switch nt {
	case ArrayType,
		AssignStmt,
		BadDecl,
		BadExpr,
		BadStmt,
		BasicLit,
		BinaryExpr,
		BlockStmt,
		BranchStmt,
		CallExpr,
		CaseClause,
		ChanType,
		CommClause,
		Comment,
		CommentGroup,
		CompositeLit,
		DeclStmt,
		DeferStmt,
		Ellipsis,
		EmptyStmt,
		ExprStmt,
		Field,
		FieldList,
		File,
		ForStmt,
		FuncDecl,
		FuncLit,
		FuncType,
		GenDecl,
		GoStmt,
		Ident,
		IfStmt,
		ImportSpec,
		IncDecStmt,
		IndexExpr,
		IndexListExpr,
		InterfaceType,
		KeyValueExpr,
		LabeledStmt,
		MapType,
		Package,
		ParenExpr,
		RangeStmt,
		ReturnStmt,
		SelectorExpr,
		SelectStmt,
		SendStmt,
		SliceExpr,
		StarExpr,
		StructType,
		SwitchStmt,
		TypeAssertExpr,
		TypeSpec,
		TypeSwitchStmt,
		UnaryExpr,
		ValueSpec:
		return true
	default:
		return false
	}
}

func (nt NodeType) IsExpr() bool {
	switch nt {
	case BadExpr,
		Ident,
		Ellipsis,
		BasicLit,
		FuncLit,
		CompositeLit,
		ParenExpr,
		SelectorExpr,
		IndexExpr,
		IndexListExpr,
		SliceExpr,
		TypeAssertExpr,
		CallExpr,
		StarExpr,
		UnaryExpr,
		BinaryExpr,
		KeyValueExpr,
		ArrayType,
		StructType,
		FuncType,
		InterfaceType,
		MapType,
		ChanType:
		return true
	default:
		return false
	}
}

func (nt NodeType) IsStmt() bool {
	switch nt {
	case BadStmt,
		DeclStmt,
		EmptyStmt,
		LabeledStmt,
		ExprStmt,
		SendStmt,
		IncDecStmt,
		AssignStmt,
		GoStmt,
		DeferStmt,
		ReturnStmt,
		BranchStmt,
		BlockStmt,
		IfStmt,
		CaseClause,
		SwitchStmt,
		TypeSwitchStmt,
		CommClause,
		SelectStmt,
		ForStmt,
		RangeStmt:
		return true
	default:
		return false
	}
}

func (nt NodeType) IsDecl() bool {
	switch nt {
	case BadDecl, GenDecl, FuncDecl:
		return true
	default:
		return false
	}
}

func (nt NodeType) Construct() any {
	if nt.IsInterfaceType() {
		return ConstructMockInterfaceTypes(nt)
	} else if nt.IsSliceType() {
		return ConstructNodeSlices(nt)
	} else {
		return ConstructBasicNodes(nt)
	}
}

func GetNodeTypeForASTNode(n ast.Node) NodeType {
	switch n.(type) {
	case *ast.ArrayType:
		return ArrayType
	case *ast.AssignStmt:
		return AssignStmt
	case *ast.BadDecl:
		return BadDecl
	case *ast.BadExpr:
		return BadExpr
	case *ast.BadStmt:
		return BadStmt
	case *ast.BasicLit:
		return BasicLit
	case *ast.BinaryExpr:
		return BinaryExpr
	case *ast.BlockStmt:
		return BlockStmt
	case *ast.BranchStmt:
		return BranchStmt
	case *ast.CallExpr:
		return CallExpr
	case *ast.CaseClause:
		return CaseClause
	case *ast.ChanType:
		return ChanType
	case *ast.CommClause:
		return CommClause
	case *ast.Comment:
		return Comment
	case *ast.CommentGroup:
		return CommentGroup
	case *ast.CompositeLit:
		return CompositeLit
	case *ast.DeclStmt:
		return DeclStmt
	case *ast.DeferStmt:
		return DeferStmt
	case *ast.Ellipsis:
		return Ellipsis
	case *ast.EmptyStmt:
		return EmptyStmt
	case *ast.ExprStmt:
		return ExprStmt
	case *ast.Field:
		return Field
	case *ast.FieldList:
		return FieldList
	case *ast.File:
		return File
	case *ast.ForStmt:
		return ForStmt
	case *ast.FuncDecl:
		return FuncDecl
	case *ast.FuncLit:
		return FuncLit
	case *ast.FuncType:
		return FuncType
	case *ast.GenDecl:
		return GenDecl
	case *ast.GoStmt:
		return GoStmt
	case *ast.Ident:
		return Ident
	case *ast.IfStmt:
		return IfStmt
	case *ast.ImportSpec:
		return ImportSpec
	case *ast.IncDecStmt:
		return IncDecStmt
	case *ast.IndexExpr:
		return IndexExpr
	case *ast.IndexListExpr:
		return IndexListExpr
	case *ast.InterfaceType:
		return InterfaceType
	case *ast.KeyValueExpr:
		return KeyValueExpr
	case *ast.LabeledStmt:
		return LabeledStmt
	case *ast.MapType:
		return MapType
	case *ast.Package:
		return Package
	case *ast.ParenExpr:
		return ParenExpr
	case *ast.RangeStmt:
		return RangeStmt
	case *ast.ReturnStmt:
		return ReturnStmt
	case *ast.SelectorExpr:
		return SelectorExpr
	case *ast.SelectStmt:
		return SelectStmt
	case *ast.SendStmt:
		return SendStmt
	case *ast.SliceExpr:
		return SliceExpr
	case *ast.StarExpr:
		return StarExpr
	case *ast.StructType:
		return StructType
	case *ast.SwitchStmt:
		return SwitchStmt
	case *ast.TypeAssertExpr:
		return TypeAssertExpr
	case *ast.TypeSpec:
		return TypeSpec
	case *ast.TypeSwitchStmt:
		return TypeSwitchStmt
	case *ast.UnaryExpr:
		return UnaryExpr
	case *ast.ValueSpec:
		return ValueSpec
	default:
		panic("Unhandled condition GetNodeTypeForASTNode")
	}
}
