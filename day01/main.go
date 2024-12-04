package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	lines, err := utils.ReadLines("day01/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	sum := 0
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Invalid number: %v", err)
		}
		sum += num
	}

	fmt.Println("Sum of numbers:", sum)
}
