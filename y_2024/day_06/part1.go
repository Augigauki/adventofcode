package main

import "fmt"

type Position struct {
	line int
	char int
}

type TraversedMap struct {
	startPos     Position
	traversedMap [][]string
	stuckInLoop  bool
}

func TraverseMap(startPos Position, guardMap [][]string) TraversedMap {
	localMap := guardMap
	traversedPoses := []Position{}
	//fmt.Println("Guard is starting at position: ", startPos)
	/* fmt.Println("Position in LocalMap: ", localMap[startPos.line][startPos.char])
	for _, line := range localMap {
		fmt.Println(line)
	} */
	direction := "^"
	patrol := true
	pos := startPos
	hitManualObstacle := false
	stuckInLoop := false
	for patrol {
		//fmt.Println("Guard is at position: ", pos)
		/* Checking if guard is moving north */
		//fmt.Println("Is patrolling? ", patrol)
		if stuckInLoop {
			fmt.Println("Guard is stuck in a loop! Breaking out!")
			break
		}
		spot := localMap[pos.line][pos.char]
		if spot == "X" {
			//fmt.Println("Guard has been here before!")
			if direction == "^" {
				if pos.line > 0 {
					nextSpot := localMap[pos.line-1][pos.char]
					if nextSpot == "#" {
						direction = ">"

					} else if nextSpot == "O" {
						direction = ">"
						if !hitManualObstacle {
							hitManualObstacle = true
						} else {
							stuckInLoop = true
							break
						}
					} else {
						pos.line = pos.line - 1
					}
				}
			} else if direction == ">" {
				if pos.char < len(localMap[pos.line])-1 {
					nextSpot := localMap[pos.line][pos.char+1]
					if nextSpot == "#" {
						direction = "v"
						/* continue */
					} else if nextSpot == "O" {
						direction = "v"
						if !hitManualObstacle {
							hitManualObstacle = true
						} else {
							stuckInLoop = true
							break
						}
					} else {
						pos.char = pos.char + 1
					}
				}
			} else if direction == "v" {
				if pos.line < len(localMap)-1 {
					nextSpot := localMap[pos.line+1][pos.char]
					if nextSpot == "#" {
						direction = "<"
					} else if nextSpot == "O" {
						direction = "<"
						if !hitManualObstacle {
							hitManualObstacle = true
						} else {
							stuckInLoop = true
							break
						}
					} else {
						pos.line = pos.line + 1
					}
				}
			} else if direction == "<" {
				if pos.char > 0 {
					nextSpot := localMap[pos.line][pos.char-1]
					if nextSpot == "#" {
						direction = "^"
					} else if nextSpot == "O" {
						direction = "^"
						if !hitManualObstacle {
							hitManualObstacle = true
						} else {
							stuckInLoop = true
							break
						}
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
				nextSpot := localMap[pos.line-1][pos.char]
				//fmt.Println("Guard is trying to move north!")
				//fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */
				if AddToTraversed(traversedPoses, pos.line, pos.char) {
					traversedPoses = append(traversedPoses, Position{pos.line, pos.char})
					localMap[pos.line][pos.char] = "X"
				}
				/* Checking if next location is open */
				if nextSpot == "." {
					//fmt.Println("Next space is free, moving to: ", pos.line, pos.char-1)
					localMap[pos.line][pos.char] = "X"
					pos.line = pos.line - 1
					/* Checking if next location is blocked */
				} else if nextSpot == "#" {
					//fmt.Println("Next space is blocked, changing direction!")
					direction = ">"
					localMap[pos.line][pos.char] = "X"
					//pos.char = pos.char + 1
					/* Checking if next location has been visited before */
				} else if nextSpot == "X" {
					//fmt.Println("Guard has been here before!")
					//pos.line = pos.line - 1
				} else if nextSpot == "O" {
					if !hitManualObstacle {
						direction = ">"
						hitManualObstacle = true
						fmt.Println("Guard hit manual obstacle!")
						localMap[pos.line][pos.char] = "X"
					} else {
						fmt.Println("Guard hit manual obstacle again!")
						stuckInLoop = true
						patrol = false
					}
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
				nextSpot := localMap[pos.line][pos.char+1]
				//fmt.Println("Guard is trying to move east!")
				//fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */
				if AddToTraversed(traversedPoses, pos.line, pos.char) {
					traversedPoses = append(traversedPoses, Position{pos.line, pos.char})
					localMap[pos.line][pos.char] = "X"
				}
				/* Checking if next location is open */
				if nextSpot == "." {
					//fmt.Println("Next space is free, moving to: ", pos.line, pos.char+1)
					localMap[pos.line][pos.char] = "X"
					//pos.char = pos.char + 1
					/* Checking if next location is blocked */
				} else if nextSpot == "#" {
					//fmt.Println("Next space is blocked, changing direction!")
					direction = "v"
					localMap[pos.line][pos.char] = "X"
					//pos.line = pos.line + 1
					/* Checking if next location has been visited before */
				} else if nextSpot == "O" {
					if !hitManualObstacle {
						direction = "v"
						hitManualObstacle = true
						fmt.Println("Guard hit manual obstacle!")
						localMap[pos.line][pos.char] = "X"
					} else {
						fmt.Println("Guard hit manual obstacle again!")
						stuckInLoop = true
						patrol = false
					}
				} else if nextSpot == "X" {
					//fmt.Println("Guard has been here before!")
					//pos.char = pos.char + 1
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
				nextSpot := localMap[pos.line+1][pos.char]
				//fmt.Println("Guard is trying to move south!")
				//fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */
				if AddToTraversed(traversedPoses, pos.line, pos.char) {
					traversedPoses = append(traversedPoses, Position{pos.line, pos.char})
					localMap[pos.line][pos.char] = "X"
				}
				/* Checking if next location is open */
				if nextSpot == "." {
					//fmt.Println("Next space is free, moving to: ", pos.line+1, pos.char)
					localMap[pos.line][pos.char] = "X"
					pos.line = pos.line + 1
					/* Checking if next location is blocked */
				} else if nextSpot == "#" {
					//fmt.Println("Next space is blocked, changing direction!")
					direction = "<"
					localMap[pos.line][pos.char] = "X"
					//pos.char = pos.char - 1
					/* Checking if next location has been visited before */
				} else if nextSpot == "O" {
					if !hitManualObstacle {
						direction = "<"
						hitManualObstacle = true
						fmt.Println("Guard hit manual obstacle!")
						localMap[pos.line][pos.char] = "X"
					} else {
						fmt.Println("Guard hit manual obstacle again!")
						stuckInLoop = true
						patrol = false
					}
				} else if nextSpot == "X" {
					//fmt.Println("Guard has been here before!")
					//pos.line = pos.line + 1
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
				nextSpot := localMap[pos.line][pos.char-1]
				//fmt.Println("Guard is trying to move west!")
				//fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */
				if AddToTraversed(traversedPoses, pos.line, pos.char) {
					traversedPoses = append(traversedPoses, Position{pos.line, pos.char})
					localMap[pos.line][pos.char] = "X"
				}
				/* Checking if next location is open */
				if nextSpot == "." {
					//fmt.Println("Next space is free, moving to: ", pos.line, pos.char-1)
					localMap[pos.line][pos.char] = "X"
					pos.char = pos.char - 1
					/* Checking if next location is blocked */
				} else if nextSpot == "#" {
					//fmt.Println("Next space is blocked, changing direction!")
					direction = "^"
					localMap[pos.line][pos.char] = "X"
					//pos.line = pos.line - 1
					/* Checking if next location has been visited before */
				} else if nextSpot == "O" {
					if !hitManualObstacle {
						direction = "^"
						hitManualObstacle = true
						fmt.Println("Guard hit manual obstacle!")
						localMap[pos.line][pos.char] = "X"
					} else {
						fmt.Println("Guard hit manual obstacle again!")
						stuckInLoop = true
						patrol = false
					}
				} else if nextSpot == "X" {
					//fmt.Println("Guard has been here before!")
					//pos.char = pos.char - 1
				}
			} else {
				fmt.Println("Guard left the area west!")
				localMap[pos.line][pos.char] = "X"
				pos.char = pos.char - 1
				patrol = false
			}
		}

	}

	/* total_x := 0
	for _, line := range localMap {
		for _, char := range line {
			if char == "X" {
				total_x++
			}
		}
		fmt.Println(line)
	}
	fmt.Println("Traversed poses: ", total_x) */
	return TraversedMap{startPos, localMap, stuckInLoop}
}

func FindStartPos(guardMap [][]string) Position {
	for i, line := range guardMap {
		for j, pos := range line {
			if pos == "^" {
				return Position{i, j}
			}
		}
	}
	return Position{-1, -1}
}

func AddToTraversed(traversedPoses []Position, line, char int) bool {

	newPos := Position{line, char}
	for _, pos := range traversedPoses {
		if pos == newPos {
			return false
		}
	}
	//traversedPoses = append(traversedPoses, newPos)
	return true
}
