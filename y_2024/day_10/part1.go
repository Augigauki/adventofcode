package main

import "fmt"

type Trailhead struct {
	Pos       Pos
	Score     Score
	Goals     []Goal
	Traversed []Pos
	AllPaths  [][]Pos
	Value     int
}

type TrailheadPos struct {
	Pos           Pos
	PossiblePaths []Pos
}

type Score struct {
	UniqueScore int
	AllScores   int
}

type Goal struct {
	Pos Pos
}

type Pos struct {
	Line int
	Char int
}

func FindTrailHeadScores(topoMap [][]int) int {
	var trailHeadScore = 0
	FindTrailHeads(topoMap)

	return trailHeadScore
}

func FindTrailHeads(topoMap [][]int) {
	var trailHeads []Trailhead
	for i, line := range topoMap {
		for j, char := range line {
			if char == 0 {
				trailHeads = append(trailHeads, Trailhead{
					Pos:       Pos{i, j},
					Score:     Score{0, 0},
					Goals:     []Goal{},
					Traversed: []Pos{},
					AllPaths:  [][]Pos{},
					Value:     0,
				})
			}
		}
	}
	/* for _, trailHead := range trailHeads {
		HikeTrail(&trailHead, trailHead.Pos, topoMap, []Pos{})
	} */

	// After traversal, print unique paths
	var totalUniqueGoalScore = 0
	var totalAllPathScore = 0
	for _, trailHead := range trailHeads {
		allPaths, uniqueGoals := HikeTrailBFS(topoMap, trailHead.Pos)
		//trailHead.AllPaths = allPaths
		fmt.Printf("Trailhead at %v found %d unique goals.\n", trailHead.Pos, len(uniqueGoals))
		totalUniqueGoalScore += len(uniqueGoals)
		fmt.Printf("Trailhead at %v found %d unique paths.\n", trailHead.Pos, len(allPaths))
		totalAllPathScore += len(allPaths)
		/* for i, path := range allPaths {
			fmt.Printf("Path %d: %v\n", i+1, path)
		} */
	}
	fmt.Printf("Total unique goals found: %d\n", totalUniqueGoalScore)
	fmt.Printf("Total unique paths found: %d\n", totalAllPathScore)
}

func FindTrailHeads2(topoMap [][]int) {
	var trailHeads []Trailhead
	for i, line := range topoMap {
		for j, char := range line {
			if char == 0 {
				trailHeads = append(trailHeads, Trailhead{Pos{i, j}, Score{0, 0}, []Goal{}, []Pos{}, [][]Pos{}, 0})
			}
		}
	}
	for _, trailHead := range trailHeads {
		HikeTrail(&trailHead, trailHead.Pos, topoMap, []Pos{})
	}
}

func HikeTrailBFS(topoMap [][]int, startPos Pos) ([][]Pos, map[Pos]bool) {
	queue := []struct {
		pos   Pos
		value int
		path  []Pos
	}{{startPos, 0, []Pos{startPos}}}

	allPaths := [][]Pos{}
	uniqueGoals := map[Pos]bool{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		//If goal is reached, store the path
		if current.value == 9 {
			allPaths = append(allPaths, append([]Pos{}, current.path...))
			uniqueGoals[current.pos] = true
			continue
		}

		//Explore all valid moves
		for _, dir := range []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			nextPos := Pos{current.pos.Line + dir.Line, current.pos.Char + dir.Char}

			//Check boundaries
			if nextPos.Line >= 0 && nextPos.Line < len(topoMap) &&
				nextPos.Char >= 0 && nextPos.Char < len(topoMap[0]) &&
				topoMap[nextPos.Line][nextPos.Char] != -1 && // Check for impassable tile
				topoMap[nextPos.Line][nextPos.Char] == current.value+1 {
				// Add the next position to the queue
				queue = append(queue, struct {
					pos   Pos
					value int
					path  []Pos
				}{
					pos:   nextPos,
					value: current.value + 1,
					path:  append(append([]Pos{}, current.path...), nextPos),
				})
			}
		}
	}

	return allPaths, uniqueGoals
}

