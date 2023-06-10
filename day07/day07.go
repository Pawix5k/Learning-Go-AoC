package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	parent   *Node
	isDir    bool
	size     int
	children map[string]*Node
}

func newNode(parent *Node, isDir bool, size int) *Node {
	node := &Node{
		parent:   parent,
		isDir:    isDir,
		size:     size,
		children: make(map[string]*Node),
	}
	return node
}

func buildTree(lines []string) Node {
	root := newNode(nil, true, 0)
	cur := root
	for _, line := range lines[1:] {
		if len(line) == 0 {
			break
		}
		fields := strings.Fields(line)
		if fields[0] == "$" && fields[1] == "cd" && fields[2] == ".." {
			cur = cur.parent
		} else if fields[0] == "$" && fields[1] == "cd" {
			cur = cur.children[fields[2]]
		} else if fields[0] == "dir" {
			cur.children[fields[1]] = newNode(cur, true, 0)
		} else if fields[0] != "$" {
			size, _ := strconv.Atoi(fields[0])
			cur.children[fields[1]] = newNode(cur, false, size)
		}
	}

	return *root
}

func calcDirSizes(node *Node, count int) (int, int) {
	if !node.isDir {
		return node.size, count
	}
	size := 0
	curCount := count
	for _, child := range node.children {
		chSize, chCount := calcDirSizes(child, 0)
		size += chSize
		curCount += chCount
	}
	if size <= 100_000 {
		curCount += size
	}
	node.size = size
	return size, curCount
}

func findSmallestLargerOrEqualN(node *Node, n int, parentSize int) int {
	if n > node.size {
		return parentSize
	}
	curSmallest := parentSize
	for _, child := range node.children {
		if childSize := findSmallestLargerOrEqualN(child, n, node.size); childSize < curSmallest {
			curSmallest = childSize
		}
	}

	return curSmallest
}

func partA() {
	utils.DownloadInput(7)
	data, err := os.ReadFile(utils.GetFilePath(7))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	root := buildTree(lines)
	_, count := calcDirSizes(&root, 0)
	fmt.Println(count)
}

func partB() {
	utils.DownloadInput(7)
	data, err := os.ReadFile(utils.GetFilePath(7))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	root := buildTree(lines)
	size, _ := calcDirSizes(&root, 0)
	needToFree := 30_000_000 - (70_000_000 - size)
	fmt.Println(findSmallestLargerOrEqualN(&root, needToFree, size))
}

func main() {
	partA()
	partB()
}
