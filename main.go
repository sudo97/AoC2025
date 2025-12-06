package main

import (
	"fmt"
	"os"
	"sudo97/AoC2025/day1"
	"sudo97/AoC2025/day2"
	"sudo97/AoC2025/day3"
	"sudo97/AoC2025/day4"
	"sudo97/AoC2025/day5"
	"sudo97/AoC2025/day6"
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
	case "day3":
		file, err := os.Open(fmt.Sprintf("%s.in", day))
		defer file.Close()
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		day3.Solution(file)
	case "day4":
		file, err := os.Open(fmt.Sprintf("%s.in", day))
		defer file.Close()
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		day4.Solution(file)
	case "day5":
		file, err := os.Open(fmt.Sprintf("%s.in", day))
		defer file.Close()
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		day5.Solution(file)
	case "day6":
		file, err := os.Open(fmt.Sprintf("%s.in", day))
		defer file.Close()
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		day6.Solution(file)
	}
}
