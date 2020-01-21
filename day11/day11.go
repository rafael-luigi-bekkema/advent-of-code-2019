package day11

import (
	"aoc/intcode"
	"aoc/utils"
	"fmt"
)

type coord struct {
	x, y int
}

func paintHull(data []int64, puzzle int, paint bool) int {
	panels := make(map[coord]int64)
	pos := coord{0, 0}
	input, output := make(chan int64), make(chan int64)
	done := make(chan bool)
	go func() {
		intcode.IntcodeComp(data, input, output)
		done <- true
	}()
	direction := 0 // 0 = up, 1 = right, 2 = down, 3 = left

	if puzzle == 2 {
		panels[pos] = 1 // Start with white panel
	}

	var minX, maxX, minY, maxY int

Outer:
	for {
		select {
		case input <- panels[pos]:
		case <-done:
			break Outer
		}

		color := <-output
		panels[pos] = color

		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.x < minX {
			minX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
		if pos.y < minY {
			minY = pos.y
		}

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

	if puzzle == 2 && paint {
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if panels[coord{x, y}] == 0 {
					utils.Print(utils.Black, " ")
				} else {
					utils.Print(utils.White, " ")
				}
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}

	return len(panels)
}

func Puzzle1() int {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	return paintHull(data, 1, false)
}

func Puzzle2(paint bool) int {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	return paintHull(data, 2, paint)
}
