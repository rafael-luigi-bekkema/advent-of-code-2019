package day06

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tt := []struct {
		orbits []string
		expect int
	}{
		{[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}, 42},
	}

	for idx, tc := range tt {
		result := puzzle1(tc.orbits)
		if result != tc.expect {
			t.Errorf("test %d: expect %d, got %d", idx+1, tc.expect, result)
		}
	}
}

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: 314247
}
