package context

import (
	"go/ast"
	"tde/internal/astw"
)

// TODO: Detect FuncLit's in code and add to Context
// FIXME: convert to BFS
func InspectWithContext(startNode ast.Node, callback func(ctx Context, node ast.Node)) {
	ctx := NewContext()
	astw.InspectTwiceWithTrace(startNode,
		func(node ast.Node, parents []ast.Node, indices []int) bool {
			ExamineEnteringNode(&ctx, node)
			callback(ctx, node)
			return true
		}, func(node ast.Node, parents []ast.Node, indices []int) {
			ExamineLeavingNode(&ctx, node)
		},
	)
}