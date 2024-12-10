package day10

import "fmt"

func isValidStep(mapData [][]int, x, y, prevHeight int) bool {
	return x >= 0 && y >= 0 && x < len(mapData) && y < len(mapData[0]) && mapData[x][y] == prevHeight+1
}

func calculateScore(mapData [][]int, startX, startY int) int {
	queue := []struct {
		x, y int
	}{{
		x: startX,
		y: startY,
	}}
	visited := make(map[string]bool)
	reachableNines := make(map[string]bool)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		posKey := fmt.Sprintf("%d,%d", curr.x, curr.y)
		if visited[posKey] {
			continue
		}
		visited[posKey] = true

		if mapData[curr.x][curr.y] == 9 {
			reachableNines[posKey] = true
			continue
		}

		directions := []struct{ dx, dy int }{
			{dx: 1, dy: 0}, {dx: -1, dy: 0}, {dx: 0, dy: 1}, {dx: 0, dy: -1},
		}
		for _, dir := range directions {
			newX, newY := curr.x+dir.dx, curr.y+dir.dy
			if isValidStep(mapData, newX, newY, mapData[curr.x][curr.y]) {
				queue = append(queue, struct {
					x, y int
				}{x: newX, y: newY})
			}
		}
	}

	return len(reachableNines)
}

func findDistinctTrails(mapData [][]int, x, y int, visited map[string]bool) int {
	type Position struct {
		x, y int
		path string
	}

	// BFS queue to explore paths
	queue := []Position{{x: x, y: y, path: fmt.Sprintf("%d,%d", x, y)}}
	distinctTrails := make(map[string]bool)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if mapData[curr.x][curr.y] == 9 {
			distinctTrails[curr.path] = true
			continue
		}

		directions := []struct{ dx, dy int }{
			{dx: 1, dy: 0}, {dx: -1, dy: 0}, {dx: 0, dy: 1}, {dx: 0, dy: -1},
		}

		for _, dir := range directions {
			nx, ny := curr.x+dir.dx, curr.y+dir.dy
			if nx >= 0 && ny >= 0 && nx < len(mapData) && ny < len(mapData[0]) &&
				mapData[nx][ny] == mapData[curr.x][curr.y]+1 {

				nextPath := curr.path + fmt.Sprintf("->%d,%d", nx, ny)
				if !visited[nextPath] {
					visited[nextPath] = true
					queue = append(queue, Position{x: nx, y: ny, path: nextPath})
				}
			}
		}
	}

	return len(distinctTrails)
}

func Solve(lines []string) (int, int) {

	// Parse the map into a 2D slice of integers
	mapData := make([][]int, len(lines))
	for i, line := range lines {
		mapData[i] = make([]int, len(line))
		for j, char := range line {
			mapData[i][j] = int(char - '0')
		}
	}

	totalScore := 0
	totalRating := 0

	// Iterate through the map to find trailheads and calculate scores
	for x := 0; x < len(mapData); x++ {
		for y := 0; y < len(mapData[x]); y++ {
			if mapData[x][y] == 0 { // A trailhead

				// part 1
				score := calculateScore(mapData, x, y)
				totalScore += score

				// part 2
				visited := make(map[string]bool)
				rating := findDistinctTrails(mapData, x, y, visited)
				totalRating += rating
			}
		}
	}

	return totalScore, totalRating
}
