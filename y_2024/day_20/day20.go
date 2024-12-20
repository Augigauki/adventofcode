package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
)

type Direction struct {
	x, y int
}

type Pos struct {
	x, y int
}

type Item struct {
	value    Pos // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

var Directions = []Direction{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

var startPos = Pos{0, 0}
var endPos = Pos{0, 0}
var width, height = 0, 0

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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

// Dijkstra's Algorithm
func dijkstra(racetrack map[Pos]string, startPos Pos) (map[Pos]int, map[Pos]Pos) {
	distances := make(map[Pos]int)
	previous := make(map[Pos]Pos) // To store the path
	pq := make(PriorityQueue, 0)

	distances[startPos] = 0
	heap.Init(&pq)
	item := &Item{value: startPos, priority: 0}
	heap.Push(&pq, item)

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item).value
		if current == endPos {
			break
		}

		for _, dir := range Directions {
			neighbor := Pos{current.x + dir.x, current.y + dir.y}
			if _, ok := racetrack[neighbor]; !ok || racetrack[neighbor] == "#" {
				continue
			}

			newDist := distances[current] + 1
			if oldDist, ok := distances[neighbor]; !ok || newDist < oldDist {
				distances[neighbor] = newDist
				previous[neighbor] = current
				if oldDist == 0 {
					heap.Push(&pq, &Item{value: neighbor, priority: newDist})
				} else {
					// Update priority if neighbor is already in the queue - this is a simplification
					for _, existingItem := range pq {
						if existingItem.value == neighbor {
							pq.update(existingItem, newDist)
							break
						}
					}
				}
			}
		}
	}

	return distances, previous
}

// Reconstruct the path from the start to a given position
func reconstructPath(previous map[Pos]Pos, end Pos) []Pos {
	var path []Pos
	current := end

	for {
		path = append([]Pos{current}, path...)
		prev, ok := previous[current]
		if !ok {
			break // No previous node, we've reached the start
		}
		current = prev
	}

	return path
}

func isWithinBounds(pos Pos) bool {
	return pos.x >= 0 && pos.x < width && pos.y >= 0 && pos.y < height
}

func findPositionInPath(path []Pos, pos Pos) int {
	for i, p := range path {
		if p == pos {
			return i
		}
	}
	return -1
}

func printMap(memSpace map[Pos]string) {
	counter := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if memSpace[Pos{x, y}] == "." {
				counter++
			}
			fmt.Print(memSpace[Pos{x, y}])
		}
		fmt.Println()
	}
	fmt.Println("Total number of empty spaces: ", counter)
}

func parseFile(fileName string) map[Pos]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error reading file: ", err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	memSpace := map[Pos]string{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)
		for x, char := range line {
			if char == 'S' {
				startPos = Pos{x, y}
				memSpace[Pos{x, y}] = "."
			} else if char == 'E' {
				endPos = Pos{x, y}
				memSpace[Pos{x, y}] = "."
			} else {
				memSpace[Pos{x, y}] = string(char)
			}
		}
		y++
	}
	height = y
	return memSpace
}

func manhattanDistBetween(p1, p2 Pos) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}
