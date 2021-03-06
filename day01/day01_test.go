package day01

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tt := []struct {
		masses []string
		expect int
	}{
		{[]string{"12"}, 2},
		{[]string{"12", "14"}, 4},
		{[]string{"1969"}, 654},
		{[]string{"100756"}, 33583},
		{[]string{"1969", "100756"}, 654 + 33583},
		{[]string{"12", "14", "1969", "100756"}, 34241},
	}

	for idx, tc := range tt {
		res := puzzle1(tc.masses)
		if res != tc.expect {
			t.Fatalf("test %d: expected %d, go %d", idx+1, tc.expect, res)
		}
	}
}

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: 3270717
}

func TestPuzzle2(t *testing.T) {
	tt := []struct {
		masses []string
		expect int
	}{
		{[]string{"14"}, 2},
		{[]string{"1969"}, 966},
		{[]string{"100756"}, 50346},
		{[]string{"100756", "1969"}, 966 + 50346},
	}

	for idx, tc := range tt {
		res := puzzle2(tc.masses)
		if res != tc.expect {
			t.Fatalf("test %d: expected %d, go %d", idx+1, tc.expect, res)
		}
	}
}

func ExamplePuzzle2() {
	fmt.Println(Puzzle2())

	// Output: 4903193
}
