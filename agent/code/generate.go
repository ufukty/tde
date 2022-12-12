package code

import (
	"GoGP/evolve/evolution"
	"GoGP/evolve/utility"
	"go/ast"
)

type NodeType int

const (
	// BadExpr
	// Ident
	// Comment
	// CommentGroup
	BasicLit NodeType = iota
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
	BasicLit:       0.0,
	Field:          0.0,
	FieldList:      0.0,
	Ellipsis:       0.0,
	FuncLit:        0.0,
	CompositeLit:   0.0,
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
	ArrayType:      0.0,
	StructType:     0.0,
	FuncType:       0.0,
	InterfaceType:  0.0,
	MapType:        0.0,
	ChanType:       0.0,
	BadStmt:        0.0,
	DeclStmt:       0.0,
	EmptyStmt:      0.0,
	LabeledStmt:    0.0,
	ExprStmt:       0.0,
	SendStmt:       0.0,
	IncDecStmt:     0.0,
	AssignStmt:     0.0,
	GoStmt:         0.0,
	DeferStmt:      0.0,
	ReturnStmt:     0.0,
	BranchStmt:     0.0,
	BlockStmt:      0.0,
	IfStmt:         0.0,
	CaseClause:     0.0,
	SwitchStmt:     0.0,
	TypeSwitchStmt: 0.0,
	CommClause:     0.0,
	SelectStmt:     0.0,
	ForStmt:        0.0,
	RangeStmt:      0.0,
	ImportSpec:     0.0,
	ValueSpec:      0.0,
	TypeSpec:       0.0,
	BadDecl:        0.0,
	GenDecl:        0.0,
	FuncDecl:       0.0,
	File:           0.0,
	Package:        0.0,
}

var orderedNodeTypes = []NodeType{
	BasicLit,
	Field,
	FieldList,
	Ellipsis,
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
	ImportSpec,
	ValueSpec,
	TypeSpec,
	BadDecl,
	GenDecl,
	FuncDecl,
	File,
	Package,
}

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
		prob  = 0.0
	)
	for _, nodeType := range orderedNodeTypes {
		prob = probabilities[nodeType]
		cumulativeProbabilities[nodeType] = prob
		total += prob
	}

	cumulativeProbabilitiesUpperBound = total
}

func PickRandomNodeType() NodeType {
	rand := utility.URandFloatForCrypto() * cumulativeProbabilitiesUpperBound
	index := evolution.BinarySearchSmallestOfGreaters(cumulativeProbabilities, rand)
	return orderedNodeTypes[index]
}

func Generate(kind NodeType) ast.Node {
	var (
		node ast.Node
	)
	switch kind {
	case BasicLit:
		d := ast.BasicLit{}
		node = &d
	case Field:
		d := ast.Field{}
		node = &d
	case FieldList:
		d := ast.FieldList{}
		node = &d
	case Ellipsis:
		d := ast.Ellipsis{}
		node = &d
	case FuncLit:
		d := ast.FuncLit{}
		node = &d
	case CompositeLit:
		d := ast.CompositeLit{}
		node = &d
	case ParenExpr:
		d := ast.ParenExpr{}
		node = &d
	case SelectorExpr:
		d := ast.SelectorExpr{}
		node = &d
	case IndexExpr:
		d := ast.IndexExpr{}
		node = &d
	case IndexListExpr:
		d := ast.IndexListExpr{}
		node = &d
	case SliceExpr:
		d := ast.SliceExpr{}
		node = &d
	case TypeAssertExpr:
		d := ast.TypeAssertExpr{}
		node = &d
	case CallExpr:
		d := ast.CallExpr{}
		node = &d
	case StarExpr:
		d := ast.StarExpr{}
		node = &d
	case UnaryExpr:
		d := ast.UnaryExpr{}
		node = &d
	case BinaryExpr:
		d := ast.BinaryExpr{}
		node = &d
	case KeyValueExpr:
		d := ast.KeyValueExpr{}
		node = &d
	case ArrayType:
		d := ast.ArrayType{}
		node = &d
	case StructType:
		d := ast.StructType{}
		node = &d
	case FuncType:
		d := ast.FuncType{}
		node = &d
	case InterfaceType:
		d := ast.InterfaceType{}
		node = &d
	case MapType:
		d := ast.MapType{}
		node = &d
	case ChanType:
		d := ast.ChanType{}
		node = &d
	case BadStmt:
		d := ast.BadStmt{}
		node = &d
	case DeclStmt:
		d := ast.DeclStmt{}
		node = &d
	case EmptyStmt:
		d := ast.EmptyStmt{}
		node = &d
	case LabeledStmt:
		d := ast.LabeledStmt{}
		node = &d
	case ExprStmt:
		d := ast.ExprStmt{}
		node = &d
	case SendStmt:
		d := ast.SendStmt{}
		node = &d
	case IncDecStmt:
		d := ast.IncDecStmt{}
		node = &d
	case AssignStmt:
		d := ast.AssignStmt{}
		node = &d
	case GoStmt:
		d := ast.GoStmt{}
		node = &d
	case DeferStmt:
		d := ast.DeferStmt{}
		node = &d
	case ReturnStmt:
		d := ast.ReturnStmt{}
		node = &d
	case BranchStmt:
		d := ast.BranchStmt{}
		node = &d
	case BlockStmt:
		d := ast.BlockStmt{}
		node = &d
	case IfStmt:
		d := ast.IfStmt{}
		node = &d
	case CaseClause:
		d := ast.CaseClause{}
		node = &d
	case SwitchStmt:
		d := ast.SwitchStmt{}
		node = &d
	case TypeSwitchStmt:
		d := ast.TypeSwitchStmt{}
		node = &d
	case CommClause:
		d := ast.CommClause{}
		node = &d
	case SelectStmt:
		d := ast.SelectStmt{}
		node = &d
	case ForStmt:
		d := ast.ForStmt{}
		node = &d
	case RangeStmt:
		d := ast.RangeStmt{}
		node = &d
	case ImportSpec:
		d := ast.ImportSpec{}
		node = &d
	case ValueSpec:
		d := ast.ValueSpec{}
		node = &d
	case TypeSpec:
		d := ast.TypeSpec{}
		node = &d
	case BadDecl:
		d := ast.BadDecl{}
		node = &d
	case GenDecl:
		d := ast.GenDecl{}
		node = &d
	case FuncDecl:
		d := ast.FuncDecl{}
		node = &d
	case File:
		d := ast.File{}
		node = &d
	case Package:
		d := ast.Package{}
		node = &d
	}
	return node
}
