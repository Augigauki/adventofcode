package main

import (
	"fmt"
	"math"
)

func part2() {
	fileName := ""
	if example {
		fileName = "example.txt"
	} else {
		fileName = "input.txt"
	}
	racetrack := parseFile(fileName)
	fmt.Println("Racetrack:")
	printMap(racetrack)

	// 1. Dijkstra's Algorithm to get shortest path
	distances, previous := dijkstra(racetrack, startPos)
	shortestPath := reconstructPath(previous, endPos)

	// 2. Iterate through path segments and check for valid cheats
	cheatSavings := make(map[int]int)
	for i := 0; i < len(shortestPath); i++ {
		for j := i + 2; j < len(shortestPath); j++ {
			p1 := shortestPath[i]
			p2 := shortestPath[j]

			manhattanDist := int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))

			// Check for cheats up to 20 steps
			if manhattanDist <= 20 && isValidCheat(p1, p2, racetrack, distances, shortestPath) {
				saved := (distances[p2] - distances[p1]) - manhattanDist
				if saved >= 2 { // Count cheats that save at least 2
					cheatSavings[saved]++
				}
			}
		}
	}

	// 3. Print cheat savings distribution (for savings >= 50)
	fmt.Println("Cheat Savings Distribution:")
	for saving, count := range cheatSavings {
		if saving >= 50 {
			fmt.Printf("There are %d cheats that save %d picoseconds.\n", count, saving)
		}
	}

	// 4. Count cheats saving at least 100 picoseconds
	threshold := 0
	if example {
		threshold = 50
	} else {
		threshold = 100
	}
	count100 := 0
	for saving := range cheatSavings {
		if saving >= threshold {
			count100++
		}
	}
	fmt.Println("Number of cheats saving at least ", threshold, " picoseconds:", count100)
}

func isValidCheat(p1, p2 Pos, racetrack map[Pos]string, distances map[Pos]int, shortestPath []Pos) bool {
	// Check if p2 is a valid position (not a wall and within bounds)
	if !isWithinBounds(p2) || racetrack[p2] == "#" {
		return false
	}

	// Perform BFS to check for a path of walls between p1 and p2
	return canReachWithWalls(p1, p2, racetrack)
}

func canReachWithWalls(start, end Pos, racetrack map[Pos]string) bool {
	queue := []Pos{start}
	visited := make(map[Pos]bool)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return true // Found a path to the end position
		}

		for _, dir := range Directions {
			next := Pos{current.x + dir.x, current.y + dir.y}

			// Check if the next position is within bounds, a wall, and not visited
			if isWithinBounds(next) && racetrack[next] == "#" && !visited[next] {
				queue = append(queue, next)
				visited[next] = true
			}
		}
	}

	return false // No path found
}
