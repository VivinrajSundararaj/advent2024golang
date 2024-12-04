package day01

import (
	"fmt"
	"sort"
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

func SolvePart1(day1Input01 []int, day1Input02 []int) int {

	// Check if lists are of the same length
	if len(day1Input01) != len(day1Input02) {
		fmt.Println("Error: The two lists must have the same number of elements.")
		return 0
	}

	// Calculate total distance
	totalDistance := calculateDistance(day1Input01, day1Input02)

	return totalDistance
}
