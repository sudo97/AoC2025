package main

import (
	"fmt"
	"os"
	"sudo97/AoC2025/day1"
	"sudo97/AoC2025/day2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide a day\n")
		return
	}
	day := os.Args[1]

	switch day {
	case "day1":
		data, err := os.ReadFile(fmt.Sprintf("%s.in", day))

		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		day1.Solution(string(data))
	case "day2":
		data, err := os.ReadFile(fmt.Sprintf("%s.in", day))

		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		day2.Solution(string(data))
	}
}
