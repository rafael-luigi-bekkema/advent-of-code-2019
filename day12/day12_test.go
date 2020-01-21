package day12

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tt := []struct {
		data          []string
		steps, expect int
	}{
		{[]string{"<x=-1, y=0, z=2>",
			"<x=2, y=-10, z=-7>",
			"<x=4, y=-8, z=8>",
			"<x=3, y=5, z=-1>"}, 10, 179},
		{[]string{"<x=-8, y=-10, z=0>",
			"<x=5, y=5, z=10>",
			"<x=2, y=-7, z=3>",
			"<x=9, y=-8, z=-3>"}, 100, 1940},
	}

	for idx, tc := range tt {
		result := puzzle1(tc.data, tc.steps)
		if tc.expect != result {
			t.Errorf("test %d: expected %d, got %d", idx+1, tc.expect, result)
		}
	}
}

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: 10845
}
