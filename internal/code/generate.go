package code

import (
	"tde/internal/utilities"

	"go/ast"
)

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

var probabilities = map[NodeType]float64{
	// first level type
	DeclStmt:     0.1,
	AssignStmt:   0.1,
	CompositeLit: 0.1, // a literal that combines other type of nodes and used as value to fill an Ident
	IfStmt:       0.0,
	ForStmt:      0.0,
	ReturnStmt:   0.0,

	// second level type
	BasicLit:  0.0, // basic literal, like string, int, float value
	Ident:     0.0, // ref a variable (to declare/assign/access)
	ArrayType: 0.0,
	MapType:   0.0,

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
	CallExpr:       0.0,
	StarExpr:       0.0,
	UnaryExpr:      0.0,
	BinaryExpr:     0.0,
	KeyValueExpr:   0.0,
	StructType:     0.0,
	FuncType:       0.0,
	InterfaceType:  0.0,
	ChanType:       0.0,
	EmptyStmt:      0.0,
	LabeledStmt:    0.0,
	ExprStmt:       0.0,
	SendStmt:       0.0,
	IncDecStmt:     0.0,
	GoStmt:         0.0,
	DeferStmt:      0.0,
	BranchStmt:     0.0,
	BlockStmt:      0.0,
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

type NodeTypeClass int

const (
	Expression = NodeTypeClass(iota)
	Statement
)

var NodeTypeClasses = map[NodeTypeClass][]NodeType{
	Expression: {
		BadExpr,
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
	},
	Statement: {
		BadStmt,
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
		RangeStmt,
	},
}

var NodeTydpeClasses = map[NodeType]NodeTypeClass{
	BadExpr:        Expression,
	Ident:          Expression,
	Ellipsis:       Expression,
	BasicLit:       Expression,
	FuncLit:        Expression,
	CompositeLit:   Expression,
	ParenExpr:      Expression,
	SelectorExpr:   Expression,
	IndexExpr:      Expression,
	IndexListExpr:  Expression,
	SliceExpr:      Expression,
	TypeAssertExpr: Expression,
	CallExpr:       Expression,
	StarExpr:       Expression,
	UnaryExpr:      Expression,
	BinaryExpr:     Expression,
	KeyValueExpr:   Expression,
	ArrayType:      Expression,
	StructType:     Expression,
	FuncType:       Expression,
	InterfaceType:  Expression,
	MapType:        Expression,
	ChanType:       Expression,
	BadStmt:        Statement,
	DeclStmt:       Statement,
	EmptyStmt:      Statement,
	LabeledStmt:    Statement,
	ExprStmt:       Statement,
	SendStmt:       Statement,
	IncDecStmt:     Statement,
	AssignStmt:     Statement,
	GoStmt:         Statement,
	DeferStmt:      Statement,
	ReturnStmt:     Statement,
	BranchStmt:     Statement,
	BlockStmt:      Statement,
	IfStmt:         Statement,
	CaseClause:     Statement,
	SwitchStmt:     Statement,
	TypeSwitchStmt: Statement,
	CommClause:     Statement,
	SelectStmt:     Statement,
	ForStmt:        Statement,
	RangeStmt:      Statement,
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

	var (
		total = 0.0
	)
	for nodeType, prob := range probabilities {
		if prob != 0.0 {
			orderedNodeTypes = append(orderedNodeTypes, nodeType)
			cumulativeProbabilities = append(cumulativeProbabilities, total)
			total += prob
		}
	}
	cumulativeProbabilitiesUpperBound = total
}

func PickRandomNodeType() NodeType {
	rand := utilities.URandFloatForCrypto() * cumulativeProbabilitiesUpperBound
	index := utilities.BinaryRangeSearch(cumulativeProbabilities, rand)
	return orderedNodeTypes[index]
}

func Generate(kind NodeType) ast.Node {
	var node ast.Node

	switch kind {
	case BasicLit:
		node = &ast.BasicLit{
			ValuePos: 0,
			Kind:     0,
			Value:    "",
		}

	case Field:
		node = &ast.Field{}

	case FieldList:
		node = &ast.FieldList{}

	case Ellipsis:
		node = &ast.Ellipsis{}

	case FuncLit:
		node = &ast.FuncLit{}

	case CompositeLit:
		node = &ast.CompositeLit{}

	case ParenExpr:
		node = &ast.ParenExpr{}

	case SelectorExpr:
		node = &ast.SelectorExpr{}

	case IndexExpr:
		node = &ast.IndexExpr{}

	case IndexListExpr:
		node = &ast.IndexListExpr{}

	case SliceExpr:
		node = &ast.SliceExpr{}

	case TypeAssertExpr:
		node = &ast.TypeAssertExpr{}

	case CallExpr:
		node = &ast.CallExpr{}

	case StarExpr:
		node = &ast.StarExpr{}

	case UnaryExpr:
		node = &ast.UnaryExpr{}

	case BinaryExpr:
		node = &ast.BinaryExpr{}

	case KeyValueExpr:
		node = &ast.KeyValueExpr{}

	case ArrayType:
		node = &ast.ArrayType{}

	case StructType:
		node = &ast.StructType{}

	case FuncType:
		node = &ast.FuncType{}

	case InterfaceType:
		node = &ast.InterfaceType{}

	case MapType:
		node = &ast.MapType{}

	case ChanType:
		node = &ast.ChanType{}

	case BadStmt:
		node = &ast.BadStmt{}

	case DeclStmt:
		node = &ast.DeclStmt{}

	case EmptyStmt:
		node = &ast.EmptyStmt{}

	case LabeledStmt:
		node = &ast.LabeledStmt{}

	case ExprStmt:
		node = &ast.ExprStmt{}

	case SendStmt:
		node = &ast.SendStmt{}

	case IncDecStmt:
		node = &ast.IncDecStmt{}

	case AssignStmt:
		node = &ast.AssignStmt{}

	case GoStmt:
		node = &ast.GoStmt{}

	case DeferStmt:
		node = &ast.DeferStmt{}

	case ReturnStmt:
		node = &ast.ReturnStmt{}

	case BranchStmt:
		node = &ast.BranchStmt{}

	case BlockStmt:
		node = &ast.BlockStmt{}

	case IfStmt:
		node = &ast.IfStmt{
			If:   0,
			Init: nil,
			Cond: nil,
			Body: Generate(BlockStmt).(*ast.BlockStmt),
			Else: nil,
		}

	case CaseClause:
		node = &ast.CaseClause{}

	case SwitchStmt:
		node = &ast.SwitchStmt{}

	case TypeSwitchStmt:
		node = &ast.TypeSwitchStmt{}

	case CommClause:
		node = &ast.CommClause{}

	case SelectStmt:
		node = &ast.SelectStmt{}

	case ForStmt:
		node = &ast.ForStmt{}

	case RangeStmt:
		node = &ast.RangeStmt{}

	case ImportSpec:
		node = &ast.ImportSpec{}

	case ValueSpec:
		node = &ast.ValueSpec{}

	case TypeSpec:
		node = &ast.TypeSpec{}

	case BadDecl:
		node = &ast.BadDecl{}

	case GenDecl:
		node = &ast.GenDecl{}

	case FuncDecl:
		node = &ast.FuncDecl{
			Doc:  &ast.CommentGroup{},
			Recv: &ast.FieldList{},
			Name: &ast.Ident{},
			Type: &ast.FuncType{},
			Body: &ast.BlockStmt{},
		}

	case File:
		node = &ast.File{}

	case Package:
		node = &ast.Package{}

	}
	return node
}
