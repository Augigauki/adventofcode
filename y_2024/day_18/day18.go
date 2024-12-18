package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

type Direction struct {
	x, y int
}

type Byte struct {
	pos Pos
}

type Me struct {
	pos  Pos
	path []Pos
	dir  Direction
}

var Directions = []Direction{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

var example = false

func main() {
	size := 0
	nSeconds := 0
	fileName := ""
	if example {
		size = 7
		nSeconds = 12
		fileName = "example.txt"
	} else {
		size = 71
		nSeconds = 1024
		fileName = "input.txt"
	}
	bytes := parseFile(fileName)
	fmt.Println(bytes)

	memSpace := createMap(size)
	memSpace = corruptMemory(memSpace, bytes, nSeconds)
	fmt.Printf("Memory space after %v nanoseconds:\n", nSeconds)
	for _, row := range memSpace {
		fmt.Println(row)
	}
	//me := Me{Pos{0, 0}, []Pos{Pos{0, 0}}, Direction{1, 0}}
	//findShortestPath(&me, memSpace, &map[Pos]bool{})
	shortestPathLength := len(memSpace) * len(memSpace[0])
	shortestPath := []Pos{}
	paths := findPaths(memSpace)
	for _, path := range paths {
		if len(path) < shortestPathLength {
			shortestPathLength = len(path)
			shortestPath = path
		}
	}
	for _, pos := range shortestPath {
		memSpace[pos.y][pos.x] = "O"
	}
	fmt.Println("Shortest path: ", shortestPathLength-1)
	for _, row := range memSpace {
		fmt.Println(row)
	}
	CorruptMore(memSpace, bytes)
}

func CorruptMore(memSpace [][]string, bytes []Byte) {
	fmt.Println("Part 2:")
	fmt.Println("Length of findPaths: ", len(findPaths(memSpace)))
	paths := len(findPaths(memSpace))
	i := 0
	if example {
		i = 12
	} else {
		i = 1024
	}
	for paths > 0 {
		memSpace = corruptMemory(memSpace, bytes, i)
		for _, row := range memSpace {
			fmt.Println(row)
		}
		paths = len(findPaths(memSpace))
		fmt.Println("Length of findPaths: ", paths)
		i++
	}
	fmt.Println("Blocking obstacle: ", bytes[i-2])
}

func findPaths(memSpace [][]string) [][]Pos {
	startPos := Pos{0, 0}
	endPos := Pos{x: len(memSpace[0]) - 1, y: len(memSpace) - 1}
	allPaths := [][]Pos{}

	queue := list.New()
	queue.PushBack(&Me{startPos, []Pos{startPos}, Direction{1, 0}})

	visited := make(map[Pos]bool)

	for queue.Len() > 0 {
		curr := queue.Front()
		queue.Remove(curr)
		me := curr.Value.(*Me)

		if me.pos == endPos {
			allPaths = append(allPaths, me.path)
			continue
		}

		if visited[me.pos] {
			continue
		}
		visited[me.pos] = true

		for _, dir := range Directions {
			nextPos := Pos{me.pos.x + dir.x, me.pos.y + dir.y}
			if isWithinBounds(nextPos, memSpace) && memSpace[nextPos.y][nextPos.x] != "#" {
				newPath := append([]Pos{}, me.path...)
				newPath = append(newPath, nextPos)
				//newMe := Me{nextPos, newPath, dir}
				queue.PushBack(&Me{nextPos, newPath, dir})
			}
		}
	}
	return allPaths

}

func isWithinBounds(pos Pos, memSpace [][]string) bool {
	return pos.y >= 0 && pos.y < len(memSpace) && pos.x >= 0 && pos.x < len(memSpace[pos.y])
}

func corruptMemory(memSpace [][]string, bytes []Byte, nanoSeconds int) [][]string {
	corruptedMemSpace := memSpace
	for i := 0; i < nanoSeconds; i++ {
		corruptedMemSpace[bytes[i].pos.y][bytes[i].pos.x] = "#"
	}
	return corruptedMemSpace
}

func createMap(size int) [][]string {
	m := make([][]string, size)
	for i := 0; i < size; i++ {
		m[i] = make([]string, size)
		for j := 0; j < size; j++ {
			m[i][j] = "."
		}
	}
	return m
}

func parseFile(fileName string) []Byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error reading file: ", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bytes := []Byte{}
	for scanner.Scan() {
		line := scanner.Text()
		poses := strings.Split(line, ",")
		x, _ := strconv.Atoi(poses[0])
		y, _ := strconv.Atoi(poses[1])
		bytes = append(bytes, Byte{Pos{x, y}})
	}
	return bytes
}
