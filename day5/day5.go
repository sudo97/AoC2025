package day5

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"sudo97/AoC2025/day2"
)

type PuzzleInput struct {
	idRange []day2.Range
	ids     []int64
}

func isInRange(val int64, rng *day2.Range) bool {
	return val >= rng.Start && val <= rng.End
}

func isInOneOfRanges(val int64, idRange []day2.Range) bool {
	for _, r := range idRange {
		if isInRange(val, &r) {
			return true
		}
	}
	return false
}

func parse(r io.Reader) PuzzleInput {
	scanner := bufio.NewScanner(r)

	idRanges := parseRanges(scanner)

	ids := parseIds(scanner)

	return PuzzleInput{idRanges, ids}
}

func parseIds(scanner *bufio.Scanner) []int64 {
	ids := []int64{}
	for scanner.Scan() {
		line := scanner.Text()
		id, _ := strconv.ParseInt(line, 10, 64)
		ids = append(ids, id)
	}
	return ids
}

func parseRanges(scanner *bufio.Scanner) []day2.Range {
	idRanges := []day2.Range{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		row := day2.ParseRange(line)
		idRanges = append(idRanges, row)
	}
	return idRanges
}

func Solution(r io.Reader) {
	input := parse(r)
	fmt.Printf("day5, part1: %d\n", part1(input))
	fmt.Printf("day5, part2: %d\n", part2(input))
}

func part1(input PuzzleInput) int {
	count := 0
	for _, id := range input.ids {
		if isInOneOfRanges(id, input.idRange) {
			count++
		}
	}
	return count
}

func isOverlapping(a, b *day2.Range) bool {
	return isInRange(a.Start, b) || isInRange(a.End, b)
}

func mergeOverlapping(a, b *day2.Range) day2.Range {
	return day2.Range{
		Start: min(b.Start, a.Start),
		End:   max(a.End, b.End),
	}
}

func part2(input PuzzleInput) int64 {
	slices.SortFunc(input.idRange, func(a, b day2.Range) int {
		if a.Start < b.Start {
			return -1
		}
		return 1
	})

	merged := []day2.Range{}

merging:
	for _, r := range input.idRange {
		for i := 0; i < len(merged); i++ {
			resIdx := merged[i]
			if isOverlapping(&r, &resIdx) {
				merged[i] = mergeOverlapping(&r, &resIdx)
				continue merging
			}
		}
		merged = append(merged, r)
	}

	var count int64 = 0

	for _, r := range merged {
		count += (r.End) - (r.Start) + 1
	}
	return count
}
