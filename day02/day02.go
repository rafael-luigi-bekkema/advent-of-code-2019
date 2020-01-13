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

func puzzle1(init []int, noun, verb int) int {
	pos := 0

	memory := make([]int, len(init))
	copy(memory, init)

	memory[1] = noun
	memory[2] = verb

For:
	for {
		switch op := memory[pos]; op {
		case 1:
			// Add
			posVal1 := memory[pos+1]
			posVal2 := memory[pos+2]
			resultPos := memory[pos+3]
			memory[resultPos] = memory[posVal1] + memory[posVal2]
		case 2:
			// Mul
			posVal1 := memory[pos+1]
			posVal2 := memory[pos+2]
			resultPos := memory[pos+3]
			memory[resultPos] = memory[posVal1] * memory[posVal2]
		case 99:
			// Done
			break For
		default:
			log.Panicf("Op code not recognized: %d", op)
		}
		pos += 4
	}
	return memory[0]
}

func puzzle2(data []int) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			res := puzzle1(data, noun, verb)
			if res == 19690720 {
				return 100*noun + verb
			}
		}
	}
	log.Panicf("no solution found")
	return 0
}

func Puzzle1() int {
	data := utils.ReadAll("./input")
	return puzzle1(parseInput(data), 12, 2)
}

func Puzzle2() int {
	data := utils.ReadAll("./input")
	return puzzle2(parseInput(data))
}
