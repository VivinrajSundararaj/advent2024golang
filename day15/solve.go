package day15

import (
	"fmt"
	"strings"
)

const (
	Wall  = '#'
	Empty = '.'
	Box   = 'O'
	Robot = '@'
)

type Point struct {
	X, Y int
}

// parseInput parses the input into a 2D warehouse grid, robot moves, and robot position
func parseInput(lines []string) ([][]rune, string, Point) {
	var warehouse [][]rune
	var robotPosition Point
	var moveSequence strings.Builder

	readingMap := true

	for y, line := range lines {
		if line == "" {
			readingMap = false
			continue
		}

		if readingMap {
			row := []rune(line)
			for x, ch := range row {
				if ch == Robot {
					robotPosition = Point{X: x, Y: y}
				}
			}
			warehouse = append(warehouse, row)
		} else {
			// Append moves into the moveSequence
			moveSequence.WriteString(line)
		}
	}

	return warehouse, moveSequence.String(), robotPosition
}

// moveRobot processes the robot's moves and updates the warehouse grid
func moveRobot(warehouse *[][]rune, robotPosition Point, moves string) {
	directions := map[rune]Point{
		'^': {0, -1}, // Up
		'v': {0, 1},  // Down
		'<': {-1, 0}, // Left
		'>': {1, 0},  // Right
	}

	for _, move := range moves {
		dir := directions[move]
		newRobotPosition := add(robotPosition, dir)

		if canMove(warehouse, newRobotPosition) {
			// Move the robot and possibly push a box
			if (*warehouse)[newRobotPosition.Y][newRobotPosition.X] == Box {
				// Check if the robot can push the box
				boxNewPosition := add(newRobotPosition, dir)
				if canMoveBox(warehouse, boxNewPosition) {
					// Move the box
					(*warehouse)[boxNewPosition.Y][boxNewPosition.X] = Box
					// Move the robot
					(*warehouse)[newRobotPosition.Y][newRobotPosition.X] = Robot
					(*warehouse)[robotPosition.Y][robotPosition.X] = Empty
					robotPosition = newRobotPosition
				}
			} else {
				// Move only the robot
				(*warehouse)[newRobotPosition.Y][newRobotPosition.X] = Robot
				(*warehouse)[robotPosition.Y][robotPosition.X] = Empty
				robotPosition = newRobotPosition
			}
		}
	}
}

// add adds two points together
func add(a, b Point) Point {
	return Point{X: a.X + b.X, Y: a.Y + b.Y}
}

// canMove checks if the robot can move to a new position
func canMove(warehouse *[][]rune, target Point) bool {
	if target.Y < 0 || target.Y >= len(*warehouse) || target.X < 0 || target.X >= len((*warehouse)[0]) {
		return false // Out of bounds
	}
	targetCell := (*warehouse)[target.Y][target.X]
	return targetCell == Empty || targetCell == Box
}

// canMoveBox checks if a box can be pushed to the given position
func canMoveBox(warehouse *[][]rune, target Point) bool {
	if target.Y < 0 || target.Y >= len(*warehouse) || target.X < 0 || target.X >= len((*warehouse)[0]) {
		return false // Out of bounds
	}
	return (*warehouse)[target.Y][target.X] == Empty
}

// calculateGPSSum computes the total GPS sum of all boxes in the warehouse
func calculateGPSSum(warehouse [][]rune) int {
	total := 0
	for y, row := range warehouse {
		for x, cell := range row {
			if cell == Box {
				// Correct GPS calculation as 100 * y + x
				total += 100*y + x
			}
		}
	}
	return total
}

// Solve the problem
func Solve(lines []string) (int, int) {
	// Step 2: Parse the input into the warehouse map, robot moves, and robot position
	warehouse, moveSequence, robotPosition := parseInput(lines)

	// Step 3: Move the robot and simulate the pushes
	moveRobot(&warehouse, robotPosition, moveSequence)

	// Step 4: Calculate and print the total GPS sum for all boxes
	totalSum := calculateGPSSum(warehouse)
	fmt.Println("Total GPS Sum:", totalSum)
	return totalSum, 0
}
