package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
)

type Point struct {
	x int
	y int
}

func getRockPoints(shape string, h int) []*Point {
	switch shape {
	case "_":
		return []*Point{{2, h + 4}, {3, h + 4}, {4, h + 4}, {5, h + 4}}
	case "+":
		return []*Point{{2, h + 5}, {3, h + 5}, {3, h + 4}, {3, h + 6}, {4, h + 5}}
	case "L":
		return []*Point{{2, h + 4}, {3, h + 4}, {4, h + 4}, {4, h + 5}, {4, h + 6}}
	case "|":
		return []*Point{{2, h + 4}, {2, h + 5}, {2, h + 6}, {2, h + 7}}
	case "o":
		return []*Point{{2, h + 4}, {3, h + 4}, {2, h + 5}, {3, h + 5}}
	}
	return []*Point{}
}

func makeHoriznotalMove(rock []*Point, solidBlocks []map[int]bool, direction string) {
	offset := 1
	if direction == "<" {
		offset = -1
	}
	poss := true
	for _, point := range rock {
		if point.x+offset < 0 || point.x+offset >= len(solidBlocks) || solidBlocks[point.x+offset][point.y] == true {
			poss = false
			break
		}
	}
	if poss {
		for _, point := range rock {
			point.x += offset
		}
	}
}

func makeVerticalMove(rock []*Point, solidBlocks []map[int]bool) (bool, int) {
	solidified := false
	for _, point := range rock {
		if point.y-1 <= 0 || solidBlocks[point.x][point.y-1] == true {
			solidified = true
			break
		}
	}
	maxH := 0
	for _, point := range rock {
		if solidified {
			solidBlocks[point.x][point.y] = true
			if point.y > maxH {
				maxH = point.y
			}
		} else {
			point.y += -1
		}
	}

	return solidified, maxH
}

type GameState struct {
	shape string
	fell  int
}

func partA(input string) {
	maxH := 0
	solidBlocks := make([]map[int]bool, 7)
	for i := range solidBlocks {
		solidBlocks[i] = make(map[int]bool)
	}
	rockTypes := [...]string{"_", "+", "L", "|", "o"}
	rocksN := 2022

	dirI := 0

	for i := 0; i < rocksN; i++ {
		rock := getRockPoints(rockTypes[i%len(rockTypes)], maxH)
		for {
			direction := string(input[dirI%len(input)])
			dirI++
			makeHoriznotalMove(rock, solidBlocks, direction)
			solidified, newMaxH := makeVerticalMove(rock, solidBlocks)
			if newMaxH > maxH {
				maxH = newMaxH
			}
			if solidified {
				break
			}
		}
	}

	fmt.Println(maxH)
}

func partB(input string) {
	maxH := 0
	solidBlocks := make([]map[int]bool, 7)
	for i := range solidBlocks {
		solidBlocks[i] = make(map[int]bool)
	}
	rockTypes := [...]string{"_", "+", "L", "|", "o"}
	rocksN := 1_000_000_000_000
	mem := map[GameState][2]int{}
	cashed := 0

	dirI := 0

	for i := 0; i < rocksN; i++ {
		rock := getRockPoints(rockTypes[i%len(rockTypes)], maxH)
		startDirI := dirI
		for {
			direction := string(input[dirI%len(input)])
			if dirI%len(input) == 0 {
				if thenState, ok := mem[GameState{rockTypes[i%len(rockTypes)], dirI - startDirI}]; cashed == 0 && ok {
					thenH, thenRocks := thenState[0], thenState[1]
					deltaRocks := i - thenRocks
					doNIters := (rocksN - i) / deltaRocks
					cashed = (maxH - thenH) * doNIters
					i += deltaRocks * doNIters
				} else {
					mem[GameState{rockTypes[i%len(rockTypes)], dirI - startDirI}] = [2]int{maxH, i}
				}
			}
			dirI++
			makeHoriznotalMove(rock, solidBlocks, direction)
			solidified, newMaxH := makeVerticalMove(rock, solidBlocks)
			if newMaxH > maxH {
				maxH = newMaxH
			}
			if solidified {
				break
			}
		}
	}

	fmt.Println(maxH + cashed)
}

func main() {
	input := utils.ReadInputString(17)
	partA(input)
	partB(input)
}
