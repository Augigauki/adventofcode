package main

type Pos struct {
	line int
	char int
}

func InBounds(gardenMap [][]string, pos Pos) bool {
	return pos.line >= 0 && pos.line < len(gardenMap) && pos.char >= 0 && pos.char < len(gardenMap[0])
}

var Directions = [][]int{
	{-1, 0}, //Up
	{0, 1},  //Right
	{1, 0},  //Down
	{0, -1}, //Left
}

var AllDirections = [][]int{
	{-1, 0},  // UP
	{-1, -1}, // UP LEFT
	{-1, 1},  // UP RIGHT
	{0, -1},  // LEFT
	{0, 1},   // RIGHT
	{1, 0},   // DOWN
	{1, -1},  // DOWN LEFT
	{1, 1},   // DOWN RIGHT
}

var North = Pos{-1, 0}
var South = Pos{1, 0}
var East = Pos{0, 1}
var West = Pos{0, -1}
var NorthWest = Pos{-1, -1}
var NorthEast = Pos{-1, 1}
var SouthWest = Pos{1, -1}
var SouthEast = Pos{1, 1}
