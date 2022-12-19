package code

import (
	"fmt"
	"tde/internal/utilities"

	"go/ast"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

func GenerateRandomLiteral() (kind token.Token, value string) {
	return (*utilities.Pick([]func() (token.Token, string){
		func() (token.Token, string) {
			return token.INT, "0"
		},
		func() (token.Token, string) {
			return token.STRING, ""
		},
		func() (token.Token, string) {
			return token.FLOAT, fmt.Sprint(utilities.URandFloatForCrypto())
		},
	}))()
}

func PickRandomNodeType() NodeType {
	rand := utilities.URandFloatForCrypto() * cumulativeProbabilitiesUpperBound
	index := utilities.BinaryRangeSearch(cumulativeProbabilities, rand)
	return orderedNodeTypes[index]
}

func GenerateRandomNumberOfInstance(nodeTypeClass NodeTypeClass) []ast.Node {
	list := []ast.Node{}
	switch nodeTypeClass {
	case Statement:

	}
	return list
}

func GenerateInstance(nodeType NodeType) ast.Node {
	switch nodeType {
	case BasicLit: // basic literal
		kind, value := GenerateRandomLiteral()
		return &ast.BasicLit{
			Kind:  kind,
			Value: value,
		}
	case FieldList:
		return &ast.FieldList{}
	case Field: // struct field
		return &ast.Field{}
	case Ellipsis:
		return &ast.Ellipsis{} // ... variadic argument spreading
	case FuncLit:
		return &ast.FuncLit{}
	case CompositeLit:
		return &ast.CompositeLit{}
	case ParenExpr:
		return &ast.ParenExpr{}
	case SelectorExpr:
		return &ast.SelectorExpr{}
	case IndexExpr:
		return &ast.IndexExpr{}
	case IndexListExpr:
		return &ast.IndexListExpr{}
	case SliceExpr:
		return &ast.SliceExpr{}
	case TypeAssertExpr:
		return &ast.TypeAssertExpr{}
	case CallExpr:
		return &ast.CallExpr{}
	case StarExpr:
		return &ast.StarExpr{}
	case UnaryExpr:
		return &ast.UnaryExpr{}
	case BinaryExpr:
		return &ast.BinaryExpr{}
	case KeyValueExpr:
		return &ast.KeyValueExpr{}
	case ArrayType:
		return &ast.ArrayType{}
	case StructType:
		return &ast.StructType{}
	case FuncType:
		return &ast.FuncType{}
	case InterfaceType:
		return &ast.InterfaceType{}
	case MapType:
		return &ast.MapType{}
	case ChanType:
		return &ast.ChanType{}
	case BadStmt:
		return &ast.BadStmt{}
	case DeclStmt:
		return &ast.DeclStmt{}
	case EmptyStmt:
		return &ast.EmptyStmt{}
	case LabeledStmt:
		return &ast.LabeledStmt{}
	case ExprStmt:
		return &ast.ExprStmt{}
	case SendStmt:
		return &ast.SendStmt{}
	case IncDecStmt:
		return &ast.IncDecStmt{}
	case AssignStmt:
		return &ast.AssignStmt{}
	case GoStmt:
		return &ast.GoStmt{}
	case DeferStmt:
		return &ast.DeferStmt{}
	case ReturnStmt:
		return &ast.ReturnStmt{}
	case BranchStmt:
		return &ast.BranchStmt{}
	case BlockStmt:
		return &ast.BlockStmt{
			Lbrace: 0,
			List:   []ast.Stmt{},
			Rbrace: 0,
		}
	case IfStmt:
		return &ast.IfStmt{}
	case CaseClause:
		return &ast.CaseClause{}
	case SwitchStmt:
		return &ast.SwitchStmt{}
	case TypeSwitchStmt:
		return &ast.TypeSwitchStmt{}
	case CommClause:
		return &ast.CommClause{}
	case SelectStmt:
		return &ast.SelectStmt{}
	case ForStmt:
		return &ast.ForStmt{}
	case RangeStmt:
		return &ast.RangeStmt{}
	case ImportSpec:
		return &ast.ImportSpec{}
	case ValueSpec:
		return &ast.ValueSpec{}
	case TypeSpec:
		return &ast.TypeSpec{}
	case BadDecl:
		return &ast.BadDecl{}
	case GenDecl:
		return &ast.GenDecl{}
	case FuncDecl:
		return &ast.FuncDecl{}
	case File:
		return &ast.File{}
	case Package:
		return &ast.Package{}
	}

	return nil
}

func GenerateRandomSubtree(nodeTypeClass NodeTypeClass) ast.Node {
	var (
		node ast.Node
		// isGenerated = false
	)
	// for !isGenerated {
	// 	node = GenerateInstance(utilities.Pick())
	// }
	return node
}

// It picks random BlockStmt amongst all subnodes, generates new stm
func NewLine(f *Function) {
	body := f.Root.Body

	parentNode := *utilities.Pick(FilterSubNodes(body, func(n ast.Node) bool {
		if _, ok := n.(*ast.BlockStmt); ok {
			return true
		}
		return false
	}))
	siblingNode := *utilities.Pick(ChildNodes(parentNode))
	isInserted := false
	astutil.Apply(parentNode, func(c *astutil.Cursor) bool {
		if !isInserted {
			if c.Node() == siblingNode {
				c.InsertBefore(GenerateRandomSubtree(Statement))
				isInserted = true
			}
		}
		return !isInserted
	}, nil)
}
