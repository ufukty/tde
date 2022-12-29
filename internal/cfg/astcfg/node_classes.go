package astcfg

import "go/ast"

type NodeTypeClass int

const (
	Expression = NodeTypeClass(iota)
	Statement
	Declaration
	Spec
)

var Dict_NodeTypeClassToNodeType = map[NodeTypeClass][]NodeType{
	Expression:  {},
	Statement:   {},
	Declaration: {},
	Spec:        {},
}

var Dict_NodeTypeToNodeTypeClass = map[NodeType]NodeTypeClass{}

func init() {
	for nodeType, constructor := range nc.Dictionary {
		node := constructor()
		switch node.(type) {
		case ast.Expr:
			Dict_NodeTypeClassToNodeType[Expression] = append(Dict_NodeTypeClassToNodeType[Expression], nodeType)
			Dict_NodeTypeToNodeTypeClass[nodeType] = Expression
		case ast.Stmt:
			Dict_NodeTypeClassToNodeType[Statement] = append(Dict_NodeTypeClassToNodeType[Statement], nodeType)
			Dict_NodeTypeToNodeTypeClass[nodeType] = Statement
		case ast.Decl:
			Dict_NodeTypeClassToNodeType[Declaration] = append(Dict_NodeTypeClassToNodeType[Declaration], nodeType)
			Dict_NodeTypeToNodeTypeClass[nodeType] = Declaration
		case ast.Spec:
			Dict_NodeTypeClassToNodeType[Spec] = append(Dict_NodeTypeClassToNodeType[Spec], nodeType)
			Dict_NodeTypeToNodeTypeClass[nodeType] = Spec
		}
	}
}
