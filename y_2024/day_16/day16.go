package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Pos struct {
	x, y int
}

type Reindeer struct {
	startPos Pos
	pos      Pos
	dir      Direction
	path     []Pos
	paths    [][]Pos
}

type Direction struct {
	x, y int
}

type BestPath struct {
	path  []Pos
	score int
}

var Directions = []Direction{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file: ", fileName)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		chars := []string{}
		for _, char := range line {
			chars = append(chars, string(char))
		}
		input = append(input, chars)
	}
	//part1(input)
	part2()
}

// var goalPaths [][]Pos
var bestPaths []BestPath

func part1(maze [][]string) {
	fmt.Println("Part 1")
	//fmt.Println("Maze:")
	racer := Reindeer{Pos{0, 0}, Pos{0, 0}, Directions[1], []Pos{}, [][]Pos{}}
	//fmt.Println("Maze: ", maze)
	//endPos := Pos{0, 0}
	for y, row := range maze {
		fmt.Println(row)
		//fmt.Println("y: ", y)
		for x, char := range row {
			//fmt.Println("x: ", x)
			//fmt.Println("Char:", char)
			if char == "S" {
				//fmt.Println("Found start position")
				racer.startPos = Pos{x, y}
				racer.pos = Pos{x, y}
			}
			if char == "E" {
				//fmt.Println("Found end position")
				//endPos = Pos{x, y}
			}
		}
	}
	maze[racer.startPos.y][racer.startPos.x] = "."
	fmt.Println("Start position:", racer.startPos)
	racer.path = append(racer.path, racer.startPos)
	explorePath(&racer, maze, map[Pos]map[Direction]int{}, 0)

	lowestScore := 1<<31 - 1
	fmt.Println("Best paths:", len(bestPaths))
	for _, track := range bestPaths {
		if track.score < lowestScore {
			lowestScore = track.score
		}
	}
	allBestPaths := [][]Pos{}
	for _, track := range bestPaths {
		if track.score == lowestScore {
			allBestPaths = append(allBestPaths, track.path)
		}
	}
	fmt.Printf("Lowest score: %v, num of best paths: %v\n", lowestScore, len(allBestPaths))
	pathMap := maze
	fmt.Println("All shortest paths:")
	for i, path := range allBestPaths {
		for _, pos := range path {
			pathMap[pos.y][pos.x] = strconv.Itoa(i + 1)
		}
	}
	allBestTiles := 0
	for _, row := range pathMap {
		fmt.Println(row)
		for _, tile := range row {
			if tile != "." && tile != "#" {
				allBestTiles++
			}
		}
	}
	fmt.Println("All best tiles: ", allBestTiles)
	//fmt.Println("Valid paths: ", validPaths)
	//fmt.Println("Shortest path: ", shortestPath)
}

func explorePath(racer *Reindeer, maze [][]string, visited map[Pos]map[Direction]int, score int) {
	if maze[racer.pos.y][racer.pos.x] == "E" {
		// Goal reached: Save the path and score
		racer.paths = append(racer.paths, racer.path)
		bestPaths = append(bestPaths, BestPath{path: racer.path, score: score})
		return
	}

	for _, dir := range Directions {
		nextPos := Pos{racer.pos.x + dir.x, racer.pos.y + dir.y}

		// Check boundaries and walls
		if nextPos.x < 0 || nextPos.x >= len(maze[0]) || nextPos.y < 0 || nextPos.y >= len(maze) {
			continue
		}
		if maze[nextPos.y][nextPos.x] == "#" {
			continue
		}

		// Calculate new score
		newScore := score + 1
		if dir != racer.dir {
			newScore += 1000
		}

		// Initialize visited state tracking
		if _, ok := visited[nextPos]; !ok {
			visited[nextPos] = make(map[Direction]int)
		}

		// Allow revisiting if the score is equal or better for the same direction
		if prevScore, ok := visited[nextPos][dir]; !ok || newScore <= prevScore {
			visited[nextPos][dir] = newScore // Update the visited map

			// Recursively explore next position
			newPath := append([]Pos{}, racer.path...)
			newPath = append(newPath, nextPos)
			newRacer := Reindeer{startPos: racer.startPos, pos: nextPos, dir: dir, path: newPath, paths: racer.paths}
			explorePath(&newRacer, maze, visited, newScore)
		}
	}
}
