package day04

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tt := []struct {
		min, max int
		expect   int
	}{
		{111111, 111111, 1},
		{223450, 223450, 0},
		{123789, 123789, 0},
		{123669, 123677, 2},
	}

	for idx, tc := range tt {
		result := puzzle1(tc.min, tc.max)
		if result != tc.expect {
			t.Errorf("test %d: expect %d, got %d", idx+1, tc.expect, result)
		}
	}
}

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: 2050
}

func TestPuzzle2(t *testing.T) {
	tt := []struct {
		min, max int
		expect   int
	}{
		{111111, 111111, 0},
		{112233, 112233, 1},
		{123444, 123444, 0},
		{111122, 111122, 1},
		{123789, 123789, 0},
		{123788, 123788, 1},
		{123669, 123677, 2},
	}

	for idx, tc := range tt {
		result := puzzle2(tc.min, tc.max)
		if result != tc.expect {
			t.Errorf("test %d: expect %d, got %d", idx+1, tc.expect, result)
		}
	}
}

func ExamplePuzzle2() {
	fmt.Println(Puzzle2())

	// Output: 1390
}
