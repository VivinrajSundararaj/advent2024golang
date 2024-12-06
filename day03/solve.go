package day03

import (
	"fmt"
	"regexp"
	"strconv"
)

// Declare the regular expressions for mul, do(), and don't() patterns
var mulPattern = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var doPattern = regexp.MustCompile(`do\(\)`)
var dontPattern = regexp.MustCompile(`don't\(\)`)

func Solve(lines []string) (int, int) {

	// Variable to keep track of the total sum of multiplications
	totalPart1 := 0

	// Iterate through each line in the input
	for _, line := range lines {
		// Loop to process each valid mul(x, y)
		matches := mulPattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			// Extract the two numbers and compute the product
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				// Add the result of the multiplication to the total
				totalPart1 += num1 * num2
			}
		}
	}

	totalPart2 := SolvePart2(lines)
	// Return the total for Part 1
	return totalPart1, totalPart2

}

// SolvePart2 function that handles Part 2 logic
func SolvePart2(lines []string) int {
	// Variable to store the sum of valid multiplications
	totalPart2 := 0
	mulEnabled := true // Start with multiplications enabled

	// Iterate over each line of input
	for _, line := range lines {
		pos := 0        // Starting position in the line
		currentPos := 0 // current position in the line

		// Process the line character by character
		for pos <= len(line) {

			// Only check do() and don't() after the first mul(X, Y) is found
			if mulPattern.MatchString(line[currentPos:pos]) {

				// Check if we encounter a mul(X,Y) operation
				match := mulPattern.FindStringSubmatch(line[currentPos:pos])

				if match != nil {
					num1, _ := strconv.Atoi(match[1])
					num2, _ := strconv.Atoi(match[2])

					if mulEnabled {
						lineLength := line[currentPos:pos]
						fmt.Println(lineLength)
						totalPart2 += num1 * num2
					}
					// Move the current position to the end of the current match
					currentPos = pos
				}
			} else if doPattern.MatchString(line[currentPos:pos]) {
				// Enable multiplications after do()
				mulEnabled = true
				// Move the current position to the end of the current match
				currentPos = pos
			} else if dontPattern.MatchString(line[currentPos:pos]) {
				// Disable multiplications after don't()
				mulEnabled = false
				// Move the current position to the end of the current match
				currentPos = pos
			} else {
				// If no matching pattern is found, move to the next character
				pos++
			}
		}
	}
	// Return the result of all valid multiplications
	return totalPart2
}
