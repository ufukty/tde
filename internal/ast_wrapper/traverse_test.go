package ast_wrapper

import (
	"testing"
)

// Pass if no panic
func Test_Traverse(t *testing.T) {
	_, astFile, _ := LoadFile("walk.go")

	appandableNodes := []TraversableNode{}

	Traverse(GetTraversableNodeForASTNode(astFile), func(tNode TraversableNode) bool {
		// if tNode.IsNil {
		// 	fmt.Printf("%-20s nil\n", tNode.ExpectedType)
		// } else {
		// 	fmt.Printf("%-20s %v\n", tNode.ExpectedType, tNode.Value)
		// }

		if tNode.IsNil || tNode.ExpectedType.IsSliceType() {
			appandableNodes = append(appandableNodes, tNode)
		}

		return true
	})

}
