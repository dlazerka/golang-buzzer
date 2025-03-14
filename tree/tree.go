package tree

import "strconv"

type Node struct {
	I Interval
	M int
	L *Node
	R *Node
}

func (head *Node) Add(node Node) {
	var subtree *Node // subtree head

	if node.I.Start < head.I.Start {
		if head.L == nil {
			head.L = &node
			subtree = &node
		} else {
			subtree = head.L
			subtree.Add(node)
		}
	} else {
		if head.R == nil {
			head.R = &node
			subtree = &node
		} else {
			subtree = head.R
			subtree.Add(node)
		}
	}

	if head.M < subtree.M {
		head.M = subtree.M
	}
}

// /////////////// Output in DOT format
func (node Node) dotId() string {
	return "\"[" + strconv.Itoa(node.I.Start) +
		", " + strconv.Itoa(node.I.End) + "]\" [label=" + strconv.Itoa(node.M) + "]"
}

func (node Node) DotNode() string {
	var result = ""
	if node.L != nil {
		result += node.dotId() + " -> " + node.L.dotId() + "\n"
		result += node.L.DotNode()
	}

	if node.R != nil {
		result += node.dotId() + " -> " + node.R.dotId() + "\n"
		result += node.R.DotNode()
	}

	return result
}
