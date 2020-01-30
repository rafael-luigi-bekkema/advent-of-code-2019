package day16

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tt := []struct {
		input, expect string
	}{
		{"80871224585914546619083218645595", "24176176"},
		{"19617804207202209144916044189917", "73745418"},
		{"69317163492948606335995924319873", "52432133"},
	}

	for idx, tc := range tt {
		result := puzzle1(tc.input)
		if result != tc.expect {
			t.Errorf("test %d: expected %q, got %q", idx+1, tc.expect, result)
		}
	}
}

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: 88323090
}

func TestPuzzle2(t *testing.T) {
	tt := []struct {
		input, expect string
	}{
		{"03036732577212944063491565474664", "84462026"},
		{"02935109699940807407585447034323", "78725270"},
		{"03081770884921959731165446850517", "53553731"},
	}

	for idx, tc := range tt {
		result := puzzle2(tc.input)
		if result != tc.expect {
			t.Errorf("test %d: expected %q, got %q", idx+1, tc.expect, result)
		}
	}
}

func ExamplePuzzle2() {
	fmt.Println(Puzzle2())

	// Output: 50077964
}
