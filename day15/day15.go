package day15

import (
	"aoc/intcode"
	"aoc/utils"
	"bufio"
	"fmt"
	"os"
	"time"
)

type coord struct {
	x, y int
}

func (c coord) move(m int64) coord {
	newPos := coord{c.x, c.y}
	switch m {
	case 0:
		newPos.y--
	case 1:
		newPos.y++
	case 2:
		newPos.x--
	case 3:
		newPos.x++
	}
	return newPos
}

const (
	Unknown = iota
	Wall
	Empty
	System
)

func render(theMap map[coord]*space, current coord) {
	var minX, maxX, minY, maxY int
	for c := range theMap {
		if c.x < minX {
			minX = c.x
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if x == minX && y != minY {
				fmt.Print("\n")
			}

			c := coord{x, y}
			if c == current {
				fmt.Print("D")
			} else if (c == coord{0, 0}) {
				fmt.Print("B")
			} else {
				point := theMap[c]
				if point == nil {
					point = &space{}
				}
				switch point.id {
				case Unknown:
					fmt.Print(" ")
				case Wall:
					fmt.Print("#")
				case Empty:
					fmt.Print(".")
				case System:
					fmt.Print("S")
				}
			}
		}
	}

	fmt.Print("\n---\n")
}

type space struct {
	id    int
	moves [4]bool
}

func manualMove() int64 {
	reader := bufio.NewReader(os.Stdin)
	var move int64

	fmt.Print("Move (j/k/l/i): ")
	text, _ := reader.ReadString('\n')
	switch text {
	case "j\n": // left
		move = 2
	case "k\n": // down
		move = 1
	case "l\n": // right
		move = 3
	case "i\n": // up
		move = 0
	default:
		move = 3
	}

	return move
}

func autoMove(lastMove int64, pos coord, theMap map[coord]*space) int64 {
	// Look around for unexplored are and prefer that
	if theMap[coord{pos.x + 1, pos.y}] == nil {
		return 3
	}
	if theMap[coord{pos.x - 1, pos.y}] == nil {
		return 2
	}
	if theMap[coord{pos.x, pos.y + 1}] == nil {
		return 1
	}
	if theMap[coord{pos.x, pos.y - 1}] == nil {
		return 0
	}

	newPos := pos.move(lastMove)
	if theMap[newPos].id != Empty {
		switch lastMove {
		case 0:
			return 3
		case 3:
			return 1
		case 1:
			return 2
		case 2:
			return 0
		}
	}

	// Otherwise keep going
	return lastMove
}

func puzzle1(init []int64) int {
	input, output := make(chan int64), make(chan int64)

	go intcode.IntcodeComp(init, input, output)

	theMap := make(map[coord]*space)
	var moves int
	// move: north (0), south (1), west (2), and east (3)
	// move: up (0), down (1), left (2), and right (3)
	var move int64
	var pos coord
Outer:
	for {
		curr, ok := theMap[pos]
		if !ok {
			curr = &space{}
			theMap[pos] = curr
		}
		curr.id = Empty

		move = autoMove(move, pos, theMap)

		input <- move + 1
		curr.moves[move] = true

		newPos := pos.move(move)

		status := <-output

		newCurr, ok := theMap[newPos]
		if !ok {
			newCurr = &space{}
			theMap[newPos] = newCurr
		}
		switch status {
		case 0: // wall
			newCurr.id = Wall
			render(theMap, pos)
		case 1: // move success
			newCurr.id = Empty
			pos = newPos
			moves++
			render(theMap, pos)
		case 2: // oxygen system reached
			newCurr.id = System
			moves++
			render(theMap, pos)
			break Outer
		}

		time.Sleep(time.Millisecond * 10)
	}

	return moves
}

func Puzzle1() {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	puzzle1(data)
}
