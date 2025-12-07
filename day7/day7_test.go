package day7

import (
	"slices"
	"strings"
	"testing"
)

const Sample = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

func Example() {
	Solution(strings.NewReader(Sample))

	// Output:
	// day7, part1: 21
	// day7, part2: 40
}

func TestSplit(t *testing.T) {
	table := []struct {
		beams     []int
		splitters []int
		output    []int
		numSplits int
	}{
		{[]int{3}, []int{}, []int{3}, 0},
		{[]int{3}, []int{3}, []int{2, 4}, 1},
		{[]int{2, 4}, []int{2, 3}, []int{1, 3, 4}, 1},
		{[]int{2, 4}, []int{2, 4}, []int{1, 3, 5}, 2},
		{[]int{2, 4, 6, 7, 8, 10, 12}, []int{2, 6, 12}, []int{1, 3, 4, 5, 7, 8, 10, 11, 13}, 3},
	}

	if len(table) == 0 {
		t.Error("Can't have empty table")
	}

	for _, tst := range table {
		res, splits := part1Step(tst.beams, tst.splitters)
		if len(res) != len(tst.output) {
			t.Errorf("want len(res)=%d, got %d", len(tst.output), len(res))
			continue
		}
		for i, v := range tst.output {
			if v != res[i] {
				t.Errorf("want res[%d]=%d, got %d", i, v, res[i])
			}
		}
		if splits != tst.numSplits {
			t.Errorf("got %d numsplits, want %d", splits, tst.numSplits)
		}
	}
}

func TestStepTrace(t *testing.T) {
	table := []struct {
		arr  []int
		sp   []int
		want []int
	}{
		{
			[]int{0, 0, 1, 0, 0},
			[]int{2},
			[]int{0, 1, 0, 1, 0},
		},
		{
			[]int{0, 1, 0, 1, 0},
			[]int{1, 3},
			[]int{1, 0, 2, 0, 1},
		},
		{
			[]int{0, 1, 0, 2, 0},
			[]int{1, 3},
			[]int{1, 0, 3, 0, 2},
		},
		{
			[]int{1, 1, 0, 2, 0},
			[]int{1, 3},
			[]int{2, 0, 3, 0, 2},
		},
		{
			[]int{0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 0, 0, 0, 0, 0},
			[]int{5, 7, 9},
			[]int{0, 0, 0, 0, 1, 0, 3, 0, 3, 0, 1, 0, 0, 0, 0},
		},
		{
			[]int{0, 0, 0, 0, 1, 0, 3, 0, 3, 0, 1, 0, 0, 0, 0},
			[]int{4, 6, 10},
			[]int{0, 0, 0, 1, 0, 4, 0, 3, 3, 1, 0, 1, 0, 0, 0},
		},
		{
			[]int{0, 0, 0, 1, 0, 4, 0, 3, 3, 1, 0, 1, 0, 0, 0},
			[]int{3, 5, 9, 11},
			[]int{0, 0, 1, 0, 5, 0, 4, 3, 4, 0, 2, 0, 1, 0, 0},
		},
		{
			[]int{0, 0, 1, 0, 5, 0, 4, 3, 4, 0, 2, 0, 1, 0, 0},
			[]int{2, 6, 12},
			[]int{0, 1, 0, 1, 5, 4, 0, 7, 4, 0, 2, 1, 0, 1, 0},
		},
	}

	for _, tst := range table {
		got := part2Step(tst.arr, tst.sp)

		if !slices.Equal(got, tst.want) {
			splitters := make([]int, len(got))
			for i := range got {
				if slices.Contains(tst.sp, i) {
					splitters[i] = 1
				} else {
					splitters[i] = 0
				}
			}
			t.Errorf("\nhad\t\t\t%v\nspl\t\t\t%v\nwant\t\t%v\ngot\t\t\t%v", tst.arr, splitters, tst.want, got)
		}
	}
}
