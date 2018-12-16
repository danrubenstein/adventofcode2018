package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var partB = flag.Bool("partB", true, "Use partB logic?")
var inputFile = flag.String("file", "inputs/day8.txt", "The path to the input")

type node struct {
	metadata []int
	children []*node
}

func parseNode(headers []string) (*node, int) {
	if len(headers) < 2 {
		fmt.Println("something is wrong")
		return nil, 0
	}

	numChildren, err := strconv.Atoi(headers[0])
	numMetadata, err := strconv.Atoi(headers[1])

	if err != nil {
		fmt.Println("err", err)
		return nil, 0
	}

	place := 2

	children := make([]*node, 0)
	for i := 0; i < numChildren; i++ {
		child, size := parseNode(headers[place:])
		place += size
		children = append(children, child)
	}

	metadata := make([]int, 0)
	for i := 0; i < numMetadata; i++ {
		meta, err := strconv.Atoi(headers[place+i])
		if err != nil {
			fmt.Println(err)
		}
		metadata = append(metadata, meta)
	}

	return &node{metadata, children}, (place + numMetadata)
}

func sumMetadata(n *node) int {
	fmt.Println(n)
	res := 0
	for _, val := range n.children {
		res += sumMetadata(val)
	}

	for _, val := range n.metadata {
		res += val
	}

	return res
}

func sumMetadataIndex(n *node) int {

	res := 0
	nChildren := len(n.children)
	if nChildren == 0 {
		for _, val := range n.metadata {
			res += val
		}
		return res
	}

	for _, val := range n.metadata {
		if val == 0 {
			continue
		} else if val > nChildren {
			continue
		} else {
			res += sumMetadataIndex(n.children[val-1])
		}
	}

	return res
}

func main() {

	contents, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("There was an error reading the file", err)
		return
	}

	inputs := strings.Split(strings.Replace(string(contents), "\n", "", -1), " ")

	tree, _ := parseNode(inputs)

	var sum int
	if !*partB {
		sum = sumMetadata(tree)
	} else {
		sum = sumMetadataIndex(tree)
	}

	fmt.Printf("sum: %d\n", sum)
}
