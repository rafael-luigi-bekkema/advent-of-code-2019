package day10

import (
	"aoc/utils"
	"fmt"
	"math"
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

// blocked determines if line of sight between two coords is blocked by another coord
func blocked(begin, end Coord, a *Asteroids) bool {
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

func puzzle1(data []string) (Coord, int) {
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

			if !blocked(roids.coord(aid), roids.coord(subidx), roids) {
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

	return roids.coord(maxid), maxcount
}

func Puzzle1() (Coord, int) {
	data := utils.ReadLines("./input")
	return puzzle1(data)
}
