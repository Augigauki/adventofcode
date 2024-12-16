package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	x, y int
}

type Direction struct {
	x, y int
}

type Item struct {
	pos      Pos
	dir      Direction
	score    int
	priority int
	index    int
	path     []Pos
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
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
	part1(input)
}

func part1(maze [][]string) {
	fmt.Println("Part 1")

	var start, end Pos
	for y, row := range maze {
		for x, char := range row {
			if char == "S" {
				start = Pos{x, y}
			} else if char == "E" {
				end = Pos{x, y}
			}
		}
	}

	maze[start.y][start.x] = "."
	maze[end.y][end.x] = "."

	score, path := findShortestPath(maze, start, end)
	fmt.Println("Lowest score:", score)
	fmt.Println("Path length:", path)
}

func findShortestPath(maze [][]string, start Pos, end Pos) (int, int) {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{
		pos:      start,
		dir:      Direction{0, 0}, // No direction at start
		score:    0,
		priority: 0,
		path:     []Pos{start},
	})

	// Visited map: Includes direction to avoid revisits with different penalties
	visited := make(map[Pos]map[Direction]int)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)

		// Debug: Print the current item being processed
		//fmt.Printf("Processing: %+v\n", current)

		// If we reach the end, return the score and path length
		if current.pos == end {
			// Length is the total steps (excluding the start position)
			pathLength := len(current.path) - 1
			return current.score, pathLength
		}

		// Check if we've already visited this position with the same direction
		if _, exists := visited[current.pos]; !exists {
			visited[current.pos] = make(map[Direction]int)
		}
		if v, ok := visited[current.pos][current.dir]; ok && v <= current.score {
			continue
		}
		visited[current.pos][current.dir] = current.score

		// Explore all possible directions
		for _, dir := range Directions {
			nextPos := Pos{current.pos.x + dir.x, current.pos.y + dir.y}

			// Skip invalid moves (out of bounds or hitting a wall)
			if nextPos.x < 0 || nextPos.x >= len(maze[0]) || nextPos.y < 0 || nextPos.y >= len(maze) {
				continue
			}
			if maze[nextPos.y][nextPos.x] == "#" {
				continue
			}

			// Calculate new score
			newScore := current.score + 1
			if dir != current.dir {
				newScore += 1000 // Penalty for direction change
			} else if current.dir == (Direction{0, 0}) {
				newScore += 1000 // First step from the start
			}

			// Skip if we've already visited this position with the same direction and lower score
			if v, ok := visited[nextPos][dir]; ok && v <= newScore {
				continue
			}

			// Append next position to the path
			newPath := append([]Pos{}, current.path...)
			newPath = append(newPath, nextPos)

			// Debug: Print the next position being added
			//fmt.Printf("Visiting: %+v, Score: %d\n", nextPos, newScore)

			// Add to priority queue
			heap.Push(pq, &Item{
				pos:      nextPos,
				dir:      dir,
				score:    newScore,
				priority: newScore,
				path:     newPath,
			})
		}
	}

	return -1, -1 // No path found
}
