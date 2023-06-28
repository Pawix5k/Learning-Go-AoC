package main

import (
	"bytes"
	"example/pawix5k/aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector [3]int
type Matrix [3]Vector

var rotations = map[string]map[int]Matrix{
	"x": {
		90:  Matrix{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}},
		270: Matrix{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}},
	},
	"y": {
		90:  Matrix{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}},
		270: Matrix{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}},
	},
	"z": {
		90:  Matrix{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}},
		270: Matrix{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}},
	},
}

type Side struct {
	pos       Vector
	dir       Vector
	start     Position
	neighbors map[Direction]*Side
}

type Cube struct {
	sides   map[Vector]*Side
	size    int
	curFace *Side
}

type Position struct {
	x int
	y int
}

type Instruction struct {
	insType   string
	steps     int
	changeDir string
}

type Direction int

const (
	right Direction = iota
	down
	left
	up
)

var blockTypes = map[string]byte{
	"illegal": 32,
	"open":    46,
	"wall":    35,
}

func parseInstructions(instructionsStr string) []Instruction {
	instructions := []Instruction{}
	number := ""
	for i := range instructionsStr {
		if utils.IsNumber(instructionsStr[i]) {
			number += string(instructionsStr[i])
		} else {
			steps, _ := strconv.Atoi(number)
			number = ""
			instructions = append(instructions, Instruction{insType: "steps", steps: steps})
			instructions = append(instructions, Instruction{insType: "dir", changeDir: string(instructionsStr[i])})
		}
	}
	if len(number) != 0 {
		steps, _ := strconv.Atoi(number)
		instructions = append(instructions, Instruction{insType: "steps", steps: steps})
	}

	return instructions
}

func findStartPoint(grid [][]byte) Position {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == blockTypes["open"] {
				return Position{j, i}
			}
		}
	}

	return Position{}
}

func changeDir(curDir Direction, changeDir string) Direction {
	if changeDir == "R" && curDir == up {
		return right
	} else if changeDir == "R" {
		return curDir + 1
	}
	if curDir == right {
		return up
	}
	return curDir - 1
}

func add(a, b Position) Position {
	return Position{a.x + b.x, a.y + b.y}
}

func getNextPos(curPos Position, curDir Direction, grid [][]byte) (Position, bool) {
	offset := Position{}
	switch curDir {
	case up:
		offset.y = -1
	case right:
		offset.x = 1
	case down:
		offset.y = 1
	case left:
		offset.x = -1
	}

	nextPosCandidate := add(curPos, offset)
	for nextPosCandidate.x < 0 || nextPosCandidate.x >= len(grid[0]) || nextPosCandidate.y < 0 || nextPosCandidate.y >= len(grid) || grid[nextPosCandidate.y][nextPosCandidate.x] == blockTypes["illegal"] {
		if nextPosCandidate.x < 0 {
			nextPosCandidate.x = len(grid[0]) - 1
		} else if nextPosCandidate.x >= len(grid[0]) {
			nextPosCandidate.x = 0
		} else if nextPosCandidate.y < 0 {
			nextPosCandidate.y = len(grid) - 1
		} else if nextPosCandidate.y >= len(grid) {
			nextPosCandidate.y = 0
		} else {
			nextPosCandidate = add(nextPosCandidate, offset)
		}
	}
	isLegal := true
	if grid[nextPosCandidate.y][nextPosCandidate.x] == blockTypes["wall"] {
		isLegal = false
	}
	return nextPosCandidate, isLegal
}

func addSpaces(grid [][]byte) [][]byte {
	maxWidth := 0
	for _, row := range grid {
		if len(row) > maxWidth {
			maxWidth = len(row)
		}
	}

	for i := range grid {
		for len(grid[i]) != maxWidth {
			grid[i] = append(grid[i], ' ')
		}
	}

	return grid
}

func getSquareSize(grid [][]byte) int {
	width := len(grid[0])
	height := len(grid)

	max := width
	if height > width {
		max = height
	}
	return max / 4
}

func isValidPlace(b byte) bool {
	if b == '#' || b == '.' {
		return true
	}
	return false
}

func matmul(v1 Vector, v2 [3]Vector) Vector {
	newVector := Vector{}
	for i := range v1 {
		num := 0
		for j := range v1 {
			num += v1[j] * v2[j][i]
		}
		newVector[i] = num
	}

	return newVector
}

