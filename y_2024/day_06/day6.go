package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	line   int
	char   int
	visits int
}

type Guard struct {
	direction string
	positions []Pos
	pos       Pos
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	guardMap := [][]string{}
	for scanner.Scan() {
		var line string = scanner.Text()
		chars := []string{}
		for _, char := range line {
			chars = append(chars, string(char))
		}
		guardMap = append(guardMap, chars)
	}
	startPos := FindStartPos2(guardMap)
	//traversedMap := TraverseMap(startPos, guardMap)
	part1 := TraverseMapWithGuard(guardMap)
	fmt.Println("New traversed map:")
	for _, line := range part1.traversedMap {
		fmt.Println(line)
	}
	totalObstacles := CountObstaclePlacements(guardMap, part1.visitedPositions, startPos)
	fmt.Println("Total obstacles: ", totalObstacles)
}
