package main

import "fmt"

func CalcFenceCostWithDiscount(gardenMap [][]string) {
	var totalCost int
	visited := map[Pos]bool{}
	fmt.Println("Garden map:")
	for i, line := range gardenMap {
		for j, char := range line {
			if !visited[Pos{i, j}] {
				plant := char
				area, perimeter := MapPlotArea(gardenMap, Pos{i, j}, visited, plant)
				cost := area * perimeter
				totalCost += cost
			}
		}
	}
	fmt.Println("Total cost:", totalCost)
}

func MapPlotSides(gardenMap [][]string, pos Pos, visited map[Pos]bool, plant string) (int, int) {
	stack := []Pos{pos}
	area := 0
	perimeter := 0

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currPos := Pos{curr.line, curr.char}
		if visited[curr] {
			continue
		}
		visited[curr] = true
		area++
		fmt.Printf("Currpos: %v, plant: %v\n", currPos, plant)
		for _, dir := range Directions {
			newPos := Pos{currPos.line + dir[0], currPos.char + dir[1]}
			if InBounds(gardenMap, newPos) {
				if gardenMap[newPos.line][newPos.char] == plant {
					if !visited[newPos] {
						stack = append(stack, newPos)
					}
				} else {
					perimeter++
				}
			} else {
				perimeter++
			}
		}
	}
	return area, perimeter
}
