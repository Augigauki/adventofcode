package main

import "fmt"

func findAccessiblePaper(paperMap [][]string) int {
	fmt.Println("Paper Map:")
	for _, line := range paperMap {
		fmt.Println(line)
	}
	cloneMap := make([][]string, len(paperMap))
	for i := range paperMap {
		cloneMap[i] = make([]string, len(paperMap[i]))
		copy(cloneMap[i], paperMap[i])
	}
	var dimensions = []int{len(paperMap[0]), len(paperMap)}
	accessiblePapers := 0
	for y := 0; y < dimensions[1]; y++ {
		for x := 0; x < dimensions[0]; x++ {
			if isAccessible(paperMap, dimensions, x, y) {
				accessiblePapers++
				cloneMap[y][x] = "x"
			}
		}
	}
	fmt.Println("Clone Map:")
	for _, line := range cloneMap {
		fmt.Println(line)
	}
	return accessiblePapers
}

func isAccessible(paperMap [][]string, dimensions []int, x int, y int) bool {
	// Check boundaries
	if x < 0 || x >= dimensions[0] || y < 0 || y >= dimensions[1] {
		return false
	}
	if paperMap[y][x] == "." {
		return false
	}
	surroundingPapers := checkSurroundings(paperMap, dimensions, x, y)
	if surroundingPapers < 4 {
		return true
	} else {
		return false
	}
}

func checkSurroundings(paperMap [][]string, dimensions []int, x int, y int) int {
	accessibleCount := 0
	directions := [][2]int{
		{0, -1},  // Up
		{0, 1},   // Down
		{-1, 0},  // Left
		{1, 0},   // Right
		{-1, -1}, // Top-left
		{1, -1},  // Top-right
		{-1, 1},  // Bottom-left
		{1, 1},   // Bottom-right
	}

	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]
		if newX < 0 || newX >= dimensions[0] || newY < 0 || newY >= dimensions[1] {
			continue
		}
		if paperMap[newY][newX] == "@" {
			accessibleCount++
		}
	}
	return accessibleCount
}
