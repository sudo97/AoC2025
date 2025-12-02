package day1

import (
	"fmt"
	"strconv"
	"strings"
)

type direction string

const (
	LEFT  direction = "L"
	RIGHT direction = "R"
)

type entry struct {
	direction direction
	num       int
}

func parseLine(s string) (*entry, error) {
	var direction direction

	switch s[0] {
	case 'L':
		direction = LEFT
	case 'R':
		direction = RIGHT
	default:
		return nil, fmt.Errorf("unknown direction %c", s[0])
	}

	num, err := strconv.ParseInt(s[1:], 10, 32)

	if err != nil {
		return nil, err
	}

	return &entry{direction: direction, num: int(num)}, nil
}

func part1(s string) int {
	counter := 0
	pos := 50
	for l := range strings.SplitSeq(s, "\n") {
		val, err := parseLine(l)
		if err != nil {
			panic(err)
		}
		turns := val.num % 100
		if val.direction == RIGHT {
			pos = (pos + turns) % 100
		} else {
			pos = (100 + pos - turns) % 100
		}
		if pos == 0 {
			counter++
		}
	}
	return counter
}

func part2(s string) int {
	counter := 0
	pos := 50
	for l := range strings.SplitSeq(s, "\n") {
		val, err := parseLine(l)
		if err != nil {
			panic(err)
		}

		turns := val.num % 100
		zeroCount := val.num / 100

		if val.direction == "R" {
			newPos := (pos + turns) % 100
			if newPos < pos {
				zeroCount += 1
			}
			pos = newPos
		} else {
			newPos := (100 + pos - turns) % 100
			if (newPos > pos || newPos == 0) && pos != 0 {
				zeroCount += 1
			}
			pos = newPos
		}
		counter += zeroCount
	}
	return counter
}

func Solution(s string) {
	fmt.Printf("day 1, part 1 is %d\n", part1(s))
	fmt.Printf("day 1, part 2 is %d\n", part2(s))
}
