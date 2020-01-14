package day03

import (
	"aoc/utils"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func puzzle1(line1, line2 string) int {
	coords := make(map[coord]bool)
	var x, y int
	for _, step := range strings.Split(line1, ",") {
		val, _ := strconv.Atoi(step[1:])

		for i := 0; i < val; i++ {
			switch step[0] {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}
			coords[coord{x, y}] = true
		}
	}

	var intersects []coord
	x = 0
	y = 0
	for _, step := range strings.Split(line2, ",") {
		val, _ := strconv.Atoi(step[1:])

		for i := 0; i < val; i++ {
			switch step[0] {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}

			coord := coord{x, y}
			if coords[coord] {
				intersects = append(intersects, coord)
			}
		}
	}

	distance := -1
	for _, coord := range intersects {
		d := int(math.Abs(float64(coord.x)) + math.Abs(float64(coord.y)))
		if distance == -1 || d < distance {
			distance = d
		}
	}

	return distance
}

func puzzle2(line1, line2 string) int {
	coords := make(map[coord]int)
	var x, y, stepCount int
	for _, step := range strings.Split(line1, ",") {
		val, _ := strconv.Atoi(step[1:])

		for i := 0; i < val; i++ {
			switch step[0] {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}
			stepCount++
			if coords[coord{x, y}] == 0 {
				coords[coord{x, y}] = stepCount
			}
		}
	}

	intersects := make(map[coord]int)
	x = 0
	y = 0
	stepCount = 0
	for _, step := range strings.Split(line2, ",") {
		val, _ := strconv.Atoi(step[1:])

		for i := 0; i < val; i++ {
			switch step[0] {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}

			stepCount++
			coord := coord{x, y}
			if sc := coords[coord]; sc > 0 && intersects[coord] == 0 {
				intersects[coord] = sc + stepCount
			}
		}
	}

	minSteps := -1
	for _, stepCount := range intersects {
		if minSteps == -1 || stepCount < minSteps {
			minSteps = stepCount
		}
	}

	return minSteps
}

func Puzzle1() int {
	lines := utils.ReadLines("./input")
	return puzzle1(lines[0], lines[1])
}

func Puzzle2() int {
	lines := utils.ReadLines("./input")
	return puzzle2(lines[0], lines[1])
}
