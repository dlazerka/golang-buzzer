package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"os/exec"
	"io/ioutil"
	tree "github.com/dlazerka/golang-buzzer/tree"
)


func main() {
	intervals := read()
	//tree := make([]node, len(tickets))

	head := tree.Node{intervals[0], intervals[0].end, nil, nil}

	for j := 1; j < len(intervals); j ++ {
		interval := intervals[j]
		n := tree.Node{interval, interval.end, nil, nil}
		head.add(n)
	}

	write(head)
}


func read() []tree.Interval {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

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

func write(head Node) {
	dot := "digraph graphname {\n" + head.dotNode() + "}"

	println(dot)

	outFileType := "png"
	command := exec.Command("dot", "-T" + outFileType)
	wc, err := command.StdinPipe()
	check(err)
	wc.Write([]byte(dot))
	wc.Close()

	out, err := command.Output()

	ioutil.WriteFile("graph." + outFileType, out, 0)
}

func check(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

