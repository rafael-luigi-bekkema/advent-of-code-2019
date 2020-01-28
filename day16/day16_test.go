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
