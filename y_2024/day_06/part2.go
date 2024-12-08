package main

import "fmt"

func CountObstaclePlacements(guardMap [][]string, visitedPositions []Pos, startPos Pos) int {

	count := 0
	for _, pos := range visitedPositions {
		if pos == startPos {
			continue
		}

		guardMap[pos.line][pos.char] = "O"
		if doesObstacleTrapGuard(guardMap, startPos) {
			count++
		}

		guardMap[pos.line][pos.char] = "X"
	}
	return count
}

func doesObstacleTrapGuard(guardMap [][]string, startPos Pos) bool {
	guard := Guard{North, []Pos{}, Pos{startPos.line, startPos.char, 0}}
	visited := make(map[string]bool)
	for {
		if !IsWithinBounds(guard.pos, guardMap) {
			break
		}

		stateKey := fmt.Sprintf("%d,%d,%s", guard.pos.line, guard.pos.char, guard.direction)
		//fmt.Printf("State key: %s\n", stateKey)
		if visited[stateKey] {
			fmt.Println("Guard is stuck in loop!")
			return true
		}
		visited[stateKey] = true

		nextPos := GetNextPos(&guard)
		if IsWithinBounds(nextPos, guardMap) && (guardMap[nextPos.line][nextPos.char] == "#" || guardMap[nextPos.line][nextPos.char] == "O") {
			TurnGuard(&guard)
		} else if IsWithinBounds(nextPos, guardMap) {
			guard.pos = nextPos
		} else {
			break
		}
	}
	//fmt.Println("No loop detected!")
	return false
}
