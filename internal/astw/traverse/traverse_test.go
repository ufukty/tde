package traverse

import (
	ast_utl "tde/internal/astw/utilities"

	"fmt"
	"testing"
)

// Pass if no panic
func Test_Traverse(t *testing.T) {
	_, astFile, _ := ast_utl.LoadFile("walk.go")

	appandableNodes := []*TraversableNode{}

	Traverse(GetTraversableNodeForASTNode(astFile), func(tNodePtr *TraversableNode) bool {
		if tNodePtr.PointsToNilSpot {
			fmt.Printf("%-20s nil\n", tNodePtr.ExpectedType)
		} else {
			fmt.Printf("%-20s %v\n", tNodePtr.ExpectedType, tNodePtr.Value)
		}

		if tNodePtr.PointsToNilSpot || tNodePtr.ExpectedType.IsSliceType() {
			appandableNodes = append(appandableNodes, tNodePtr)
		}

		return true
	})

}
