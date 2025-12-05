package day1

import "testing"

func TestParting(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult entry
	}{
		{"L123", entry{LEFT, 123}},
		{"R321", entry{RIGHT, 321}},
	}

	for _, testCase := range testCases {
		line, err := parseLine(testCase.input)
		if err != nil {
			t.Errorf("Should parse without errors, instead got %s", err)
			continue
		}
		if line.direction != testCase.expectedResult.direction {
			t.Errorf("Should parse as %s for the input %s, instead got %s", testCase.expectedResult.direction, testCase.input, line.direction)
		}
		if line.num != testCase.expectedResult.num {
			t.Errorf("Should parse as %d for the input: %s, instead got %d", testCase.expectedResult.num, testCase.input, line.num)
		}
	}

	_, err := parseLine("U123")
	if err == nil {
		t.Errorf("Should return error when direction is unknown")
	}
	_, err = parseLine("RABC")
	if err == nil {
		t.Errorf("Should return error when number of turns is not a number")
	}
}

func TestPart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	expected := 3

	res := part1(input)

	if res != expected {
		t.Errorf("expected %d, but got %d", expected, res)
	}
}

func TestPart2(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	expected := 6

	res := part2(input)

	if res != expected {
		t.Errorf("expected %d, but got %d", expected, res)
	}
}

func Example() {
	sample := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	Solution(sample)
	// Output:
	// day 1, part 1 is 3
	// day 1, part 2 is 6
}
