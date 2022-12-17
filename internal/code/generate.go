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

type NodeClass int

const (
	Expression = NodeClass(iota)
	Statement
)

var NodeTypeClasses = map[NodeClass][]NodeType{
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

var NodeTydpeClasses = map[NodeType]NodeClass{
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
	// TODO:
	case BasicLit:
		node = &ast.BasicLit{}

	// TODO:
	case Field:
		node = &ast.Field{}

	// TODO:
	case FieldList:
		node = &ast.FieldList{}

	// TODO:
	case Ellipsis:
		node = &ast.Ellipsis{}

	// TODO:
	case FuncLit:
		node = &ast.FuncLit{}

	// TODO:
	case CompositeLit:
		node = &ast.CompositeLit{}

	// TODO:
	case ParenExpr:
		node = &ast.ParenExpr{}

	// TODO:
	case SelectorExpr:
		node = &ast.SelectorExpr{}

	// TODO:
	case IndexExpr:
		node = &ast.IndexExpr{}

	// TODO:
	case IndexListExpr:
		node = &ast.IndexListExpr{}

	// TODO:
	case SliceExpr:
		node = &ast.SliceExpr{}

	// TODO:
	case TypeAssertExpr:
		node = &ast.TypeAssertExpr{}

	// TODO:
	case CallExpr:
		node = &ast.CallExpr{}

	// TODO:
	case StarExpr:
		node = &ast.StarExpr{}

	// TODO:
	case UnaryExpr:
		node = &ast.UnaryExpr{}

	// TODO:
	case BinaryExpr:
		node = &ast.BinaryExpr{}

	// TODO:
	case KeyValueExpr:
		node = &ast.KeyValueExpr{}

	// TODO:
	case ArrayType:
		node = &ast.ArrayType{}

	// TODO:
	case StructType:
		node = &ast.StructType{}

	// TODO:
	case FuncType:
		node = &ast.FuncType{}

	// TODO:
	case InterfaceType:
		node = &ast.InterfaceType{}

	// TODO:
	case MapType:
		node = &ast.MapType{}

	// TODO:
	case ChanType:
		node = &ast.ChanType{}

	// TODO:
	case BadStmt:
		node = &ast.BadStmt{}

	// TODO:
	case DeclStmt:
		node = &ast.DeclStmt{}

	// TODO:
	case EmptyStmt:
		node = &ast.EmptyStmt{}

	// TODO:
	case LabeledStmt:
		node = &ast.LabeledStmt{}

	// TODO:
	case ExprStmt:
		node = &ast.ExprStmt{}

	// TODO:
	case SendStmt:
		node = &ast.SendStmt{}

	// TODO:
	case IncDecStmt:
		node = &ast.IncDecStmt{}

	// TODO:
	case AssignStmt:
		node = &ast.AssignStmt{}

	// TODO:
	case GoStmt:
		node = &ast.GoStmt{}

	// TODO:
	case DeferStmt:
		node = &ast.DeferStmt{}

	// TODO:
	case ReturnStmt:
		node = &ast.ReturnStmt{}

	// TODO:
	case BranchStmt:
		node = &ast.BranchStmt{}

	// TODO:
	case BlockStmt:
		node = &ast.BlockStmt{}

	// TODO:
	case IfStmt:
		node = &ast.IfStmt{
			If:   0,
			Init: nil,
			Cond: nil,
			Body: Generate(BlockStmt).(*ast.BlockStmt),
			Else: nil,
		}

	// TODO:
	case CaseClause:
		node = &ast.CaseClause{}

	// TODO:
	case SwitchStmt:
		node = &ast.SwitchStmt{}

	// TODO:
	case TypeSwitchStmt:
		node = &ast.TypeSwitchStmt{}

	// TODO:
	case CommClause:
		node = &ast.CommClause{}

	// TODO:
	case SelectStmt:
		node = &ast.SelectStmt{}

	// TODO:
	case ForStmt:
		node = &ast.ForStmt{}

	// TODO:
	case RangeStmt:
		node = &ast.RangeStmt{}

	// TODO:
	case ImportSpec:
		node = &ast.ImportSpec{}

	// TODO:
	case ValueSpec:
		node = &ast.ValueSpec{}

	// TODO:
	case TypeSpec:
		node = &ast.TypeSpec{}

	// TODO:
	case BadDecl:
		node = &ast.BadDecl{}

	// TODO:
	case GenDecl:
		node = &ast.GenDecl{}

	// TODO:
	case FuncDecl:
		node = &ast.FuncDecl{
			Doc:  &ast.CommentGroup{},
			Recv: &ast.FieldList{},
			Name: &ast.Ident{},
			Type: &ast.FuncType{},
			Body: &ast.BlockStmt{},
		}

	// TODO:
	case File:
		node = &ast.File{}

	// TODO:
	case Package:
		node = &ast.Package{}

	}
	return node
}
