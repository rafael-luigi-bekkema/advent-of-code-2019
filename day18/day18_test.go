package day18

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tt := []struct {
		data   string
		expect int
	}{
		{`#########
#b.A.@.a#
#########`, 8},
		{`########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`, 86},
		{`########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`, 132},
		{`########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`, 81},
		{`#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`, 136},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := puzzle1(tc.data)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: 5
}
