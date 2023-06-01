package main

import (
	"fmt"
	"os"
	"strings"
)

var symbolToInt = map[string]int{
	"X": 0,
	"Y": 1,
	"Z": 2,
	"A": 0,
	"B": 1,
	"C": 2,
}

var resultPoints = map[int]int{
	0: 3,
	1: 6,
	2: 0,
}

var bonusPoints = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var intToSymbol = map[int]string{
	0: "X",
	1: "Y",
	2: "Z",
}

var offset = map[string]int{
	"X": 2,
	"Y": 0,
	"Z": 1,
}

func unpackLine(line string) (string, string) {
	symbols := strings.Fields(line)
	return symbols[0], symbols[1]
}

func getPoints(opponent, player string) int {
	resultToInt := (symbolToInt[player] - symbolToInt[opponent]) % 3
	if resultToInt < 0 {
		resultToInt += 3
	}
	points := resultPoints[resultToInt] + bonusPoints[player]
	return points
}

func partA() {
	data, err := os.ReadFile("day02/input.txt")
	if err != nil {
		panic(err)
	}
	dataSplit := strings.Split(string(data), "\r\n")
	sumPoints := 0
	for _, line := range dataSplit {
		sumPoints += getPoints(unpackLine(line))
	}

	fmt.Println(sumPoints)
}

func partB() {
	data, err := os.ReadFile("day02/input.txt")
	if err != nil {
		panic(err)
	}
	dataSplit := strings.Split(string(data), "\r\n")
	sumPoints := 0
	for _, line := range dataSplit {
		opponent, result := unpackLine(line)
		player := intToSymbol[(symbolToInt[opponent]+offset[result])%3]
		sumPoints += getPoints(opponent, player)
	}

	fmt.Println(sumPoints)
}

func main() {
	partA()
	partB()
}
