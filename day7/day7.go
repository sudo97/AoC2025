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
	fmt.Printf("day7, part1: %d\n", part1(reader))
}

func part1(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var beams []int

	if scanner.Scan() {
		beams = parseInitBeam(scanner.Text())
	} else {
		panic("No init beam")
	}

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitters := parseSplitterLine(line)

		newBeams, splits := split(beams, splitters)

		count += splits
		beams = newBeams
	}

	return count
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
