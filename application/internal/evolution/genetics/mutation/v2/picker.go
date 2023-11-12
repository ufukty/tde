package mutation

import (
	"go/ast"
)

//go:generate stringer -type Mutation
type Mutation int

const (
	MLiteralShuffle = Mutation(iota)
	MTokenShuffle
	// BlockStmt
	MBlockStmtNewLine
	MBlockStmtCreateVariableDeclaration
	MBlockStmtSwapTwoLines
	MBlockStmtConditionalize
	MBlockStmtLoopify
	// Ifstmt
	MIfStmtRegenerateCond
	MIfStmtRegenerateBody
	MIfStmtRegenerateElse
	MIfStmtAddElif
	MIfStmtRemoveElif
	MIfStmtRegenerateElifCondition
	MIfStmtRegenerateElifBody
	// ForStmt
	MForStmtRegenerateCond
	MFor
	MBlockStmtForRange
	// Datastructures
	MDefine
	// Expressions
	MExprRegenerate
	MExprCallPackageFunction
	MExprRangeOverArrayWithLoopCounter
)

type NodeContext struct {
	LoopCounters [][]*ast.Ident
}

func (nc *NodeContext) Inspect(n ast.Node) {
	switch n := n.(type) {
	case *ast.ForStmt:
		if n, ok := n.Init.(*ast.AssignStmt); ok {
			loopCounters := []*ast.Ident{}
			for _, lh := range n.Lhs {
				if lh, ok := lh.(*ast.Ident); ok {
					loopCounters = append(loopCounters, lh)
				}
			}
			if len(loopCounters) > 0 {
				nc.LoopCounters = append(nc.LoopCounters, loopCounters)
			}
		}
	}
}

func PickMutationFor(n *ast.FuncDecl) []Mutation {
	var cursor ast.Node = n
	var nodeContext = NodeContext{}
	for {
		switch cursor.(type) {
		case *ast.BlockStmt:
			return []Mutation{
				MBlockStmtConditionalize,
				MBlockStmtNewLine,
				MBlockStmtCreateVariableDeclaration,
				MBlockStmtSwapTwoLines,
			}
		case *ast.IfStmt:
			return []Mutation{
				MIfStmtRegenerateCond,
				MIfStmtRegenerateBody,
				MIfStmtRegenerateElse,
				MIfStmtAddElif,
				MIfStmtRemoveElif,
				MIfStmtRegenerateElifCondition,
				MIfStmtRegenerateElifBody,
			}
		case *ast.ForStmt:
			nodeContext.Inspect(n)
			return []Mutation{}
		}
	}
}
