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

func ExamplePuzzle2() {
	layerPrinter(Puzzle2(), 25, false)

	// Output:
	//X__X__XX___XX__XXXX__XX__
	//X_X__X__X_X__X_X____X__X_
	//XX___X____X____XXX__X____
	//X_X__X____X_XX_X____X____
	//X_X__X__X_X__X_X____X__X_
	//X__X__XX___XXX_XXXX__XX__
}
