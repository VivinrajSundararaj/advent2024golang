package day04

func Solve(lines []string) (int, int) {

	word := "XMAS"
	wordLength := len(word)
	part1 := 0
	rows := len(lines)
	cols := len(lines[0])

	// Helper function to check for word in a given position (r, c) and direction (dr, dc)
	checkDirection := func(r, c, dr, dc int) bool {
		for i := 0; i < wordLength; i++ {
			newR := r + i*dr
			newC := c + i*dc
			if newR < 0 || newR >= rows || newC < 0 || newC >= cols || lines[newR][newC] != word[i] {
				return false
			}
		}
		return true
	}

	// Check all positions in the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// Check horizontal left to right
			if checkDirection(r, c, 0, 1) {
				part1++
			}
			// Check horizontal right to left
			if checkDirection(r, c, 0, -1) {
				part1++
			}
			// Check vertical top to bottom
			if checkDirection(r, c, 1, 0) {
				part1++
			}
			// Check vertical bottom to top
			if checkDirection(r, c, -1, 0) {
				part1++
			}
			// Check diagonal top-left to bottom-right
			if checkDirection(r, c, 1, 1) {
				part1++
			}
			// Check diagonal bottom-right to top-left
			if checkDirection(r, c, -1, -1) {
				part1++
			}
			// Check diagonal top-right to bottom-left
			if checkDirection(r, c, 1, -1) {
				part1++
			}
			// Check diagonal bottom-left to top-right
			if checkDirection(r, c, -1, 1) {
				part1++
			}
		}
	}
	part2 := 0
	return part1, part2
}

// func isValidPosition(grid [][]rune, x, y int) bool {
// 	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
// }

// func checkXmasPattern(grid [][]rune, x, y int) bool {
// 	// Define the directions to check the diagonals: (-1, -1), (-1, 1), (1, -1), (1, 1)
// 	directions := [][2]int{
// 		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
// 	}

// 	// Check if the position (x, y) is valid and can form an X-MAS pattern
// 	for _, dir := range directions {
// 		dx, dy := dir[0], dir[1]

// 		// Ensure the diagonals don't go out of bounds and check the pattern for M.S at appropriate positions
// 		if isValidPosition(grid, x+dx, y+dy) &&
// 			isValidPosition(grid, x-dx, y-dy) &&
// 			isValidPosition(grid, x+2*dx, y+2*dy) &&
// 			isValidPosition(grid, x-2*dx, y-2*dy) {

// 			if grid[x][y] == 'A' &&
// 				grid[x+dx][y+dy] == 'M' &&
// 				grid[x-dx][y-dy] == 'M' &&
// 				grid[x+2*dx][y+2*dy] == 'S' &&
// 				grid[x-2*dx][y-2*dy] == 'S' {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// func findXmasPatterns(grid [][]rune) int {
// 	count := 0

// 	// Iterate through every position in the grid
// 	for r := 0; r < len(grid); r++ {
// 		for c := 0; c < len(grid[0]); c++ {
// 			// We look for 'A' and check if it can form X-MAS pattern.
// 			if grid[r][c] == 'A' && checkXmasPattern(grid, r, c) {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }

// func convertToRuneGrid(grid []string) [][]rune {
// 	runeGrid := make([][]rune, len(grid))
// 	for i, line := range grid {
// 		runeGrid[i] = []rune(line)
// 	}
// 	return runeGrid
// }

// func countXMAS(gridLines []string) int {

// 	grid := convertToRuneGrid(gridLines)
// 	result := findXmasPatterns(grid)

// 	return result
// }
