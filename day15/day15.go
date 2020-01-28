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
	Station
	Oxygen
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
				case Station:
					fmt.Print("S")
				case Oxygen:
					fmt.Print("O")
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

type bot struct {
	mapp          map[coord]*space
	pos           coord
	input, output chan int64
	mode          int64
	lastMove      int64
	station       coord
}

func (b *bot) countSteps(pos, from coord, depth int) int {
	var minCount = -1
	for i := int64(0); i < 4; i++ {
		newPos := pos.move(i)
		if from == newPos {
			continue
		}
		curr := b.mapp[newPos]
		if curr.id == Station {
			return depth
		}
		if curr.id == Wall {
			continue
		}
		if curr.id == Empty {
			count := b.countSteps(pos.move(i), pos, depth+1)
			if count == -1 {
				continue
			}
			if minCount == -1 || count < minCount {
				minCount = count
			}
		}
	}
	return minCount
}

func (b *bot) fillOxygen() int {
	depth := 0
	tips := []coord{b.station}
	for {
		var newTips []coord
		for _, tip := range tips {
			for i := int64(0); i < 4; i++ {
				newPos := tip.move(i)
				curr := b.mapp[newPos]
				if curr.id == Empty {
					curr.id = Oxygen
					newTips = append(newTips, newPos)
				}
			}
		}
		if len(newTips) == 0 {
			return depth
		}
		depth++
		tips = newTips
	}
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
		b.mapp[newPos].id = Station
		b.pos = newPos
		b.station = newPos
	}
	return response
}

func puzzle1(init []int64) int {
	input, output := make(chan int64), make(chan int64)

	go intcode.IntcodeComp(init, input, output)

	theMap := make(map[coord]*space)

	// move: north (0), south (1), west (2), and east (3)
	// move: up (0), down (1), left (2), and right (3)
	b := &bot{mapp: theMap, input: input, output: output}
	b.mapp[coord{}] = &space{id: Empty}

	for {
		if b.autoMove() {
			return b.countSteps(coord{}, coord{}, 1)
		}
	}
}

func puzzle2(init []int64) int {
	input, output := make(chan int64), make(chan int64)

	go intcode.IntcodeComp(init, input, output)

	theMap := make(map[coord]*space)

	// move: north (0), south (1), west (2), and east (3)
	// move: up (0), down (1), left (2), and right (3)
	b := &bot{mapp: theMap, input: input, output: output}
	b.mapp[coord{}] = &space{id: Empty}

	for {
		if b.autoMove() {
			res := b.fillOxygen()
			//render(b.mapp, b.pos)
			return res
		}
	}
}

func Puzzle1() int {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	return puzzle1(data)
}

func Puzzle2() int {
	data := intcode.ParseInput(utils.ReadAll("./input"))
	return puzzle2(data)
}
