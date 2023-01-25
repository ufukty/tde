package astw

import (
	"go/ast"

	"golang.org/x/exp/slices"
)

// Retraces the trace returned by InspectWithTrace or InspectTwiceWithTrace and calls the callback at each node.
func Retrace(root ast.Node, trace []int, callback func(node ast.Node, depth int)) {
	InspectWithTrace(root, func(currentNode ast.Node, parentTrace []ast.Node, childIndexTrace []int) bool {
		// fmt.Printf(
		// 	"\n>> %-20s %-50s %-20s %-60s\n",
		// 	fmt.Sprint(reflect.TypeOf(currentNode)),
		// 	fmt.Sprint(reflect.ValueOf(currentNode)),
		// 	fmt.Sprint(childIndexTrace),
		// 	fmt.Sprint(parentTrace),
		// )
		if currentNode == nil {
			return false // doesn't matter
		}
		if slices.Compare(trace[:len(childIndexTrace)], childIndexTrace) == 0 {
			callback(currentNode, len(childIndexTrace))
			return true
		}
		return false
	})
}
