package main

import (
	"fmt"
	"os"
	"strings"
)

func partA() {
	data, err := os.ReadFile("day03/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\r\n")
	sum := 0
	for _, line := range lines {
		l := len(line)
		comp := make(map[byte]bool)
		for i := 0; i < l/2; i++ {
			comp[line[i]] = true
		}
		for i := l / 2; i < l; i++ {
			_, ok := comp[line[i]]
			if ok {
				if line[i] >= 'a' && line[i] <= 'z' {
					sum += int(line[i]-'a') + 1
				}
				if line[i] >= 'A' && line[i] <= 'Z' {
					sum += int(line[i]-'A') + 27
				}
				break
			}
		}
	}
	fmt.Println(sum)
}

func partB() {
	data, err := os.ReadFile("day03/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\r\n")
	sum := 0
	i := 0
	for i < len(lines) {
		items := make(map[byte]bool)
		for j := 0; j < len(lines[i]); j++ {
			items[lines[i][j]] = false
		}
		for j := 0; j < len(lines[i+1]); j++ {
			if _, ok := items[lines[i+1][j]]; ok {
				items[lines[i+1][j]] = true
			}
		}
		for j := 0; j < len(lines[i+2]); j++ {
			if res, _ := items[lines[i+2][j]]; res {
				b := lines[i+2][j]
				if b >= 'a' && b <= 'z' {
					sum += int(b-'a') + 1
				}
				if b >= 'A' && b <= 'Z' {
					sum += int(b-'A') + 27
				}
				break
			}
		}
		i += 3
	}
	fmt.Println(sum)
}

func main() {
	partA()
	partB()
}
