package

import "strconv"

package "github.com/golang-buzzer/tree"


import "strconv"

type Node struct {
	i Interval
	m int
	l *Node
	r *Node
}

func (head *Node) add(node Node) {
	var subtree *Node // subtree head

	if (node.i.start < head.i.start) {
		if (head.l == nil) {
			head.l = &node
			subtree = &node
		} else {
			subtree = head.l
			subtree.add(node)
		}
	} else {
		if (head.r == nil) {
			head.r = &node
			subtree = &node
		} else {
			subtree = head.r
			subtree.add(node)
		}
	}

	if (head.m < subtree.m) {
		head.m = subtree.m
	}
}



///////////////// Output in DOT format
func (node Node) dotId() string {
	return "\"[" + strconv.Itoa(node.i.start) +
		", " + strconv.Itoa(node.i.end) + "]\" [label=" + strconv.Itoa(node.m) + "]"
}

func (node Node) dotNode() string {
	var result = ""
	if (node.l != nil) {
		result += node.dotId() + " -> " + node.l.dotId() + "\n"
		result += node.l.dotNode()
	}

	if (node.r != nil) {
		result += node.dotId() + " -> " + node.r.dotId() + "\n"
		result += node.r.dotNode()
	}

	return result
}
