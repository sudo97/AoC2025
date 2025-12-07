package day7

import (
	"fmt"
	"io"
	"slices"
)

func Solution(reader io.Reader) {
	beams, splitters, width := parseInput(reader)
	fmt.Printf("day7, part1: %d\n", part1(beams, splitters))
	fmt.Printf("day7, part2: %d\n", part2(beams, splitters, width))
}

func part1(beams []int, splitters [][]int) int {
	count := 0

	for _, sp := range splitters {
		newBeams, splits := part1Step(beams, sp)

		count += splits
		beams = newBeams
	}

	return count
}

func part2(beams []int, splitters [][]int, width int) int {
	arr := make([]int, width)

	for _, beam := range beams {
		arr[beam]++
	}

	for _, sp := range splitters {
		arr = part2Step(arr, sp)
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func part2Step(arr []int, splitters []int) []int {
	newArr := make([]int, len(arr))

	for sp, v := range arr {
		if slices.Contains(splitters, sp) {
			newArr[sp-1] += v
			newArr[sp+1] += v
		} else {
			newArr[sp] += v
		}
	}
	return newArr
}

func part1Step(beams []int, splitters []int) ([]int, int) {
	res := make([]int, 0, len(beams))
	numSplits := 0

	for _, beam := range beams {
		if slices.Contains(splitters, beam) {
			numSplits += 1
			if !slices.Contains(res, beam-1) {
				res = append(res, beam-1)
			}
			if !slices.Contains(res, beam+1) {
				res = append(res, beam+1)
			}
		} else {
			if !slices.Contains(res, beam) {
				res = append(res, beam)
			}
		}
	}

	return res, numSplits
}
