package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []byte
type Instruction [3]int

func (s *Stack) Push(b byte) {
	*s = append(*s, b)
}

func (s *Stack) BulkPush(slice []byte) {
	*s = append(*s, slice...)
}

func (s *Stack) Pop() byte {
	lastIndex := len(*s) - 1
	e := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return e
}

func (s *Stack) PopN(n int) []byte {
	lastIndex := len(*s)
	slice := (*s)[lastIndex-n:]
	*s = (*s)[:lastIndex-n]
	return slice
}

func ParseInstruction(s string) Instruction {
	s = s[5:]
	n, rest, _ := strings.Cut(s, " from ")
	nInt, _ := strconv.Atoi(n)
	from, to, _ := strings.Cut(rest, " to ")
	fromInt, _ := strconv.Atoi(from)
	toInt, _ := strconv.Atoi(to)

	return Instruction{nInt, fromInt - 1, toInt - 1}
}

func partA() {
	data, err := os.ReadFile("day05/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\r\n")
	stacksN := len(lines[0])/4 + 1
	towers := make([][]byte, stacksN)
	stacks := make([]Stack, stacksN)
	instructions := make([]Instruction, 0, len(lines)-stacksN-2)

	i := 0
	for {
		if !strings.Contains(lines[i], "[") {
			break
		}
		for j := 1; j < len(lines[i]); j += 4 {
			if b := lines[i][j]; b != ' ' {
				towers[j/4] = append(towers[j/4], b)
			}
		}
		i += 1
	}

	for _, line := range lines[i:] {
		if !strings.Contains(line, "move") {
			continue
		}
		instructions = append(instructions, ParseInstruction(line))
	}

	for i, tower := range towers {
		for j := len(tower) - 1; j >= 0; j-- {
			stacks[i].Push(tower[j])
		}
	}

	for _, ins := range instructions {
		for i := 0; i < ins[0]; i++ {
			stacks[ins[2]].Push(stacks[ins[1]].Pop())
		}
	}

	res := make([]byte, stacksN)
	for i, stack := range stacks {
		res[i] = stack.Pop()
	}

	fmt.Println(string(res))
}

func partB() {
	data, err := os.ReadFile("day05/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\r\n")
	stacksN := len(lines[0])/4 + 1
	towers := make([][]byte, stacksN)
	stacks := make([]Stack, stacksN)
	instructions := make([]Instruction, 0, len(lines)-stacksN-2)

	i := 0
	for {
		if !strings.Contains(lines[i], "[") {
			break
		}
		for j := 1; j < len(lines[i]); j += 4 {
			if b := lines[i][j]; b != ' ' {
				towers[j/4] = append(towers[j/4], b)
			}
		}
		i += 1
	}

	for _, line := range lines[i:] {
		if !strings.Contains(line, "move") {
			continue
		}
		instructions = append(instructions, ParseInstruction(line))
	}

	for i, tower := range towers {
		for j := len(tower) - 1; j >= 0; j-- {
			stacks[i].Push(tower[j])
		}
	}

	for _, ins := range instructions {
		stacks[ins[2]].BulkPush(stacks[ins[1]].PopN(ins[0]))
	}

	res := make([]byte, stacksN)
	for i, stack := range stacks {
		res[i] = stack.Pop()
	}

	fmt.Println(string(res))
}

func main() {
	partA()
	partB()
}
