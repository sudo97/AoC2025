package day6

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type PuzzleInput struct {
	numbers [][]int
	ops     []string
}

func parseNumLine(s string) []int {
	res := []int{}
	for v := range strings.SplitSeq(s, " ") {

		if len(v) == 0 {
			continue
		}

		v, _ := strconv.Atoi(v)

		res = append(res, v)
	}
	return res
}

func parseOp(inp string) []string {
	split := strings.Split(inp, " ")

	res := []string{}

	for _, item := range split {
		if len(item) != 0 {
			res = append(res, item)
		}
	}
	return res
}

func shouldParseOp(line string) bool {
	return line[0] == '+' || line[0] == '*'
}

func parse(reader io.Reader) *PuzzleInput {
	scanner := bufio.NewScanner(reader)

	inp := PuzzleInput{}

	for scanner.Scan() {
		line := scanner.Text()
		if shouldParseOp(line) {
			inp.ops = parseOp(line)
			return &inp
		}
		inp.numbers = append(inp.numbers, parseNumLine(line))
	}

	panic("No ops")
}

func Solution(reader io.Reader) {
	puzzleInput := parse(reader)
	fmt.Printf("day6, part1: %d\n", part1(puzzleInput))
}

func part1(puzzleInput *PuzzleInput) int {
	count := 0
	for i := 0; i < len(puzzleInput.ops); i++ {
		op := puzzleInput.ops[i]
		switch op {
		case "+":
			for j := 0; j < len(puzzleInput.numbers); j++ {
				count += puzzleInput.numbers[j][i]
			}
		case "*":
			prod := 1
			for j := 0; j < len(puzzleInput.numbers); j++ {
				prod *= puzzleInput.numbers[j][i]
			}
			count += prod
		}
	}
	return count
}
