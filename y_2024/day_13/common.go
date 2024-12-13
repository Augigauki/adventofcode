package main

type Pos struct {
	line int
	char int
}

type ClawMachine struct {
	A     Button
	B     Button
	Prize Pos
	Pos   Pos
}

type MathMachine struct {
	Ax, Ay int
	Bx, By int
	PrizeX int
	PrizeY int
}

type Button struct {
	lineIncr int
	charIncr int
	counter  int
}
