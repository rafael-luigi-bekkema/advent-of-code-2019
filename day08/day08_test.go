package day08

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tt := []struct {
		input         string
		width, height int
		expect        int
	}{
		{"123456789012", 3, 2, 1},
		{"100221210021", 2, 2, 4},
	}

	for idx, tc := range tt {
		result := puzzle1(tc.input, tc.width, tc.height)
		if result != tc.expect {
			t.Errorf("test %d: expected %d, got %d", idx+1, tc.expect, result)
		}
	}
}

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: 1584
}