func HikeTrail(trailHead *Trailhead, startPos Pos, topoMap [][]int, currentPath []Pos) {
	// Add current position to the path
	currentPath = append(currentPath, trailHead.Pos)

	if trailHead.Value == 9 {
		fmt.Println("\nGOAL!!! 🚀")

		// Ensure the path is added as a copy
		trailHead.AllPaths = append(trailHead.AllPaths, append([]Pos{}, currentPath...))

		fmt.Printf("Unique paths found so far: %d\n", len(trailHead.AllPaths))
		return // End recursion for this path
	}

	// Explore all valid directions
	directions := []Pos{
		{-1, 0}, // North
		{0, 1},  // East
		{1, 0},  // South
		{0, -1}, // West
	}

	for _, dir := range directions {
		nextLine := trailHead.Pos.Line + dir.Line
		nextChar := trailHead.Pos.Char + dir.Char

		// Check boundaries
		if nextLine >= 0 && nextLine < len(topoMap) && nextChar >= 0 && nextChar < len(topoMap[0]) {
			// Check if the next position is valid (current value + 1)
			if topoMap[nextLine][nextChar] == trailHead.Value+1 {
				// Move to the next position
				newTrailHead := *trailHead // Copy the trailhead for recursive call
				newTrailHead.Pos = Pos{nextLine, nextChar}
				newTrailHead.Value++

				HikeTrail(&newTrailHead, startPos, topoMap, currentPath) // Recursive call
			}
		}
	}
}

func HikeTrail2(trailHead *Trailhead, startPos Pos, topoMap [][]int) {
	possiblePaths := findAllNextPaths(trailHead, topoMap)
	if trailHead.Value == 9 {
		fmt.Println("\nGOAL!!! 🚀")
		trailHead.Traversed = append(trailHead.Traversed, trailHead.Pos)
		trailHead.Score.AllScores++
		if len(trailHead.Goals) == 0 {
			trailHead.Score.UniqueScore++
		} else {
			for _, goal := range trailHead.Goals {
				var alreadyCounted = false
				if goal.Pos.Line == trailHead.Pos.Line && goal.Pos.Char == trailHead.Pos.Char {
					alreadyCounted = true
				}
				if !alreadyCounted {
					trailHead.Score.UniqueScore++
				}
			}
		}
		trailHead.Goals = append(trailHead.Goals, Goal{Pos{trailHead.Pos.Line, trailHead.Pos.Char}})
		fmt.Printf("Trailhead scores: Unique: %v -  All: %v\n", trailHead.Score.UniqueScore, trailHead.Score.AllScores)
		/* fmt.Println("Trailhead path: ")
		for _, path := range trailHead.Traversed {
			fmt.Printf("Line: %v, Char: %v\n", path.Pos.Line, path.Pos.Char)
		} */
		victoryPath := []Pos{}
		for _, path := range trailHead.Traversed {
			victoryPath = append(victoryPath, path)
		}
		trailHead.AllPaths = append(trailHead.AllPaths, victoryPath)
		fmt.Println("All paths: ")
		for _, path := range trailHead.AllPaths {
			fmt.Println(path)
		}
		if trailHead.Score.AllScores > 4 {
			return
		}
		ResetTrailHead(trailHead, startPos)
		return
	}
	//possiblePaths = findAllNextPaths(trailHead, topoMap)
	fmt.Println("Possible paths: ", possiblePaths)
	fmt.Printf("Trailhead pos: %v, value: %v\n", trailHead.Pos, trailHead.Value)
	//fmt.Printf("Trailhead scores: Unique: %v -  All: %v\n", trailHead.Score.UniqueScore, trailHead.Score.AllScores)

	if trailHead.Pos.Line > 0 {
		/* currPos := Pos{trailHead.Pos.Line, trailHead.Pos.Char}
		nextPos := Pos{trailHead.Pos.Line - 1, trailHead.Pos.Char} */
		if topoMap[trailHead.Pos.Line-1][trailHead.Pos.Char] == trailHead.Value+1 {
			/* value := topoMap[trailHead.Pos.Line-1][trailHead.Pos.Char]
			var alreadyTraversed = checkIfAlreadyTraversed(trailHead, currPos, nextPos, value)
			fmt.Println("Already traversed north: ", alreadyTraversed)
			if !alreadyTraversed {

			} */
			fmt.Println("MOVING NORTH")
			trailHead.Traversed = append(trailHead.Traversed, trailHead.Pos)
			trailHead.Value++
			trailHead.Pos.Line--
			HikeTrail(trailHead, startPos, topoMap, []Pos{})
		}
	}
	if trailHead.Pos.Char < len(topoMap[trailHead.Pos.Line])-1 {
		if topoMap[trailHead.Pos.Line][trailHead.Pos.Char+1] == trailHead.Value+1 {
			fmt.Println("MOVING EAST")
			trailHead.Traversed = append(trailHead.Traversed, trailHead.Pos)
			trailHead.Value++
			trailHead.Pos.Char++
			HikeTrail(trailHead, startPos, topoMap, []Pos{})
		}
	}
	if trailHead.Pos.Line < len(topoMap)-1 {
		if topoMap[trailHead.Pos.Line+1][trailHead.Pos.Char] == trailHead.Value+1 {
			fmt.Println("MOVING SOUTH")
			trailHead.Traversed = append(trailHead.Traversed, trailHead.Pos)
			trailHead.Value++
			trailHead.Pos.Line++
			HikeTrail(trailHead, startPos, topoMap, []Pos{})
		}
	}
	if trailHead.Pos.Char > 0 {
		if topoMap[trailHead.Pos.Line][trailHead.Pos.Char-1] == trailHead.Value+1 {

			fmt.Println("MOVING WEST")
			trailHead.Traversed = append(trailHead.Traversed, trailHead.Pos)
			trailHead.Value++
			trailHead.Pos.Char--
			HikeTrail(trailHead, startPos, topoMap, []Pos{})
		}

	}
}

