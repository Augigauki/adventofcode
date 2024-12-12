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
