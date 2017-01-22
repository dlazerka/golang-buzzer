package main

import (
	"fmt"
	//"io/ioutil"
	"bufio"
	"strconv"
	"os"
	//graphviz "github.com/awalterschulze/gographviz"
)

type Interval struct {
	start, end int
}

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

func main() {
	intervals := read()
	//tree := make([]node, len(tickets))

	head := Node{intervals[0], intervals[0].end, nil, nil}

	for j := 1; j < len(intervals); j ++ {
		interval := intervals[j]
		n := Node{interval, interval.end, nil, nil}
		head.add(n)
	}

	fmt.Println("digraph graphname {")
	head.drawChildren()
	fmt.Println("}")

	//graphviz.foo()

	//fmt.Println(head)
}

func (node Node) dotId() string {
	return "\"[" + strconv.Itoa(node.i.start) + ", " + strconv.Itoa(node.i.end) + "] *" + strconv.Itoa(node.m) + "\""
}

func (node Node) drawChildren() {
	if (node.l != nil) {
		fmt.Println(node.dotId() + " -> " + node.l.dotId())
		node.l.drawChildren()
	}

	if (node.r != nil) {
		fmt.Println(node.dotId() + " -> " + node.r.dotId())
		node.r.drawChildren()
	}
}

func read() []Interval {
	f, err := os.Open("input.txt")
	check(err)

	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanWords)
	scan.Scan()

	word := scan.Text()
	ticketsCount, err := strconv.Atoi(word)
	check(err)

	var tickets = make([]Interval, ticketsCount)

	for i := 0; i < ticketsCount; i++ {
		scan.Scan()
		start, err := strconv.Atoi(scan.Text())
		check(err)

		scan.Scan()
		end, err := strconv.Atoi(scan.Text())
		check(err)

		tickets[i] = Interval{start, end}
	}
	return tickets
}

func check(err error) {
	if err != nil {
		fmt.Print(err)
	}
}
