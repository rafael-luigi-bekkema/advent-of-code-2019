package day12

import (
	"aoc/utils"
	"fmt"
	"math"
)

type Coord struct {
	x, y, z int
}

func (c Coord) String() string {
	return fmt.Sprintf("x=% 3d, y=% 3d, z=% 3d", c.x, c.y, c.z)
}

type Moon struct {
	pos, vel Coord
}

func (moon *Moon) String() string {
	return fmt.Sprintf("pos: %v, vel: %v", moon.pos, moon.vel)
}

func (moon *Moon) adjustVel(submoon *Moon) {
	moon.vel.x += diff(moon.pos.x, submoon.pos.x)
	moon.vel.y += diff(moon.pos.y, submoon.pos.y)
	moon.vel.z += diff(moon.pos.z, submoon.pos.z)
}

func (moon *Moon) adjustPos() {
	moon.pos.x += moon.vel.x
	moon.pos.y += moon.vel.y
	moon.pos.z += moon.vel.z
}

func (moon *Moon) totalEnergy() int {
	total := (math.Abs(float64(moon.pos.x)) + math.Abs(float64(moon.pos.y)) + math.Abs(float64(moon.pos.z))) *
		(math.Abs(float64(moon.vel.x)) + math.Abs(float64(moon.vel.y)) + math.Abs(float64(moon.vel.z)))
	return int(total)
}

func diff(a, b int) int {
	if b > a {
		return 1
	} else if b < a {
		return -1
	}
	return 0
}

func puzzle1(data []string, steps int) int {
	moons := make([]*Moon, len(data))
	for idx, item := range data {
		var moon Moon
		_, _ = fmt.Sscanf(item, "<x=%d, y=%d, z=%d>", &moon.pos.x, &moon.pos.y, &moon.pos.z)
		moons[idx] = &moon
	}

	for i := 0; i < steps; i++ {
		// Adjust velocities
		for _, moon := range moons {
			for _, submoon := range moons {
				moon.adjustVel(submoon)
			}
		}

		// Adjust positions
		for _, moon := range moons {
			moon.adjustPos()
		}
	}

	var total int
	for _, moon := range moons {
		total += moon.totalEnergy()
	}
	return total
}

func Puzzle1() int {
	data := utils.ReadLines("./input")
	return puzzle1(data, 1000)
}
