package main

import "fmt"

var totalObstacles int = 0

func TraverseWithObstacles(startPos Position, guardMap [][]string) {
	localMap := guardMap
	checkedPoses := []Position{}
	//startPos := FindStartPos(guardMap)
	/* fmt.Println("Guard is starting at position: ", startPos)
	fmt.Println("Position in LocalMap: ", localMap[startPos.line][startPos.char])
	for _, line := range localMap {
		fmt.Println(line)
	} */
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
					nextSpot := localMap[pos.line-1][pos.char]
					if nextSpot == "#" {
						direction = ">"
						/* continue */
					} else if nextSpot == "X" {
						obstacleMap := localMap
						obstacleMap[pos.line-1][pos.char] = "O"
						if TraverseMap(startPos, obstacleMap).stuckInLoop {
							if AddToChecked(checkedPoses, pos.line-1, pos.char) {
								checkedPoses = append(checkedPoses, Position{pos.line - 1, pos.char})
								totalObstacles++
								fmt.Println("\n::::OBSTACLE MAP")
								for _, line := range obstacleMap {
									fmt.Println(line)
								}
							} else {
								fmt.Println("Not stuck in loop with obstacle at: ", pos.line-1, pos.char)
							}
						}
						obstacleMap[pos.line-1][pos.char] = "X"
						pos.line = pos.line - 1
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
					} else if nextSpot == "X" {
						obstacleMap := localMap
						obstacleMap[pos.line][pos.char+1] = "O"
						if TraverseMap(startPos, obstacleMap).stuckInLoop {
							if AddToChecked(checkedPoses, pos.line, pos.char+1) {
								checkedPoses = append(checkedPoses, Position{pos.line, pos.char + 1})
								totalObstacles++
								fmt.Println("\n::::OBSTACLE MAP")
								for _, line := range obstacleMap {
									fmt.Println(line)
								}
							}
						}
						obstacleMap[pos.line][pos.char+1] = "X"
						pos.char = pos.char + 1
					} else {
						pos.char = pos.char + 1
					}
				}
			} else if direction == "v" {
				if pos.line < len(localMap)-1 {
					nextSpot := localMap[pos.line+1][pos.char]
					if nextSpot == "#" {
						direction = "<"
						/* continue */
					} else if nextSpot == "X" {
						obstacleMap := localMap
						obstacleMap[pos.line+1][pos.char] = "O"
						if TraverseMap(startPos, obstacleMap).stuckInLoop {
							if AddToChecked(checkedPoses, pos.line+1, pos.char) {
								checkedPoses = append(checkedPoses, Position{pos.line + 1, pos.char})
								totalObstacles++
								fmt.Println("\n::::OBSTACLE MAP")
								for _, line := range obstacleMap {
									fmt.Println(line)
								}
							}
						}
						obstacleMap[pos.line+1][pos.char] = "X"
						pos.line = pos.line + 1
					} else {
						pos.line = pos.line + 1
					}
				}
			} else if direction == "<" {
				if pos.char > 0 {
					nextSpot := localMap[pos.line][pos.char-1]
					if nextSpot == "#" {
						direction = "^"
						/* continue */
					} else if nextSpot == "O" {
						obstacleMap := localMap
						obstacleMap[pos.line][pos.char-1] = "O"
						if TraverseMap(startPos, obstacleMap).stuckInLoop {
							if AddToChecked(checkedPoses, pos.line, pos.char-1) {
								checkedPoses = append(checkedPoses, Position{pos.line, pos.char - 1})
								totalObstacles++
								fmt.Println("\n::::OBSTACLE MAP")
								for _, line := range obstacleMap {
									fmt.Println(line)
								}
							}
						}
						obstacleMap[pos.line][pos.char-1] = "X"
						pos.char = pos.char - 1
					} else {
						pos.char = pos.char - 1
					}
				}
			}
		}
		if direction == "^" {
			/* Checking if guard isn't at northmost edge */
			if pos.line > 0 {
				//fmt.Println("Guard is trying to move north!")
				//fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */

				/* Checking if next location is open */
				if localMap[pos.line-1][pos.char] == "#" {
					//fmt.Println("Next space is blocked, changing direction!")
					direction = ">"
					/* Checking if next location has been visited before */
				} else if localMap[pos.line-1][pos.char] == "X" {
					//fmt.Println("Next spot is X - testing swapping with O")
					//fmt.Println("Guard is at position: ", pos)
					obstacleMap := localMap
					obstacleMap[pos.line-1][pos.char] = "O"
					if TraverseMap(startPos, obstacleMap).stuckInLoop {
						if AddToChecked(checkedPoses, pos.line-1, pos.char) {
							checkedPoses = append(checkedPoses, Position{pos.line - 1, pos.char})
							totalObstacles++
							fmt.Println("\n::::OBSTACLE MAP")
							for _, line := range obstacleMap {
								fmt.Println(line)
							}
						} else {
							fmt.Println("Not stuck in loop with obstacle at: ", pos.line-1, pos.char)
						}
					}
					obstacleMap[pos.line-1][pos.char] = "X"
					pos.line = pos.line - 1
				}
			} else {
				//fmt.Println("Guard left the area north!")
				localMap[pos.line][pos.char] = "X"
				pos.line = pos.line - 1
				patrol = false
			}
		}
		if direction == ">" {
			/* Checking if guard isn't at eastmost edge */
			if pos.char < len(localMap[pos.line])-1 {
				//fmt.Println("Guard is trying to move east!")
				//fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */

				/* Checking if next location is open */
				if localMap[pos.line][pos.char+1] == "#" {
					//	fmt.Println("Next space is blocked, changing direction!")
					direction = "v"
					localMap[pos.line][pos.char] = "X"
					//pos.line = pos.line + 1
					/* Checking if next location has been visited before */
				} else if localMap[pos.line][pos.char+1] == "X" {
					obstacleMap := localMap
					obstacleMap[pos.line][pos.char+1] = "O"
					if TraverseMap(startPos, obstacleMap).stuckInLoop {
						if AddToChecked(checkedPoses, pos.line, pos.char+1) {
							checkedPoses = append(checkedPoses, Position{pos.line, pos.char + 1})
							totalObstacles++
							fmt.Println("\n::::OBSTACLE MAP")
							for _, line := range obstacleMap {
								fmt.Println(line)
							}
						}
					}
					obstacleMap[pos.line][pos.char+1] = "X"
					pos.char = pos.char + 1
				}
			} else {
				//	fmt.Println("Guard left the area east!")
				localMap[pos.line][pos.char] = "X"
				pos.char = pos.char + 1
				patrol = false
			}
		}
		if direction == "v" {
			/* Checking if guard isn't at southmost edge */
			if pos.line < len(localMap)-1 {
				//	fmt.Println("Guard is trying to move south!")
				//	fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */

				/* Checking if next location is open */
				if localMap[pos.line+1][pos.char] == "#" {
					//		fmt.Println("Next space is blocked, changing direction!")
					direction = "<"
					localMap[pos.line][pos.char] = "X"
					//pos.char = pos.char - 1
					/* Checking if next location has been visited before */
				} else if localMap[pos.line+1][pos.char] == "X" {
					obstacleMap := localMap
					obstacleMap[pos.line+1][pos.char] = "O"
					if TraverseMap(startPos, obstacleMap).stuckInLoop {
						if AddToChecked(checkedPoses, pos.line+1, pos.char) {
							checkedPoses = append(checkedPoses, Position{pos.line + 1, pos.char})
							totalObstacles++
							fmt.Println("\n::::OBSTACLE MAP")
							for _, line := range obstacleMap {
								fmt.Println(line)
							}
						}

					}
					obstacleMap[pos.line+1][pos.char] = "X"
					pos.line = pos.line + 1
				}
			} else {
				//fmt.Println("Guard left the area south!")
				localMap[pos.line][pos.char] = "X"
				pos.line = pos.line + 1
				patrol = false
			}
		}
		if direction == "<" {
			/* Checking if guard isn't at westmost edge */
			if pos.char > 0 {
				//fmt.Println("Guard is trying to move west!")
				//fmt.Println("Guard is at position: ", pos)
				/* Checking if position has been traversed already */

				if localMap[pos.line][pos.char-1] == "#" {
					//fmt.Println("Next space is blocked, changing direction!")
					direction = "^"
					localMap[pos.line][pos.char] = "X"
					//pos.line = pos.line - 1
					/* Checking if next location has been visited before */
				} else if localMap[pos.line][pos.char-1] == "X" {
					obstacleMap := localMap
					obstacleMap[pos.line][pos.char-1] = "O"
					if TraverseMap(startPos, obstacleMap).stuckInLoop {
						if AddToChecked(checkedPoses, pos.line, pos.char-1) {
							checkedPoses = append(checkedPoses, Position{pos.line, pos.char - 1})
							totalObstacles++
							fmt.Println("\n::::OBSTACLE MAP")
							for _, line := range obstacleMap {
								fmt.Println(line)
							}
						}
					}
					obstacleMap[pos.line][pos.char-1] = "X"
					pos.char = pos.char - 1
				}
			} else {
				//fmt.Println("Guard left the area west!")
				localMap[pos.line][pos.char] = "X"
				pos.char = pos.char - 1
				patrol = false
			}
		}
	}
	fmt.Println("Total Obstacles: ", totalObstacles)
}

func AddToChecked(checkedPoses []Position, line, char int) bool {
	fmt.Println("Checked poses: ", checkedPoses)
	newPos := Position{line, char}
	for _, pos := range checkedPoses {
		if pos == newPos {
			fmt.Println("Position has been checked before! Returning false")
			return false
		}
	}
	fmt.Printf("Position %v has not been checked before! Returning true\n", newPos)
	return true
}
