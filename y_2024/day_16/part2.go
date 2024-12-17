package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type State struct {
	pos   Pos
	dir   Direction
	score int
	path  []Pos
}

func part2() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file: ", fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maze := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		chars := []string{}
		for _, char := range line {
			chars = append(chars, string(char))
		}
		maze = append(maze, chars)
	}

	findBestPaths(maze)
}

func findBestPaths(maze [][]string) {
	start, end := findStartAndEnd(maze)
	queue := []State{{start, Direction{0, 0}, 0, []Pos{start}}}
	visited := make(map[Pos]map[Direction]int) // Tracks scores for (position, direction)

	bestScore := 1 << 31
	var bestPaths [][]Pos

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check if we reached the goal
		if current.pos == end {
			if current.score < bestScore {
				bestScore = current.score
				bestPaths = [][]Pos{current.path} // Reset best paths
			} else if current.score == bestScore {
				bestPaths = append(bestPaths, current.path) // Add this path
			}
			continue
		}

		// Explore all possible directions
		for _, dir := range Directions {
			nextPos := Pos{current.pos.x + dir.x, current.pos.y + dir.y}
			if !isValid(nextPos, maze) {
				continue
			}

			newScore := current.score
			if dir != current.dir {
				newScore += 1000 // Turning cost
			}
			newScore += 1 // Step cost

			// Track visited positions with direction and score
			if _, ok := visited[nextPos]; !ok {
				visited[nextPos] = make(map[Direction]int)
			}
			// Only skip if the previous score is strictly better
			if prevScore, ok := visited[nextPos][dir]; ok && prevScore < newScore {
				continue
			}

			visited[nextPos][dir] = newScore
			newPath := append([]Pos{}, current.path...)
			newPath = append(newPath, nextPos)

			queue = append(queue, State{nextPos, dir, newScore, newPath})
		}
	}

	// Step 1: Combine all positions in the best paths into a single set
	tileSet := make(map[Pos]bool)
	for _, path := range bestPaths {
		for _, pos := range path {
			tileSet[pos] = true // Add tile to the set
		}
	}

	// Step 2: Copy the maze and mark the tiles
	markedMaze := make([][]string, len(maze))
	for i, row := range maze {
		markedMaze[i] = make([]string, len(row))
		copy(markedMaze[i], row) // Copy the maze
	}

	// Step 3: Mark all positions in the tile set, including Start and End
	for pos := range tileSet {
		if markedMaze[pos.y][pos.x] == "S" || markedMaze[pos.y][pos.x] == "E" {
			markedMaze[pos.y][pos.x] = "O"
		} else {
			markedMaze[pos.y][pos.x] = "O"
		}
	}

	// Step 4: Count the number of unique tiles in the best paths
	count := 0
	for _, row := range markedMaze {
		for _, char := range row {
			if char == "O" {
				count++
			}
		}
	}

	// Step 5: Print results
	fmt.Println("Best paths: ", len(bestPaths))
	fmt.Println("Lowest score:", bestScore)
	fmt.Println("Number of unique tiles in best paths:", count)

	// Step 6: Print the updated maze
	for _, row := range markedMaze {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func findStartAndEnd(maze [][]string) (Pos, Pos) {
	var start, end Pos
	for y, row := range maze {
		for x, char := range row {
			if char == "S" {
				start = Pos{x, y}
			}
			if char == "E" {
				end = Pos{x, y}
			}
		}
	}
	return start, end
}

func isValid(pos Pos, maze [][]string) bool {
	return pos.x >= 0 && pos.x < len(maze[0]) && pos.y >= 0 && pos.y < len(maze) && maze[pos.y][pos.x] != "#"
}
