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

func Solution(s string) {
	puzzleInput := parse(strings.NewReader(s))
	fmt.Printf("day6, part1: %d\n", part1(puzzleInput))

	transposed, _ := transpose(s)
	fmt.Printf("%s\n", transposed)
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

func transpose(s string) (string, error) {
	lines := strings.Split(s, "\n")
	linesAsRunes := make([][]rune, len(lines))

	for i, l := range lines {
		linesAsRunes[i] = []rune(l)
	}

	if !allLengthEqual(linesAsRunes) {
		return "", fmt.Errorf("input is not an equal length rows")
	}

	result := []rune("")

	for i := len(linesAsRunes[0]) - 1; i >= 0; i-- {
		for j := 0; j < len(linesAsRunes); j++ {
			c := linesAsRunes[j][i]
			result = append(result, c)
		}

		if i > 0 {
			result = append(result, '\n')
		}
	}

	return string(result), nil
}

func allLengthEqual[T any](matrix [][]T) bool {
	if len(matrix) == 0 {
		return true
	}
	lastLen := len(matrix[0])
	for _, row := range matrix {
		if lastLen != len(row) {
			return false
		}
		lastLen = len(row)
	}
	return true
}
