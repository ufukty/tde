package traverse

import (
	"tde/internal/astw/astwutl"

	"fmt"
	"testing"
)

// Pass if no panic
func Test_Traverse(t *testing.T) {
	_, astFile, err := astwutl.LoadFile("testdata/walk.go")
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}

	appandableNodes := []*Node{}
	Once(NewNode(astFile), func(n *Node) bool {
		if n.IsNil {
			fmt.Printf("%-20s nil\n", n.Expected)
		} else {
			fmt.Printf("%-20s %v %s\n", n.Expected, n.ref.Get(), n.Constraints)
		}

		if n.IsNil || n.Expected.IsSliceType() {
			appandableNodes = append(appandableNodes, n)
		}

		return true
	})

}
