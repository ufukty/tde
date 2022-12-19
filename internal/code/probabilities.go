package code

var probabilities = map[NodeType]float64{
	// first level type
	DeclStmt:   0.1,
	AssignStmt: 0.1,
	IfStmt:     0.0,
	ForStmt:    0.0,
	ReturnStmt: 0.0,
	BlockStmt:  0.0,

	CallExpr: 0.0, // function call

	// variable type
	Ident:        0.0, // ref a variable (to declare/assign/access)
	BasicLit:     0.0, // basic literal, like string, int, float value
	CompositeLit: 0.1, // a literal that combines other type of nodes and used as value to fill an Ident
	ArrayType:    0.0,
	MapType:      0.0,
	// FuncType:  0.0,
	// ChanType:  0.0,

	// third level type

	// unfunctional
	BadExpr:      0.0,
	BadStmt:      0.0,
	BadDecl:      0.0,
	Comment:      0.0,
	CommentGroup: 0.0,

	// todo (automatic function definition)

	// rest
	Field:          0.0, // struct rel.
	FieldList:      0.0, // struct rel.
	Ellipsis:       0.0, // variadic function
	FuncLit:        0.0, // function literal, eg. lambda functions
	ParenExpr:      0.0,
	SelectorExpr:   0.0,
	IndexExpr:      0.0,
	IndexListExpr:  0.0,
	SliceExpr:      0.0,
	TypeAssertExpr: 0.0,
	StarExpr:       0.0,
	UnaryExpr:      0.0,
	BinaryExpr:     0.0,
	KeyValueExpr:   0.0,
	StructType:     0.0,
	InterfaceType:  0.0,
	EmptyStmt:      0.0,
	LabeledStmt:    0.0,
	ExprStmt:       0.0,
	SendStmt:       0.0,
	IncDecStmt:     0.0,
	GoStmt:         0.0,
	DeferStmt:      0.0,
	BranchStmt:     0.0,

	CaseClause:     0.0,
	SwitchStmt:     0.0,
	TypeSwitchStmt: 0.0,
	CommClause:     0.0,
	SelectStmt:     0.0,
	RangeStmt:      0.0,
	ImportSpec:     0.0,
	ValueSpec:      0.0,
	TypeSpec:       0.0,
	GenDecl:        0.0,
	FuncDecl:       0.0,
	File:           0.0,
	Package:        0.0,
}
