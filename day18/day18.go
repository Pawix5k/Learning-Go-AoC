package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Cube struct {
	x int
	y int
	z int
}

func parseLine(line string) Cube {
	fields := strings.Split(line, ",")
	xStr, yStr, zStr := fields[0], fields[1], fields[2]
	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)
	z, _ := strconv.Atoi(zStr)

	return Cube{x, y, z}
}

func countExposedSides(cube Cube, cubes map[Cube]bool) int {
	total := 6
	if _, ok := cubes[Cube{cube.x - 1, cube.y, cube.z}]; ok {
		total--
	}
	if _, ok := cubes[Cube{cube.x + 1, cube.y, cube.z}]; ok {
		total--
	}
	if _, ok := cubes[Cube{cube.x, cube.y - 1, cube.z}]; ok {
		total--
	}
	if _, ok := cubes[Cube{cube.x, cube.y + 1, cube.z}]; ok {
		total--
	}
	if _, ok := cubes[Cube{cube.x, cube.y, cube.z - 1}]; ok {
		total--
	}
	if _, ok := cubes[Cube{cube.x, cube.y, cube.z + 1}]; ok {
		total--
	}

	return total
}

func getNeighbors(base Cube) []Cube {
	return []Cube{
		{base.x - 1, base.y, base.z}, {base.x + 1, base.y, base.z},
		{base.x, base.y - 1, base.z}, {base.x, base.y + 1, base.z},
		{base.x, base.y, base.z - 1}, {base.x, base.y, base.z + 1},
	}
}

func partA(lines []string) {
	cubes := map[Cube]bool{}
	for _, line := range lines {
		cubes[parseLine(line)] = true
	}
	totalSides := 0
	for cube := range cubes {
		totalSides += countExposedSides(cube, cubes)
	}

	fmt.Println(totalSides)
}

func partB(lines []string) {
	cubes := map[Cube]bool{}
	outer := map[Cube]bool{}
	inner := map[Cube]bool{}
	fCube := parseLine(lines[0])
	minX, maxX, minY, maxY, minZ, maxZ := fCube.x, fCube.x, fCube.y, fCube.y, fCube.z, fCube.z
	for _, line := range lines {
		cube := parseLine(line)
		cubes[cube] = true
		if cube.x < minX {
			minX = cube.x
		}
		if cube.x > maxX {
			maxX = cube.x
		}
		if cube.y < minY {
			minY = cube.y
		}
		if cube.y > maxY {
			maxY = cube.y
		}
		if cube.z < minZ {
			minZ = cube.z
		}
		if cube.z > maxZ {
			maxZ = cube.z
		}
	}

	queue := []Cube{{minX - 1, minY - 1, minZ - 1}}
	outer[Cube{minX - 1, minY - 1, minZ - 1}] = true
	for len(queue) != 0 {
		newQueue := []Cube{}
		for _, cube := range queue {
			for _, neighbor := range getNeighbors(cube) {
				inRange := neighbor.x >= minX-1 && neighbor.x <= maxX+1 &&
					neighbor.y >= minY-1 && neighbor.y <= maxY+1 &&
					neighbor.z >= minZ-1 && neighbor.z <= maxZ+1
				if inRange && !outer[neighbor] && !cubes[neighbor] {
					newQueue = append(newQueue, neighbor)
					outer[neighbor] = true
				}
			}
			queue = newQueue
		}
	}

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				cube := Cube{x, y, z}
				if !outer[cube] && !cubes[cube] {
					inner[cube] = true
				}
			}
		}
	}

	totalSides := 0
	for cube := range cubes {
		totalSides += countExposedSides(cube, cubes)
	}
	for cube := range inner {
		totalSides -= countExposedSides(cube, inner)
	}

	fmt.Println(totalSides)
}

func main() {
	lines := utils.ReadInputStringLines(18, "\n")
	partA(lines)
	partB(lines)
}
