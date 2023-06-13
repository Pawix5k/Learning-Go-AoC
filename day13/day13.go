package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	parent   *Node
	isInt    bool
	val      int
	children []*Node
}

func parseString(s string) *Node {
	dummyHead := &Node{}
	cur := dummyHead
	curVal := ""
	for _, ch := range s {
		if ch == '[' {
			new := &Node{parent: cur, isInt: false}
			cur.children = append(cur.children, new)
			cur = new
		} else if ch == ']' {
			if len(curVal) != 0 {
				valInt, _ := strconv.Atoi(curVal)
				new := &Node{parent: cur, isInt: true, val: valInt}
				cur.children = append(cur.children, new)
				curVal = ""
			}
			cur = cur.parent
		} else if utils.IsNumber(byte(ch)) {
			curVal += string(ch)
		} else {
			if len(curVal) != 0 {
				valInt, _ := strconv.Atoi(curVal)
				new := &Node{parent: cur, isInt: true, val: valInt}
				cur.children = append(cur.children, new)
				curVal = ""
			}
		}
	}

	return dummyHead.children[0]
}

func compareTrees(node1, node2 Node) int {
	if node1.isInt && node2.isInt {
		if node1.val < node2.val {
			return 1
		} else if node1.val > node2.val {
			return -1
		}
	} else if !node1.isInt && !node2.isInt {
		i := 0
		for i < len(node1.children) || i < len(node2.children) {
			if i >= len(node1.children) {
				return 1
			} else if i >= len(node2.children) {
				return -1
			} else {
				res := compareTrees(*node1.children[i], *node2.children[i])
				if res != 0 {
					return res
				}
			}
			i++
		}
	} else if node1.isInt {
		new := Node{children: []*Node{&node1}}
		res := compareTrees(new, node2)
		if res != 0 {
			return res
		}
	} else if node2.isInt {
		new := Node{children: []*Node{&node2}}
		res := compareTrees(node1, new)
		if res != 0 {
			return res
		}
	}

	return 0
}

func partA(text string) {
	pairs := strings.Split(text, "\n\n")

	sum := 0

	for i, pair := range pairs {
		pairSplit := strings.Split(pair, "\n")
		s1, s2 := pairSplit[0], pairSplit[1]
		t1, t2 := parseString(s1), parseString(s2)
		if compareTrees(*t1, *t2) == 1 {
			sum += i + 1
		}
	}

	fmt.Println(sum)
}

func partB(text string) {
	lines := strings.Fields(text)

	packets := make([]*Node, 0)
	divider1 := parseString("[[2]]")
	divider2 := parseString("[[6]]")

	for _, line := range lines {
		packets = append(packets, parseString(line))
	}
	packets = append(packets, divider1)
	packets = append(packets, divider2)

	sort.Slice(packets, func(i, j int) bool {
		return compareTrees(*packets[i], *packets[j]) == 1
	})

	Idiv1 := 0
	Idiv2 := 0

	for i, packet := range packets {
		if Idiv1 == 0 && compareTrees(*packet, *divider1) == 0 {
			Idiv1 = i + 1
		}
		if compareTrees(*packet, *divider2) == 0 {
			Idiv2 = i + 1
			break
		}
	}

	fmt.Println(Idiv1 * Idiv2)
}

func main() {
	lines := utils.ReadInputString(13)
	partA(lines)
	partB(lines)
}
