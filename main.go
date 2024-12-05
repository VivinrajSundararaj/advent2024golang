package main

import (
	"fmt"
	"os"

	"github.com/VivinrajSundararaj/advent2024golang/day01"
	"github.com/VivinrajSundararaj/advent2024golang/day02"
	"github.com/VivinrajSundararaj/advent2024golang/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}
	day := os.Args[1]
	fmt.Printf("\n--- DAY %s ---\n", day)
	part1, part2 := 0, 0
	switch day {
	case "01":
		dayInput01, dayInput02 := utils.ReadFile("day" + day + "/input.txt")
		part1, part2 = day01.Solve(dayInput01, dayInput02)
	case "02":
		lines, err := utils.ReadLines("day" + day + "/input.txt")
		if err != nil {
			fmt.Println("Error: Issue encountered while parsing the input.")
			os.Exit(1)
		}
		part1, part2 = day02.Solve(lines)
	default:
		fmt.Printf("Day %s is not implemented yet.\n", day)
	}
	fmt.Printf("\n--- RESULT FOR DAY %s ---\n", day)
	fmt.Printf("Part 1: %d \n", part1)
	fmt.Printf("Part 2: %d \n", part2)
}
