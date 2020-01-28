package day15

import (
	"aoc/intcode"
	"aoc/utils"
	"fmt"
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

	fmt.Printf("\n")
	fmt.Printf("---\n")
}

type space struct {
	id        int
	moves     [4]bool
	moveCount int
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
		//move = manualMove()

		input <- move + 1
		curr.moves[move] = true

		newPos := pos.move(move)

		status := <-output

		newCurr, ok := theMap[newPos]
		if !ok {
			newCurr = &space{}
			theMap[newPos] = newCurr
		}
		if status == 1 || status == 2 {
			if newCurr.id == 0 {
				moves++
				newCurr.moveCount = moves
			} else {
				moves = newCurr.moveCount
			}
		}
		switch status {
		case 0: // wall
			newCurr.id = Wall
			//render(theMap, pos)
		case 1: // move success
			newCurr.id = Empty
			pos = newPos
			//render(theMap, pos)
		case 2: // oxygen system reached
			newCurr.id = System
			//render(theMap, pos)
			break Outer
		}
	}

	return moves
}

type bot struct {
	mapp          map[coord]*space
	pos           coord
	input, output chan int64
	mode          int64
	lastMove      int64
}

func (b *bot) unexploredMove(pos coord) int64 {
	if b.mapp[coord{pos.x + 1, pos.y}] == nil {
		return 3
	}
	if b.mapp[coord{pos.x - 1, pos.y}] == nil {
		return 2
	}
	if b.mapp[coord{pos.x, pos.y + 1}] == nil {
		return 1
	}
	if b.mapp[coord{pos.x, pos.y - 1}] == nil {
		return 0
	}
	return -1
}

func (b *bot) emptyMove() int64 {
	if b.mapp[coord{b.pos.x + 1, b.pos.y}].id == Empty && b.unexploredMove(b.pos.move(3)) != -1 {
		return 3
	}
	if b.mapp[coord{b.pos.x - 1, b.pos.y}].id == Empty && b.unexploredMove(b.pos.move(2)) != -1 {
		return 2
	}
	if b.mapp[coord{b.pos.x, b.pos.y + 1}].id == Empty && b.unexploredMove(b.pos.move(1)) != -1 {
		return 1
	}
	if b.mapp[coord{b.pos.x, b.pos.y - 1}].id == Empty && b.unexploredMove(b.pos.move(0)) != -1 {
		return 0
	}
	panic("no empty")
}

func (b *bot) closestUnexplored(pos, from coord, depth int) int {
	for i := int64(0); i < 4; i++ {
		newPos := pos.move(i)
		if from == newPos {
			continue
		}
		curr := b.mapp[newPos]
		// Has unexplored edge so go here
		if curr == nil {
			return depth
		}
		if curr.id == Wall {
			continue
		}
		if curr.id == Empty {
			move := b.closestUnexplored(newPos, pos, depth+1)
			if move != -1 {
				return depth
			}
		}
	}
	return -1
}

func (b *bot) findUnexplored(pos coord) int64 {
	var minDepth = -1
	var minMove int64 = -1
	for i := int64(0); i < 4; i++ {
		newPos := pos.move(i)
		curr := b.mapp[newPos]
		if curr == nil {
			return i
		}
		if curr.id != Wall {
			depth := b.closestUnexplored(pos.move(i), pos, 0)
			if depth == -1 {
				continue
			}
			if minDepth == -1 || depth < minDepth {
				minDepth = depth
				minMove = i
			}
		}
	}
	return minMove
}

func (b *bot) autoMove() bool {
	move := b.findUnexplored(b.pos)
	if move == -1 {
		return true
	}
	b.move(move)
	return false
}

func (b *bot) move(move int64) int64 {
	b.input <- move + 1
	response := <-b.output
	newPos := b.pos.move(move)
	newCurr, ok := b.mapp[newPos]
	if !ok {
		newCurr = &space{}
		b.mapp[newPos] = newCurr
	}
	switch response {
	case 0:
		b.mapp[newPos].id = Wall
	case 1:
		b.mapp[newPos].id = Empty
		b.pos = newPos
	case 2:
		b.mapp[newPos].id = System
		b.pos = newPos
	}
	return response
}

func puzzle2(init []int64) {
	input, output := make(chan int64), make(chan int64)

	go intcode.IntcodeComp(init, input, output)

	theMap := make(map[coord]*space)

	// move: north (0), south (1), west (2), and east (3)
	// move: up (0), down (1), left (2), and right (3)
	b := &bot{mapp: theMap, input: input, output: output}
	b.mapp[coord{}] = &space{id: Empty}

	for {
		if b.autoMove() {
			render(b.mapp, b.pos)
			break
		}
	}
}

func Puzzle1() int {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	return puzzle1(data)
}

func Puzzle2() {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	puzzle2(data)
}
