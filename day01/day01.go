package day01

import (
	"aoc/utils"
	"log"
	"math"
	"strconv"
)

func calcFuel(mass int) int {
	val := int(math.Floor(float64(mass)/3) - 2)
	if val < 0 {
		return 0
	}
	return val
}

func puzzle1(lines []string) int {
	totalFuel := 0
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Panicf("Not a number: %s", line)
		}

		fuel := calcFuel(mass)
		totalFuel += fuel
	}

	return totalFuel
}

func puzzle2(lines []string) int {
	totalFuel := 0
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Panicf("Not a number: %s", line)
		}

		fuel := calcFuel(mass)
		totalFuel += fuel

		for {
			fuel = calcFuel(fuel)
			if fuel == 0 {
				break
			}
			totalFuel += fuel
		}
	}

	return totalFuel
}

func Puzzle1() int {
	data := utils.ReadLines("./input")
	return puzzle1(data)
}

func Puzzle2() int {
	data := utils.ReadLines("./input")
	return puzzle2(data)
}
