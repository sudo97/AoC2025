package day3

import (
	"strings"
	"testing"
)

func TestJoltage(t *testing.T) {
	t.Run("Two digits case", func(t *testing.T) {
		table := []struct {
			input  []int
			output int64
		}{
			{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, 98},
			{[]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, 89},
			{[]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 78},
			{[]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 92},
		}

		for _, test := range table {
			res := Joltage(test.input, 2)
			if res != test.output {
				t.Errorf("for %v, expected %d, got %d", test.input, test.output, res)
			}
		}
	})

	t.Run("Three digits case", func(t *testing.T) {
		table := []struct {
			input  []int
			output int64
		}{
			{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, 987},
			{[]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, 819},
			{[]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 478},
			{[]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 921},
		}

		for _, test := range table {
			res := Joltage(test.input, 3)
			if res != test.output {
				t.Errorf("for %v, expected %d, got %d", test.input, test.output, res)
			}
		}
	})

	t.Run("12 digits case", func(t *testing.T) {
		table := []struct {
			input  []int
			output int64
		}{
			{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, 987654321111},
			{[]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, 811111111119},
			{[]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 434234234278},
			{[]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 888911112111},
		}

		for _, test := range table {
			res := Joltage(test.input, 12)
			if res != test.output {
				t.Errorf("for %v, expected %d, got %d", test.input, test.output, res)
			}
		}
	})
}

func TestCalcSum(t *testing.T) {
	test := [][]int{{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
		{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
		{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
		{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
	}

	if calcSum(test, 2) != 357 {
		t.Errorf("expected 357, got %d", calcSum(test, 2))
	}
}

func TestParseLine(t *testing.T) {
	table := []struct {
		input  string
		output []int
	}{
		{"987654321111111", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}},
	}

	for _, test := range table {
		res := ParseLine(test.input)
		if len(res) != len(test.output) {
			t.Errorf("For %s expected len %d, got len %d", test.input, len(test.output), len(res))
			continue
		}
		for i, v := range res {
			if v != test.output[i] {
				t.Errorf("res[%d] expected %d, got %d", i, test.output[i], v)
			}
		}
	}
}

func TestReadAndParseLines(t *testing.T) {
	input := `987654321111111
123456789
987654321
55555
`

	reader := strings.NewReader(input)
	_, err := ReadAndParseLines(reader)
	if err != nil {
		t.Errorf("Expected no errors")
	}
}
