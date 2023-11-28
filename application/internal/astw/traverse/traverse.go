package traverse

func onceHelper(n *Node, callback func(n *Node) bool) {
	if cont := callback(n); !cont {
		return
	}
	for _, s := range n.Subnodes() {
		onceHelper(s, callback)
	}
}

// There are 2 differences compared to WalkWithNil:
//  1. Sends the information of expected type for nil values
//  2. Threats slices are individual nodes, their items will have isolated indices from sibling nodes.
func Once(n *Node, callback func(n *Node) bool) {
	onceHelper(n, callback)
}

func twiceHelper(n *Node, pre func(n *Node) bool, post func(n *Node)) {
	if pre != nil {
		if cont := pre(n); !cont {
			return
		}
	}
	for _, s := range n.Subnodes() {
		twiceHelper(s, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// There are 2 differences compared to WalkWithNil:
//  1. Sends the information of expected type for nil values
//  2. Threats slices are individual nodes, their items will have isolated indices from sibling nodes.
func Twice(n *Node, pre func(n *Node) bool, post func(n *Node)) {
	twiceHelper(n, pre, post)
}
