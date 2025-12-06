package day6

import (
	"strings"
	"testing"
)

func TestParseNumLine(t *testing.T) {
	line := `1 2 123`
	expected := []int{1, 2, 123}
	got := parseNumLine(line)

	if len(expected) != len(got) {
		t.Errorf("length not equal")
	}

	for i, v := range expected {
		if v != got[i] {
			t.Errorf("Expected arr[%d]==%d, got %d", i, v, got[i])
		}
	}
}

func TestShouldParseOp(t *testing.T) {
	line := `+ *`
	expected := true
	got := shouldParseOp(line)
	if expected != got {
		t.Errorf("Expected %t, got %t", expected, got)
	}
}

func TestParseInput(t *testing.T) {
	inp := `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `
	expected := &PuzzleInput{
		numbers: [][]int{
			{123, 328, 51, 64},
			{45, 64, 387, 23},
			{6, 98, 215, 314},
		},
		ops: []string{"*", "+", "*", "+"},
	}
	reader := strings.NewReader(inp)
	got := parse(reader)
	if len(got.numbers) != len(expected.numbers) {
		t.Errorf("numbers is wrong")
	}
	for i, val := range expected.numbers {
		if len(got.numbers[i]) != len(val) {
			t.Errorf("input numbers rows do not match, expected %v, got %v", val, got.numbers[i])
		}
		for j, val := range val {
			if val != got.numbers[i][j] {
				t.Errorf("got[%d][%d]=%d, want=%d", i, j, got.numbers[i][j], val)
			}
		}
	}

	if len(got.ops) != len(expected.ops) {
		t.Errorf("ops length got=%d, want=%d", len(got.ops), len(expected.ops))
	}

	for i, val := range expected.ops {
		if val != got.ops[i] {
			t.Errorf("got[%d]=%s, want=%s", i, got.ops[i], val)
		}
	}
}

func TestParseOp(t *testing.T) {
	inp := `*   +   *   +  `
	res := parseOp(inp)
	if len(res) != 4 {
		t.Errorf("Not 4, got %d", len(res))
	}
}

func Example() {
	sample := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	Solution(sample)

	// TODO: full input has 4 rows of numbers, sample has 3.
	// I solved the puzzle, but this test fails

	// Output:
	// day6, part1: 4277556
	// day6, part2: 3263827
}

func TestTranspose(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"a", "a"},
		{"ab", "b\na"},
		{"ab\ncd", "bd\nac"},
		{"abc\ndef\nghi", "cfi\nbeh\nadg"},
	}

	for _, v := range table {
		res, _ := transpose(v.input)
		if res != v.output {
			t.Errorf("got\n===\n%s\n===\nwant\n===\n%s\n===", res, v.output)
		}
	}
}
