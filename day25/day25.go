package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
)

var SnafuToInt = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'-': -1,
	'=': -2,
}

var IntToSnafu = map[int]byte{
	0:  '0',
	1:  '1',
	2:  '2',
	-1: '-',
	-2: '=',
}

func parseSNAFU(chs []byte) int {
	num := 0
	for i := range chs {
		num = num * 5
		num += SnafuToInt[chs[i]]
	}

	return num
}

func convertToSnafu(num int) string {
	multiplier := 1
	for num >= multiplier*3 {
		multiplier = multiplier * 5
	}

	s := ""

	for multiplier > 1 {
		x := (num+2*multiplier+multiplier/2)/multiplier - 2
		ch := IntToSnafu[x]
		s = s + string(ch)
		num -= x * multiplier
		multiplier = multiplier / 5
	}
	s = s + string(IntToSnafu[num/multiplier])

	return s
}

func partA(lines [][]byte) {
	sum := 0
	for _, line := range lines {
		num := parseSNAFU(line)
		sum += num
	}
	fmt.Println(convertToSnafu(sum))
}

func main() {
	lines := utils.ReadInputByteLines(25, "\n")
	partA(lines)
}
