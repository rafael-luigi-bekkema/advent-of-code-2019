package day14

import (
	"aoc/utils"
	"fmt"
	"math"
	"strings"
)

type res struct {
	name string
	qty  uint
}

func fromString(s string) res {
	var r res
	_, _ = fmt.Sscanf(s, "%d %s", &r.qty, &r.name)
	return r
}

type out struct {
	qty       uint
	rs        []res
	leftovers uint
}

func oreCalc(name string, qty uint, rcns map[string]*out) uint {
	r := rcns[name]
	var count uint

	// Take from leftovers
	if r.leftovers > 0 {
		if qty > r.leftovers {
			qty -= r.leftovers
			r.leftovers = 0
		} else {
			r.leftovers -= qty
			return count
		}
	}

	var mul uint = 1
	if qty > r.qty {
		mul = uint(math.Ceil(float64(qty) / float64(r.qty)))
	}

	r.leftovers += mul*r.qty - qty

	for _, inpr := range r.rs {
		if inpr.name == "ORE" {
			count += inpr.qty * mul
		} else {
			count += oreCalc(inpr.name, inpr.qty*mul, rcns)
		}
	}
	return count
}

func puzzle1(data []string) uint {
	rcns := make(map[string]*out)

	// Parse data into map
	for _, line := range data {
		io := strings.Split(line, " => ")
		rs := strings.Split(io[0], ", ")

		outRes := fromString(io[1])
		o := out{outRes.qty, make([]res, len(rs)), 0}

		for idx, resource := range rs {
			o.rs[idx] = fromString(resource)
		}

		rcns[outRes.name] = &o
	}

	count := oreCalc("FUEL", 1, rcns)
	return count
}

func Puzzle1() uint {
	data := utils.ReadLines("./input")
	return puzzle1(data)
}
