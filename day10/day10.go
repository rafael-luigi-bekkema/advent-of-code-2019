package day10

import (
	"aoc/utils"
	"fmt"
	"math"
	"sort"
)

type Coord struct {
	x, y int
}

type Asteroids struct {
	data  []bool
	width int
}

func (a *Asteroids) isHit(c Coord) bool {
	idx := c.y*a.width + c.x
	if idx < 0 {
		fmt.Println(idx)
	}
	return a.data[idx]
}

func (a *Asteroids) coord(idx int) Coord {
	return Coord{idx % a.width, idx / a.width}
}

func (a *Asteroids) idx(c Coord) int {
	return c.y*a.width + c.x
}

// blocked determines if line of sight between two coords is blocked by another coord
func (a *Asteroids) blocked(begin, end Coord) bool {
	// If coords are on the same line
	// check that line for blocks
	if begin.y == end.y {
		if begin.x > end.x {
			begin, end = end, begin
		}
		for i := begin.x + 1; i < end.x; i++ {
			if a.isHit(Coord{i, begin.y}) {
				return true
			}
		}
		return false
	}

	if begin.y > end.y {
		begin, end = end, begin
	}

	// x movement for every y step
	diffx := float64(end.x-begin.x) / float64(end.y-begin.y)

	for curry := begin.y; curry <= end.y; curry++ {
		xpos := float64(begin.x) + float64(curry-begin.y)*diffx
		// Check if this row has a coordinate that falls on the line between begin and end
		if math.Floor(xpos) == xpos {
			newcoord := Coord{int(xpos), curry}
			if newcoord == end || newcoord == begin {
				continue
			}
			if a.isHit(newcoord) {
				return true
			}
		}
	}

	return false
}

func puzzle1(data []string) (Coord, int, *Asteroids) {
	width := len(data[0])
	height := len(data)
	roids := &Asteroids{
		data:  make([]bool, width*height),
		width: width,
	}
	sighted := make(map[int]int)

	// Build map
	for by, row := range data {
		for bx, val := range row {
			if val == '#' {
				roids.data[bx+(by*width)] = true
			}
		}
	}

	// For every asteroid
	for aid, a := range roids.data {
		if !a {
			continue
		}

		// Check all other a
		for subidx, suba := range roids.data {
			if !suba || subidx == aid {
				continue
			}

			if !roids.blocked(roids.coord(aid), roids.coord(subidx)) {
				sighted[aid]++
			}
		}
	}

	var maxid, maxcount int
	for idx, s := range sighted {
		if s > maxcount {
			maxid = idx
			maxcount = s
		}
	}

	_ = maxid

	return roids.coord(maxid), maxcount, roids
}

func puzzle2(data []string) int {
	station, _, roids := puzzle1(data)
	stationIdx := station.y*roids.width + station.x

	type ang struct {
		c     Coord
		angle float64
	}

	roidCount := 0
	for _, a := range roids.data {
		if a {
			roidCount++
		}
	}

	blasted := 0

	for {
		var angles []ang
		newGrid := make([]bool, len(roids.data))
		copy(newGrid, roids.data)
		for idx, a := range roids.data {
			if !a || idx == stationIdx {
				continue
			}

			c := roids.coord(idx)
			if !roids.blocked(station, c) {
				var opposite, adjacent, offset int
				if c.x >= station.x && station.y > c.y {
					opposite = c.x - station.x
					adjacent = station.y - c.y
				} else if c.x > station.x && station.y <= c.y {
					opposite = c.y - station.y
					adjacent = c.x - station.x
					offset = 90
				} else if c.x <= station.x && station.y < c.y {
					opposite = station.x - c.x
					adjacent = c.y - station.y
					offset = 180
				} else { // c.x < station.x && station.y >= c.y
					opposite = station.y - c.y
					adjacent = station.x - c.x
					offset = 270
				}
				rad := math.Atan2(float64(opposite), float64(adjacent))
				deg := rad * (180 / math.Pi)
				angle := float64(offset) + deg
				angles = append(angles, ang{c, angle})
			}
		}

		sort.Slice(angles, func(i, j int) bool {
			return angles[i].angle < angles[j].angle
		})

		for _, angle := range angles {
			newGrid[roids.idx(angle.c)] = false
			blasted++
			if blasted == 200 {
				//fmt.Printf("====> Pew %03d: vaporized %v\n", blasted, angle.c)
				return angle.c.x*100 + angle.c.y
			} else {
				//fmt.Printf("Pew %03d: vaporized %v\n", blasted, angle.c)
			}
		}

		if roidCount-blasted == 1 {
			break
		}

		roids.data = newGrid
	}

	panic("no result found")
}

func Puzzle1() (Coord, int) {
	data := utils.ReadLines("./input")
	c, num, _ := puzzle1(data)
	return c, num
}

func Puzzle2() int {
	data := utils.ReadLines("./input")
	return puzzle2(data)
}
