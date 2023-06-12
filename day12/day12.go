package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
)

type Point struct {
	x int
	y int
}

func isValidMoveUp(visited [][]int, lines [][]byte, point Point, dx, dy int) (bool, bool) {
	x := point.x + dx
	y := point.y + dy

	if x < 0 || x >= len(lines) || y < 0 || y >= len(lines[0]) {
		return false, false
	}

	if lines[x][y] == 'E' && lines[point.x][point.y] == 'z' {
		return true, true
	}

	if visited[x][y] == 0 && (lines[x][y] < lines[point.x][point.y] || lines[x][y]-lines[point.x][point.y] <= 1) || lines[x][y] == 'a' && lines[point.x][point.y] == 'S' {
		return true, false
	}

	return false, false
}

func isValidMoveDown(visited [][]int, lines [][]byte, point Point, dx, dy int) (bool, bool) {
	x := point.x + dx
	y := point.y + dy

	if x < 0 || x >= len(lines) || y < 0 || y >= len(lines[0]) {
		return false, false
	}

	if (lines[x][y] == 'a' || lines[x][y] == 'S') && lines[point.x][point.y] == 'b' {
		return true, true
	}

	if lines[point.x][point.y] == 'E' {
		if lines[x][y] == 'z' {
			return true, false
		} else {
			return false, false
		}
	}

	if visited[x][y] == 0 && lines[x][y] >= lines[point.x][point.y]-1 {
		return true, false
	}

	return false, false
}

func findStart(lines [][]byte, marker byte) Point {
	for x, row := range lines {
		for y, b := range row {
			if b == marker {
				return Point{x, y}
			}
		}
	}
	return Point{-1, -1}
}

func partA(lines [][]byte) {
	start := findStart(lines, 'S')
	curStep := 1
	recentlyVisited := []Point{start}
	visited := make([][]int, len(lines))
	for i := range visited {
		visited[i] = make([]int, len(lines[0]))
	}

	dxy := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for {
		curRecentlyVisited := []Point{}
		for _, point := range recentlyVisited {
			for _, offset := range dxy {
				valid, finish := isValidMoveUp(visited, lines, point, offset.x, offset.y)
				if finish {
					fmt.Println(curStep)
					return
				}
				if valid {
					curRecentlyVisited = append(curRecentlyVisited, Point{point.x + offset.x, point.y + offset.y})
					visited[point.x+offset.x][point.y+offset.y] = curStep
				}
			}
		}
		recentlyVisited = curRecentlyVisited
		curStep++
	}
}

func partB(lines [][]byte) {
	start := findStart(lines, 'E')
	curStep := 1
	recentlyVisited := []Point{start}
	visited := make([][]int, len(lines))
	for i := range visited {
		visited[i] = make([]int, len(lines[0]))
	}

	dxy := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for {
		curRecentlyVisited := []Point{}
		for _, point := range recentlyVisited {
			for _, offset := range dxy {
				valid, finish := isValidMoveDown(visited, lines, point, offset.x, offset.y)
				if finish {
					fmt.Println(curStep)
					return
				}
				if valid {
					curRecentlyVisited = append(curRecentlyVisited, Point{point.x + offset.x, point.y + offset.y})
					visited[point.x+offset.x][point.y+offset.y] = curStep
				}
			}
		}
		recentlyVisited = curRecentlyVisited
		curStep++
	}
}

func main() {
	lines := utils.ReadInputByteLines(12, "\n")
	partA(lines)
	partB(lines)
}
