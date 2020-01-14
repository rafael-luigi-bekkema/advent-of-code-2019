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

func Puzzle1() int {
	lines := utils.ReadLines("./input")
	return puzzle1(lines[0], lines[1])
}
