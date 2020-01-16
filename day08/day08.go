package day08

import (
	"aoc/utils"
	"fmt"
)

func puzzle1(input string, width, height int) int {
	pixelsPerLayer := width * height
	layerCount := len(input) / pixelsPerLayer
	layers := make([]map[int32]int, layerCount)
	layer := -1
	for idx, digit := range input {
		if idx%pixelsPerLayer == 0 {
			// New layer
			layer++
			layers[layer] = make(map[int32]int)
		}
		layers[layer][digit]++
	}

	layerIdx, minLayer0s := -1, -1
	for idx, layer := range layers {
		if idx == 0 || layer['0'] < minLayer0s {
			minLayer0s = layer['0']
			layerIdx = idx
		}
	}

	ones := layers[layerIdx]['1']
	twos := layers[layerIdx]['2']

	return ones * twos
}

func Puzzle1() int {
	data := utils.ReadAll("./input")
	return puzzle1(data, 25, 6)
}

func puzzle2(input string, width, height int) []int32 {
	pixelsPerLayer := width * height
	layer := make([]int32, pixelsPerLayer)

	// Make initial layer transparent
	for idx := range layer {
		layer[idx] = '2'
	}

	// Decode picture
	for idx, digit := range input {
		lidx := idx % pixelsPerLayer

		// If current pixel is transparent
		// overwrite with this layer's pixel value
		if layer[lidx] == '2' {
			layer[lidx] = digit
		}
	}

	return layer
}

var (
	Black = "\033[1;40m%s\033[0m"
	White = "\033[1;47m%s\033[0m"
)

func layerPrinter(layer []int32, width int, color bool) {
	for idx, digit := range layer {
		if idx != 0 && idx%width == 0 {
			fmt.Print("\n")
		}

		switch digit {
		case '0': // Black
			if color {
				fmt.Print(fmt.Sprintf(Black, " "))
			} else {
				fmt.Print("_")
			}
		case '1':
			if color {
				fmt.Print(fmt.Sprintf(White, " "))
			} else {
				fmt.Print("X")
			}
		}
	}
}

func Puzzle2() []int32 {
	data := utils.ReadAll("./input")
	return puzzle2(data, 25, 6)
}
