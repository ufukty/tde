package astcfg

import "tde/internal/utilities"

func PickStatemenNodetType() NodeType {
	return *utilities.Pick(Dict_NodeTypeClassToNodeType[Statement])
}

func PickExpressionNodeType() NodeType {
	return *utilities.Pick(Dict_NodeTypeClassToNodeType[Expression])
}

func PickRandomNodeType() NodeType {
	rand := utilities.URandFloatForCrypto() * cumulativeProbabilitiesUpperBound
	index := utilities.BinaryRangeSearch(cumulativeProbabilities, rand)
	return orderedNodeTypes[index]
}
