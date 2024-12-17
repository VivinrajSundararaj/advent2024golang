package day15

import (
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

// moveRobot processes the robot's moves and updates the warehouse grid.
func moveRobot(warehouse *[][]rune, robotPosition *Point, moves string) {
	// Direction mappings
	directions := map[rune]Point{
		'^': {0, -1}, // Up
		'v': {0, 1},  // Down
		'<': {-1, 0}, // Left
		'>': {1, 0},  // Right
	}

	for _, move := range moves {
		dir := directions[move]
		newRobotPosition := add(*robotPosition, dir)

		// Check if the robot can move to the new position
		if canMove(warehouse, newRobotPosition) {
			// If the space is empty, just move the robot
			if (*warehouse)[newRobotPosition.Y][newRobotPosition.X] == Empty {
				(*warehouse)[newRobotPosition.Y][newRobotPosition.X] = Robot
				(*warehouse)[(*robotPosition).Y][(*robotPosition).X] = Empty
				*robotPosition = newRobotPosition
			} else if (*warehouse)[newRobotPosition.Y][newRobotPosition.X] == Box {
				// If the robot moves into a box, check if it can push the box or a chain of boxes
				if canMoveSingleBox(warehouse, newRobotPosition, dir) {
					// Move the single box one step forward
					boxPos := newRobotPosition
					nextPos := add(boxPos, dir)
					(*warehouse)[nextPos.Y][nextPos.X] = Box
					(*warehouse)[boxPos.Y][boxPos.X] = Empty

					// The robot moves to the position where the box was
					(*warehouse)[boxPos.Y][boxPos.X] = Robot
					(*warehouse)[(*robotPosition).Y][(*robotPosition).X] = Empty
					*robotPosition = newRobotPosition
				} else if canMoveMultipleBoxes(warehouse, newRobotPosition, dir) {
					// Move multiple consecutive boxes forward
					boxesToMove := []Point{newRobotPosition}
					currentPos := newRobotPosition
					// Collect all the consecutive boxes
					for {
						currentPos = add(currentPos, dir)
						if currentPos.Y < 0 || currentPos.Y >= len(*warehouse) || currentPos.X < 0 || currentPos.X >= len((*warehouse)[0]) {
							break
						}
						if (*warehouse)[currentPos.Y][currentPos.X] != Box {
							break
						}
						boxesToMove = append(boxesToMove, currentPos)
					}

					// Move the boxes in the chain one step forward
					for i := len(boxesToMove) - 1; i >= 0; i-- {
						boxPos := boxesToMove[i]
						// Move the box one step in the direction
						nextPos := add(boxPos, dir)
						(*warehouse)[nextPos.Y][nextPos.X] = Box
						(*warehouse)[boxPos.Y][boxPos.X] = Empty
					}

					// After moving all boxes, the robot moves to the position of the first box
					(*warehouse)[newRobotPosition.Y][newRobotPosition.X] = Robot
					(*warehouse)[(*robotPosition).Y][(*robotPosition).X] = Empty
					*robotPosition = newRobotPosition
				}
			}
		}
	}
}

// add adds two Points (coordinates).
func add(p1, p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

// canMove checks if the robot can move to the target position (either empty or box).
func canMove(warehouse *[][]rune, target Point) bool {
	// Check if the target is within bounds and is either empty or a box
	if target.Y < 0 || target.Y >= len(*warehouse) || target.X < 0 || target.X >= len((*warehouse)[0]) {
		return false // Out of bounds
	}
	return (*warehouse)[target.Y][target.X] == Empty || (*warehouse)[target.Y][target.X] == Box
}

// canMoveSingleBox checks if the robot can move a single box and if there's space behind it.
func canMoveSingleBox(warehouse *[][]rune, boxPosition Point, target Point) bool {
	// Check if the position behind the box is empty
	spaceBehindBox := add(boxPosition, target)
	if spaceBehindBox.Y < 0 || spaceBehindBox.Y >= len(*warehouse) || spaceBehindBox.X < 0 || spaceBehindBox.X >= len((*warehouse)[0]) {
		return false // Out of bounds
	}

	// Ensure that the space behind the box is empty or contains the robot
	if (*warehouse)[spaceBehindBox.Y][spaceBehindBox.X] == Empty || (*warehouse)[spaceBehindBox.Y][spaceBehindBox.X] == Robot {
		return true
	}

	return false
}

// canMoveMultipleBoxes checks if the robot can move multiple consecutive boxes and if there is enough space behind all boxes.
func canMoveMultipleBoxes(warehouse *[][]rune, firstBoxPosition Point, direction Point) bool {
	// Check for a chain of boxes in the given direction
	currentPos := firstBoxPosition
	boxesToMove := []Point{currentPos}

	// Collect all the consecutive boxes in the chain
	for {
		// Move to the next position in the direction
		currentPos = add(currentPos, direction)
		if currentPos.Y < 0 || currentPos.Y >= len(*warehouse) || currentPos.X < 0 || currentPos.X >= len((*warehouse)[0]) {
			break
		}
		if (*warehouse)[currentPos.Y][currentPos.X] != Box {
			break
		}
		boxesToMove = append(boxesToMove, currentPos)
	}

	// Check if there is enough space after the last box in the chain
	lastBoxPosition := boxesToMove[len(boxesToMove)-1]
	spaceAfterLastBox := add(lastBoxPosition, direction)

	// Directly return the result of canMove check
	return canMove(warehouse, spaceAfterLastBox)
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

// func printWarehouse(warehouse [][]rune) {
// 	for _, row := range warehouse {
// 		fmt.Println(string(row))
// 	}
// }

// Solve the problem
func Solve(lines []string) (int, int) {
	// Part1
	warehouse, moveSequence, robotPosition := parseInput(lines)

	moveRobot(&warehouse, &robotPosition, moveSequence)

	// fmt.Println("\nFinal state:")
	// printWarehouse(warehouse)

	// Step 4: Calculate and print the total GPS sum for all boxes
	totalSum := calculateGPSSum(warehouse)
	return totalSum, 0
}
