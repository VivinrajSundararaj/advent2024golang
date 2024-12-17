package day14

import (
	"math"
	"strconv"
	"strings"
)

const (
	Width    = 101
	Height   = 103
	Seconds  = 100
	GridMidX = Width / 2
	GridMidY = Height / 2
)

// Robot represents a robot's position and velocity
type Robot struct {
	PosX, PosY           int
	VelocityX, VelocityY int
}

func ParseInput(lines []string) []Robot {
	var robots []Robot
	for _, line := range lines {
		// Example input: p=0,4 v=3,-3
		parts := strings.Split(line, " ")
		posPart := strings.TrimPrefix(parts[0], "p=")
		velPart := strings.TrimPrefix(parts[1], "v=")

		posParts := strings.Split(posPart, ",")
		velParts := strings.Split(velPart, ",")

		posX, _ := strconv.Atoi(posParts[0])
		posY, _ := strconv.Atoi(posParts[1])
		velX, _ := strconv.Atoi(velParts[0])
		velY, _ := strconv.Atoi(velParts[1])

		robots = append(robots, Robot{
			PosX: posX, PosY: posY,
			VelocityX: velX, VelocityY: velY,
		})
	}
	return robots
}

// SimulateRobots simulates robot positions after 'Seconds' time with edge wrapping
func SimulateRobots(robots []Robot, width, height, seconds int) [][]int {
	// Create a grid to track robot counts
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	// Simulate each robot's motion
	for _, robot := range robots {
		// Calculate the robot's position after 'seconds' time
		newX := (robot.PosX + robot.VelocityX*seconds) % width
		newY := (robot.PosY + robot.VelocityY*seconds) % height

		// Handle negative modulo by adding grid dimensions
		if newX < 0 {
			newX += width
		}
		if newY < 0 {
			newY += height
		}

		// Increment the grid position
		grid[newY][newX]++
	}

	return grid
}

// CountQuadrants counts robots in each quadrant and computes the safety factor
func CountQuadrants(grid [][]int, width, height int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			count := grid[y][x]

			// Skip empty tiles
			if count == 0 {
				continue
			}

			// Determine the quadrant, excluding the middle lines
			if x > GridMidX && y < GridMidY {
				q1 += count // Top-right
			} else if x < GridMidX && y < GridMidY {
				q2 += count // Top-left
			} else if x < GridMidX && y > GridMidY {
				q3 += count // Bottom-left
			} else if x > GridMidX && y > GridMidY {
				q4 += count // Bottom-right
			}
		}
	}

	// Calculate the safety factor as the product of quadrant counts
	return q1 * q2 * q3 * q4
}

func AdvanceTime(robots []Robot, seconds int) []Robot {
	advanced := make([]Robot, len(robots))
	for i, r := range robots {
		advanced[i] = Robot{
			PosX:      r.PosX + r.VelocityX*seconds,
			PosY:      r.PosY + r.VelocityY*seconds,
			VelocityX: r.VelocityX,
			VelocityY: r.VelocityY,
		}
	}
	return advanced
}

// BoundingBox computes the bounding box around the robots
func BoundingBox(robots []Robot) (int, int, int, int) {
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt

	for _, r := range robots {
		if r.PosX < minX {
			minX = r.PosX
		}
		if r.PosX > maxX {
			maxX = r.PosX
		}
		if r.PosY < minY {
			minY = r.PosY
		}
		if r.PosY > maxY {
			maxY = r.PosY
		}
	}
	return minX, maxX, minY, maxY
}

// Area calculates the area of the bounding box
func Area(minX, maxX, minY, maxY int) int {
	return (maxX - minX + 1) * (maxY - minY + 1)
}

// Solve the problem
func Solve(lines []string) (int, int) {

	// Parse input to get robot positions and velocities
	robots := ParseInput(lines)

	// Simulate robots after the specified time
	grid := SimulateRobots(robots, Width, Height, Seconds)

	// Count robots in quadrants and calculate the safety factor
	safetyFactor := CountQuadrants(grid, Width, Height)

	// Find the smallest christmas tree
	minArea := math.MaxInt
	bestTime := 0

	for t := 0; t < 100000; t++ { // Larger upper limit
		advancedRobots := AdvanceTime(robots, t)
		minX, maxX, minY, maxY := BoundingBox(advancedRobots)
		area := Area(minX, maxX, minY, maxY)

		if area < minArea {
			minArea = area
			bestTime = t
		}
	}
	return safetyFactor, bestTime
}
