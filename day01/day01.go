package day01

import (
	"aoc/utils"
	"log"
	"math"
	"strconv"
)

func puzzle1(lines []string) int {
	totalFuel := 0
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Panicf("Not a number: %s", line)
		}

		fuel := int(math.Floor(float64(mass)/3) - 2)
		totalFuel += fuel
	}

	return totalFuel
}

func Puzzle1() int {
	data := utils.ReadLines("./input")
	return puzzle1(data)
}
