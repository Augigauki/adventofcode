package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day6v2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Last line contains operators
	operatorLine := lines[len(lines)-1]
	numberLines := lines[:len(lines)-1]

	// Find operator positions and types
	var opPositions []int
	var operators []rune

	for i, char := range operatorLine {
		if char == '*' || char == '+' {
			opPositions = append(opPositions, i)
			operators = append(operators, char)
		}
	}

	total := 0

	// Find max line length to know where data ends
	maxLen := 0
	for _, line := range numberLines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// Process each operator
	for opIdx := 0; opIdx < len(opPositions); opIdx++ {
		// Determine column range: start at operator position, end before next operator (or end of line)
		startCol := opPositions[opIdx] // Include the operator column!
		endCol := maxLen               // Go up to maxLen (inclusive)
		if opIdx < len(opPositions)-1 {
			endCol = opPositions[opIdx+1] // Or up to next operator (exclusive)
		}

		fmt.Printf("\nOperator '%c' at position %d (columns %d to %d)\n",
			operators[opIdx], opPositions[opIdx], startCol, endCol-1)

		// Extract numbers by reading vertically in each column
		var numbers []int
		isLastOp := (opIdx == len(opPositions)-1)
		endLimit := endCol
		if !isLastOp {
			endLimit = endCol - 1 // Don't include next operator column
		}
		for col := startCol; col <= endLimit; col++ {
			numStr := ""
			for row := 0; row < len(numberLines); row++ {
				if col < len(numberLines[row]) && numberLines[row][col] != ' ' {
					numStr += string(numberLines[row][col])
				}
			}
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, num)
				fmt.Printf("  Col %d: %s -> %d\n", col, numStr, num)
			}
		}

		// Apply operator
		result := 0
		if operators[opIdx] == '+' {
			for _, n := range numbers {
				result += n
			}
		} else { // '*'
			result = 1
			for _, n := range numbers {
				result *= n
			}
		}

		fmt.Printf("  Numbers: %v, Result: %d\n", numbers, result)
		total += result
	}

	fmt.Println("Final Score:", total)
}
