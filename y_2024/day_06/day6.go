package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	startPos := FindStartPos(guardMap)
	traversedMap := TraverseMap(startPos, guardMap)
	/* fmt.Println("\n::  TRAVERSED MAP  ::")
	for _, line := range traversedMap.traversedMap {
		fmt.Println(line)
	} */
	TraverseWithObstacles(traversedMap.startPos, traversedMap.traversedMap)
}
