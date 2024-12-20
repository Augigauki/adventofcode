package main

import (
	"fmt"
	"math"
)

func part1() {
	fileName := ""
	if example {
		fileName = "example.txt"
	} else {
		fileName = "input.txt"
	}
	racetrack := parseFile(fileName)
	fmt.Println("Racetrack:")
	printMap(racetrack)

	// 1. Flood Fill (BFS) to get distances from the start
	distances := bfs(racetrack, startPos)

	// 2. Iterate through point pairs and check for savings
	count := 0
	usedCheats := make(map[string]bool) // Keep track of unique cheat start/end pairs

	for p1, dist1 := range distances {
		for p2, dist2 := range distances {
			if dist1 >= dist2 {
				continue // Avoid duplicates and consider only p2 further from start than p1
			}

			manhattanDist := math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y))

			if manhattanDist == 2 && (dist2-dist1)-int(manhattanDist) >= 100 && isWithinTwoWalls(p1, p2, racetrack) {
				// Create a unique key for this cheat (start and end positions)
				cheatKey := fmt.Sprintf("%d,%d-%d,%d", p1.x, p1.y, p2.x, p2.y)

				// Check if this cheat has already been counted
				if !usedCheats[cheatKey] {
					count++
					usedCheats[cheatKey] = true // Mark this cheat as counted
				}
			}
		}
	}

	fmt.Println("Number of cheats saving at least 100 picoseconds:", count)
}

func isWithinTwoWalls(p1 Pos, p2 Pos, racetrack map[Pos]string) bool {
	dx := p2.x - p1.x
	dy := p2.y - p1.y

	// Check for direct horizontal or vertical movement
	if dx == 0 { // Vertical movement
		for y := min(p1.y, p2.y) + 1; y < max(p1.y, p2.y); y++ {
			if racetrack[Pos{p1.x, y}] != "#" {
				return false // Not a wall
			}
		}
		return true
	} else if dy == 0 { // Horizontal movement
		for x := min(p1.x, p2.x) + 1; x < max(p1.x, p2.x); x++ {
			if racetrack[Pos{x, p1.y}] != "#" {
				return false // Not a wall
			}
		}
		return true
	}

	// Check for diagonal movement (Manhattan distance of 2)
	if abs(dx)+abs(dy) == 2 {
		// Check the two intermediate positions for walls
		return racetrack[Pos{p1.x + sign(dx), p1.y}] == "#" && racetrack[Pos{p1.x, p1.y + sign(dy)}] == "#"
	}

	return false
}

// Helper functions for min, max, and sign
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Breadth-First Search (BFS) to calculate distances from startPos
func bfs(racetrack map[Pos]string, startPos Pos) map[Pos]int {
	distances := make(map[Pos]int)
	queue := []Pos{startPos}
	distances[startPos] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range Directions {
			next := Pos{current.x + dir.x, current.y + dir.y}

			if _, ok := racetrack[next]; !ok || racetrack[next] == "#" {
				continue // Skip invalid or wall positions
			}

			if _, ok := distances[next]; !ok {
				distances[next] = distances[current] + 1
				queue = append(queue, next)
			}
		}
	}

	return distances
}
