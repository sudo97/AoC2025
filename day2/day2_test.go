package day2

import (
	"testing"
)

func TestIsInvalidPart1(t *testing.T) {
	tests := []struct {
		val      int64
		expected bool
	}{
		{55, true},
		{6464, true},
		{123123, true},
		{121212, false},
		{32, false},
		{101, false},
		{321123, false},
	}

	testIsInvalid(t, tests, isInvalidPart1)
}

func TestIsInvalidPart2(t *testing.T) {
	tests := []struct {
		val      int64
		expected bool
	}{
		{66, true},
		{67, false},
		{666, true},
		{676, false},
		{667, false},
		{121212, true},
		{6666, true},
		{6767, true},
		{55, true},
		{6464, true},
		{123123, true},
		{32, false},
		{101, false},
		{321123, false},
		{1111111, true},
		{12341234, true},
		{123123123, true},
		{1212121212, true},
		{2121212118, false},
		{2121212119, false},
		{2121212120, false},
		{2121212121, true},
		{2121212122, false},
		{2121212123, false},
		{2121212124, false},
	}

	testIsInvalid(t, tests, isInvalidPart2)
}

func testIsInvalid(t *testing.T, tests []struct {
	val      int64
	expected bool
}, isInvalid func(int64) bool) {
	t.Helper()

	for _, test := range tests {
		res := isInvalid(test.val)
		if res != test.expected {
			t.Errorf("Expected %t, but got %t, for value %d", test.expected, res, test.val)
		}
	}
}

func TestSumOfInvalidInRange(t *testing.T) {
	tests := []struct {
		start       int64
		end         int64
		expectedSum int64
	}{
		{11, 22, 11 + 22},
		{95, 115, 99},
		{998, 1012, 1010},
		{1188511880, 1188511890, 1188511885},
		{222220, 222224, 222222},
		{1698522, 1698528, 0},
		{446443, 446449, 446446},
		{38593856, 38593862, 38593859},
	}

	for _, test := range tests {
		res := sumOfInvalidInRange(test.start, test.end, isInvalidPart1)
		if res != test.expectedSum {
			t.Errorf("Expected %d, but got %d, for range %d-%d", test.expectedSum, res, test.start, test.end)
		}
	}
}

func TestLoopAndSumP2(t *testing.T) {
	tests := []struct {
		start       int64
		end         int64
		expectedSum int64
	}{
		{11, 22, 11 + 22},
		{95, 115, 99 + 111},
		{998, 1012, 999 + 1010},
		{1188511880, 1188511890, 1188511885},
		{222220, 222224, 222222},
		{1698522, 1698528, 0},
		{446443, 446449, 446446},
		{38593856, 38593862, 38593859},
		{565653, 565659, 565656},
		{824824821, 824824827, 824824824},
		{2121212118, 2121212124, 2121212121},
	}

	for _, test := range tests {
		res := sumOfInvalidInRange(test.start, test.end, isInvalidPart2)
		if res != test.expectedSum {
			t.Errorf("Expected %d, but got %d, for range %d-%d", test.expectedSum, res, test.start, test.end)
		}
	}
}

func TestPart2(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	res := part2(input)
	var expected int64 = 4174379265
	if expected != res {
		t.Errorf("result is not right for the sample input, got %d, want %d", res, expected)
	}
}

func TestPart1(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	res := part1(input)
	var expected int64 = 1227775554
	if expected != res {
		t.Errorf("result is not right for the sample input, got %d, want %d", res, expected)
	}
}
