package day02

import (
	"aoc/utils"
	"log"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	vals := strings.Split(input, ",")
	res := make([]int, len(vals))
	for idx, strval := range vals {
		num, err := strconv.Atoi(strval)
		if err != nil {
			panic(err)
		}
		res[idx] = num
	}
	return res
}

func puzzle1(program []int) int {
	pos := 0

For:
	for {
		switch op := program[pos]; op {
		case 1:
			// Add
			posVal1 := program[pos+1]
			posVal2 := program[pos+2]
			resultPos := program[pos+3]
			program[resultPos] = program[posVal1] + program[posVal2]
		case 2:
			// Mul
			posVal1 := program[pos+1]
			posVal2 := program[pos+2]
			resultPos := program[pos+3]
			program[resultPos] = program[posVal1] * program[posVal2]
		case 99:
			// Done
			break For
		default:
			log.Panicf("Op code not recognized: %d", op)
		}
		pos += 4
	}
	return program[0]
}

func Puzzle1() int {
	data := utils.ReadAll("./input")
	program := parseInput(data)
	program[1] = 12
	program[2] = 2

	return puzzle1(program)
}