func rotate(pos Vector, dir Vector, d Direction) (Vector, Vector) {
	if d == right || d == left {
		dirAxis := "x"
		val := dir[0]
		if dir[1] != 0 {
			dirAxis = "y"
			val = dir[1]
		} else if dir[2] != 0 {
			dirAxis = "z"
			val = dir[2]
		}
		valIsOne := val == 1
		dIsW := d == left

		rot := rotations[dirAxis][270]
		if valIsOne != dIsW {
			rot = rotations[dirAxis][90]
		}

		return matmul(pos, rot), dir
	}

	axis := "x"
	crossVal := pos[1]*dir[2] - pos[2]*dir[1]
	if pos[1] == 0 && dir[1] == 0 {
		axis = "y"
		crossVal = pos[2]*dir[0] - pos[0]*dir[2]
	} else if pos[2] == 0 && dir[2] == 0 {
		axis = "z"
		crossVal = pos[0]*dir[1] - pos[1]*dir[0]
	}

	crossValLtZero := crossVal < 0
	dIsN := d == up

	rot := rotations[axis][90]
	if crossValLtZero != dIsN {
		rot = rotations[axis][270]
	}

	return matmul(pos, rot), matmul(dir, rot)
}

func getCube(grid [][]byte) Cube {
	squareSize := getSquareSize(grid)
	start := findStartPoint(grid)
	startPos := Vector{0, 0, 1}
	top := &Side{pos: startPos, dir: [3]int{1, 0, 0}, start: start}
	cube := Cube{sides: map[Vector]*Side{}, size: squareSize}

	var makeCubeRecursively func(Side)
	makeCubeRecursively = func(prevSide Side) {
		cube.sides[prevSide.pos] = &prevSide
		gridPos := prevSide.start

		new := Position{gridPos.x, gridPos.y - squareSize}
		if new.x >= 0 && new.x < len(grid[0]) && new.y >= 0 && new.y < len(grid) && isValidPlace(grid[new.y][new.x]) {
			newPos, newDir := rotate(prevSide.pos, prevSide.dir, up)
			side := Side{pos: newPos, dir: newDir, start: new}
			if _, ok := cube.sides[side.pos]; !ok {
				makeCubeRecursively(side)
			}
		}
		new = Position{gridPos.x + squareSize, gridPos.y}
		if new.x >= 0 && new.x < len(grid[0]) && new.y >= 0 && new.y < len(grid) && isValidPlace(grid[new.y][new.x]) {
			newPos, newDir := rotate(prevSide.pos, prevSide.dir, right)
			side := Side{pos: newPos, dir: newDir, start: new}
			if _, ok := cube.sides[side.pos]; !ok {
				makeCubeRecursively(side)
			}
		}
		new = Position{gridPos.x, gridPos.y + squareSize}
		if new.x >= 0 && new.x < len(grid[0]) && new.y >= 0 && new.y < len(grid) && isValidPlace(grid[new.y][new.x]) {
			newPos, newDir := rotate(prevSide.pos, prevSide.dir, down)
			side := Side{pos: newPos, dir: newDir, start: new}
			if _, ok := cube.sides[side.pos]; !ok {
				makeCubeRecursively(side)
			}
		}
		new = Position{gridPos.x - squareSize, gridPos.y}
		if new.x >= 0 && new.x < len(grid[0]) && new.y >= 0 && new.y < len(grid) && isValidPlace(grid[new.y][new.x]) {
			newPos, newDir := rotate(prevSide.pos, prevSide.dir, left)
			side := Side{pos: newPos, dir: newDir, start: new}
			if _, ok := cube.sides[side.pos]; !ok {
				makeCubeRecursively(side)
			}
		}
	}
	makeCubeRecursively(*top)

	for v := range cube.sides {
		side := cube.sides[v]
		sideU, _ := rotate(side.pos, side.dir, up)
		sideR, _ := rotate(side.pos, side.dir, right)
		sideD, _ := rotate(side.pos, side.dir, down)
		sideL, _ := rotate(side.pos, side.dir, left)

		side.neighbors = map[Direction]*Side{
			up:    cube.sides[sideU],
			right: cube.sides[sideR],
			down:  cube.sides[sideD],
			left:  cube.sides[sideL],
		}
	}

	cube.curFace = cube.sides[startPos]

	return cube
}

