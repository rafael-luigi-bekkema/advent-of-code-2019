package day07

import (
	"aoc/utils"
	"log"
	"strconv"
	"strings"
	"sync"
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

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	var res [][]int

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func intcodeComp(init []int, input <-chan int, output chan<- int) {
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
			memory[memory[pos+1]] = <-input
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
}

func puzzle1(init []int) int {
	outputSignal := 0
	inpChan := make(chan int, 2)
	outChan := make(chan int, 1)
	for idx, phaseSettings := range permutations([]int{0, 1, 2, 3, 4}) {
		input := 0
		for _, phaseSetting := range phaseSettings {
			inpChan <- phaseSetting
			inpChan <- input
			intcodeComp(init, inpChan, outChan)
			input = <-outChan
		}
		if idx == 0 || input > outputSignal {
			outputSignal = input
		}
	}
	return outputSignal
}

func Puzzle1() int {
	return puzzle1(parseInput(utils.ReadAll("./input")))
}

type amp struct {
	input  chan int
	output chan int
}

func puzzle2(init []int) int {
	var mutex sync.Mutex
	outputSignal := 0
	for _, phaseSettings := range permutations([]int{5, 6, 7, 8, 9}) {
		running := 5
		amps := make([]*amp, 5)
		// Start up each Amp and give it the phaseSetting as input
		for idx, phaseSetting := range phaseSettings {
			amp := &amp{
				input:  make(chan int),
				output: make(chan int),
			}
			amps[idx] = amp
			go func() {
				intcodeComp(init, amp.input, amp.output)
				mutex.Lock()
				running--
				mutex.Unlock()
			}()
			amp.input <- phaseSetting
		}

		// Loop until the amps quit
		input := 0
	Outer:
		for {
			for _, amp := range amps {
				amp.input <- input
				input = <-amp.output

				if outputSignal == 0 || input > outputSignal {
					outputSignal = input
				}

				if running == 0 {
					break Outer
				}
			}
		}
	}
	return outputSignal
}

func Puzzle2() int {
	return puzzle2(parseInput(utils.ReadAll("./input")))
}
