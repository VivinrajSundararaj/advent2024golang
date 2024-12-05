package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func calculateDistance(left, right []int) int {
	// Sort both lists to align numbers in increasing order
	sort.Ints(left)
	sort.Ints(right)

	// Initialize total distance
	totalDistance := 0

	// Calculate the distance for each pair
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff // Take absolute difference
		}
		totalDistance += diff
	}

	return totalDistance
}

func similarityScore(left, right []int) int {
	var similarityScore int

	for _, g1Val := range left {
		g2Count := 0

		for _, g2Val := range right {
			if g1Val == g2Val {
				g2Count++
			}
		}

		elementSimilarity := g1Val * g2Count
		similarityScore += elementSimilarity
	}

	return similarityScore
}

// ConvertToIntArrays takes lines of string input and splits them into two integer arrays.
func ConvertToIntArrays(lines []string) ([]int, []int, error) {
	var left []int
	var right []int

	for _, line := range lines {
		// Split the line into parts based on whitespace
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		// Convert the parts to integers
		leftVal, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in line: %s", parts[0])
		}
		rightVal, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in line: %s", parts[1])
		}

		// Append to the respective slices
		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	return left, right, nil
}

func Solve(lines []string) (int, int) {

	// Convert the input lines to integer arrays

	left, right, err := ConvertToIntArrays(lines)
	if err != nil {
		fmt.Printf("Error processing lines: %v\n", err)
		return 0, 0
	}

	// Check if lists are of the same length
	if len(left) != len(right) {
		fmt.Println("Error: The two lists must have the same number of elements.")
		return 0, 0
	}

	// Part 1: Calculate total distance
	totalDistance := calculateDistance(left, right)
	fmt.Printf("Total Distance: %d\n", totalDistance)

	// Part 2: similarity score
	similarityScore := similarityScore(left, right)
	fmt.Printf("Similarity score: %d\n", similarityScore)

	return totalDistance, similarityScore
}
