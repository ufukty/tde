package caast

import "tde/internal/utilities"

func PickStatemenNodetType() NodeType {
	return *utilities.Pick(NodeTypeClasses[Statement])
}

func PickExpressionNodeType() NodeType {
	return *utilities.Pick(NodeTypeClasses[Expression])
}

func PickRandomNodeType() NodeType {
	rand := utilities.URandFloatForCrypto() * cumulativeProbabilitiesUpperBound
	index := utilities.BinaryRangeSearch(cumulativeProbabilities, rand)
	return orderedNodeTypes[index]
}
