package day02

import (
	"fmt"
	"testing"
)

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: 7210630
}

func TestPuzzle1(t *testing.T) {
	tt := []struct {
		program []int
		expect  int
	}{
		{[]int{1, 0, 0, 0, 99}, 2},
		{[]int{2, 3, 0, 3, 99}, 2},
		{[]int{2, 4, 4, 5, 99, 0}, 2},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 30},
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, 3500},
	}

	for idx, tc := range tt {
		result := puzzle1(tc.program, tc.program[1], tc.program[2])
		if result != tc.expect {
			t.Errorf("test %d: expected %d, got %d", idx+1, tc.expect, result)
		}
	}
}

func ExamplePuzzle2() {
	fmt.Println(Puzzle2())

	// Output: 3892
}
