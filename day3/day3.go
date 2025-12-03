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

func calcSum(allLines [][]int, numDigits int) int64 {
	var sum int64 = 0
	for _, v := range allLines {
		sum += Joltage(v, numDigits)
	}
	return sum
}

func Solution(reader io.Reader) {
	allLines, err := ReadAndParseLines(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("part1: %d\n", calcSum(allLines, 2))
	fmt.Printf("part2: %d\n", calcSum(allLines, 12))
}
