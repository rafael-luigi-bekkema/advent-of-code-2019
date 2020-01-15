package day05

import (
	"aoc/utils"
	"fmt"
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

func puzzle1(init []int, inputs []int) int {
	pos := 0

	memory := make([]int, len(init))
	copy(memory, init)

For:
	for {
		value := strconv.Itoa(memory[pos])
		startFrom := 0
		if len(value) > 2 {
			startFrom = len(value) - 2
		}
		opcode, _ := strconv.Atoi(value[startFrom:])

		getVal := func(offset int) int {
			immediateMode := false
			if len(value) >= offset+2 {
				idx := len(value) - (offset + 2)
				mode := value[idx]
				immediateMode = mode == '1'
			}

			if immediateMode {
				return memory[pos+offset]
			} else {
				return memory[memory[pos+offset]]
			}
		}

		switch opcode {
		case 1:
			// Add
			memory[memory[pos+3]] = getVal(1) + getVal(2)
			pos += 4
		case 2:
			// Mul
			memory[memory[pos+3]] = getVal(1) * getVal(2)
			pos += 4
		case 3:
			//Input
			inpVal := inputs[0]
			memory[memory[pos+1]] = inpVal
			inputs = inputs[1:]
			pos += 2
		case 4:
			//Output
			fmt.Println(getVal(1))
			pos += 2
		case 99:
			// Done
			break For
		default:
			log.Panicf("Op code not recognized: %d", opcode)
		}
	}

	return 0
}

func Puzzle1() int {
	data := utils.ReadAll("./input")
	return puzzle1(parseInput(data), []int{1})
}

func puzzle2(init []int, inputs []int) []int {
	pos := 0

	memory := make([]int, len(init))
	copy(memory, init)

	result := make([]int, 0, 1)

For:
	for {
		value := strconv.Itoa(memory[pos])
		startFrom := 0
		if len(value) > 2 {
			startFrom = len(value) - 2
		}
		opcode, _ := strconv.Atoi(value[startFrom:])

		getVal := func(offset int) int {
			immediateMode := false
			if len(value) >= offset+2 {
				idx := len(value) - (offset + 2)
				mode := value[idx]
				immediateMode = mode == '1'
			}

			if immediateMode {
				return memory[pos+offset]
			} else {
				return memory[memory[pos+offset]]
			}
		}

		switch opcode {
		case 1:
			// Add
			memory[memory[pos+3]] = getVal(1) + getVal(2)
			pos += 4
		case 2:
			// Mul
			memory[memory[pos+3]] = getVal(1) * getVal(2)
			pos += 4
		case 3:
			//Input
			inpVal := inputs[0]
			memory[memory[pos+1]] = inpVal
			inputs = inputs[1:]
			pos += 2
		case 4:
			//Output
			result = append(result, getVal(1))
			pos += 2
		case 5:
			if getVal(1) != 0 {
				pos = getVal(2)
			} else {
				pos += 3
			}
		case 6:
			if getVal(1) == 0 {
				pos = getVal(2)
			} else {
				pos += 3
			}
		case 7:
			val := 0
			if getVal(1) < getVal(2) {
				val = 1
			}
			memory[memory[pos+3]] = val
			pos += 4
		case 8:
			val := 0
			if getVal(1) == getVal(2) {
				val = 1
			}
			memory[memory[pos+3]] = val
			pos += 4
		case 99:
			// Done
			break For
		default:
			log.Panicf("Op code not recognized: %d", opcode)
		}
	}

	return result
}

func Puzzle2() []int {
	data := utils.ReadAll("./input")
	return puzzle2(parseInput(data), []int{5})
}
