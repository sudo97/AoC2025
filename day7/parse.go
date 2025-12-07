package day7

import (
	"bufio"
	"io"
)

func parseInput(r io.Reader) ([]int, [][]int, int) {
	scanner := bufio.NewScanner(r)

	var beams []int
	var width int

	if scanner.Scan() {
		line := scanner.Text()

		width = len(line)

		beams = parseInitBeam(line)
	}

	splitters := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		splitters = append(splitters, parseSplitterLine(line))
	}

	return beams, splitters, width
}

func parseInitBeam(s string) []int {
	return findAllEntries(s, 'S')
}

func parseSplitterLine(s string) []int {
	return findAllEntries(s, '^')
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
