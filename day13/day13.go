package day13

import (
	"aoc/intcode"
	"aoc/utils"
)

func puzzle1(init []int64) int {
	input, output := make(chan int64), make(chan int64)
	go intcode.IntcodeComp(init, input, output)

	var count int
	for range output {
		<-output
		id := <-output

		if id == 2 { // block
			count++
		}
	}
	return count
}

func Puzzle1() int {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	return puzzle1(data)
}

func puzzle2(init []int64) int {
	input, output := make(chan int64), make(chan int64, 5)
	end := make(chan int64, 5)
	paddleChange, ballChange := make(chan int64, 5), make(chan int64, 5)
	go func() {
		intcode.IntcodeComp(init, input, output)
	}()

	pixels := make([]int64, 1008)
	width, _ := 42, 24

	//paint := func() {
	//	for idx, pixel := range pixels {
	//		switch pixel {
	//		case 0: // Empty
	//			fmt.Print(" ")
	//		case 1: // Wall
	//			fmt.Print("|")
	//		case 2: // Block
	//			fmt.Print("D")
	//		case 3: // Paddle
	//			fmt.Print("_")
	//		case 4: // Ball
	//			fmt.Print("o")
	//		}
	//
	//		if (idx+1)%width == 0 {
	//			fmt.Print("\n")
	//		}
	//	}
	//
	//	fmt.Printf("Score: %d\n", score)
	//}

	go func() {
		var score int64
		count := 0
		for x := range output {
			y := <-output
			id := <-output

			if x == -1 && y == 0 {
				// id == score
				score = id
				end <- score
			} else {
				pixels[int(y)*width+int(x)] = id
				switch id { // Paddle
				case 3:
					paddleChange <- x
				case 4:
					ballChange <- x
				}
			}

			count++
		}
		close(end)
	}()

	var paddleX, ballX, score int64
	move := func() {
		var joystick int64
		if paddleX > ballX {
			joystick = -1
		} else if paddleX < ballX {
			joystick = 1
		} else {
			joystick = 0
		}
		input <- joystick
	}

	for {
		select {
		case x := <-ballChange:
			ballX = x
			if paddleX != 0 {
				move()
			}
		case x := <-paddleChange:
			if paddleX == 0 {
				paddleX = x
				move()
			} else {
				paddleX = x
			}
		case s, ok := <-end:
			if !ok {
				return int(score)
			}

			score = s
		}
	}
}

func Puzzle2() int {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	data[0] = 2 // Insert coins
	return puzzle2(data)
}
