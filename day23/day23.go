package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
)

type Position struct {
	x int
	y int
}

type Elf struct {
	pos      Position
	proposed Position
}

func (e *Elf) updateProposed(offset int, elves map[Position]bool) {
	choosenDir := -1
	for i := offset; i < offset+4; i++ {
		viable := true
		for _, cand := range regions[i%4] {
			if _, ok := elves[add(e.pos, cand)]; ok {
				viable = false
			}
		}
		if viable {
			choosenDir = i % 4
			break
		}
	}

	free := true
	for _, region := range regions {
		for _, space := range region {
			if _, ok := elves[add(e.pos, space)]; ok {
				free = false
				break
			}
		}
	}

	if choosenDir == -1 || free {
		e.proposed = e.pos
	} else {
		e.proposed = add(e.pos, dirs[choosenDir])
	}
}

var regions = [4][3]Position{
	{{-1, -1}, {0, -1}, {1, -1}},
	{{-1, 1}, {0, 1}, {1, 1}},
	{{-1, -1}, {-1, 0}, {-1, 1}},
	{{1, -1}, {1, 0}, {1, 1}},
}

var dirs = [4]Position{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func add(a, b Position) Position {
	return Position{a.x + b.x, a.y + b.y}
}

func parseElves(input [][]byte) map[*Elf]bool {
	elves := make(map[*Elf]bool)
	for i := range input {
		for j := range input[i] {
			if input[i][j] == '#' {
				pos := Position{j, i}
				elf := Elf{pos: pos, proposed: pos}
				elves[&elf] = true
			}
		}
	}

	return elves
}

func getEnclosingArea(elves map[*Elf]bool) int {
	var minX, maxX, minY, maxY int

	for elf := range elves {
		minX, maxX, minY, maxY = elf.pos.x, elf.pos.x, elf.pos.y, elf.pos.y
		break
	}

	for elf := range elves {
		if elf.pos.x < minX {
			minX = elf.pos.x
		}
		if elf.pos.x > maxX {
			maxX = elf.pos.x
		}
		if elf.pos.y < minY {
			minY = elf.pos.y
		}
		if elf.pos.y > maxY {
			maxY = elf.pos.y
		}
	}

	return (maxX - minX + 1) * (maxY - minY + 1)
}

func partA(input [][]byte) {
	elves := parseElves(input)
	occupied := map[Position]bool{}
	turns := 10
	offset := 0

	for elf := range elves {
		occupied[elf.pos] = true
	}

	for i := 0; i < turns; i++ {
		proposed := map[Position]int{}
		if offset >= 4 {
			offset = offset % 4
		}

		for elf := range elves {
			elf.updateProposed(offset, occupied)
			proposed[elf.proposed]++
		}

		occupied = make(map[Position]bool)

		for elf := range elves {
			if proposed[elf.proposed] == 1 {
				elf.pos = elf.proposed
			}
			occupied[elf.pos] = true
		}
		offset++
	}

	fmt.Println(getEnclosingArea(elves) - len(elves))
}

func partB(input [][]byte) {
	elves := parseElves(input)
	occupied := map[Position]bool{}
	offset := 0

	for elf := range elves {
		occupied[elf.pos] = true
	}

	i := 1
	for {
		proposed := map[Position]int{}
		if offset >= 4 {
			offset = offset % 4
		}

		for elf := range elves {
			elf.updateProposed(offset, occupied)
			proposed[elf.proposed]++
		}

		occupied = make(map[Position]bool)
		stationary := 0

		for elf := range elves {
			if elf.pos == elf.proposed {
				stationary++
			}
			if proposed[elf.proposed] == 1 {
				elf.pos = elf.proposed
			}
			occupied[elf.pos] = true
		}

		if stationary == len(elves) {
			fmt.Println(i)
			return
		}
		offset++
		i++
	}
}

func main() {
	input := utils.ReadInputByteLines(23, "\n")
	partA(input)
	partB(input)
}
