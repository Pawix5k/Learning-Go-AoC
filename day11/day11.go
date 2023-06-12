package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Value struct {
	isSelf bool
	value  int
}

type Operation struct {
	operator string
	val1     Value
	val2     Value
}

func (o Operation) calculate(old int) int {
	v1 := o.val1.value
	if o.val1.isSelf {
		v1 = old
	}
	v2 := o.val2.value
	if o.val2.isSelf {
		v2 = old
	}
	res := 0
	switch o.operator {
	case "+":
		res = v1 + v2
	case "*":
		res = v1 * v2
	}
	return res
}

type Monkey struct {
	items     []int
	operation Operation
	test      int
	ifTrue    int
	ifFalse   int
	inspected int
}

type MonkeyB struct {
	items     []map[int]int
	operation Operation
	test      int
	ifTrue    int
	ifFalse   int
	inspected int
}

type monkeysByInspected []*Monkey

func (m monkeysByInspected) Len() int {
	return len(m)
}

func (m monkeysByInspected) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m monkeysByInspected) Less(i, j int) bool {
	return m[i].inspected < m[j].inspected
}

type monkeysBByInspected []*MonkeyB

func (m monkeysBByInspected) Len() int {
	return len(m)
}

func (m monkeysBByInspected) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m monkeysBByInspected) Less(i, j int) bool {
	return m[i].inspected < m[j].inspected
}

func (m *Monkey) clearItems() {
	m.items = make([]int, 0)
}

func (m *MonkeyB) clearItems() {
	m.items = []map[int]int{}
}

func (m *Monkey) resolveItem(itemIndex int) (int, int) {
	val := m.items[itemIndex]
	newVal := m.operation.calculate(val) / 3
	if newVal%m.test == 0 {
		return newVal, m.ifTrue
	}
	return newVal, m.ifFalse
}

func (m *MonkeyB) resolveItem(itemIndex int) (map[int]int, int) {
	for k, v := range m.items[itemIndex] {
		m.items[itemIndex][k] = m.operation.calculate(v) % k
	}
	if m.items[itemIndex][m.test] == 0 {
		return m.items[itemIndex], m.ifTrue
	}
	return m.items[itemIndex], m.ifFalse
}

func parseOperation(s string) Operation {
	pieces := strings.Fields(strings.TrimPrefix(s, "  Operation: new = "))
	val1Str, operator, val2Str := pieces[0], pieces[1], pieces[2]
	val1 := Value{isSelf: true}
	if val1Str != "old" {
		val1.isSelf = false
		val1.value, _ = strconv.Atoi(val1Str)
	}
	val2 := Value{isSelf: true}
	if val2Str != "old" {
		val2.isSelf = false
		val2.value, _ = strconv.Atoi(val2Str)
	}

	return Operation{operator, val1, val2}
}

func parseMonkey(monkeyStr string) *Monkey {
	lines := strings.Split(monkeyStr, "\n")
	itemsStr := strings.TrimPrefix(lines[1], "  Starting items: ")
	itemStrs := strings.Split(itemsStr, ", ")
	items := make([]int, len(itemStrs))
	for i, v := range itemStrs {
		vInt, _ := strconv.Atoi(v)
		items[i] = vInt
	}

	operation := parseOperation(lines[2])
	test, _ := strconv.Atoi(strings.TrimPrefix(lines[3], "  Test: divisible by "))
	ifTrue, _ := strconv.Atoi(strings.TrimPrefix(lines[4], "    If true: throw to monkey "))
	ifFalse, _ := strconv.Atoi(strings.TrimPrefix(lines[5], "    If false: throw to monkey "))

	return &Monkey{items, operation, test, ifTrue, ifFalse, 0}
}

func parseMonkeyB(monkeyStr string) ([]int, *MonkeyB) {
	lines := strings.Split(monkeyStr, "\n")
	itemsStr := strings.TrimPrefix(lines[1], "  Starting items: ")
	itemStrs := strings.Split(itemsStr, ", ")
	items := make([]int, len(itemStrs))
	for i, v := range itemStrs {
		vInt, _ := strconv.Atoi(v)
		items[i] = vInt
	}

	operation := parseOperation(lines[2])
	test, _ := strconv.Atoi(strings.TrimPrefix(lines[3], "  Test: divisible by "))
	ifTrue, _ := strconv.Atoi(strings.TrimPrefix(lines[4], "    If true: throw to monkey "))
	ifFalse, _ := strconv.Atoi(strings.TrimPrefix(lines[5], "    If false: throw to monkey "))

	return items, &MonkeyB{operation: operation, test: test, ifTrue: ifTrue, ifFalse: ifFalse, inspected: 0}
}

func resolveTurn(monkeys []*Monkey, i int) {
	for itemIndex := range monkeys[i].items {
		newVal, newMonkeyIndex := monkeys[i].resolveItem(itemIndex)
		monkeys[newMonkeyIndex].items = append(monkeys[newMonkeyIndex].items, newVal)
		monkeys[i].inspected++
	}
	monkeys[i].clearItems()
}

func resolveTurnB(monkeys []*MonkeyB, i int) {
	for itemIndex := range monkeys[i].items {
		newVal, newMonkeyIndex := monkeys[i].resolveItem(itemIndex)
		monkeys[newMonkeyIndex].items = append(monkeys[newMonkeyIndex].items, newVal)
		monkeys[i].inspected++
	}
	monkeys[i].clearItems()
}

func partA(input string) {
	monkeysStrs := strings.Split(input, "\n\n")

	var monkeys []*Monkey
	for _, monkeyStr := range monkeysStrs {
		monkeys = append(monkeys, parseMonkey(monkeyStr))
	}

	rounds := 20

	for round := 0; round < rounds; round++ {
		for i := range monkeys {
			resolveTurn(monkeys, i)
		}
	}

	sort.Sort(monkeysByInspected(monkeys))

	fmt.Println(monkeys[len(monkeys)-1].inspected * monkeys[len(monkeys)-2].inspected)
}

func partB(input string) {
	monkeysStrs := strings.Split(input, "\n\n")

	var monkeys []*MonkeyB
	tempItems := make(map[int][]int)
	for i, monkeyStr := range monkeysStrs {
		itemsInt, monkey := parseMonkeyB(monkeyStr)
		monkeys = append(monkeys, monkey)
		tempItems[i] = itemsInt
	}

	tests := make([]int, len(monkeys))
	for i := range monkeys {
		tests[i] = monkeys[i].test
	}

	for i := range monkeys {
		for j := range tempItems[i] {
			item := make(map[int]int)
			for _, k := range tests {
				item[k] = tempItems[i][j] % k
			}
			monkeys[i].items = append(monkeys[i].items, item)
		}
	}

	rounds := 10000

	for round := 0; round < rounds; round++ {
		for i := range monkeys {
			resolveTurnB(monkeys, i)
		}
	}

	sort.Sort(monkeysBByInspected(monkeys))

	fmt.Println(monkeys[len(monkeys)-1].inspected * monkeys[len(monkeys)-2].inspected)
}

func main() {
	input := utils.ReadInputString(11)
	partA(input)
	partB(input)
}
