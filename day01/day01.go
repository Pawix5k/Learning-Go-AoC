package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partA() {
	data, err := os.ReadFile("day01/input.txt")
	if err != nil {
		panic(err)
	}

	dataSplit := strings.Split(string(data), "\r\n\r\n")

	max := 0
	for _, elfData := range dataSplit {
		elfCalories := 0
		for _, calories := range strings.Split(elfData, "\r\n") {
			caloriesInt, _ := strconv.Atoi(calories)
			elfCalories += caloriesInt
		}
		if elfCalories > max {
			max = elfCalories
		}
	}

	fmt.Println(max)
}

func pushToArray(slice []int, number int) {
	if number > slice[0] {
		slice[0] = number
	}
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			slice[i], slice[i-1] = slice[i-1], slice[i]
		}
	}
}

func partB() {
	data, err := os.ReadFile("day01/input.txt")
	if err != nil {
		panic(err)
	}

	dataSplit := strings.Split(string(data), "\r\n\r\n")

	top3 := []int{0, 0, 0}
	for _, elfData := range dataSplit {
		elfCalories := 0
		for _, calories := range strings.Split(elfData, "\r\n") {
			caloriesInt, _ := strconv.Atoi(calories)
			elfCalories += caloriesInt
		}
		pushToArray(top3, elfCalories)
	}

	max := 0
	for _, v := range top3 {
		max += v
	}

	fmt.Println(max)
}

func main() {
	partA()
	partB()
}
