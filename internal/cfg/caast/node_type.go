package caast

type NodeType int

const (
	BasicLit NodeType = iota
	BadExpr
	Comment
	CommentGroup
	Ident
	Field
	FieldList
	Ellipsis
	FuncLit
	CompositeLit
	ParenExpr
	SelectorExpr
	IndexExpr
	IndexListExpr
	SliceExpr
	TypeAssertExpr
	CallExpr
	StarExpr
	UnaryExpr
	BinaryExpr
	KeyValueExpr
	ArrayType
	StructType
	FuncType
	InterfaceType
	MapType
	ChanType
	BadStmt
	DeclStmt
	EmptyStmt
	LabeledStmt
	ExprStmt
	SendStmt
	IncDecStmt
	AssignStmt
	GoStmt
	DeferStmt
	ReturnStmt
	BranchStmt
	BlockStmt
	IfStmt
	CaseClause
	SwitchStmt
	TypeSwitchStmt
	CommClause
	SelectStmt
	ForStmt
	RangeStmt
	ImportSpec
	ValueSpec
	TypeSpec
	BadDecl
	GenDecl
	FuncDecl
	File
	Package
)

var stringRepresentation = map[NodeType]string{
	BasicLit:       "BasicLit",
	BadExpr:        "BadExpr",
	Comment:        "Comment",
	CommentGroup:   "CommentGroup",
	Ident:          "Ident",
	Field:          "Field",
	FieldList:      "FieldList",
	Ellipsis:       "Ellipsis",
	FuncLit:        "FuncLit",
	CompositeLit:   "CompositeLit",
	ParenExpr:      "ParenExpr",
	SelectorExpr:   "SelectorExpr",
	IndexExpr:      "IndexExpr",
	IndexListExpr:  "IndexListExpr",
	SliceExpr:      "SliceExpr",
	TypeAssertExpr: "TypeAssertExpr",
	CallExpr:       "CallExpr",
	StarExpr:       "StarExpr",
	UnaryExpr:      "UnaryExpr",
	BinaryExpr:     "BinaryExpr",
	KeyValueExpr:   "KeyValueExpr",
	ArrayType:      "ArrayType",
	StructType:     "StructType",
	FuncType:       "FuncType",
	InterfaceType:  "InterfaceType",
	MapType:        "MapType",
	ChanType:       "ChanType",
	BadStmt:        "BadStmt",
	DeclStmt:       "DeclStmt",
	EmptyStmt:      "EmptyStmt",
	LabeledStmt:    "LabeledStmt",
	ExprStmt:       "ExprStmt",
	SendStmt:       "SendStmt",
	IncDecStmt:     "IncDecStmt",
	AssignStmt:     "AssignStmt",
	GoStmt:         "GoStmt",
	DeferStmt:      "DeferStmt",
	ReturnStmt:     "ReturnStmt",
	BranchStmt:     "BranchStmt",
	BlockStmt:      "BlockStmt",
	IfStmt:         "IfStmt",
	CaseClause:     "CaseClause",
	SwitchStmt:     "SwitchStmt",
	TypeSwitchStmt: "TypeSwitchStmt",
	CommClause:     "CommClause",
	SelectStmt:     "SelectStmt",
	ForStmt:        "ForStmt",
	RangeStmt:      "RangeStmt",
	ImportSpec:     "ImportSpec",
	ValueSpec:      "ValueSpec",
	TypeSpec:       "TypeSpec",
	BadDecl:        "BadDecl",
	GenDecl:        "GenDecl",
	FuncDecl:       "FuncDecl",
	File:           "File",
	Package:        "Package",
}

func (n NodeType) String() string {
	return stringRepresentation[n]
}

var orderedNodeTypes = []NodeType{}

var (
	cumulativeProbabilities           []float64
	cumulativeProbabilitiesUpperBound float64
)

func init() {
	calculateCumulativeProbabilities()
}

func calculateCumulativeProbabilities() {
	cumulativeProbabilities = []float64{}
	var total = 0.0
	for nodeType, prob := range probabilities {
		if prob != 0.0 {
			orderedNodeTypes = append(orderedNodeTypes, nodeType)
			cumulativeProbabilities = append(cumulativeProbabilities, total)
			total += prob
		}
	}
	cumulativeProbabilitiesUpperBound = total
}
