package day4

import (
	"strings"
	"testing"
)

func TestHasLessThanFourNeighbors(t *testing.T) {
	matrix := [][]bool{
		{false, false, true},
		{false, true, false},
		{false, false, false},
	}

	if !HasLessThanFourNeighbors(matrix, 1, 1) {
		t.Errorf("Expected true, got false\n")
	}
}

const sample = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestParser(t *testing.T) {
	table := []struct {
		inp      string
		expected [][]bool
	}{
		{`.@`, [][]bool{{false, true}}},
		{`@.`, [][]bool{{true, false}}},
		{"@.\n.@", [][]bool{{true, false}, {false, true}}},
	}

	for _, test := range table {

		r := strings.NewReader(test.inp)

		expected := test.expected

		result := parse(r)

		if len(result) != len(expected) {
			t.Errorf("matrices are unequal in height, want=%d, got=%d", len(expected), len(result))
		}

		for i, row := range expected {
			if len(row) != len(result[i]) {
				t.Errorf("matrices are unequal in width, want=%d, got=%d", len(row), len(result[i]))
			}
			for j, cell := range row {
				if cell != result[i][j] {
					t.Errorf("at %d,%d want=%t, got=%t", i, j, cell, result[i][j])
				}
			}
		}
	}
}

func TestPart1(t *testing.T) {
	simplerSample := `...
@..
...`
	r := parse(strings.NewReader(simplerSample))

	expected := 1

	if part1(r) != expected {
		t.Errorf("Expected %d, got=%d", expected, part1(r))
	}
}

func TestPart1Sample(t *testing.T) {
	simplerSample := sample
	r := parse(strings.NewReader(simplerSample))

	expected := 13

	if part1(r) != expected {
		t.Errorf("Expected %d, got=%d", expected, part1(r))
	}
}

func TestPart2Sample(t *testing.T) {
	simplerSample := sample
	r := parse(strings.NewReader(simplerSample))

	expected := 43

	if part2(r) != expected {
		t.Errorf("Expected %d, got=%d", expected, part2(r))
	}
}
