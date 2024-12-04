package main

import (
	"fmt"

	"github.com/VivinrajSundararaj/advent2024golang/day01"
	"github.com/VivinrajSundararaj/advent2024golang/utils"
)

func main() {
	fmt.Println("--- DAY 01 ---")
	day1Input01, day1Input02 := utils.ReadFile("day01/input.txt")
	part1, part2 := day01.Solve(day1Input01, day1Input02)
	fmt.Printf("Result - Part 1: %d, Part 2: %d ", part1, part2)
}
