package main

type Pos struct {
	line int
	char int
}

func InBounds(gardenMap [][]string, pos Pos) bool {
	return pos.line >= 0 && pos.line < len(gardenMap) && pos.char >= 0 && pos.char < len(gardenMap[0])
}

var Directions = [][]int{
	/* North */
	{-1, 0},
	/* East */
	{0, 1},
	/* South */
	{1, 0},
	/* West */
	{0, -1},
}
