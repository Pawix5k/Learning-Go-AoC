package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type State struct {
	inventory [4]int
	robots    [4]int
	blocks    [4]bool
}

func parseLine(line string) map[string]int {
	fields := strings.Fields(line)
	oreRobotOre, _ := strconv.Atoi(fields[6])
	clayRobotOre, _ := strconv.Atoi(fields[12])
	obsidianRobotOre, _ := strconv.Atoi(fields[18])
	obsidianRobotClay, _ := strconv.Atoi(fields[21])
	geodeRobotOre, _ := strconv.Atoi(fields[27])
	geodeRobotObsidian, _ := strconv.Atoi(fields[30])

	costs := map[string]int{
		"ore-ore":        oreRobotOre,
		"clay-ore":       clayRobotOre,
		"obsidian-ore":   obsidianRobotOre,
		"obsidian-clay":  obsidianRobotClay,
		"geode-ore":      geodeRobotOre,
		"geode-obsidian": geodeRobotObsidian,
	}

	return costs
}

func getPossibleStates(inventory, robots [4]int, blocks [4]bool, prices map[string]int) []State {
	possStates := []State{}
	newInventory := inventory
	newRobots := robots

	couldBuyGeodeRobot := inventory[0] >= prices["geode-ore"] && inventory[2] >= prices["geode-obsidian"]
	couldBuyObsidianRobot := inventory[0] >= prices["obsidian-ore"] && inventory[1] >= prices["obsidian-clay"]
	couldBuyClayRobot := inventory[0] >= prices["clay-ore"]
	couldBuyOreRobot := inventory[0] >= prices["ore-ore"]

	if !blocks[3] && couldBuyGeodeRobot {
		newRobots = robots
		newRobots[3]++
		newInventory = inventory
		newInventory[0] -= prices["geode-ore"]
		newInventory[2] -= prices["geode-obsidian"]
		for i := range inventory {
			newInventory[i] += robots[i]
		}
		possStates = append(possStates, State{newInventory, newRobots, [4]bool{}})
	}
	if !blocks[2] && couldBuyObsidianRobot {
		newRobots = robots
		newRobots[2]++
		newInventory = inventory
		newInventory[0] -= prices["obsidian-ore"]
		newInventory[1] -= prices["obsidian-clay"]
		for i := range inventory {
			newInventory[i] += robots[i]
		}
		possStates = append(possStates, State{newInventory, newRobots, [4]bool{}})
	}
	if !blocks[1] && couldBuyClayRobot {
		newRobots = robots
		newRobots[1]++
		newInventory = inventory
		newInventory[0] -= prices["clay-ore"]
		for i := range inventory {
			newInventory[i] += robots[i]
		}
		possStates = append(possStates, State{newInventory, newRobots, [4]bool{}})
	}
	if !blocks[0] && couldBuyOreRobot {
		newRobots = robots
		newRobots[0]++
		newInventory = inventory
		newInventory[0] -= prices["ore-ore"]
		for i := range inventory {
			newInventory[i] += robots[i]
		}
		possStates = append(possStates, State{newInventory, newRobots, [4]bool{}})
	}

	newInventory = inventory
	for i := range inventory {
		newInventory[i] += robots[i]
	}
	newRobots = robots
	possStates = append(possStates, State{newInventory, newRobots, [4]bool{couldBuyOreRobot, couldBuyClayRobot, couldBuyObsidianRobot, couldBuyGeodeRobot}})

	return possStates
}

func getMaxGeodes(stepsLeft int, inventory, robots [4]int, blocks [4]bool, prices map[string]int, mem map[[2][4]int]int) int {
	curMax := 0
	var getMaxGeodesInBranch func(int, [4]int, [4]int, [4]bool, map[string]int, map[[2][4]int]int) int
	getMaxGeodesInBranch = func(stepsLeft int, inventory, robots [4]int, blocks [4]bool, prices map[string]int, mem map[[2][4]int]int) int {
		if stepsLeft == 2 {
			sum := inventory[3] + 2*robots[3]
			if inventory[0] >= prices["geode-ore"] && inventory[2] >= prices["geode-obsidian"] {
				sum++
			}
			if sum > curMax {
				curMax = sum
			}
			return sum
		}
		optimisticMax := inventory[3] + (1+stepsLeft)/2*stepsLeft + robots[3]*stepsLeft
		if optimisticMax <= curMax {
			return 0
		}
		if thenStepsLeft, ok := mem[[2][4]int{inventory, robots}]; ok {
			if thenStepsLeft >= stepsLeft {
				return 0
			}
		}
		maxGeodes := 0
		for _, newState := range getPossibleStates(inventory, robots, blocks, prices) {
			possMax := getMaxGeodesInBranch(stepsLeft-1, newState.inventory, newState.robots, newState.blocks, prices, mem)

			if possMax > maxGeodes {
				maxGeodes = possMax
			}
		}
		mem[[2][4]int{inventory, robots}] = stepsLeft
		return maxGeodes
	}

	getMaxGeodesInBranch(stepsLeft, inventory, robots, blocks, prices, mem)

	return curMax
}

func partA(lines []string) {
	maxTotal := 0
	for i, line := range lines {
		prices := parseLine(line)
		stepsLeft := 24
		mem := map[[2][4]int]int{}

		max := getMaxGeodes(stepsLeft, [4]int{0, 0, 0, 0}, [4]int{1, 0, 0, 0}, [4]bool{}, prices, mem)
		maxTotal += (i + 1) * max
	}
	fmt.Println(maxTotal)
}

func partB(lines []string) {

	maxTotal := 1
	for _, line := range lines[:3] {
		prices := parseLine(line)
		stepsLeft := 32
		mem := map[[2][4]int]int{}

		max := getMaxGeodes(stepsLeft, [4]int{0, 0, 0, 0}, [4]int{1, 0, 0, 0}, [4]bool{}, prices, mem)
		maxTotal *= max
	}
	fmt.Println(maxTotal)
}

func main() {
	lines := utils.ReadInputStringLines(19, "\n")
	partA(lines)
	partB(lines)
}
