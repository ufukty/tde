package utilities

import (
	"go/ast"
)

func Compare(a, b ast.Node) bool {
	childrenA := ChildNodes(a)
	childrenB := ChildNodes(b)

	if len(childrenA) != len(childrenB) {
		return false
	}

	for i := 0; i < len(childrenA); i++ {
		if childrenA[i] != childrenB[i] {
			return false
		}
	}

	for i := 0; i < len(childrenA); i++ {
		if subComparison := Compare(childrenA[i], childrenB[i]); !subComparison {
			return false
		}
	}

	return true
}
