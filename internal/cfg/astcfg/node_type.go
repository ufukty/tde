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
	// Custom types:
	IdentType
)

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
