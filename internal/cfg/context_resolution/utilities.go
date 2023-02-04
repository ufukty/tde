package context_resolution

import (
	"golang.org/x/exp/slices"
)

//go:generate stringer -type=NodeRelationToInsertionPoint

type NodeRelationToInsertionPoint int

const (
	ITSELF = NodeRelationToInsertionPoint(iota)
	ANCESTOR
	ON_PATH
	WONT_ACCESS
)

func CompareIndicesTraces(current, insertionPoint []int) NodeRelationToInsertionPoint {
	if len(current) > len(insertionPoint) {
		return WONT_ACCESS
	}

	if len(current) == 0 {
		return ANCESTOR
	}

	isInSameScope := slices.Compare(current[:len(current)-1], insertionPoint[:len(current)-1]) == 0
	if !isInSameScope {
		return WONT_ACCESS
	}

	c, ip := current[len(current)-1], insertionPoint[len(current)-1]
	if c < ip {
		return ON_PATH
	} else if c > ip {
		return WONT_ACCESS
	} else {
		if len(current) == len(insertionPoint) {
			return ITSELF
		} else {
			return ANCESTOR
		}
	}
}
