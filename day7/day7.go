package day7

import (
	"fmt"
	"io"
	"slices"
)

func Solution(reader io.Reader) {
	beams, splitters, width := parseInput(reader)
	fmt.Printf("day7, part1: %d\n", part1(beams, splitters, width))
	fmt.Printf("day7, part2: %d\n", part2(beams, splitters, width))
}

func part1(beams []int, splitters [][]int, width int) int {
	arr := make([]int, width)
	for _, beam := range beams {
		arr[beam]++
	}
	count := 0

	for _, sp := range splitters {
		newBeams, splits := part2Step(arr, sp)

		count += splits
		arr = newBeams
	}

	return count
}

func part2(beams []int, splitters [][]int, width int) int {
	arr := make([]int, width)

	for _, beam := range beams {
		arr[beam]++
	}

	for _, sp := range splitters {
		arr, _ = part2Step(arr, sp)
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func part2Step(arr []int, splitters []int) ([]int, int) {
	newArr := make([]int, len(arr))
	count := 0

	for sp, v := range arr {
		if slices.Contains(splitters, sp) {
			newArr[sp-1] += v
			newArr[sp+1] += v
		} else {
			newArr[sp] += v
		}
	}
	return newArr, count
}
