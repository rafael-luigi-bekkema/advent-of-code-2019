package day04

import "strconv"

func testPass(pass int) bool {
	strPass := strconv.Itoa(pass)
	if len(strPass) != 6 {
		return false
	}

	var prevChar int32
	dupeDigit := false
	for idx, char := range strPass {
		if char == prevChar {
			dupeDigit = true
		}
		if idx > 0 && char < prevChar {
			return false
		}
		prevChar = char
	}
	return dupeDigit
}

func puzzle1(min, max int) int {
	count := 0
	for pass := min; pass <= max; pass++ {
		if testPass(pass) {
			count++
		}
	}
	return count
}

func Puzzle1() int {
	return puzzle1(128392, 643281)
}

func testPass2(pass int) bool {
	strPass := strconv.Itoa(pass)
	if len(strPass) != 6 {
		return false
	}

	var prevChar int32
	counts := make(map[int32]int)
	for idx, char := range strPass {
		counts[char]++
		if idx > 0 && char < prevChar {
			return false
		}
		prevChar = char
	}
	for _, count := range counts {
		if count == 2 {
			return true
		}
	}
	return false
}

func puzzle2(min, max int) int {
	count := 0
	for pass := min; pass <= max; pass++ {
		if testPass2(pass) {
			count++
		}
	}
	return count
}

func Puzzle2() int {
	return puzzle2(128392, 643281)
}
