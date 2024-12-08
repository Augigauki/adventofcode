package main

import "fmt"

func TraverseWithObstaclesAndGuard(guard Guard, guardMap [][]string) {

}

func isOutOfBounds(pos Pos, guardMap [][]string) bool {
	return pos.char < 0 || pos.char >= len(guardMap[0]) || pos.line < 0 || pos.line >= len(guardMap)
}

func getNextPos(g Guard) Pos {
	if g.direction == "^" {
		return Pos{line: g.pos.line - 1, char: g.pos.char, visits: g.pos.visits}
	}
	if g.direction == ">" {
		return Pos{line: g.pos.line, char: g.pos.char + 1, visits: g.pos.visits}
	}
	if g.direction == "v" {
		return Pos{line: g.pos.line + 1, char: g.pos.char, visits: g.pos.visits}
	}
	if g.direction == "<" {
		return Pos{line: g.pos.line, char: g.pos.char - 1, visits: g.pos.visits}
	}
	return Pos{-1, -1, -1}
}

func turn(g Guard) {
	if g.direction == "^" {
		g.direction = ">"
	} else if g.direction == ">" {
		g.direction = "v"
	} else if g.direction == "v" {
		g.direction = "<"
	} else if g.direction == "<" {
		g.direction = "^"
	}
}

func move(g Guard) {
	newPos := getNextPos(g)
	fmt.Println(newPos)

}

func recordPath(g Guard) {

}
