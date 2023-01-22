package ast_wrapper

import (
	"fmt"
	"testing"
)

func Test_Traverse(t *testing.T) {
	_, astFile, _ := LoadFile("walk.go")

	Traverse(astFile, func(n TraversableNode, trace []TraverseTraceItem) bool {

		fmt.Println(n.ExpectedType, n.IsNil)
		return true
	})
}
