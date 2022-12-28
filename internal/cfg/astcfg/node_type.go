package astcfg

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
)

var stringRepresentation = map[NodeType]string{
	ArrayType:      "ArrayType",
	AssignStmt:     "AssignStmt",
	BadDecl:        "BadDecl",
	BadExpr:        "BadExpr",
	BadStmt:        "BadStmt",
	BasicLit:       "BasicLit",
	BinaryExpr:     "BinaryExpr",
	BlockStmt:      "BlockStmt",
	BranchStmt:     "BranchStmt",
	CallExpr:       "CallExpr",
	CaseClause:     "CaseClause",
	ChanType:       "ChanType",
	CommClause:     "CommClause",
	Comment:        "Comment",
	CommentGroup:   "CommentGroup",
	CompositeLit:   "CompositeLit",
	DeclStmt:       "DeclStmt",
	DeferStmt:      "DeferStmt",
	Ellipsis:       "Ellipsis",
	EmptyStmt:      "EmptyStmt",
	ExprStmt:       "ExprStmt",
	Field:          "Field",
	FieldList:      "FieldList",
	File:           "File",
	ForStmt:        "ForStmt",
	FuncDecl:       "FuncDecl",
	FuncLit:        "FuncLit",
	FuncType:       "FuncType",
	GenDecl:        "GenDecl",
	GoStmt:         "GoStmt",
	Ident:          "Ident",
	IfStmt:         "IfStmt",
	ImportSpec:     "ImportSpec",
	IncDecStmt:     "IncDecStmt",
	IndexExpr:      "IndexExpr",
	IndexListExpr:  "IndexListExpr",
	InterfaceType:  "InterfaceType",
	KeyValueExpr:   "KeyValueExpr",
	LabeledStmt:    "LabeledStmt",
	MapType:        "MapType",
	Package:        "Package",
	ParenExpr:      "ParenExpr",
	RangeStmt:      "RangeStmt",
	ReturnStmt:     "ReturnStmt",
	SelectorExpr:   "SelectorExpr",
	SelectStmt:     "SelectStmt",
	SendStmt:       "SendStmt",
	SliceExpr:      "SliceExpr",
	StarExpr:       "StarExpr",
	StructType:     "StructType",
	SwitchStmt:     "SwitchStmt",
	TypeAssertExpr: "TypeAssertExpr",
	TypeSpec:       "TypeSpec",
	TypeSwitchStmt: "TypeSwitchStmt",
	UnaryExpr:      "UnaryExpr",
	ValueSpec:      "ValueSpec",
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
