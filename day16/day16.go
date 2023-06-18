package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Vertice struct {
	name     string
	flow     int
	edgesStr []string
	edges    []*Vertice
}

func parseLine(line string) *Vertice {
	splits := strings.SplitN(line, " ", 10)
	name := splits[1]
	flow, _ := strconv.Atoi(splits[4][5 : len(splits[4])-1])
	edges := strings.Split(splits[len(splits)-1], ", ")

	return &Vertice{name: name, flow: flow, edgesStr: edges}
}

func getGraphData(lines []string) ([][]int, map[string]int, map[string]int) {
	valves := map[string]*Vertice{}
	for _, line := range lines {
		vertice := parseLine(line)
		valves[vertice.name] = vertice
	}

	valvesOfInterest := []*Vertice{}

	for _, valve := range valves {
		for _, edge := range valve.edgesStr {
			valve.edges = append(valve.edges, valves[edge])
		}
		if valve.name == "AA" || valve.flow > 0 {
			valvesOfInterest = append(valvesOfInterest, valve)
		}
	}

	for i, valve := range valvesOfInterest {
		if valve.name == "AA" && i != 0 {
			valvesOfInterest[0], valvesOfInterest[i] = valvesOfInterest[i], valvesOfInterest[0]
			break
		}
	}
	valvesIndexes := map[string]int{}
	valveFlows := map[string]int{}
	for i, valve := range valvesOfInterest {
		valvesIndexes[valve.name] = i
		valveFlows[valve.name] = valve.flow
	}

	adjacencyMatrix := make([][]int, len(valvesOfInterest))
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, len(valvesOfInterest))
	}

	for _, valve := range valvesOfInterest {
		depth := 0
		visited := map[string]bool{valve.name: true}
		queue := []*Vertice{valve}
		for len(queue) > 0 {
			newQueue := []*Vertice{}
			for _, curValve := range queue {
				if curValve.name == "AA" || curValve.flow > 0 {
					i1 := valvesIndexes[valve.name]
					i2 := valvesIndexes[curValve.name]
					adjacencyMatrix[i1][i2] = depth
				}
				for _, edge := range curValve.edges {
					if _, ok := visited[edge.name]; !ok {
						newQueue = append(newQueue, edge)
						visited[edge.name] = true
					}
				}
			}
			queue = newQueue
			depth++
		}
	}

	return adjacencyMatrix, valvesIndexes, valveFlows
}

func findMaxPressure(visited, toVisit []string, limit int, adjacencyMatrix [][]int, valveIndexes, valveFlows map[string]int) int {
	maxResult := 0

	var findMaxPressureForBranch func([]string, []string, int, int)
	findMaxPressureForBranch = func(visited, toVisit []string, curResult int, curTurn int) {
		if curTurn > limit || len(toVisit) == 0 {
			if curResult > maxResult {
				maxResult = curResult
			}
			return
		}
		for i, next := range toVisit {
			i1 := valveIndexes[visited[len(visited)-1]]
			i2 := valveIndexes[next]
			dist := adjacencyMatrix[i1][i2]
			possible := curTurn+dist+1 < limit
			if !possible {
				if curResult > maxResult {
					maxResult = curResult
				}
				continue
			}
			visitedCopy := make([]string, len(visited))
			copy(visitedCopy, visited)
			visitedCopy = append(visitedCopy, next)

			toVisitCopy := []string{}
			toVisitCopy = append(toVisitCopy, toVisit[:i]...)
			toVisitCopy = append(toVisitCopy, toVisit[i+1:]...)

			profit := (limit - curTurn - dist - 1) * valveFlows[next]

			findMaxPressureForBranch(visitedCopy, toVisitCopy, curResult+profit, curTurn+dist+1)
		}
	}

	findMaxPressureForBranch(visited, toVisit, 0, 0)

	return maxResult
}

func partA(lines []string) {
	adjacencyMatrix, valveIndexes, valveFlows := getGraphData(lines)

	toVisit := []string{}
	for valveName := range valveFlows {
		if valveName != "AA" {
			toVisit = append(toVisit, valveName)
		}
	}

	result := findMaxPressure([]string{"AA"}, toVisit, 30, adjacencyMatrix, valveIndexes, valveFlows)
	fmt.Println(result)
}

func partB(lines []string) {
	adjacencyMatrix, valveIndexes, valveFlows := getGraphData(lines)

	valvesWithPosFlow := []string{}
	for valveName := range valveFlows {
		if valveName != "AA" {
			valvesWithPosFlow = append(valvesWithPosFlow, valveName)
		}
	}

	combinations := utils.IntPow(2, len(valveFlows)-2)
	i := 0

	result := 0

	for i < combinations {
		a, b := []string{}, []string{}
		num := i
		j := 0
		for num > 0 {
			nextNum := num / 2
			bit := num % 2
			if bit == 1 {
				a = append(a, valvesWithPosFlow[j])
			} else {
				b = append(b, valvesWithPosFlow[j])
			}
			j++
			num = nextNum
		}
		if j < len(valvesWithPosFlow) {
			b = append(b, valvesWithPosFlow[j:]...)
		}
		res1 := findMaxPressure([]string{"AA"}, a, 26, adjacencyMatrix, valveIndexes, valveFlows)
		res2 := findMaxPressure([]string{"AA"}, b, 26, adjacencyMatrix, valveIndexes, valveFlows)
		if res1+res2 > result {
			result = res1 + res2
		}
		i++
	}

	fmt.Println(result)
}

func main() {
	lines := utils.ReadInputStringLines(16, "\n")
	partA(lines)
	partB(lines)
}
