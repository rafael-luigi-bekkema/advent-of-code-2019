package day09

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntcodeComp(t *testing.T) {
	tt := []struct {
		init, expect []int64
	}{
		{[]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}},
		{[]int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, []int64{1219070632396864}},
		{[]int64{104, 1125899906842624, 99}, []int64{1125899906842624}},
	}

	for idx, tc := range tt {
		input := make(chan int64)
		output := make(chan int64)
		go intcodeComp(tc.init, input, output)
		result := make([]int64, 0, len(tc.expect))
		for val := range output {
			result = append(result, val)
		}
		if !reflect.DeepEqual(result, tc.expect) {
			t.Errorf("test %d: expected %v, got %v", idx+1, tc.expect, result)
		}
	}
}

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: [2775723069]
}

func ExamplePuzzle2() {
	fmt.Println(Puzzle2())

	// Output: [49115]
}
