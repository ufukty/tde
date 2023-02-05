package types

//go:generate stringer -type=NodeType

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
	TypeExpr

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
	case Expr, Stmt, Decl, Spec, TypeExpr:
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
		ChanType,
		TypeExpr:
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
