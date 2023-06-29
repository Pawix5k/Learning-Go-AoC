package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
)

type Position struct {
	x int
	y int
}

type Blizzard struct {
	pos     Position
	nextPos Position
	move    Position
}

func (b *Blizzard) GetNextMove(grid Grid) Position {
	cand := add(b.pos, b.move)
	if cand.x < grid.minX {
		cand.x = grid.maxX
	} else if cand.x > grid.maxX {
		cand.x = grid.minX
	}
	if cand.y < grid.minY {
		cand.y = grid.maxY
	} else if cand.y > grid.maxY {
		cand.y = grid.minY
	}
	return cand
}

func (b *Blizzard) Update(grid Grid) {
	b.pos = b.nextPos
	b.nextPos = b.GetNextMove(grid)
}

type Grid struct {
	start Position
	end   Position
	minX  int
	maxX  int
	minY  int
	maxY  int
}

func (g Grid) IsInRange(pos Position) bool {
	if pos.x >= g.minX && pos.x <= g.maxX && pos.y >= g.minY && pos.y <= g.maxY {
		return true
	}
	if pos == g.start || pos == g.end {
		return true
	}
	return false
}

func add(a, b Position) Position {
	return Position{a.x + b.x, a.y + b.y}
}

var moves = map[byte]Position{
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
	'^': {0, -1},
}

func parseInput(input [][]byte) (map[*Blizzard]bool, Grid) {
	height := len(input)
	width := len(input[0])

	grid := Grid{minX: 1, maxX: width - 2, minY: 1, maxY: height - 2}
	start := Position{grid.minX, grid.minY - 1}
	end := Position{grid.maxX, grid.maxY + 1}
	grid.start = start
	grid.end = end

	blizzards := make(map[*Blizzard]bool)

	for i := range input {
		for j := range input[i] {
			if v, ok := moves[input[i][j]]; ok {
				blizzard := Blizzard{pos: Position{j, i}, move: v}
				blizzard.nextPos = blizzard.GetNextMove(grid)
				blizzards[&blizzard] = true
			}
		}
	}
	return blizzards, grid
}

func partA(input [][]byte) {
	blizzards, grid := parseInput(input)
	step := 1
	possible := map[Position]bool{grid.start: true}

	for {
		unavailable := make(map[Position]bool)
		nextPossible := make(map[Position]bool)

		for blizzard := range blizzards {
			unavailable[blizzard.nextPos] = true
		}

		for pos := range possible {
			if _, ok := unavailable[pos]; !ok {
				nextPossible[pos] = true
			}
			for _, vector := range moves {
				cand := add(pos, vector)
				if _, ok := unavailable[cand]; !ok && grid.IsInRange(cand) {
					nextPossible[cand] = true
				}
			}
		}

		if _, ok := nextPossible[grid.end]; ok {
			fmt.Println(step)
			return
		}

		step++
		for blizzard := range blizzards {
			blizzard.Update(grid)
		}
		possible = nextPossible
	}
}

func partB(input [][]byte) {
	blizzards, grid := parseInput(input)
	step := 1
	checkpointsN := 0
	checkpoints := map[bool]Position{
		false: grid.start,
		true:  grid.end,
	}
	curCheckpointStart := true

	possible := map[Position]bool{grid.start: true}

	for {
		unavailable := make(map[Position]bool)
		nextPossible := make(map[Position]bool)

		for blizzard := range blizzards {
			unavailable[blizzard.nextPos] = true
		}

		for pos := range possible {
			if _, ok := unavailable[pos]; !ok {
				nextPossible[pos] = true
			}
			for _, vector := range moves {
				cand := add(pos, vector)
				if _, ok := unavailable[cand]; !ok && grid.IsInRange(cand) {
					nextPossible[cand] = true
				}
			}
		}

		if _, ok := nextPossible[checkpoints[curCheckpointStart]]; ok {
			if checkpointsN == 2 {
				fmt.Println(step)
				return
			}
			nextPossible = make(map[Position]bool)
			nextPossible[checkpoints[curCheckpointStart]] = true
			checkpointsN++
			curCheckpointStart = !curCheckpointStart
		}

		step++
		for blizzard := range blizzards {
			blizzard.Update(grid)
		}
		possible = nextPossible
	}
}

func main() {
	input := utils.ReadInputByteLines(24, "\n")
	partA(input)
	partB(input)
}
