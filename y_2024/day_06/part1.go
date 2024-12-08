package main

import (
	"fmt"
)

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
	width := len(localMap[0])
	height := len(localMap)
	limit := (width * height) * 4
	allObstacles := 0
	for _, line := range localMap {
		for _, char := range line {
			if char == "#" {
				allObstacles++
			}
		}
	}
	//fmt.Println("All obstacles: ", allObstacles)
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
	encounteredObstacles := []Position{}
	steps := 0
	for patrol {

		//fmt.Println("Guard is at position: ", pos)
		/* Checking if guard is moving north */
		//fmt.Println("Is patrolling? ", patrol)
		/* if len(encounteredObstacles) >= allObstacles*2 {
			fmt.Println("Found a # loop! Breaking out")
			stuckInLoop = true
			break
		} */
		fmt.Println("Steps and limit: ", steps, limit)
		if steps > limit {
			fmt.Println("Guard has taken too many steps! Breaking out!")
			stuckInLoop = true
			break
		}
		//fmt.Println("Encountered obstacles: ", encounteredObstacles)
		if stuckInLoop {
			stuckInLoop = true
			fmt.Println("Guard is stuck in a loop! Breaking out!")
			break
		}

		//fmt.Println("Encountered obstacles: ", encounteredObstacles)
		spot := localMap[pos.line][pos.char]
		if spot == "X" {
			//fmt.Println("Guard has been here before!")
			if direction == "^" {
				if pos.line > 0 {
					nextSpot := localMap[pos.line-1][pos.char]
					if nextSpot == "#" {
						direction = ">"
						encounteredObstacles = append(encounteredObstacles, Position{pos.line - 1, pos.char})
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
						steps++
					}
				}
			} else if direction == ">" {
				if pos.char < len(localMap[pos.line])-1 {
					nextSpot := localMap[pos.line][pos.char+1]
					if nextSpot == "#" {
						direction = "v"
						encounteredObstacles = append(encounteredObstacles, Position{pos.line, pos.char + 1})
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
						steps++
					}
				}
			} else if direction == "v" {
				if pos.line < len(localMap)-1 {
					nextSpot := localMap[pos.line+1][pos.char]
					if nextSpot == "#" {
						direction = "<"
						encounteredObstacles = append(encounteredObstacles, Position{pos.line + 1, pos.char})
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
						steps++
					}
				}
			} else if direction == "<" {
				if pos.char > 0 {
					nextSpot := localMap[pos.line][pos.char-1]
					if nextSpot == "#" {
						direction = "^"
						encounteredObstacles = append(encounteredObstacles, Position{pos.line, pos.char - 1})
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
						steps++
					}
				}
			}
			/* continue */
		}
		if direction == "^" {
			if pos.line > 0 {
				nextSpot := localMap[pos.line-1][pos.char]
				//fmt.Println("Guard is trying to move north!")
				//fmt.Println("Guard is at position: ", pos)
				if AddToTraversed(traversedPoses, pos.line, pos.char) {
					traversedPoses = append(traversedPoses, Position{pos.line, pos.char})
					localMap[pos.line][pos.char] = "X"
				}
				if nextSpot == "." {
					//fmt.Println("Next space is free, moving to: ", pos.line, pos.char-1)
					localMap[pos.line][pos.char] = "X"
					pos.line = pos.line - 1
					steps++
				} else if nextSpot == "#" {
					//fmt.Println("Next space is blocked, changing direction!")
					direction = ">"
					localMap[pos.line][pos.char] = "X"
					encounteredObstacles = append(encounteredObstacles, Position{pos.line - 1, pos.char})
					//pos.char = pos.char + 1
				} else if nextSpot == "X" {
					//fmt.Println("Guard has been here before!")
					//pos.line = pos.line - 1
				} else if nextSpot == "O" {
					if !hitManualObstacle {
						direction = ">"
						hitManualObstacle = true
						//fmt.Println("Guard hit manual obstacle!")
						localMap[pos.line][pos.char] = "X"
					} else {
						fmt.Println("Guard hit manual obstacle again! Direction: ", direction)
						fmt.Println("Guard is at position: ", pos)
						fmt.Println("Next spot is: ", pos.line-1, pos.char)
						stuckInLoop = true
						patrol = false
					}
				}
			} else {
				//fmt.Println("Guard left the area north!")
				localMap[pos.line][pos.char] = "X"
				pos.line = pos.line - 1
				steps++
				patrol = false
			}
		}
		if direction == ">" {
			if pos.char < len(localMap[pos.line])-1 {
				nextSpot := localMap[pos.line][pos.char+1]
				//fmt.Println("Guard is trying to move east!")
				//fmt.Println("Guard is at position: ", pos)
				if AddToTraversed(traversedPoses, pos.line, pos.char) {
					traversedPoses = append(traversedPoses, Position{pos.line, pos.char})
					localMap[pos.line][pos.char] = "X"
				}
				/* Checking if next location is open */
				if nextSpot == "." {
					//fmt.Println("Next space is free, moving to: ", pos.line, pos.char+1)
					localMap[pos.line][pos.char] = "X"
					pos.char = pos.char + 1
					steps++
					/* Checking if next location is blocked */
				} else if nextSpot == "#" {
					//fmt.Println("Next space is blocked, changing direction!")
					direction = "v"
					localMap[pos.line][pos.char] = "X"
					encounteredObstacles = append(encounteredObstacles, Position{pos.line, pos.char + 1})
					//pos.line = pos.line + 1
					/* Checking if next location has been visited before */
				} else if nextSpot == "O" {
					if !hitManualObstacle {
						direction = "v"
						hitManualObstacle = true
						//fmt.Println("Guard hit manual obstacle!")
						localMap[pos.line][pos.char] = "X"
					} else {
						fmt.Println("Guard hit manual obstacle again! Direction: ", direction)
						fmt.Println("Guard is at position: ", pos)
						fmt.Println("Next spot is: ", pos.line, pos.char+1)
						stuckInLoop = true
						patrol = false
					}
				} else if nextSpot == "X" {
					//fmt.Println("Guard has been here before!")
					//pos.char = pos.char + 1
				}
			} else {
				//fmt.Println("Guard left the area east!")
				localMap[pos.line][pos.char] = "X"
				pos.char = pos.char + 1
				steps++
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
					steps++
					/* Checking if next location is blocked */
				} else if nextSpot == "#" {
					//encounteredObstacles = append(encounteredObstacles, Position{pos.line + 1, pos.char})
					//fmt.Println("Next space is blocked, changing direction!")
					direction = "<"
					localMap[pos.line][pos.char] = "X"
					encounteredObstacles = append(encounteredObstacles, Position{pos.line + 1, pos.char})
					//pos.char = pos.char - 1
					/* Checking if next location has been visited before */
				} else if nextSpot == "O" {
					if !hitManualObstacle {
						direction = "<"
						hitManualObstacle = true
						//fmt.Println("Guard hit manual obstacle!")
						localMap[pos.line][pos.char] = "X"
					} else {
						fmt.Println("Guard hit manual obstacle again! Direction: ", direction)
						fmt.Println("Guard is at position: ", pos)
						fmt.Println("Next spot is: ", pos.line+1, pos.char)
						stuckInLoop = true
						patrol = false
					}
				} else if nextSpot == "X" {
					//fmt.Println("Guard has been here before!")
					//pos.line = pos.line + 1
				}
			} else {
				//fmt.Println("Guard left the area south!")
				localMap[pos.line][pos.char] = "X"
				pos.line = pos.line + 1
				steps++
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
					steps++
					/* Checking if next location is blocked */
				} else if nextSpot == "#" {
					//fmt.Println("Next space is blocked, changing direction!")
					direction = "^"
					localMap[pos.line][pos.char] = "X"
					encounteredObstacles = append(encounteredObstacles, Position{pos.line, pos.char - 1})
					//pos.line = pos.line - 1
					/* Checking if next location has been visited before */
				} else if nextSpot == "O" {
					if !hitManualObstacle {
						direction = "^"
						hitManualObstacle = true
						//fmt.Println("Guard hit manual obstacle!")
						localMap[pos.line][pos.char] = "X"
					} else {
						fmt.Println("Guard hit manual obstacle again! Direction: ", direction)
						fmt.Println("Guard is at position: ", pos)
						fmt.Println("Next pos is: ", pos.line, pos.char-1)
						stuckInLoop = true
						patrol = false
					}
				} else if nextSpot == "X" {
					//fmt.Println("Guard has been here before!")
					//pos.char = pos.char - 1
				}
			} else {
				//fmt.Println("Guard left the area west!")
				localMap[pos.line][pos.char] = "X"
				pos.char = pos.char - 1
				steps++
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
