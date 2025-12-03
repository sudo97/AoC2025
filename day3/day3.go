package day3

import (
	"bufio"
	"fmt"
	"io"
)

func Joltage(inp []int, numDigits int) int64 {
	res := make([]int, numDigits)

	lastMaxInx := 0
	for currInd := range numDigits {

		shift := numDigits - currInd - 1

		for i := lastMaxInx; i < len(inp)-shift; i++ {
			if inp[i] > inp[lastMaxInx] {
				lastMaxInx = i
			}
		}

		res[currInd] = inp[lastMaxInx]
		lastMaxInx++
	}

	return intoNumber(res)
}

func intoNumber(inp []int) int64 {
	var res int64 = 0

	for _, v := range inp {
		res = res*10 + int64(v)
	}

	return res
}

func ParseLine(input string) []int {
	result := make([]int, 0, len(input))

	for _, char := range input {
		if char >= '0' && char <= '9' {
			result = append(result, int(char-'0'))
		}
	}

	return result
}

// Function to read and parse all lines from a reader
func ReadAndParseLines(reader io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(reader)
	var allLines [][]int

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		parsedLine := ParseLine(line)

		allLines = append(allLines, parsedLine)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}

	return allLines, nil
}

func calcSum(allLines [][]int) (int64, int64) {
	var sum1 int64 = 0
	var sum2 int64 = 0
	for _, v := range allLines {
		sum1 += Joltage(v, 2)
		sum2 += Joltage(v, 12)
	}
	return sum1, sum2
}

func Solution(reader io.Reader) {
	allLines, err := ReadAndParseLines(reader)
	if err != nil {
		panic(err)
	}
	p1, p2 := calcSum(allLines)
	fmt.Printf("part1: %d\npart2: %d\n", p1, p2)
}
