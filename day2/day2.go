package day2

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

func isInvalidPart1(value int64) bool {
	str := fmt.Sprintf("%d", value)
	l := len(str)
	if l%2 != 0 {
		return false
	}

	midPoint := l / 2

	left, right := str[:midPoint], str[midPoint:]

	return left == right
}

func isInvalidPart2(value int64) bool {
	str := fmt.Sprintf("%d", value)

	for stepSize := 1; stepSize <= len(str)/2; stepSize++ {
		if len(str)%stepSize != 0 {
			continue
		}
		first := str[:stepSize]

		res := true

		for j := stepSize; j+stepSize <= len(str); j += stepSize {
			notFirst := str[j : j+stepSize]
			res = res && notFirst == first
		}

		if res {
			return res
		}
	}
	return false
}

func sumOfInvalidInRange(start, end int64, isInvalid func(value int64) bool) int64 {
	var res int64 = 0

	for i := start; i <= end; i++ {
		if isInvalid(i) {
			res += i
		}
	}

	return res
}

type Range struct {
	start int64
	end   int64
}

func ranges(input string) iter.Seq[Range] {
	return func(yield func(Range) bool) {
		for r := range strings.SplitSeq(input, ",") {
			row := strings.Split(r, "-")
			start, _ := strconv.ParseInt(row[0], 10, 64)
			end, _ := strconv.ParseInt(row[1], 10, 64)
			if !yield(Range{start, end}) {
				return
			}
		}
	}
}

func loopAndSum(input string, isInvalid func(value int64) bool) int64 {
	var sum int64 = 0
	for r := range ranges(input) {
		sum += sumOfInvalidInRange(r.start, r.end, isInvalid)
	}
	return sum
}

func part1(input string) int64 {
	return loopAndSum(input, isInvalidPart1)
}

func part2(input string) int64 {
	return loopAndSum(input, isInvalidPart2)
}

func Solution(input string) {
	fmt.Printf("day2, part1: %d\n", part1(input))
	fmt.Printf("day2, part2: %d\n", part2(input))
}
