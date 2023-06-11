package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Point [2]int

func parseLine(line string) (string, int) {
	lineSplit := strings.Split(line, " ")
	dir, nStr := lineSplit[0], lineSplit[1]
	n, _ := strconv.Atoi(nStr)
	return dir, n
}

func resolveLeading(leading Point, dir string) Point {
	switch dir {
	case "U":
		leading[1]++
	case "D":
		leading[1]--
	case "L":
		leading[0]--
	case "R":
		leading[0]++
	}

	return leading
}

func resolveTrailing(leading, trailing Point) Point {
	distX := leading[0] - trailing[0]
	distY := leading[1] - trailing[1]

	if (distX == -2 || distX == 2) && (distY == -1 || distY == 1) {
		trailing[0] += distX / 2
		trailing[1] += distY
	} else if (distX == -1 || distX == 1) && (distY == -2 || distY == 2) {
		trailing[0] += distX
		trailing[1] += distY / 2
	} else if distX == -2 || distX == 2 || distY == -2 || distY == 2 {
		trailing[0] += distX / 2
		trailing[1] += distY / 2
	}

	return trailing
}

func partA(lines []string) {
	visited := make(map[Point]bool)
	head, tail := Point{}, Point{}

	for _, line := range lines {
		dir, n := parseLine(line)
		for i := 0; i < n; i++ {
			head = resolveLeading(head, dir)
			tail = resolveTrailing(head, tail)
			visited[tail] = true
		}
	}

	fmt.Println(len(visited))
}

func partB(lines []string) {
	visited := make(map[Point]bool)
	rope := make([]Point, 10)

	for _, line := range lines {
		dir, n := parseLine(line)
		for i := 0; i < n; i++ {
			rope[0] = resolveLeading(rope[0], dir)
			for i := range rope[:9] {
				rope[i+1] = resolveTrailing(rope[i], rope[i+1])
			}
			visited[rope[9]] = true
		}
	}

	fmt.Println(len(visited))
}

func main() {
	lines := utils.ReadInputStringLines(9, "\n")
	partA(lines)
	partB(lines)
}
