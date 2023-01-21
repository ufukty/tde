package ast_wrapper

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
	Slice
	SpecSlice
	StmtSlice
)

func (nt NodeType) IsSliceType() bool {
	switch nt {
	case ExprSlice, StmtSlice, DeclSlice, SpecSlice, IdentSlice:
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
