package day5

import (
	"strings"
	"sudo97/AoC2025/day2"
	"testing"
)

const Sample = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func Example() {
	r := strings.NewReader(Sample)
	Solution(r)

	// Output:
	// day5, part1: 3
	// day5, part2: 14
}

func TestIsInRange(t *testing.T) {
	table := []struct {
		rng    day2.Range
		value  int64
		output bool
	}{
		{day2.Range{Start: 3, End: 5}, 1, false},
		{day2.Range{Start: 3, End: 5}, 3, true},
		{day2.Range{Start: 3, End: 5}, 5, true},
		{day2.Range{Start: 3, End: 5}, 6, false},
	}

	for _, tst := range table {
		res := isInRange(tst.value, &tst.rng)
		if res != tst.output {
			t.Errorf("for %v and %d, got=%t, want=%t", tst.rng, tst.value, res, tst.output)
		}
	}
}

func TestRangesParser(t *testing.T) {
	input := `3-5
19-23

1
2
3`
	expectedRanges := []day2.Range{
		{Start: 3, End: 5},
		{Start: 19, End: 23},
	}

	expectedIds := []int64{1, 2, 3}

	r := strings.NewReader(input)

	res := parse(r)

	if len(res.idRange) != len(expectedRanges) {
		t.Errorf("Length is different")
	}
	for i, val := range expectedRanges {
		if val != res.idRange[i] {
			t.Errorf("not ok")
		}
	}

	if len(res.ids) != len(expectedIds) {
		t.Errorf("Length is different")
	}
	for i, val := range expectedRanges {
		if val != res.idRange[i] {
			t.Errorf("not ok")
		}
	}
}

func TestIsOverlapping(t *testing.T) {
	table := []struct {
		a   *day2.Range
		b   *day2.Range
		res bool
	}{
		{&day2.Range{Start: 3, End: 5}, &day2.Range{Start: 10, End: 12}, false},
		{&day2.Range{Start: 10, End: 12}, &day2.Range{Start: 3, End: 5}, false},
		{&day2.Range{Start: 3, End: 5}, &day2.Range{Start: 4, End: 12}, true},
		{&day2.Range{Start: 16, End: 20}, &day2.Range{Start: 12, End: 18}, true},
		{&day2.Range{Start: 12, End: 18}, &day2.Range{Start: 16, End: 20}, true},
	}

	for _, tst := range table {
		res := isOverlapping(tst.a, tst.b)
		if res != tst.res {
			t.Errorf("Expected %t, got %t, for %v and %v", tst.res, res, tst.a, tst.b)
		}
	}
}

func TestMergeOverlapping(t *testing.T) {
	table := []struct {
		a   *day2.Range
		b   *day2.Range
		res *day2.Range
	}{
		{&day2.Range{Start: 4, End: 12}, &day2.Range{Start: 3, End: 5}, &day2.Range{Start: 3, End: 12}},
	}

	for _, tst := range table {
		res := mergeOverlapping(tst.a, tst.b)
		if res != *tst.res {
			t.Errorf("Expected %v, got %v, for %v and %v", tst.res, res, tst.a, tst.b)
		}
	}
}
