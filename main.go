package main

import (
	"fmt"

	"github.com/VivinrajSundararaj/advent2024golang/day01"
	"github.com/VivinrajSundararaj/advent2024golang/utils"
)

func main() {
	fmt.Println("--- DAY 01 ---")
	day1Input01, day1Input02 := utils.ReadFile("day01/input.txt")
	fmt.Println("Result Part 1:", day01.SolvePart1(day1Input01, day1Input02))
}
