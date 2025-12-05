package day4

import (
	"bufio"
	"fmt"
	"io"
)

func HasLessThanFourNeighbors(matrix [][]bool, i, j int) bool {
	count := 0

	arr := [8]struct {
		x int
		y int
	}{
		{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1},
		{i, j - 1}, {i, j + 1},
		{i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1},
	}

	for _, v := range arr {
		if v.x < 0 || v.x >= len(matrix) {
			continue
		}
		if v.y < 0 || v.y >= len(matrix[v.x]) {
			continue
		}
		if matrix[v.x][v.y] {
			count++
		}
	}

	return count < 4
}

func parse(r io.Reader) [][]bool {
	scanner := bufio.NewScanner(r)

	result := [][]bool{}
	// lineNumber := 0
	for scanner.Scan() {
		// lineNumber++
		line := scanner.Text()
		row := make([]bool, len(line))
		for i, c := range line {
			if c == '@' {
				row[i] = true
			}
		}
		result = append(result, row)
	}

	return result
}

func itemsToDelete(matrix [][]bool) [][2]int {
	itemsToDelete := [][2]int{}
	for i, row := range matrix {
		for j, c := range row {
			if c && HasLessThanFourNeighbors(matrix, i, j) {
				itemsToDelete = append(itemsToDelete, [2]int{i, j})
			}
		}
	}
	return itemsToDelete
}

func part1(matrix [][]bool) int {
	return len(itemsToDelete(matrix))
}

func part2(matrix [][]bool) int {
	count := 0

	for {
		itemsToDelete := itemsToDelete(matrix)
		if len(itemsToDelete) == 0 {
			break
		}
		count += len(itemsToDelete)
		for _, item := range itemsToDelete {
			matrix[item[0]][item[1]] = false
		}
	}

	return count
}

func Solution(r io.Reader) {
	data := parse(r)
	fmt.Printf("day4, part1: %d\n", part1(data))
	fmt.Printf("day4, part2: %d\n", part2(data))
}
