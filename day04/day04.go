package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partA() {
	data, err := os.ReadFile("day04/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\r\n")
	i := 0
	n := 0
	for _, line := range lines {
		rangesStr := strings.Split(line, ",")
		r1Str := strings.Split(rangesStr[0], "-")
		start1, _ := strconv.Atoi(r1Str[0])
		end1, _ := strconv.Atoi(r1Str[1])

		r2Str := strings.Split(rangesStr[1], "-")
		start2, _ := strconv.Atoi(r2Str[0])
		end2, _ := strconv.Atoi(r2Str[1])

		if start1 >= start2 && end1 <= end2 || start1 <= start2 && end1 >= end2 {
			n++
		}

		i += 2
	}
	fmt.Println(n)
}

func partB() {
	data, err := os.ReadFile("day04/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\r\n")
	i := 0
	n := 0
	for _, line := range lines {
		rangesStr := strings.Split(line, ",")
		r1Str := strings.Split(rangesStr[0], "-")
		start1, _ := strconv.Atoi(r1Str[0])
		end1, _ := strconv.Atoi(r1Str[1])

		r2Str := strings.Split(rangesStr[1], "-")
		start2, _ := strconv.Atoi(r2Str[0])
		end2, _ := strconv.Atoi(r2Str[1])

		if start1 >= start2 && end1 <= end2 ||
			start1 <= start2 && end1 >= end2 ||
			start1 >= start2 && start1 <= end2 ||
			end1 >= start2 && end1 < end2 {
			n++
		}

		i += 2
	}
	fmt.Println(n)
}

func main() {
	partA()
	partB()
}
