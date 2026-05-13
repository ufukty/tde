package traverse

import (
	"fmt"
	"testing"

	"tde/internal/ast/parse"
)

// Pass if no panic
func Test_Traverse(t *testing.T) {
	_, astFile, err := parse.File("testdata/walk.go")
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

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
