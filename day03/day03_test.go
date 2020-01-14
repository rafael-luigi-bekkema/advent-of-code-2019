package day03

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	dd := []struct {
		line1, line2 string
		expect       int
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}

	for idx, tc := range dd {
		result := puzzle1(tc.line1, tc.line2)
		if result != tc.expect {
			t.Errorf("test %d: expected %d, got %d", idx+1, tc.expect, result)
		}
	}
}

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: 1211
}
