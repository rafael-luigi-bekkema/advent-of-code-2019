package day17

import (
	"aoc/intcode"
	"aoc/utils"
)

func puzzle1(data []int64) int {
	input, output := make(chan int64), make(chan int64)
	go intcode.IntcodeComp(data, input, output)

	var grid []int64
	var width, count int
	for val := range output {
		// fmt.Printf("%c", val)

		if val == '\n' {
			if width == 0 {
				width = count
			}
			continue
		}
		grid = append(grid, val)

		count++
	}

	var total int
	for idx, val := range grid {
		x := idx % width
		y := idx / width

		if idx%width == 0 {
			// fmt.Println()
		}

		skip := val != '#' || y == 0 || x == 0 || x == width-1 || y == (len(grid)/width)-1

		// Check up,down,left and right
		if !skip && grid[(y-1)*width+x] == '#' && grid[(y+1)*width+x] == '#' && grid[y*width+x-1] == '#' && grid[y*width+x+1] == '#' {
			// fmt.Printf("intersect: %d,%d\n", x, y)
			total += x * y
			// fmt.Printf("%c", 'O')
			continue
		}

		// fmt.Printf("%c", val)
	}

	return total
}

func puzzle2(data []int64) int64 {
	data[0] = 2
	input, output := make(chan int64), make(chan int64)
	go intcode.IntcodeComp(data, input, output)

	feed := func(inp string) {
		for _, c := range inp {
			input <- int64(c)
		}
		input <- 10
	}

	go func() {
		feed("A,B,A,C,A,B,C,C,A,B") // Movement routine
		feed("R,8,L,10,R,8")        // A
		feed("R,12,R,8,L,8,L,12")   // B
		feed("L,12,L,10,L,8")       // C
		feed("n")                   // No video feed
	}()

	var last int64
	for c := range output {
		last = c
		// fmt.Printf("%c", c)
	}

	return last
}

func Puzzle1() int {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	return puzzle1(data)
}

func Puzzle2() int64 {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	return puzzle2(data)
}

//A: R,8,L,10,R,8
//B: R,12,R,8,L,8,L,12
//C: L,12,L,10,L,8

//A,B,A,C,A,B,C,C,A,B

//R,8,L,10,R,8,R,12,R,8,L,8,L,12,R,8,L,10,R,8,L,12,L,10,L,8,R,8,L,10,R,8,R,12,R,8,L,8,L,12,L,12,L,10
//L,8,L,12,L,10,L,8,R,8,L,10,R,8,R,12,R,8,L,8,L,12
