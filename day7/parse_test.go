package day7

import (
	"slices"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	inp := `..S..
..^..`
	exp := []int{2}
	beams, splitters, width := parseInput(strings.NewReader(inp))
	if !slices.Equal(beams, exp) {
		t.Errorf("want %v, got %v", exp, beams)
	}
	if !slices.EqualFunc(splitters, [][]int{exp}, func(s1 []int, s2 []int) bool { return slices.Equal(s1, s2) }) {
		t.Errorf("want %v, got %v", exp, splitters)
	}

	if width != 5 {
		t.Errorf("Nah")
	}
}

func TestParseInitBeam(t *testing.T) {
	table := []struct {
		input  string
		output []int
	}{
		{"", []int{}},
		{"......................", []int{}},
		{"....S.................", []int{4}},
		{"....SS................", []int{4, 5}},
	}
	if len(table) == 0 {
		t.Error("Can't have empty table")
	}

	for _, tst := range table {
		res := parseInitBeam(tst.input)
		if !slices.Equal(res, tst.output) {
			t.Errorf("got %v, want %v", res, tst.output)
		}
	}
}

func TestParseSplitterLine(t *testing.T) {
	table := []struct {
		input  string
		output []int
	}{
		{"", []int{}},
		{"......................", []int{}},
		{"....^.................", []int{4}},
		{"....^^................", []int{4, 5}},
	}
	if len(table) == 0 {
		t.Error("Can't have empty table")
	}

	for _, tst := range table {
		res := parseSplitterLine(tst.input)
		if !slices.Equal(res, tst.output) {
			t.Errorf("got %v, want %v", res, tst.output)
		}
	}
}
