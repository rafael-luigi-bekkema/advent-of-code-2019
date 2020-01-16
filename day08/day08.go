package day08

import "aoc/utils"

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
