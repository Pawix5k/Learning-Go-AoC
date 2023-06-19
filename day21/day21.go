package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type ExpEntry struct {
	expType   string
	name      string
	left      string
	right     string
	operation string
	val       int
}

type Exp struct {
	name     string
	expType  string
	parent   *Exp
	val      int
	left     *Exp
	right    *Exp
	operator string
}

func (e *Exp) getVal() int {
	if e.expType == "op" {
		return operations[e.operator](e.left.getVal(), e.right.getVal())
	}
	return e.val
}

var operations = map[string]func(int, int) int{
	"+": func(a, b int) int {
		return a + b
	},
	"-": func(a, b int) int {
		return a - b
	},
	"*": func(a, b int) int {
		return a * b
	},
	"/": func(a, b int) int {
		return a / b
	},
}

func parseLine(line string) ExpEntry {
	fields := strings.Fields(line)
	name := fields[0][:len(fields[0])-1]
	if len(fields) == 2 {
		val, _ := strconv.Atoi(fields[1])
		return ExpEntry{name: name, expType: "int", val: val}
	}
	left, opStr, right := fields[1], fields[2], fields[3]
	return ExpEntry{name: name, expType: "op", left: left, right: right, operation: opStr}
}

func buildExpTree(expName string, expressionEntries map[string]ExpEntry, parent *Exp) *Exp {
	expEntry := expressionEntries[expName]
	if expEntry.expType == "int" {
		return &Exp{name: expEntry.name, expType: expEntry.expType, parent: parent, val: expEntry.val}
	}
	op := &Exp{name: expEntry.name, expType: expEntry.expType, parent: parent, operator: expEntry.operation}
	op.left = buildExpTree(expEntry.left, expressionEntries, op)
	op.right = buildExpTree(expEntry.right, expressionEntries, op)

	return op
}

func getNode(node *Exp, name string) *Exp {
	if node == nil {
		return nil
	}
	if node.name == name {
		return node
	}

	if res := getNode(node.left, name); res != nil {
		return res
	}
	if res := getNode(node.right, name); res != nil {
		return res
	}

	return nil
}

func reverseTree(node, prevNode *Exp) *Exp {

	if prevNode == nil {
		node.expType = "op"
	}
	side := "left"
	if node.parent.left.name != node.name {
		side = "right"
	}

	isLast := false

	oldLeft := node.parent.left
	oldRight := node.parent.right

	oldOperator := node.parent.operator

	if node.parent.parent.name == "root" {
		sideFromRoot := node.parent.parent.left
		if node.parent.name == node.parent.parent.left.name {
			sideFromRoot = node.parent.parent.right
		}
		node.parent = sideFromRoot
		isLast = true
	}

	switch {
	case side == "left" && oldOperator == "+":
		node.operator = "-"
		node.left = node.parent
		node.right = oldRight
	case side == "right" && oldOperator == "+":
		node.operator = "-"
		node.right = oldLeft
		node.left = node.parent
	case side == "left" && oldOperator == "-":
		node.operator = "+"
		node.right = oldRight
		node.left = node.parent
	case side == "right" && oldOperator == "-":
		node.operator = "-"
		node.left = oldLeft
		node.right = node.parent
	case side == "left" && oldOperator == "*":
		node.operator = "/"
		node.right = oldRight
		node.left = node.parent
	case side == "right" && oldOperator == "*":
		node.operator = "/"
		node.right = oldLeft
		node.left = node.parent
	case side == "left" && oldOperator == "/":
		node.operator = "*"
		node.right = oldRight
		node.left = node.parent
	case side == "right" && oldOperator == "/":
		node.operator = "/"
		node.left = oldLeft
		node.right = node.parent
	}
	if !isLast {
		node.parent = reverseTree(node.parent, node)
	}

	return node
}

func partA(lines []string) {
	expressionEntries := map[string]ExpEntry{}

	for _, line := range lines {
		expEntry := parseLine(line)
		expressionEntries[expEntry.name] = expEntry
	}

	root := buildExpTree("root", expressionEntries, nil)

	fmt.Println(root.getVal())
}

func partB(lines []string) {
	expressionEntries := map[string]ExpEntry{}

	for _, line := range lines {
		expEntry := parseLine(line)
		expressionEntries[expEntry.name] = expEntry
	}

	root := buildExpTree("root", expressionEntries, nil)
	humn := getNode(root, "humn")

	reverseTree(humn, nil)

	fmt.Println(humn.getVal())
}

func main() {
	lines := utils.ReadInputStringLines(21, "\n")
	partA(lines)
	partB(lines)
}
