package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"strconv"
)

type Node struct {
	val  int
	prev *Node
	next *Node
}

func parseLine(line string) *Node {
	val, _ := strconv.Atoi(line)
	return &Node{val: val}
}

func moveNode(node *Node, listLength int) {
	steps := node.val
	steps = steps % (listLength - 1)
	var dir string
	if steps > 0 {
		dir = "right"
	} else if steps < 0 {
		dir = "left"
		steps = -steps
	} else {
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	i := 0
	cur := node
	for i < steps {
		if dir == "right" {
			cur = cur.next
		} else {
			cur = cur.prev
		}
		i++
	}
	if dir == "left" {
		cur = cur.prev
	}

	node.next = cur.next
	node.prev = cur
	cur.next.prev = node
	cur.next = node
}

func partA(lines []string) {
	nodes := make([]*Node, len(lines))

	for i, line := range lines {
		node := parseLine(line)
		nodes[i] = node
	}

	var zero *Node

	for i, node := range nodes {
		if i == 0 {
			node.prev = nodes[len(nodes)-1]
		} else {
			node.prev = nodes[i-1]
		}

		if i == len(nodes)-1 {
			node.next = nodes[0]
		} else {
			node.next = nodes[i+1]
		}

		if node.val == 0 {
			zero = node
		}
	}

	for _, node := range nodes {
		moveNode(node, len(nodes))
	}

	cur := zero
	sum := 0

	for i := 0; i <= 3000; i++ {
		if i > 0 && i%1000 == 0 {
			sum += cur.val
		}
		cur = cur.next
	}

	fmt.Println(sum)
}

func partB(lines []string) {
	nodes := make([]*Node, len(lines))

	for i, line := range lines {
		node := parseLine(line)
		node.val *= 811589153
		nodes[i] = node
	}

	var zero *Node

	for i, node := range nodes {
		if i == 0 {
			node.prev = nodes[len(nodes)-1]
		} else {
			node.prev = nodes[i-1]
		}

		if i == len(nodes)-1 {
			node.next = nodes[0]
		} else {
			node.next = nodes[i+1]
		}

		if node.val == 0 {
			zero = node
		}
	}

	for i := 0; i < 10; i++ {
		for _, node := range nodes {
			moveNode(node, len(nodes))
		}
	}

	cur := zero
	sum := 0

	for i := 0; i <= 3000; i++ {
		if i > 0 && i%1000 == 0 {
			sum += cur.val
		}
		cur = cur.next
	}

	fmt.Println(sum)
}

func main() {
	lines := utils.ReadInputStringLines(20, "\n")
	partA(lines)
	partB(lines)
}
