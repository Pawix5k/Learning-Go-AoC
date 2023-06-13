package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func parseLine(line string) []Point {
	pointStrs := strings.Split(line, " -> ")
	points := make([]Point, len(pointStrs))
	for i, pointStr := range pointStrs {
		pointStrSplit := strings.Split(pointStr, ",")
		x, _ := strconv.Atoi(pointStrSplit[0])
		y, _ := strconv.Atoi(pointStrSplit[1])
		points[i] = Point{x, y}
	}

	return points
}

func resolveFall(stable map[Point]bool, pos Point) (Point, bool) {
	down := Point{pos.x, pos.y + 1}
	if stable[down] != true {
		return down, false
	}
	left := Point{pos.x - 1, pos.y + 1}
	if stable[left] != true {
		return left, false
	}
	right := Point{pos.x + 1, pos.y + 1}
	if stable[right] != true {
		return right, false
	}
	return pos, true
}

func resolveFallWithFloor(stable map[Point]bool, pos Point, floor int) (Point, bool) {
	down := Point{pos.x, pos.y + 1}
	if down.y == floor {
		return pos, true
	}
	if stable[down] != true {
		return down, false
	}
	left := Point{pos.x - 1, pos.y + 1}
	if stable[left] != true {
		return left, false
	}
	right := Point{pos.x + 1, pos.y + 1}
	if stable[right] != true {
		return right, false
	}
	return pos, true
}

func partA(lines []string) {
	start := Point{500, 0}
	solid := map[Point]bool{}

	lowest := start.y

	for _, line := range lines {
		wall := parseLine(line)
		for i := 0; i < len(wall)-1; i++ {
			if wall[i].x == wall[i+1].x {
				x := wall[i].x
				y1, y2 := wall[i].y, wall[i+1].y
				if y1 > y2 {
					y1, y2 = y2, y1
				}

				for y := y1; y <= y2; y++ {
					solid[Point{x, y}] = true
					if y > lowest {
						lowest = y
					}
				}
			} else {
				y := wall[i].y

				if y > lowest {
					lowest = y
				}

				x1, x2 := wall[i].x, wall[i+1].x
				if x1 > x2 {
					x1, x2 = x2, x1
				}

				for x := x1; x <= x2; x++ {
					solid[Point{x, y}] = true
				}
			}
		}
	}

	particles := 0

	for {
		pos := start
		for {
			newPos, isStable := resolveFall(solid, pos)
			if newPos.y >= lowest {
				fmt.Println(particles)
				return
			}
			if isStable {
				solid[newPos] = true
				break
			}
			pos = newPos
		}
		particles++
	}
}

func partB(lines []string) {
	start := Point{500, 0}
	solid := map[Point]bool{}

	lowest := start.y

	for _, line := range lines {
		wall := parseLine(line)
		for i := 0; i < len(wall)-1; i++ {
			if wall[i].x == wall[i+1].x {
				x := wall[i].x
				y1, y2 := wall[i].y, wall[i+1].y
				if y1 > y2 {
					y1, y2 = y2, y1
				}

				for y := y1; y <= y2; y++ {
					solid[Point{x, y}] = true
					if y > lowest {
						lowest = y
					}
				}
			} else {
				y := wall[i].y

				if y > lowest {
					lowest = y
				}

				x1, x2 := wall[i].x, wall[i+1].x
				if x1 > x2 {
					x1, x2 = x2, x1
				}

				for x := x1; x <= x2; x++ {
					solid[Point{x, y}] = true
				}
			}
		}
	}

	lowest += 2

	particles := 0

	for {
		pos := start
		for {
			newPos, isStable := resolveFallWithFloor(solid, pos, lowest)
			if isStable && newPos == start {
				fmt.Println(particles + 1)
				return
			} else if isStable {
				solid[newPos] = true
				break
			}
			pos = newPos
		}
		particles++
	}
}

func main() {
	lines := utils.ReadInputStringLines(14, "\n")
	partA(lines)
	partB(lines)
}
