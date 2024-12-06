package day06

var directions = [][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func Solve(lines []string) (int, int) {
	rows := len(lines)
	cols := len(lines[0])

	// Find the initial position and direction of the guard
	var guardX, guardY, dir int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			switch lines[r][c] {
			case '^':
				guardX, guardY, dir = r, c, 0
			case '>':
				guardX, guardY, dir = r, c, 1
			case 'v':
				guardX, guardY, dir = r, c, 2
			case '<':
				guardX, guardY, dir = r, c, 3
			}
		}
	}

	// Set to track visited positions
	visited := make(map[[2]int]bool)
	visited[[2]int{guardX, guardY}] = true

	for {
		// Calculate the next position in the current direction
		nextX := guardX + directions[dir][0]
		nextY := guardY + directions[dir][1]

		// Check if the next position is out of bounds
		if nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols {
			break
		}

		// Check if the next position is blocked
		if lines[nextX][nextY] == '#' {
			// Turn right (90 degrees clockwise)
			dir = (dir + 1) % 4
		} else {
			// Move forward
			guardX, guardY = nextX, nextY
			visited[[2]int{guardX, guardY}] = true
		}
	}
	part2 := SolvePart2(lines)
	return len(visited), part2
}

func simulateWithObstruction(grid []string, obstructionX, obstructionY int) (bool, map[[2]int]bool) {
	rows := len(grid)
	cols := len(grid[0])

	// Copy grid to avoid modifying the original
	modifiedGrid := make([][]rune, rows)
	for i := range grid {
		modifiedGrid[i] = []rune(grid[i])
	}

	// Place obstruction
	modifiedGrid[obstructionX][obstructionY] = '#'

	// Find the initial position and direction of the guard
	var guardX, guardY, dir int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			switch modifiedGrid[r][c] {
			case '^':
				guardX, guardY, dir = r, c, 0
			case '>':
				guardX, guardY, dir = r, c, 1
			case 'v':
				guardX, guardY, dir = r, c, 2
			case '<':
				guardX, guardY, dir = r, c, 3
			}
		}
	}

	// Set to track visited positions
	visited := make(map[[2]int]bool)
	visited[[2]int{guardX, guardY}] = true

	// Track if a loop is detected
	path := make([][3]int, 0) // Stores position and direction

	for {
		// Calculate the next position in the current direction
		nextX := guardX + directions[dir][0]
		nextY := guardY + directions[dir][1]

		// Check if the next position is out of bounds
		if nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols {
			return false, visited
		}

		// Check if the next position is blocked
		if modifiedGrid[nextX][nextY] == '#' {
			// Turn right (90 degrees clockwise)
			dir = (dir + 1) % 4
		} else {
			// Move forward
			guardX, guardY = nextX, nextY
			pos := [3]int{guardX, guardY, dir}

			// Check if we've already seen this position and direction
			for _, p := range path {
				if p == pos {
					return true, visited // Loop detected
				}
			}

			// Add to path and visited
			path = append(path, pos)
			visited[[2]int{guardX, guardY}] = true
		}
	}
}

func SolvePart2(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])

	// Find the guard's starting position
	var startX, startY int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '^' || grid[r][c] == '>' || grid[r][c] == 'v' || grid[r][c] == '<' {
				startX, startY = r, c
			}
		}
	}

	validPositions := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// Skip if it's already blocked, out of bounds, or the starting position
			if grid[r][c] == '#' || (r == startX && c == startY) {
				continue
			}

			// Simulate with obstruction
			if loopDetected, _ := simulateWithObstruction(grid, r, c); loopDetected {
				validPositions++
			}
		}
	}

	return validPositions
}
