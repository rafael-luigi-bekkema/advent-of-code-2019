package day16

import (
	"aoc/utils"
)

func parseInput(input string) []int {
	res := make([]int, len(input))
	for idx, strval := range input {
		res[idx] = int(strval - 48)
	}
	return res
}

func sequence(pos int) []int {
	base := []int{0, 1, 0, -1}
	seq := make([]int, pos*4)
	for idx, val := range base {
		for i := 0; i < pos; i++ {
			seq[idx*pos+i] = val
		}
	}
	return seq
}

func puzzle1(input string) string {
	arr := parseInput(input)

	for i := 0; i < 100; i++ {
		newArr := make([]int, len(arr))
		for idx := range arr {
			seq := sequence(idx + 1)
			var newVal int
			for idx2, val := range arr {
				seqVal := seq[(idx2+1)%len(seq)]
				newVal += val * seqVal
			}
			// Keep only the ones
			if newVal < 0 {
				newVal *= -1
			}
			if newVal >= 10 || newVal <= -10 {
				newVal = newVal % (newVal / 10 * 10)
			}
			newArr[idx] = newVal
		}
		arr = newArr
	}

	strArr := make([]int32, 8)
	for i := 0; i < 8; i++ {
		strArr[i] = int32(arr[i] + 48)
	}
	return string(strArr)
}

func Puzzle1() string {
	return puzzle1(utils.ReadAll("./input"))
}
