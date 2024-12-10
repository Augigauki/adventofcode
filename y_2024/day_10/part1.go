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