func checkIfAlreadyTraversed(trailHead *Trailhead, pos Pos, nextPos Pos, value int) bool {
	hasTraversed := false
	if len(trailHead.AllPaths) == 0 {
		return hasTraversed
	}
	samePath := true
	for _, prevPath := range trailHead.AllPaths {
		if len(prevPath) == 0 {
			return hasTraversed
		}
		if len(trailHead.Traversed) == 0 {
			return hasTraversed
		}
		if len(prevPath) <= value {
			return hasTraversed
		} else {

			for len(trailHead.Traversed) < value {

			}

		}
		if samePath {
			fmt.Println("Paths are similar so far...")
			if prevPath[value+1] == nextPos {
				hasTraversed = true
				return hasTraversed
			}
		}

		/* if len(prevPath) != len(trailHead.Traversed) {
			continue
		} */

		/* for _, path := range prevPath {
			if path != pos {
				samePath = false
			}
		} */
	}
	if !samePath {
		for _, pos := range trailHead.Traversed {
			if pos == nextPos {
				fmt.Println("Already traversed this spot")
				hasTraversed = true
				return hasTraversed
			}
		}
	}
	return hasTraversed
}

func findAllNextPaths(trailHead *Trailhead, topoMap [][]int) []Pos {
	possiblePaths := []Pos{}
	fmt.Println("Finding all possible next paths for spot...")
	if trailHead.Pos.Line > 0 {
		if topoMap[trailHead.Pos.Line-1][trailHead.Pos.Char] == trailHead.Value+1 {
			fmt.Println("North is a possible path")
			possiblePaths = append(possiblePaths, Pos{trailHead.Pos.Line - 1, trailHead.Pos.Char})
		}
	}
	if trailHead.Pos.Char < len(topoMap[trailHead.Pos.Line])-1 {
		if topoMap[trailHead.Pos.Line][trailHead.Pos.Char+1] == trailHead.Value+1 {
			fmt.Println("East is a possible path")
			possiblePaths = append(possiblePaths, Pos{trailHead.Pos.Line, trailHead.Pos.Char + 1})
		}
	}
	if trailHead.Pos.Line < len(topoMap)-1 {
		if topoMap[trailHead.Pos.Line+1][trailHead.Pos.Char] == trailHead.Value+1 {
			fmt.Println("South is a possible path")
			possiblePaths = append(possiblePaths, Pos{trailHead.Pos.Line + 1, trailHead.Pos.Char})
		}
	}
	if trailHead.Pos.Char > 0 {
		if topoMap[trailHead.Pos.Line][trailHead.Pos.Char-1] == trailHead.Value+1 {
			fmt.Println("West is a possible path")
			possiblePaths = append(possiblePaths, Pos{trailHead.Pos.Line, trailHead.Pos.Char - 1})
		}
	}
	return possiblePaths

}

func ResetTrailHead(trailHead *Trailhead, startPos Pos) {
	trailHead.Value = 0
	trailHead.Pos = startPos
	trailHead.Traversed = []Pos{}
}
