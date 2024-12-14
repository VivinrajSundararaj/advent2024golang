package day11

import (
	"strconv"
	"strings"
)

// Split number into two halves for even-length stones
func splitNumber(num int) (int, int) {
	str := strconv.Itoa(num)
	mid := len(str) / 2
	left, _ := strconv.Atoi(str[:mid])
	right, _ := strconv.Atoi(str[mid:])
	return left, right
}

// Count stones without full simulation
func countStones(stones []int, blinks int) int {
	memo := make(map[int]int) // Cache for individual stone growth

	var computeGrowth func(stone, blinksLeft int) int
	computeGrowth = func(stone, blinksLeft int) int {
		// Base case: no more blinks
		if blinksLeft == 0 {
			return 1
		}

		// Check if result is cached
		key := stone*1000 + blinksLeft // Unique key for stone+blinks
		if val, exists := memo[key]; exists {
			return val
		}

		// Apply transformation rules
		var count int
		if stone == 0 {
			count = computeGrowth(1, blinksLeft-1) // Rule 1
		} else if len(strconv.Itoa(stone))%2 == 0 {
			left, right := splitNumber(stone) // Rule 2
			count = computeGrowth(left, blinksLeft-1) + computeGrowth(right, blinksLeft-1)
		} else {
			count = computeGrowth(stone*2024, blinksLeft-1) // Rule 3
		}

		// Cache the result
		memo[key] = count
		return count
	}

	// Compute total stones
	totalCount := 0
	for _, stone := range stones {
		totalCount += computeGrowth(stone, blinks)
	}

	return totalCount
}

func parseInput(lines []string) []int {
	var stones []int
	for _, numStr := range strings.Fields(lines[0]) {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			stones = append(stones, num)
		}
	}
	return stones
}

func Solve(lines []string) (int, int) {

	stones := parseInput(lines)
	part1Blinks := 25
	part2Blinks := 75

	part1Stones := countStones(stones, part1Blinks)
	part2Stones := countStones(stones, part2Blinks)

	return part1Stones, part2Stones

}
