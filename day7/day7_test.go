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
