package day7

import (
	"bufio"
	"fmt"
	"io"
	"slices"
)

func split(beams []int, splitters []int) ([]int, int) {
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

func Solution(reader io.Reader) {
	beams, splitters := parseInput(reader)
	fmt.Printf("day7, part1: %d\n", part1(beams, splitters))
	fmt.Printf("day7, part2: %d\n", part2(reader))
}

func part2(io.Reader) int {
	return 40
}

func part1(beams []int, splitters [][]int) int {
	count := 0

	for _, sp := range splitters {
		newBeams, splits := split(beams, sp)

		count += splits
		beams = newBeams
	}

	return count
}

func parseInput(r io.Reader) ([]int, [][]int) {
	scanner := bufio.NewScanner(r)

	var beams []int

	if scanner.Scan() {
		beams = parseInitBeam(scanner.Text())
	} else {
		panic("No init beam")
	}

	splitters := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		splitters = append(splitters, parseSplitterLine(line))
	}

	return beams, splitters
}

func findAllEntries(s string, c byte) []int {
	res := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			res = append(res, i)
		}
	}
	return res
}

func parseInitBeam(s string) []int {
	return findAllEntries(s, 'S')
}

func parseSplitterLine(s string) []int {
	return findAllEntries(s, '^')
}
