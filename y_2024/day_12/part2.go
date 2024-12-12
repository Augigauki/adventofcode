package main

import "fmt"

/* type Plant struct {
	Pos Pos
} */

type Plot struct {
	Seed   string
	Plants []Pos
	Area   int
	Sides  int
}

func CalcFenceCostWithDiscount(gardenMap [][]string) int {
	var totalCost int
	visited := make(map[Pos]bool)
	plots := []Plot{}

	for i, line := range gardenMap {
		for j, char := range line {
			if !visited[Pos{i, j}] {
				plant := char
				plot := FindPlots(gardenMap, Pos{i, j}, visited, plant)
				plot.Sides = CalculateSides(plot, gardenMap)
				fmt.Println("Walls: ", plot.Sides, " for plant: ", plant)
				cost := plot.Area * plot.Sides
				plots = append(plots, plot)
				totalCost += cost
			}
		}
	}
	fmt.Println("Total cost:", totalCost)
	return totalCost
}

func FindPlots(gardenMap [][]string, pos Pos, visited map[Pos]bool, plant string) Plot {
	stack := []Pos{pos}
	area := 0
	plot := Plot{Seed: plant}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[curr] {
			continue
		}

		plot.Plants = append(plot.Plants, curr)
		visited[curr] = true
		area++

		for _, dir := range Directions {
			newPos := Pos{curr.line + dir[0], curr.char + dir[1]}
			if InBounds(gardenMap, newPos) && gardenMap[newPos.line][newPos.char] == plant && !visited[newPos] {
				stack = append(stack, newPos)
			}
		}
	}

	plot.Area = area
	return plot
}

func CalculateSides(plot Plot, gardenMap [][]string) int {
	visitedBoundary := make(map[Pos]bool)
	sides := 0

	for _, plant := range plot.Plants {
		for _, dir := range Directions {
			newPos := Pos{plant.line + dir[0], plant.char + dir[1]}
			if !InBounds(gardenMap, newPos) || gardenMap[newPos.line][newPos.char] != plot.Seed {
				// Check if this boundary segment has already been counted
				if !visitedBoundary[Pos{plant.line, plant.char}] {
					sides++
					TraverseBoundary(Pos{plant.line, plant.char}, gardenMap, plot.Seed, visitedBoundary)
				}
			}
		}
	}
	return sides
}

func TraverseBoundary(start Pos, gardenMap [][]string, seed string, visited map[Pos]bool) {
	queue := []Pos{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current] = true
		for _, dir := range Directions {
			newPos := Pos{current.line + dir[0], current.char + dir[1]}
			if InBounds(gardenMap, newPos) && gardenMap[newPos.line][newPos.char] != seed && !visited[newPos] {
				queue = append(queue, newPos)
			}
		}
	}
}
