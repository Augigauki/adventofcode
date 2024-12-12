package main

import "fmt"

type Plot struct {
	Seed   string
	Plants []Plant
	Area   int
	Sides  int
}

type Plant struct {
	Pos              Pos
	OutOfBoundsPoses []Pos
	Corners          int
}

func CalcFenceCostWithDiscount(gardenMap [][]string) int {
	var totalCost int
	visited := make(map[Pos]bool)
	//plots := []Plot{}

	for i, line := range gardenMap {
		for j, char := range line {
			if !visited[Pos{i, j}] {
				seed := char
				plot := FindPlots(gardenMap, Pos{i, j}, visited, seed)
				corners := 0
				for _, plant := range plot.Plants {
					corners += plant.Corners
				}
				fmt.Println("Walls: ", corners, " for plant: ", seed)
				cost := plot.Area * corners
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
		outOfBounds := SumOutOfBounds(curr, gardenMap, plant)
		/* Count corners for position */
		corners := 0
		cardinalAndDiagonal := GetCardinalAndDiagonal(gardenMap, curr, plant)
		//fmt.Println("\nCalculating corners for plant: ", curr, " with seed ", plant)
		corners = CalcCorners(curr, cardinalAndDiagonal)

		plot.Plants = append(plot.Plants, Plant{Pos: curr, OutOfBoundsPoses: outOfBounds, Corners: corners})
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

func CalcCorners(pos Pos, allDirections map[Pos]bool) int {

	corners := 0
	northPos := Pos{pos.line - 1, pos.char}
	southPos := Pos{pos.line + 1, pos.char}
	eastPos := Pos{pos.line, pos.char + 1}
	westPos := Pos{pos.line, pos.char - 1}
	northEastPos := Pos{pos.line - 1, pos.char + 1}
	northWestPos := Pos{pos.line - 1, pos.char - 1}
	southEastPos := Pos{pos.line + 1, pos.char + 1}
	southWestPos := Pos{pos.line + 1, pos.char - 1}
	north := allDirections[northPos]
	south := allDirections[southPos]
	east := allDirections[eastPos]
	west := allDirections[westPos]
	northEast := allDirections[northEastPos]
	northWest := allDirections[northWestPos]
	southEast := allDirections[southEastPos]
	southWest := allDirections[southWestPos]

	if !north && !east {
		corners++
	}
	if !north && !west {
		corners++
	}
	if !south && !east {
		corners++
	}
	if !south && !west {
		corners++
	}
	if !northEast && (north && east) {
		corners++
	}
	if !northWest && (north && west) {
		corners++
	}
	if !southEast && (south && east) {
		corners++
	}
	if !southWest && (south && west) {
		corners++
	}
	return corners
}

func GetCardinalAndDiagonal(gardenMap [][]string, pos Pos, seed string) map[Pos]bool {
	cardinalAndDiagonal := make(map[Pos]bool)

	for _, dir := range AllDirections {
		newPos := Pos{pos.line + dir[0], pos.char + dir[1]}
		if InBounds(gardenMap, newPos) && gardenMap[newPos.line][newPos.char] == seed {
			cardinalAndDiagonal[newPos] = true
		} else {
			cardinalAndDiagonal[newPos] = false
		}
	}
	return cardinalAndDiagonal
}

func SumOutOfBounds(pos Pos, gardenMap [][]string, seed string) []Pos {
	outOfBounds := []Pos{}
	for _, dir := range Directions {
		newPos := Pos{pos.line + dir[0], pos.char + dir[1]}
		if !InBounds(gardenMap, newPos) || gardenMap[newPos.line][newPos.char] != seed {

			outOfBounds = append(outOfBounds, newPos)
		}
	}
	return outOfBounds
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
