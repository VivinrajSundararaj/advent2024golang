package main

import (
	"fmt"
	"os"

	"github.com/VivinrajSundararaj/advent2024golang/day01"
	"github.com/VivinrajSundararaj/advent2024golang/day02"
	"github.com/VivinrajSundararaj/advent2024golang/day03"
	"github.com/VivinrajSundararaj/advent2024golang/day04"
	"github.com/VivinrajSundararaj/advent2024golang/day06"
	"github.com/VivinrajSundararaj/advent2024golang/day10"
	"github.com/VivinrajSundararaj/advent2024golang/day11"
	"github.com/VivinrajSundararaj/advent2024golang/day12"
	"github.com/VivinrajSundararaj/advent2024golang/day13"
	"github.com/VivinrajSundararaj/advent2024golang/day14"
	"github.com/VivinrajSundararaj/advent2024golang/day15"
	"github.com/VivinrajSundararaj/advent2024golang/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day := os.Args[1]

	lines, err := utils.ReadLines("day" + day + "/input.txt")
	if err != nil {
		fmt.Println("Error: Issue encountered while parsing the input.")
		os.Exit(1)
	}

	// Map of day numbers to their corresponding Solve functions
	solvers := map[string]func([]string) (int, int){
		"01": day01.Solve,
		"02": day02.Solve,
		"03": day03.Solve,
		"04": day04.Solve,
		"06": day06.Solve,
		"10": day10.Solve,
		"11": day11.Solve,
		"12": day12.Solve,
		"13": day13.Solve,
		"14": day14.Solve,
		"15": day15.Solve,
	}

	// Lookup the appropriate Solve function based on the day
	solveFunc, found := solvers[day]
	if !found {
		fmt.Printf("Error: No solution implemented for day %s.\n", day)
		os.Exit(1)
	}

	// Call the Solve function and display results
	part1, part2 := solveFunc(lines)

	fmt.Printf("\n--- RESULT FOR DAY %s ---\n", day)
	fmt.Printf("Part 1: %d \n", part1)
	fmt.Printf("Part 2: %d \n", part2)
}
