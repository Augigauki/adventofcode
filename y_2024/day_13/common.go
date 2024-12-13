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

type Button struct {
	lineIncr int
	charIncr int
	counter  int
}
