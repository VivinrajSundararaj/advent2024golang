package day13

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const PrizeOffset = 10000000000000 // Offset for Part 2

type Machine struct {
	AButtonX, AButtonY int
	BButtonX, BButtonY int
	PrizeX, PrizeY     int
}

func ParseMachinesPart2(lines []string) ([]Machine, error) {
	var machines []Machine

	for i := 0; i < len(lines); i += 4 {
		if i+2 >= len(lines) {
			return nil, fmt.Errorf("incomplete machine configuration at line %d", i+1)
		}

		// Parse Button A
		aLine := strings.TrimPrefix(lines[i], "Button A: ")
		aParts := strings.Split(aLine, ", ")
		aX, err := strconv.Atoi(strings.TrimPrefix(aParts[0], "X+"))
		if err != nil {
			return nil, fmt.Errorf("error parsing AButtonX at line %d: %v", i+1, err)
		}
		aY, err := strconv.Atoi(strings.TrimPrefix(aParts[1], "Y+"))
		if err != nil {
			return nil, fmt.Errorf("error parsing AButtonY at line %d: %v", i+1, err)
		}

		// Parse Button B
		bLine := strings.TrimPrefix(lines[i+1], "Button B: ")
		bParts := strings.Split(bLine, ", ")
		bX, err := strconv.Atoi(strings.TrimPrefix(bParts[0], "X+"))
		if err != nil {
			return nil, fmt.Errorf("error parsing BButtonX at line %d: %v", i+2, err)
		}
		bY, err := strconv.Atoi(strings.TrimPrefix(bParts[1], "Y+"))
		if err != nil {
			return nil, fmt.Errorf("error parsing BButtonY at line %d: %v", i+2, err)
		}

		// Parse Prize location
		pLine := strings.TrimPrefix(lines[i+2], "Prize: ")
		pParts := strings.Split(pLine, ", ")
		pX, err := strconv.Atoi(strings.TrimPrefix(pParts[0], "X="))
		if err != nil {
			return nil, fmt.Errorf("error parsing PrizeX at line %d: %v", i+3, err)
		}
		pY, err := strconv.Atoi(strings.TrimPrefix(pParts[1], "Y="))
		if err != nil {
			return nil, fmt.Errorf("error parsing PrizeY at line %d: %v", i+3, err)
		}

		// Apply PrizeOffset for Part 2
		pX += PrizeOffset
		pY += PrizeOffset

		machines = append(machines, Machine{
			AButtonX: aX, AButtonY: aY,
			BButtonX: bX, BButtonY: bY,
			PrizeX: pX, PrizeY: pY,
		})
	}

	return machines, nil
}

func ParseMachinesPart1(lines []string) ([]Machine, error) {
	var machines []Machine

	for i := 0; i < len(lines); i += 4 {
		// Ensure there are enough lines for a complete machine configuration
		if i+2 >= len(lines) {
			return nil, fmt.Errorf("incomplete machine configuration at line %d", i+1)
		}

		// Parse Button A configuration
		aLine := strings.TrimPrefix(lines[i], "Button A: ")
		aParts := strings.Split(aLine, ", ")
		aX, err := strconv.Atoi(strings.TrimPrefix(aParts[0], "X+"))
		if err != nil {
			return nil, fmt.Errorf("error parsing AButtonX at line %d: %v", i+1, err)
		}
		aY, err := strconv.Atoi(strings.TrimPrefix(aParts[1], "Y+"))
		if err != nil {
			return nil, fmt.Errorf("error parsing AButtonY at line %d: %v", i+1, err)
		}

		// Parse Button B configuration
		bLine := strings.TrimPrefix(lines[i+1], "Button B: ")
		bParts := strings.Split(bLine, ", ")
		bX, err := strconv.Atoi(strings.TrimPrefix(bParts[0], "X+"))
		if err != nil {
			return nil, fmt.Errorf("error parsing BButtonX at line %d: %v", i+2, err)
		}
		bY, err := strconv.Atoi(strings.TrimPrefix(bParts[1], "Y+"))
		if err != nil {
			return nil, fmt.Errorf("error parsing BButtonY at line %d: %v", i+2, err)
		}

		// Parse Prize location
		pLine := strings.TrimPrefix(lines[i+2], "Prize: ")
		pParts := strings.Split(pLine, ", ")
		pX, err := strconv.Atoi(strings.TrimPrefix(pParts[0], "X="))
		if err != nil {
			return nil, fmt.Errorf("error parsing PrizeX at line %d: %v", i+3, err)
		}
		pY, err := strconv.Atoi(strings.TrimPrefix(pParts[1], "Y="))
		if err != nil {
			return nil, fmt.Errorf("error parsing PrizeY at line %d: %v", i+3, err)
		}

		// Append machine to the list
		machines = append(machines, Machine{
			AButtonX: aX, AButtonY: aY,
			BButtonX: bX, BButtonY: bY,
			PrizeX: pX, PrizeY: pY,
		})
	}

	return machines, nil
}

