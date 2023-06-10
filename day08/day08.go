package main

import (
	"bytes"
	"example/pawix5k/aoc/utils"
	"fmt"
	"os"
)

const day = 8

var inputPath = utils.GetFilePath(day)

func partA() {
	utils.DownloadInput(day)
	data, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	length := len(lines) - 2

	visibility := make([][]bool, length)
	for i := 0; i < length; i++ {
		visibility[i] = make([]bool, length)
	}

	for i := 0; i < length; i++ {
		lowN := lines[0][i+1]
		lowS := lines[length+1][i+1]
		lowW := lines[i+1][0]
		lowE := lines[i+1][length+1]

		for j := 0; j < length; j++ {
			if lines[j+1][i+1] > lowN {
				lowN = lines[j+1][i+1]
				visibility[j][i] = true
			}
			if lines[length-j][i+1] > lowS {
				lowS = lines[length-j][i+1]
				visibility[length-j-1][i] = true
			}
			if lines[i+1][j+1] > lowW {
				lowW = lines[i+1][j+1]
				visibility[i][j] = true
			}
			if lines[i+1][length-j] > lowE {
				lowE = lines[i+1][length-j]
				visibility[i][length-j-1] = true
			}
		}
	}

	count := 4*length + 4
	for _, i := range visibility {
		for _, j := range i {
			if j {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	partA()
}
