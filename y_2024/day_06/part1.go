package main

import (
	"fmt"
)

const (
	North = "^"
	East  = ">"
	South = "v"
	West  = "<"
)

type Part1 struct {
	traversedMap     [][]string
	visitedPositions []Pos
}

func TraverseMapWithGuard(guardMap [][]string) Part1 {
	startPos := FindStartPos2(guardMap)

	guard := Guard{North, []Pos{}, Pos{startPos.line, startPos.char, 0}}
	fmt.Println("Guard: ", guard)
	visited := make(map[Pos]bool)
	visitedPoses := []Pos{}

	for {
		if !IsWithinBounds(guard.pos, guardMap) {
			break
		}
		//currentSpot := guardMap[guard.pos.line][guard.pos.char]
		visited[guard.pos] = true
		nextPos := GetNextPos(&guard)
		if IsWithinBounds(nextPos, guardMap) && guardMap[nextPos.line][nextPos.char] == "#" {
			TurnGuard(&guard)
		} else if IsWithinBounds(nextPos, guardMap) {
			guard.pos = nextPos
		} else {
			break
		}
	}
	for pos := range visited {
		guardMap[pos.line][pos.char] = "X"
		visitedPoses = append(visitedPoses, pos)
	}
	return Part1{guardMap, visitedPoses}
}

func IsWithinBounds(pos Pos, guardMap [][]string) bool {
	return pos.line >= 0 && pos.line < len(guardMap) && pos.char >= 0 && pos.char < len(guardMap[pos.line])
}

func GetNextPos(guard *Guard) Pos {
	pos := guard.pos
	switch guard.direction {
	case North:
		return Pos{pos.line - 1, pos.char, 0}
	case East:
		return Pos{pos.line, pos.char + 1, 0}
	case South:
		return Pos{pos.line + 1, pos.char, 0}
	case West:
		return Pos{pos.line, pos.char - 1, 0}
	default:
		return pos
	}
}

func MoveGuard(guard *Guard) {
	switch guard.direction {
	case North:
		guard.pos.line--
	case East:
		guard.pos.char++
	case South:
		guard.pos.line++
	case West:
		guard.pos.char--
	}
}

func TurnGuard(guard *Guard) {
	switch guard.direction {
	case North:
		guard.direction = East
	case East:
		guard.direction = South
	case South:
		guard.direction = West
	case West:
		guard.direction = North
	}
}

func FindStartPos2(guardMap [][]string) Pos {
	for i, line := range guardMap {
		for j, pos := range line {
			if pos == "^" {
				return Pos{i, j, 0}
			}
		}
	}
	return Pos{-1, -1, -1}
}
