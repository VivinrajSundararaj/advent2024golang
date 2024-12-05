package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Parse a single line of levels into a slice of integers
func parseLevels(line string) []int {
	parts := strings.Fields(line)
	levels := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Error parsing number: %s\n", part)
			os.Exit(1)
		}
		levels[i] = num
	}
	return levels
}

// Check if a report is safe
func isSafe(report []int) bool {
	if len(report) < 2 {
		return false // A safe report must have at least 2 levels
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		// Check if the difference is within the allowed range [1, 3]
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
		// Update increasing/decreasing flags
		if diff > 0 {
			isDecreasing = false
		}
		if diff < 0 {
			isIncreasing = false
		}
	}

	// The report is safe if it's strictly increasing or strictly decreasing
	return isIncreasing || isDecreasing
}

// Check if a report can be made safe by removing one level
func isSafeWithDampener(report []int) bool {

	// Try removing each level one by one and check if the remaining report is safe
	for i := 0; i < len(report); i++ {
		// Create a copy of the report without the i-th level
		modifiedReport := append([]int{}, report[:i]...)         // Copy elements before i
		modifiedReport = append(modifiedReport, report[i+1:]...) // Add elements after i
		if isSafe(modifiedReport) {
			return true
		}
	}

	// If no single removal makes the report safe, return false
	return false
}

func Solve(lines []string) (int, int) {
	safeCount := 0
	safeWithDampenerCount := 0

	for _, line := range lines {
		report := parseLevels(line)

		// Check safety with and without dampener
		if isSafe(report) {
			safeCount++
			safeWithDampenerCount++ // If already safe, count it for both
		} else if isSafeWithDampener(report) {
			safeWithDampenerCount++ // Only count for dampener if not inherently safe
		}
	}

	return safeCount, safeWithDampenerCount
}
