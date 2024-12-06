package main

import "fmt"

type Position struct {
	line int
	char int
}

func TraverseMap(guardMap [][]string) {
	localMap := guardMap
	traversedPoses := []Position{}
	startPos := findStartPos(guardMap)
	fmt.Println("Guard is starting at position: ", startPos)
	fmt.Println("Position in LocalMap: ", localMap[startPos.line][startPos.char])
	for _, line := range localMap {
		fmt.Println(line)
	}
	direction := "^"
	patrol := true
	pos := startPos
	for patrol {
		//fmt.Println("Guard is at position: ", pos)
		/* Checking if guard is moving north */
		//fmt.Println("Is patrolling? ", patrol)
		spot := localMap[pos.line][pos.char]
		if spot == "X" {
			//fmt.Println("Guard has been here before!")
			if direction == "^" {
				if pos.line > 0 {
					if localMap[pos.line-1][pos.char] == "#" {
						direction = ">"
						/* continue */
					} else {
						pos.line = pos.line - 1
					}
				}
			} else if direction == ">" {
				if pos.char < len(localMap[pos.line])-1 {
					if localMap[pos.line][pos.char+1] == "#" {
						direction = "v"
						/* continue */
					} else {
						pos.char = pos.char + 1
					}
				}
			} else if direction == "v" {
				if pos.line < len(localMap)-1 {
					if localMap[pos.line+1][pos.char] == "#" {
						direction = "<"
						/* continue */
					} else {
						pos.line = pos.line + 1
					}
				}
			} else if direction == "<" {
				if pos.char > 0 {
					if localMap[pos.line][pos.char-1] == "#" {
						direction = "^"
						/* continue */
					} else {
						pos.char = pos.char - 1
					}
				}
			}
			/* continue */
		}
		if direction == "^" {
			/* Checking if guard isn't at northmost edge */
			if pos.line > 0 {
				fmt.Println("Guard is trying to move north!")
				fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */
				if addToTraversed(traversedPoses, pos.line, pos.char) {
					traversedPoses = append(traversedPoses, Position{pos.line, pos.char})
					localMap[pos.line][pos.char] = "X"
				}
				/* Checking if next location is open */
				if localMap[pos.line-1][pos.char] == "." {
					fmt.Println("Next space is free, moving to: ", pos.line, pos.char-1)
					localMap[pos.line][pos.char] = "X"
					pos.line = pos.line - 1
					/* Checking if next location is blocked */
				} else if localMap[pos.line-1][pos.char] == "#" {
					fmt.Println("Next space is blocked, changing direction!")
					direction = ">"
					localMap[pos.line][pos.char] = "X"
					//pos.char = pos.char + 1
					/* Checking if next location has been visited before */
				} else if localMap[pos.line-1][pos.char] == "X" {
					fmt.Println("Guard has been here before!")
					pos.line = pos.line - 1
				}
			} else {
				fmt.Println("Guard left the area north!")
				localMap[pos.line][pos.char] = "X"
				pos.line = pos.line - 1
				patrol = false
			}
		}
		if direction == ">" {
			/* Checking if guard isn't at eastmost edge */
			if pos.char < len(localMap[pos.line])-1 {
				fmt.Println("Guard is trying to move east!")
				fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */
				if addToTraversed(traversedPoses, pos.line, pos.char) {
					traversedPoses = append(traversedPoses, Position{pos.line, pos.char})
					localMap[pos.line][pos.char] = "X"
				}
				/* Checking if next location is open */
				if localMap[pos.line][pos.char+1] == "." {
					fmt.Println("Next space is free, moving to: ", pos.line, pos.char+1)
					localMap[pos.line][pos.char] = "X"
					//pos.char = pos.char + 1
					/* Checking if next location is blocked */
				} else if localMap[pos.line][pos.char+1] == "#" {
					fmt.Println("Next space is blocked, changing direction!")
					direction = "v"
					localMap[pos.line][pos.char] = "X"
					//pos.line = pos.line + 1
					/* Checking if next location has been visited before */
				} else if localMap[pos.line][pos.char+1] == "X" {
					fmt.Println("Guard has been here before!")
					pos.char = pos.char + 1
				}
			} else {
				fmt.Println("Guard left the area east!")
				localMap[pos.line][pos.char] = "X"
				pos.char = pos.char + 1
				patrol = false
			}
		}
		if direction == "v" {
			/* Checking if guard isn't at southmost edge */
			if pos.line < len(localMap)-1 {
				fmt.Println("Guard is trying to move south!")
				fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */
				if addToTraversed(traversedPoses, pos.line, pos.char) {
					traversedPoses = append(traversedPoses, Position{pos.line, pos.char})
					localMap[pos.line][pos.char] = "X"
				}
				/* Checking if next location is open */
				if localMap[pos.line+1][pos.char] == "." {
					fmt.Println("Next space is free, moving to: ", pos.line+1, pos.char)
					localMap[pos.line][pos.char] = "X"
					pos.line = pos.line + 1
					/* Checking if next location is blocked */
				} else if localMap[pos.line+1][pos.char] == "#" {
					fmt.Println("Next space is blocked, changing direction!")
					direction = "<"
					localMap[pos.line][pos.char] = "X"
					//pos.char = pos.char - 1
					/* Checking if next location has been visited before */
				} else if localMap[pos.line+1][pos.char] == "X" {
					fmt.Println("Guard has been here before!")
					pos.line = pos.line + 1
				}
			} else {
				fmt.Println("Guard left the area south!")
				localMap[pos.line][pos.char] = "X"
				pos.line = pos.line + 1
				patrol = false
			}
		}
		if direction == "<" {
			/* Checking if guard isn't at westmost edge */
			if pos.char > 0 {
				fmt.Println("Guard is trying to move west!")
				fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */
				if addToTraversed(traversedPoses, pos.line, pos.char) {
					traversedPoses = append(traversedPoses, Position{pos.line, pos.char})
					localMap[pos.line][pos.char] = "X"
				}
				/* Checking if next location is open */
				if localMap[pos.line][pos.char-1] == "." {
					fmt.Println("Next space is free, moving to: ", pos.line, pos.char-1)
					localMap[pos.line][pos.char] = "X"
					pos.char = pos.char - 1
					/* Checking if next location is blocked */
				} else if localMap[pos.line][pos.char-1] == "#" {
					fmt.Println("Next space is blocked, changing direction!")
					direction = "^"
					localMap[pos.line][pos.char] = "X"
					//pos.line = pos.line - 1
					/* Checking if next location has been visited before */
				} else if localMap[pos.line][pos.char-1] == "X" {
					fmt.Println("Guard has been here before!")
					pos.char = pos.char - 1
				}
			} else {
				fmt.Println("Guard left the area west!")
				localMap[pos.line][pos.char] = "X"
				pos.char = pos.char - 1
				patrol = false
			}
		}

	}
	fmt.Printf("\n::TRAVERSED MAP::\n")
	total_x := 0
	for _, line := range localMap {
		for _, char := range line {
			if char == "X" {
				total_x++
			}
		}
		fmt.Println(line)
	}
	fmt.Println("Traversed poses: ", total_x)
}

func findStartPos(guardMap [][]string) Position {
	for i, line := range guardMap {
		for j, pos := range line {
			if pos == "^" {
				return Position{i, j}
			}
		}
	}
	return Position{-1, -1}
}

func addToTraversed(traversedPoses []Position, line, char int) bool {
	newPos := Position{line, char}
	for _, pos := range traversedPoses {
		if pos == newPos {
			return false
		}
	}
	//traversedPoses = append(traversedPoses, newPos)
	return true
}
