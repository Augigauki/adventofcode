package main

type Pos struct {
	x, y int
}

type Reindeer struct {
	startPos Pos
	pos      Pos
	dir      Direction
	paths    []Path
}

type Path struct {
	positions []Pos
	score     int
}

type Direction struct {
	x, y int
	name string
}

var Directions = []Direction{{0, -1, "N"}, {1, 0, "E"}, {0, 1, "S"}, {-1, 0, "W"}}
