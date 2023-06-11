package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseLine(line string) (op string, i int) {
	lineSplit := strings.Fields(line)
	if len(lineSplit) == 1 {
		return lineSplit[0], 0
	}
	i, _ = strconv.Atoi(lineSplit[1])
	return lineSplit[0], i
}

func partA(lines []string) {
	findCycles := [6]int{}
	for i := range findCycles {
		findCycles[i] = 20 + i*40
	}
	curCycleIndex := 0
	curCycle := 1
	curReg := 1
	sum := 0

	for _, line := range lines {
		op, i := parseLine(line)

		curCycle++
		if curCycle == findCycles[curCycleIndex] {
			sum += curCycle * curReg
			curCycleIndex++
			if curCycleIndex == len(findCycles) {
				break
			}
		}

		if op == "addx" {
			curCycle++
			curReg += i

			if curCycle == findCycles[curCycleIndex] {
				sum += curCycle * curReg
				curCycleIndex++
				if curCycleIndex == len(findCycles) {
					break
				}
			}
		}
	}

	fmt.Println(sum)
}

func partB(lines []string) {
	screen := make([]bool, 240)
	curCycle := 1
	curReg := 1

	for _, line := range lines {
		op, i := parseLine(line)
		screen = paintPixel(screen, curCycle, curReg)
		curCycle++

		if op == "addx" {
			screen = paintPixel(screen, curCycle, curReg)
			curCycle++
			curReg += i
		}
	}

	drawableScreen := make([]string, 6)
	dark := " "
	light := "X"
	for i, v := range screen {
		row := i / 40
		pixel := dark
		if v {
			pixel = light
		}
		drawableScreen[row] += pixel
	}

	fmt.Println("")
	for _, line := range drawableScreen {
		fmt.Println(line)
	}
	fmt.Println("")
}

func paintPixel(screen []bool, cycle, registry int) []bool {
	screenIndex := cycle - 1
	horizontalReg := registry % 40
	horizontalScr := screenIndex % 40

	if dist := horizontalScr - horizontalReg; dist >= -1 && dist <= 1 {
		screen[screenIndex] = true
	}

	return screen
}

func main() {
	lines := utils.ReadInputStringLines(10, "\n")
	partA(lines)
	partB(lines)
}
