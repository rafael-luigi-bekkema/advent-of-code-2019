package day16

import (
	"aoc/utils"
	"strconv"
)

func parseInput(input string) []int {
	res := make([]int, len(input))
	for idx, strval := range input {
		res[idx] = int(strval - 48)
	}
	return res
}

func algo(arr []int) []int {
	base := []int{0, 1, 0, -1}

	for i := 0; i < 100; i++ {
		newArr := make([]int, len(arr))
		for idx := range arr {
			var newVal int
			for idx2, val := range arr {
				seqVal := base[((idx2+1)/(idx+1))%4]
				newVal += val * seqVal
			}
			// Keep only the ones
			if newVal < 0 {
				newVal *= -1
			}
			if newVal >= 10 {
				newVal = newVal % (newVal / 10 * 10)
			}
			newArr[idx] = newVal
		}
		arr = newArr
	}
	return arr
}

func puzzle1(input string) string {
	arr := parseInput(input)
	arr = algo(arr)

	strArr := make([]int32, 8)
	for i := 0; i < 8; i++ {
		strArr[i] = int32(arr[i] + 48)
	}
	return string(strArr)
}

func Puzzle1() string {
	return puzzle1(utils.ReadAll("./input"))
}

func mulArray(inputArr []int, mul int) []int {
	arr := make([]int, mul*len(inputArr))
	for i := 0; i < mul; i++ {
		copy(arr[i*len(inputArr):], inputArr)
	}
	return arr
}

func ones(newVal int) int {
	if newVal < 0 {
		newVal *= -1
	}
	if newVal >= 10 {
		newVal = newVal % (newVal / 10 * 10)
	}
	return newVal
}

func puzzle2(input string) string {
	arr := mulArray(parseInput(input), 10000)
	baseIdx, _ := strconv.Atoi(input[:7])

	newArr := arr[baseIdx:]

	/**
	Two realizations:
	1. Indeces before start index can be ignored
	2. Because start index is always over half, the second half of the patter 0 and -1 don't come into play.
	   The current index is always 0, and the rest are 1
	*/

	for i := 0; i < 100; i++ {
		str := make([]int, len(newArr))
		var total int
		for e := range newArr {
			if e == 0 {
				total = 0
				for _, f := range newArr {
					total += f
				}
			} else {
				total -= newArr[e-1]
			}
			str[e] = ones(total)
		}
		newArr = str
	}

	var strArr [8]int32
	for i := 0; i < 8; i++ {
		strArr[i] = int32(newArr[i] + 48)
	}
	return string(strArr[:])
}

func Puzzle2() string {
	return puzzle2(utils.ReadAll("./input"))
}
