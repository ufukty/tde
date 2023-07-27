package traverse

import (
	"tde/internal/astw/astwutl"

	"fmt"
	"testing"
)

// Pass if no panic
func Test_Traverse(t *testing.T) {
	_, astFile, _ := astwutl.LoadFile("walk.go")

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
