package day06

import (
	"aoc/utils"
	"strings"
)

func puzzle1(orbits []string) int {
	orbitMap := make(map[string]string)

	for _, orbit := range orbits {
		parts := strings.Split(orbit, ")")
		orbitee := parts[0]
		orbiter := parts[1]
		orbitMap[orbiter] = orbitee
	}

	countOrbits := func(orbiter string) int {
		count := 0
		for {
			if orbitee, ok := orbitMap[orbiter]; ok {
				count++
				orbiter = orbitee
				continue
			}
			break
		}
		return count
	}

	count := 0
	for orbiter := range orbitMap {
		count += countOrbits(orbiter)
	}
	return count
}

func Puzzle1() int {
	return puzzle1(utils.ReadLines("./input"))
}
