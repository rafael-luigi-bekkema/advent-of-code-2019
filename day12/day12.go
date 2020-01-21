package day12

import (
	"aoc/utils"
	"fmt"
	"math"
)

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

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

type History struct {
	pos, vel [4]int
}

func history(moons []*Moon) (History, History, History) {
	return History{[4]int{moons[0].pos.x, moons[1].pos.x, moons[2].pos.x, moons[3].pos.x},
			[4]int{moons[0].vel.x, moons[1].vel.x, moons[2].vel.x, moons[3].vel.x}},
		History{[4]int{moons[0].pos.y, moons[1].pos.y, moons[2].pos.y, moons[3].pos.y},
			[4]int{moons[0].vel.y, moons[1].vel.y, moons[2].vel.y, moons[3].vel.y}},
		History{[4]int{moons[0].pos.z, moons[1].pos.z, moons[2].pos.z, moons[3].pos.z},
			[4]int{moons[0].vel.z, moons[1].vel.z, moons[2].vel.z, moons[3].vel.z}}
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

func puzzle2(data []string) int {
	moons := make([]*Moon, len(data))
	for idx, item := range data {
		var moon Moon
		_, _ = fmt.Sscanf(item, "<x=%d, y=%d, z=%d>", &moon.pos.x, &moon.pos.y, &moon.pos.z)
		moons[idx] = &moon
	}

	oxh, oyh, ozh := history(moons)

	var count int
	var gotx, goty, gotz bool
	var loopx, loopy, loopz int
	for {
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

		count++
		xh, yh, zh := history(moons)
		if xh == oxh && !gotx {
			loopx = count
			gotx = true
		}
		if yh == oyh && !goty {
			loopy = count
			goty = true
		}
		if zh == ozh && !gotz {
			loopz = count
			gotz = true
		}

		if gotx && goty && gotz {
			break
		}
	}

	return lcm(loopx, loopy, loopz)
}

func Puzzle1() int {
	data := utils.ReadLines("./input")
	return puzzle1(data, 1000)
}

func Puzzle2() int {
	data := utils.ReadLines("./input")
	return puzzle2(data)
}
