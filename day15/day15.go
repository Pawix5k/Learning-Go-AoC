package main

import (
	"example/pawix5k/aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Result int

const (
	Merged Result = iota
	Smaller
	Larger
)

type Interval struct {
	isEmpty bool
	min     int
	max     int
}

type IntervalsSet struct {
	vals []Interval
}

func (intset *IntervalsSet) checkIfContains(x int) bool {
	for _, interval := range intset.vals {
		if x >= interval.min && x <= interval.max {
			return true
		} else if x > interval.max {
			return false
		}
	}
	return false
}

func (intset *IntervalsSet) getSum() int {
	sum := 0
	for _, interval := range intset.vals {
		width := interval.max - interval.min + 1
		sum += width
	}
	return sum
}

func (intset *IntervalsSet) Add(newInterval Interval) {
	if len(intset.vals) == 0 {
		intset.vals = append(intset.vals, newInterval)
		return
	}
	for i, interval := range intset.vals {
		resolvedInterval, result := resolveIntervals(newInterval, interval)
		if result == Smaller {
			intset.vals = insert(intset.vals, newInterval, i)
			break
		} else if result == Larger && i == len(intset.vals)-1 {
			intset.vals = append(intset.vals, newInterval)
			break
		} else if result == Merged && i == len(intset.vals)-1 {
			intset.vals = intset.vals[:i]
			intset.Add(resolvedInterval)
			break
		} else if result == Merged {
			intset.vals = append(intset.vals[:i], intset.vals[i+1:]...)
			intset.Add(resolvedInterval)
			break
		}
	}
}

type Point struct {
	x int
	y int
}

type Sensor struct {
	pos       Point
	beaconPos Point
}

func insert(a []Interval, interval Interval, index int) []Interval {
	a = append(a[:index+1], a[index:]...)
	a[index] = interval
	return a
}

func resolveIntervals(int1, int2 Interval) (Interval, Result) {
	int2Contains := int1.min >= int2.min && int1.max <= int2.max
	if int2Contains {
		return int2, Merged
	}
	int1Contains := int2.min >= int1.min && int2.max <= int1.max
	if int1Contains {
		return int1, Merged
	}
	int1Smaller := int1.max < int2.min-1
	if int1Smaller {
		return int1, Smaller
	}
	int2Smaller := int2.max < int1.min-1
	if int2Smaller {
		return int2, Larger
	}

	min := int1.min
	if int2.min < min {
		min = int2.min
	}
	max := int1.max
	if int2.max > max {
		max = int2.max
	}
	return Interval{false, min, max}, Merged
}

func getIntervalAtY(sensor Sensor, y int) Interval {
	sensorDistance := utils.Abs(sensor.pos.x-sensor.beaconPos.x) + utils.Abs(sensor.pos.y-sensor.beaconPos.y)
	distanceFromLine := utils.Abs(y - sensor.pos.y)

	if distanceFromLine > sensorDistance {
		return Interval{isEmpty: true}
	}
	dx := utils.Abs(distanceFromLine - sensorDistance)

	return Interval{false, sensor.pos.x - dx, sensor.pos.x + dx}
}

func getTruncatedIntervalAtY(sensor Sensor, y int, boundingBox Point) Interval {
	sensorDistance := utils.Abs(sensor.pos.x-sensor.beaconPos.x) + utils.Abs(sensor.pos.y-sensor.beaconPos.y)
	distanceFromLine := utils.Abs(y - sensor.pos.y)

	if distanceFromLine > sensorDistance {
		return Interval{isEmpty: true}
	}
	dx := utils.Abs(distanceFromLine - sensorDistance)

	min := sensor.pos.x - dx
	if boundingBox.x > min {
		min = boundingBox.x
	}
	max := sensor.pos.x + dx
	if boundingBox.y < max {
		max = boundingBox.y
	}
	if min > max {
		return Interval{isEmpty: true}
	}
	return Interval{false, sensor.pos.x - dx, sensor.pos.x + dx}
}

func parseLine(line string) Sensor {
	fields := strings.Fields(line)
	sensorStrX := fields[2][2 : len(fields[2])-1]
	sensorStrY := fields[3][2 : len(fields[3])-1]
	beaconStrX := fields[8][2 : len(fields[8])-1]
	beaconStrY := fields[9][2:len(fields[9])]
	sensorX, _ := strconv.Atoi(sensorStrX)
	sensorY, _ := strconv.Atoi(sensorStrY)
	beaconX, _ := strconv.Atoi(beaconStrX)
	beaconY, _ := strconv.Atoi(beaconStrY)

	return Sensor{Point{sensorX, sensorY}, Point{beaconX, beaconY}}
}

func partA(lines []string) {
	y := 2000000
	sensors := make([]Sensor, len(lines))
	beaconsAtY := map[Point]bool{}

	for i, line := range lines {
		sensor := parseLine(line)
		if sensor.beaconPos.y == y {
			beaconsAtY[sensor.beaconPos] = true
		}
		sensors[i] = sensor
	}

	intset := IntervalsSet{}

	for _, sensor := range sensors {
		interval := getIntervalAtY(sensor, y)
		if !interval.isEmpty {
			intset.Add(interval)
		}
	}

	sum := intset.getSum()
	for k := range beaconsAtY {
		if intset.checkIfContains(k.x) {
			sum--
		}
	}

	fmt.Println(sum)
}

func partB(lines []string) {
	boundingBox := Point{0, 4_000_000}
	sensors := make([]Sensor, len(lines))

	for i, line := range lines {
		sensor := parseLine(line)
		sensors[i] = sensor
	}

	var x, y int

	for height := boundingBox.x; height <= boundingBox.y; height++ {
		intset := IntervalsSet{}

		for _, sensor := range sensors {
			interval := getTruncatedIntervalAtY(sensor, height, boundingBox)
			if !interval.isEmpty {
				intset.Add(interval)
			}
		}

		if len(intset.vals) != 1 {
			x = intset.vals[0].max + 1
			y = height
			break
		}
	}

	fmt.Println(x*4_000_000 + y)
}

func main() {
	lines := utils.ReadInputStringLines(15, "\n")
	partA(lines)
	partB(lines)
}
