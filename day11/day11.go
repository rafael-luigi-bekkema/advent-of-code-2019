package day11

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) []int64 {
	vals := strings.Split(input, ",")
	res := make([]int64, len(vals))
	for idx, strval := range vals {
		num, err := strconv.ParseInt(strval, 10, 64)
		if err != nil {
			panic(err)
		}
		res[idx] = num
	}
	return res
}

func intcodeComp(init []int64, input <-chan int64, output chan<- int64) {
	var pos, relBase int64

	memory := make([]int64, len(init))
	copy(memory, init)

For:
	for {
		value := strconv.FormatInt(memory[pos], 10)
		startFrom := 0
		if len(value) > 2 {
			startFrom = len(value) - 2
		}
		opcode, _ := strconv.Atoi(value[startFrom:])

		getVal := func(offset int64) int64 {
			mode := uint8('0')
			if int64(len(value)) >= offset+2 {
				idx := int64(len(value)) - (offset + 2)
				mode = value[idx]
			}

			var readPos int64

			switch mode {
			case '0': // Position mode
				readPos = memory[pos+offset]
			case '1': // Immediate mode
				readPos = pos + offset
			case '2': // Relative mode
				readPos = relBase + memory[pos+offset]
			default:
				panic(fmt.Sprintf("unknown mode: %q in op: %q", mode, value))
			}
			if readPos >= int64(len(memory)) {
				return 0
			}
			return memory[readPos]
		}

		setVal := func(offset int64, newVal int64) {
			mode := uint8('0')
			if int64(len(value)) >= offset+2 {
				idx := int64(len(value)) - (offset + 2)
				mode = value[idx]
			}

			var targetPos int64
			switch mode {
			case '0': // Position mode
				targetPos = memory[pos+offset]
			case '2': // Relative mode
				targetPos = relBase + memory[pos+offset]
			default:
				panic(fmt.Sprintf("unknown mode: %q in op: %q", mode, value))
			}

			if targetPos >= int64(len(memory)) {
				// Extend memory
				newSize := targetPos - int64(len(memory)) + 1
				memory = append(memory, make([]int64, newSize)...)
			}

			memory[targetPos] = newVal
		}

		switch opcode {
		case 1:
			// Add
			setVal(3, getVal(1)+getVal(2))
			pos += 4
		case 2:
			// Mul
			setVal(3, getVal(1)*getVal(2))
			pos += 4
		case 3:
			//Input
			setVal(1, <-input)
			pos += 2
		case 4:
			//Output
			output <- getVal(1)
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
			var val int64
			if getVal(1) < getVal(2) {
				val = 1
			}
			setVal(3, val)
			pos += 4
		case 8:
			var val int64
			if getVal(1) == getVal(2) {
				val = 1
			}
			setVal(3, val)
			pos += 4
		case 9: // adjusts the relative base
			relBase += getVal(1)
			pos += 2
		case 99:
			// Done
			break For
		default:
			panic(fmt.Sprintf("op code not recognized: %d", opcode))
		}
	}
	close(output)
}

type coord struct {
	x, y int
}

func puzzle1(data []int64) int {
	panels := make(map[coord]int64)
	pos := coord{0, 0}
	input, output := make(chan int64), make(chan int64)
	done := make(chan bool)
	go func() {
		intcodeComp(data, input, output)
		done <- true
	}()
	direction := 0 // 0 = up, 1 = right, 2 = down, 3 = left

Outer:
	for {
		select {
		case input <- panels[pos]:
		case <-done:
			break Outer
		}

		color := <-output
		panels[pos] = color

		turn := <-output
		if turn == 0 {
			direction = (direction + 3) % 4
		} else {
			direction = (direction + 1) % 4
		}

		switch direction {
		case 0:
			pos = coord{pos.x, pos.y - 1}
		case 1:
			pos = coord{pos.x + 1, pos.y}
		case 2:
			pos = coord{pos.x, pos.y + 1}
		case 3:
			pos = coord{pos.x - 1, pos.y}
		}
	}

	return len(panels)
}

func Puzzle1() int {
	data := parseInput(utils.ReadAll("./input"))
	return puzzle1(data)
}