func getPosAndDirOnAnotherFace(nextPosCandidate Position, nextDirCandidate Direction, cube Cube) (Position, Direction, *Side) {
	dirOnSide := up
	dist := nextPosCandidate.x - cube.curFace.start.x
	if nextPosCandidate.x >= cube.curFace.start.x+cube.size {
		dirOnSide = right
		dist = nextPosCandidate.y - cube.curFace.start.y
	} else if nextPosCandidate.y >= cube.curFace.start.y+cube.size {
		dirOnSide = down
		dist = cube.curFace.start.x + cube.size - nextPosCandidate.x - 1
	} else if nextPosCandidate.x < cube.curFace.start.x {
		dirOnSide = left
		dist = cube.curFace.start.y + cube.size - nextPosCandidate.y - 1
	}

	newSide := cube.curFace.neighbors[dirOnSide]
	dirFromNeighbor := up
	newPos := Position{newSide.start.x + cube.size - dist - 1, newSide.start.y}

	if newSide.neighbors[right].pos == cube.curFace.pos {
		dirFromNeighbor = right
		newPos = Position{newSide.start.x + cube.size - 1, newSide.start.y + cube.size - dist - 1}
	} else if newSide.neighbors[down].pos == cube.curFace.pos {
		dirFromNeighbor = down
		newPos = Position{newSide.start.x + dist, newSide.start.y + cube.size - 1}
	} else if newSide.neighbors[left].pos == cube.curFace.pos {
		dirFromNeighbor = left
		newPos = Position{newSide.start.x, newSide.start.y + dist}
	}

	newDir := nextDirCandidate + 2 + dirFromNeighbor - dirOnSide
	if newDir > 3 {
		newDir = newDir % 4
	} else if newDir < 0 {
		newDir += 4
	}

	return newPos, newDir, newSide
}

func getNextPosOnCube(pos Position, dir Direction, grid [][]byte, cube Cube) (Position, Direction, *Side, bool) {
	offset := Position{}
	switch dir {
	case up:
		offset.y = -1
	case right:
		offset.x = 1
	case down:
		offset.y = 1
	case left:
		offset.x = -1
	}

	nextDirCandidate, nextFace := dir, cube.curFace
	nextPosCandidate := add(pos, offset)
	posNotInFace := nextPosCandidate.x < cube.curFace.start.x || nextPosCandidate.x >= cube.curFace.start.x+cube.size || nextPosCandidate.y < cube.curFace.start.y || nextPosCandidate.y >= cube.curFace.start.y+cube.size
	if posNotInFace {
		nextPosCandidate, nextDirCandidate, nextFace = getPosAndDirOnAnotherFace(nextPosCandidate, nextDirCandidate, cube)
	}
	isLegal := true
	if grid[nextPosCandidate.y][nextPosCandidate.x] == blockTypes["wall"] {
		isLegal = false
	}
	return nextPosCandidate, nextDirCandidate, nextFace, isLegal
}

func partA(input []byte) {
	split := bytes.Split(input, []byte("\n\n"))
	grid, instructionsStr := bytes.Split(split[0], []byte("\n")), string(split[1])

	grid = addSpaces(grid)

	instructionsStr = strings.TrimSpace(instructionsStr)

	instructions := parseInstructions(instructionsStr)

	curPos := findStartPoint(grid)
	curDir := right

	for _, instruction := range instructions {
		if instruction.insType == "dir" {
			curDir = changeDir(curDir, instruction.changeDir)
			continue
		}
		for j := 0; j < instruction.steps; j++ {
			nextPos, isLegal := getNextPos(curPos, curDir, grid)
			if !isLegal {
				break
			}
			curPos = nextPos
		}
	}

	fmt.Println(1000*(curPos.y+1) + 4*(curPos.x+1) + int(curDir))
}

func partB(input []byte) {
	split := bytes.Split(input, []byte("\n\n"))
	grid, instructionsStr := bytes.Split(split[0], []byte("\n")), string(split[1])
	grid = addSpaces(grid)
	instructionsStr = strings.TrimSpace(instructionsStr)
	instructions := parseInstructions(instructionsStr)

	cube := getCube(grid)
	pos := findStartPoint(grid)
	dir := right

	for _, instruction := range instructions {
		if instruction.insType == "dir" {
			dir = changeDir(dir, instruction.changeDir)
			continue
		}
		for j := 0; j < instruction.steps; j++ {
			nextPos, nextDir, newSide, isLegal := getNextPosOnCube(pos, dir, grid, cube)
			if !isLegal {
				break
			}
			pos, dir = nextPos, nextDir
			cube.curFace = newSide
		}
	}

	fmt.Println(1000*(pos.y+1) + 4*(pos.x+1) + int(dir))
}

func main() {
	utils.DownloadInput(22)
	input, _ := os.ReadFile(utils.GetFilePath(22))
	partA(input)
	partB(input)
}