// FindMinimumTokens calculates the minimum tokens required to win a prize for a given machine.
func FindMinimumTokensPart1(machine Machine) (int, bool) {
	minTokens := -1
	//bestA, bestB := 0, 0
	found := false

	// Try all combinations of A and B within bounds
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			// Check if the current combination solves both equations
			xSolved := a*machine.AButtonX + b*machine.BButtonX
			ySolved := a*machine.AButtonY + b*machine.BButtonY
			if xSolved == machine.PrizeX && ySolved == machine.PrizeY {
				tokens := a*3 + b
				if !found || tokens < minTokens {
					minTokens = tokens
					// bestA = a
					// bestB = b
					found = true
				}
			}
		}
	}

	return minTokens, found
}

// ExtendedGCD computes the GCD of a and b, and returns the coefficients (x, y) such that ax + by = gcd(a, b)
func ExtendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := ExtendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

// SolveDiophantine solves the equation A * x + B * y = target using the extended Euclidean algorithm
func SolveDiophantine(A, B, target int) (int, int, bool) {
	// Get the gcd of A and B
	gcd, x, y := ExtendedGCD(A, B)

	// If target is not divisible by gcd(A, B), then there is no solution
	if target%gcd != 0 {
		return 0, 0, false
	}

	// Scale the solution to match the target
	scale := target / gcd
	return x * scale, y * scale, true
}

// AdjustForNonNegative adjusts the solution to make x and y non-negative
func AdjustForNonNegative(x, y, A, B int) (int, int) {
	// Shift the solution to ensure both x and y are non-negative
	t := int(math.Max(float64((-x)/B), float64((-y)/A)))
	x += t * B
	y -= t * A
	return x, y
}

// FindMinimumTokens calculates the minimal tokens to win the prize for the given machine
func FindMinimumTokensPart2(machine Machine) (int, bool) {
	// Solve for X axis
	x, y, xOK := SolveDiophantine(machine.AButtonX, machine.BButtonX, machine.PrizeX)
	if !xOK {
		return 0, false
	}

	// Solve for Y axis
	yX, yY, yOK := SolveDiophantine(machine.AButtonY, machine.BButtonY, machine.PrizeY)
	if !yOK {
		return 0, false
	}

	// Ensure both x and y are non-negative
	x, y = AdjustForNonNegative(x, y, machine.AButtonX, machine.BButtonX)
	yX, yY = AdjustForNonNegative(yX, yY, machine.AButtonY, machine.BButtonY)

	fmt.Println("Print X:", x, y)
	fmt.Println("Print Y:", yX, yY)

	// Calculate the minimal tokens required
	aCost := 3
	bCost := 1

	// Calculate the number of presses required to reach the prize position
	aPresses := x
	bPresses := y

	// Ensure both x and y are non-negative
	if aPresses < 0 || bPresses < 0 {
		return 0, false
	}

	// Calculate the minimal tokens required
	tokens := aPresses*aCost + bPresses*bCost
	return tokens, true
}

// Solve the problem
func Solve(lines []string) (int, int) {
	// Parse machines
	machinesPart1, err := ParseMachinesPart1(lines)
	if err != nil {
		fmt.Println("Error parsing machines:", err)
	}

	machinesPart2, err := ParseMachinesPart2(lines)
	if err != nil {
		fmt.Println("Error parsing machines:", err)
	}

	totalPrizesPart1 := 0
	totalTokensPart1 := 0

	// Process each machine for Part1
	for _, machine := range machinesPart1 {
		minTokens, found := FindMinimumTokensPart1(machine)
		if found {
			totalPrizesPart1++
			totalTokensPart1 += minTokens
		}
	}

	totalPrizesPart2 := 0
	totalTokensPart2 := 0

	for i, machine := range machinesPart2 {
		// Calculate tokens for each machine's prize
		tokens, found := FindMinimumTokensPart2(machine)
		if found {
			// Print the result for each machine with a valid solution
			fmt.Printf("Machine %d: Tokens=%d\n", i+1, tokens)
			totalTokensPart2 += tokens
			totalPrizesPart2++
		} else {
			// If no valid solution, print no solution
			fmt.Printf("Machine %d: No solution found\n", i+1)
		}
	}

	return totalTokensPart1, totalPrizesPart2
}
