package main

import "fmt"

func removeAllPaper(paperMap [][]string) int {
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
				paperMap[y][x] = "."
				// Restart scanning from the beginning
				x = -1
				y = 0
			}
		}
	}
	fmt.Println("Clone Map:")
	for _, line := range cloneMap {
		fmt.Println(line)
	}
	return accessiblePapers
}
