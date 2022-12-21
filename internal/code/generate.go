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
		return nil
	}
	return list
}

func GenerateRandomNumberOfStatements() []ast.Stmt {
	stmts := []ast.Stmt{}
	for _, n := range GenerateRandomNumberOfInstance(Statement) {
		if n, ok := n.(ast.Stmt); ok {
			stmts = append(stmts, n)
		} else {
			panic("looks like GenerateRandomNumberOfInstance() returned another type of node then given (Statement)")
		}
	}
	return stmts
}

func GenerateReturnStatement() []ast.Expr {
	//
}

func GenerateBooleanExpression() ast.Expr {
	//
}

func GenerateBlockStatement() *ast.BlockStmt {
	//
}

type Stack struct {
	vars      [][]ast.BasicLit
	constants [][]ast.BasicLit
}

func (s *Stack) Recurse() {}

func (s *Stack) Return() {
	s.vars = s.vars[:len(s.vars)-1]
	s.constants = s.constants[:len(s.vars)-1]
}

func NodeGenerator(node ast.Node, nodeType NodeType) ast.Node {
	var (
		stack       = Stack{}
		isPerformed = false
	)

	ast.Inspect(node, func(n ast.Node) bool {

		if !isPerformed {
			return false
		}

		if n == nil {
			stack.Return()
		}

		if utilities.Coin() {

		}

		return true
	})

	switch nodeType {

	// MARK: Expressions and types

	// case Field: // struct field
	// 	return &ast.Field{}
	// case FieldList:
	// 	return &ast.FieldList{}
	// case Ellipsis:
	// 	return &ast.Ellipsis{} // ... variadic argument spreading
	case CompositeLit:
		return &ast.CompositeLit{
			Type:       nil,
			Lbrace:     0,
			Elts:       []ast.Expr{},
			Rbrace:     0,
			Incomplete: false,
		}
	case BasicLit: // basic literal
		kind, value := GenerateRandomLiteral()
		return &ast.BasicLit{
			Kind:  kind,
			Value: value,
		}
	// case FuncLit:
	// 	return &ast.FuncLit{}
	case ParenExpr:
		return &ast.ParenExpr{}
	case SelectorExpr:
		return &ast.SelectorExpr{Sel: &ast.Ident{}}
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

	// MARK: Types

	case ArrayType:
		return &ast.ArrayType{}
	case StructType:
		return &ast.StructType{}
	// case FuncType:
	// 	return &ast.FuncType{}
	// case InterfaceType:
	// 	return &ast.InterfaceType{}
	case MapType:
		return &ast.MapType{}
	// case ChanType:
	// 	return &ast.ChanType{}

	// MARK: Statements

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
	// case GoStmt:
	// 	return &ast.GoStmt{}
	// case DeferStmt:
	// 	return &ast.DeferStmt{}
	case ReturnStmt:
		return &ast.ReturnStmt{Results: GenerateReturnStatement()}
	// case BranchStmt:
	// 	return &ast.BranchStmt{}
	case BlockStmt:
		return &ast.BlockStmt{List: GenerateRandomNumberOfStatements()}
	case IfStmt:
		return &ast.IfStmt{
			Cond: GenerateBooleanExpression(),
			Body: GenerateBlockStatement(),
			Else: GenerateBlockStatement(),
		}
	// case CaseClause:
	// 	return &ast.CaseClause{}
	// case SwitchStmt:
	// 	return &ast.SwitchStmt{}
	// case TypeSwitchStmt:
	// 	return &ast.TypeSwitchStmt{}
	// case CommClause:
	// 	return &ast.CommClause{}
	// case SelectStmt:
	// 	return &ast.SelectStmt{}
	case ForStmt:
		return &ast.ForStmt{}
	case RangeStmt:
		return &ast.RangeStmt{}

	// MARK: Declarations/Spec

	case ImportSpec:
		return &ast.ImportSpec{}
	case ValueSpec:
		return &ast.ValueSpec{}
	case TypeSpec:
		return &ast.TypeSpec{}

		// MARK: Declarations/Other

		// case BadDecl:
		// 	return &ast.BadDecl{}
		// case GenDecl:
		// 	return &ast.GenDecl{}
		// case FuncDecl:
		// 	return &ast.FuncDecl{}

		// MARK: Files & packages

		// case File:
		// 	return &ast.File{}
		// case Package:
		// 	return &ast.Package{}
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

// It picks random BlockStmt amongst all subnodes, generates new Statement
func NewLine(f *ast.FuncDecl) {
	body := f.Body

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
				newStmt := GenerateRandomSubtree(Statement)
				c.InsertBefore(newStmt)
				isInserted = true
			}
		}
		return !isInserted
	}, nil)
}
