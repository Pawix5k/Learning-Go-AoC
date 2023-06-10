package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"os"
)

func partA(day int) {
	data, err := os.ReadFile(utils.GetFilePath(day))
	if err != nil {
		panic(err)
	}
	l, r := 0, 0
	counter := make(map[byte]int)
	uniqueN := 0
	for {
		if counter[data[r]] == 0 {
			uniqueN++
		}
		counter[data[r]] += 1
		r++

		if r-l > 4 {
			if counter[data[l]] == 1 {
				uniqueN--
			}
			counter[data[l]]--
			l++
		}

		if uniqueN == 4 {
			fmt.Println(r)
			break
		}
	}
}

func partB(day int) {
	data, err := os.ReadFile(utils.GetFilePath(day))
	if err != nil {
		panic(err)
	}
	l, r := 0, 0
	counter := make(map[byte]int)
	uniqueN := 0
	for {
		if counter[data[r]] == 0 {
			uniqueN++
		}
		counter[data[r]] += 1
		r++

		if r-l > 14 {
			if counter[data[l]] == 1 {
				uniqueN--
			}
			counter[data[l]]--
			l++
		}

		if uniqueN == 14 {
			fmt.Println(r)
			break
		}
	}
}

func main() {
	day := 6
	utils.DownloadInput(day)
	partA(day)
	partB(day)
}
