package day12

import "strings"

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Directions: up, down, left, right

// Flood fill function to explore each region
func floodFill(grid [][]rune, visited [][]bool, x, y int, plant rune) (int, int) {
	rows, cols := len(grid), len(grid[0])
	area := 0
	perimeter := 0
	stack := [][2]int{{x, y}}

	for len(stack) > 0 {
		cx, cy := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]

		if visited[cx][cy] {
			continue
		}
		visited[cx][cy] = true
		area++

		// For each plot, check the four directions (up, down, left, right)
		for _, dir := range directions {
			nx, ny := cx+dir[0], cy+dir[1]
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] != plant {
				// Boundary or different plant type contributes to perimeter
				perimeter++
			} else if !visited[nx][ny] {
				// If it is a valid plot of the same type and not yet visited, add to stack
				stack = append(stack, [2]int{nx, ny})
			}
		}
	}
	return area, perimeter
}

// Function to calculate the fencing cost for Part 1 (based on perimeter)
func calculateFencingCost(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	totalCost := 0

	// Process each plot in the grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				// Calculate the area and perimeter for Part 1
				area, perimeter := floodFill(grid, visited, i, j, grid[i][j])
				totalCost += area * perimeter
			}
		}
	}

	return totalCost
}

// Parse input into a grid of garden plots
func parseInput(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(strings.TrimSpace(line))
	}
	return grid
}

// Solve the problem
func Solve(lines []string) (int, int) {
	grid := parseInput(lines)
	totalCostPart1 := calculateFencingCost(grid)
	return totalCostPart1, 0
}
