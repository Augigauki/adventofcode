package main

import (
	"fmt"
	"strings"
)

func GetTotalGpsSum(goodsMap [][]string, bot Bot) {
	//fmt.Println("Part 1")
	fmt.Println("Part 2")

	var PosMap = make(map[Pos]string)
	wideMap := expandMap(goodsMap)
	fmt.Println("Wide map: ")
	for i, line := range wideMap {
		fmt.Println(line)
		for j, char := range line {
			if char == "@" {
				bot.Pos = Pos{i, j}
			}
		}
	}
	width := len(wideMap[0])
	height := len(wideMap)
	for i, line := range wideMap {
		for j, char := range line {
			PosMap[Pos{i, j}] = char
		}
	}

	MoveBotWiderMap(PosMap, bot, width, height)
	//fmt.Println("PosMap: ", PosMap)
}

func expandMap(origMap [][]string) [][]string {
	fmt.Println("Expanding map...")
	wideMap := [][]string{}
	for i, line := range origMap {
		for j, char := range line {
			if char == "#" {
				origMap[i][j] = "##"
			}
			if char == "." {
				origMap[i][j] = ".."
			}
			if char == "O" {
				origMap[i][j] = "[]"
			}
			if char == "@" {
				origMap[i][j] = "@."
			}
		}
	}
	for _, line := range origMap {
		var wideLine []string
		for _, char := range line {
			chars := strings.Split(char, "")
			wideLine = append(wideLine, chars...)
		}
		wideMap = append(wideMap, wideLine)
	}
	return wideMap
}

func MoveBotWiderMap(posMap map[Pos]string, bot Bot, width, height int) {
	fmt.Println("Moving bot in wider map")
	/* fmt.Println("PosMap:")
	for pos, value := range posMap {
		fmt.Println(pos, value)
	} */
	for _, mov := range bot.Movements {
		fmt.Println("\nBot pos: ", bot.Pos)
		fmt.Println("Bot mov: ", mov)
		nextPos := Pos{bot.Pos.Char + mov.Dir[1], bot.Pos.Line + mov.Dir[0]}
		fmt.Println("Next pos: ", nextPos)
		if posMap[nextPos] == "." {
			fmt.Println("Next pos: ", nextPos, " value: ", posMap[nextPos])
			fmt.Println("Next pos is empty, moving bot!")
			posMap[bot.Pos] = "."
			posMap[nextPos] = "@"
			bot.Pos = nextPos
			//printMap(posMap, width, height)

		} else if posMap[nextPos] == "#" {
			fmt.Println("Next pos is wall")

		} else if posMap[nextPos] == "[" || posMap[nextPos] == "]" {
			fmt.Println("Next pos is a box!")
			posMap = CheckIfWideBoxCanBeMoved(posMap, mov, &bot)

		} else {
			fmt.Println("Something is wrong with next pos...")
			fmt.Printf("Next pos: %v, value: %v\n", nextPos, posMap[nextPos])
		}
		printMap(posMap, width, height)
	}
	printMap(posMap, width, height)
	sum := 0
	for boxPos, value := range posMap {
		if value == "[" {
			gps := 0
			//fmt.Println("Box at pos: ", boxPos)
			gps = (100 * boxPos.Char) + boxPos.Line
			sum += gps
		}
	}
	fmt.Println("Sum of GPS coordinates: ", sum)
}

func MoveBot(posMap map[Pos]string, bot Bot, width, height int) {
	fmt.Println("Moving bot")
	for _, mov := range bot.Movements {
		fmt.Println("\nBot pos: ", bot.Pos)
		fmt.Println("Bot mov: ", mov)
		nextPos := Pos{bot.Pos.Char + mov.Dir[1], bot.Pos.Line + mov.Dir[0]}
		fmt.Println("Next pos: ", nextPos)
		if posMap[nextPos] == "." {
			fmt.Println("Next pos: ", nextPos, " value: ", posMap[nextPos])
			fmt.Println("Next pos is empty, moving bot!")
			posMap[bot.Pos] = "."
			posMap[nextPos] = "@"
			bot.Pos = nextPos
			//printMap(posMap, width, height)

		} else if posMap[nextPos] == "#" {
			fmt.Println("Next pos is wall")

		} else if posMap[nextPos] == "O" {
			fmt.Println("Next pos is a box!")
			posMap = CheckIfBoxCanBeMoved(posMap, mov, &bot)

		} else {
			fmt.Println("Something is wrong with next pos...")
			fmt.Printf("Next pos: %v, value: %v\n", nextPos, posMap[nextPos])
		}
		//printMap(posMap, width, height)
	}
	printMap(posMap, width, height)
	sum := 0
	//fmt.Println("PosMap: ", posMap)
	for boxPos, value := range posMap {
		if value == "O" {
			gps := 0
			//fmt.Println("Box at pos: ", boxPos)
			gps = (100 * boxPos.Char) + boxPos.Line
			sum += gps
		}
	}
	fmt.Println("Sum of GPS coordinates: ", sum)
}

func CheckIfWideBoxCanBeMoved(posMap map[Pos]string, dir Direction, bot *Bot) map[Pos]string {
	fmt.Println("Checking if wide box can be moved...")

	// Identify the position of the first box part
	origPos := bot.Pos
	nextPos := Pos{bot.Pos.Char + dir.Dir[1], bot.Pos.Line + dir.Dir[0]}

	// Identify all parts of the wide box
	var boxParts []Pos
	if posMap[nextPos] == "[" {
		// Horizontal wide box (left-to-right)
		boxParts = []Pos{nextPos, Pos{nextPos.Char + 1, nextPos.Line}}
	} else if posMap[nextPos] == "]" {
		// Horizontal wide box (right-to-left)
		boxParts = []Pos{nextPos, Pos{nextPos.Char - 1, nextPos.Line}}
	} else if posMap[nextPos] == "." {
		fmt.Println("No box to move.")
		return posMap
	} else {
		fmt.Println("Unexpected state: Box not detected!")
		return posMap
	}

	// Check if all parts of the box can move
	var newBoxPositions []Pos
	for _, part := range boxParts {
		newPos := Pos{part.Char + dir.Dir[1], part.Line + dir.Dir[0]}
		if posMap[newPos] != "." {
			fmt.Println("Wide box cannot be moved. Obstacle detected.")
			return posMap
		}
		newBoxPositions = append(newBoxPositions, newPos)
	}

	// Move the wide box
	for _, part := range boxParts {
		posMap[part] = "." // Clear the old box positions
	}
	for i, newPos := range newBoxPositions {
		if i == 0 {
			posMap[newPos] = "[" // First part of the box
		} else {
			posMap[newPos] = "]" // Second part of the box
		}
	}

	// Move the robot
	posMap[origPos] = "."                                               // Clear the robot's old position
	bot.Pos = Pos{bot.Pos.Char + dir.Dir[1], bot.Pos.Line + dir.Dir[0]} // Update robot's position
	posMap[bot.Pos] = "@"                                               // Place the robot at the new position

	return posMap
}

func findNextLineWideBoxes(posMap map[Pos]string, botPos Pos, dir Direction) map[Pos]string {
	fmt.Println("Finding next line wide boxes...")
	nextPos := Pos{botPos.Char + dir.Dir[1], botPos.Line + dir.Dir[0]}
	boxesMap := make(map[Pos]string)
	boxesMap[nextPos] = posMap[nextPos]
	leftPos := Pos{nextPos.Char - 1, nextPos.Line}
	rightPos := Pos{nextPos.Char + 1, nextPos.Line}
	for posMap[leftPos] == "[" || posMap[leftPos] == "]" {
		boxesMap[leftPos] = posMap[leftPos]
		leftPos = Pos{leftPos.Char - 1, leftPos.Line}
	}
	for posMap[rightPos] == "[" || posMap[rightPos] == "]" {
		boxesMap[rightPos] = posMap[rightPos]
		rightPos = Pos{rightPos.Char + 1, rightPos.Line}
	}
	return boxesMap
}

func checkNextLine(posMap map[Pos]string, boxesMap map[Pos]string, line int) bool {
	fmt.Println("Checking next line...")
	nextLineBlocked := false
	for pos := range boxesMap {
		if posMap[Pos{pos.Char, line}] == "#" {
			nextLineBlocked = true
		}
	}
	return nextLineBlocked
}

func CheckIfBoxCanBeMoved(posMap map[Pos]string, dir Direction, bot *Bot) map[Pos]string {
	fmt.Println("Checking if box can be moved...")
	origPos := bot.Pos
	nextPos := Pos{bot.Pos.Char + dir.Dir[1], bot.Pos.Line + dir.Dir[0]}
	origNextPos := nextPos
	boxesMap := make(map[Pos]string)
	for posMap[nextPos] == "O" {
		nextPos = Pos{nextPos.Char + dir.Dir[1], nextPos.Line + dir.Dir[0]}
		boxesMap[nextPos] = "O"
	}
	fmt.Println("Done checking boxes")
	fmt.Println("Boxes map: ", boxesMap)
	if posMap[nextPos] == "." {
		fmt.Println("Box can be moved")
		i := len(boxesMap)
		for boxPos := range boxesMap {
			fmt.Println("boxPos: ", boxPos)
			if i == 0 {
				posMap[boxPos] = "O"
			} else {
				posMap[boxPos] = "O"
			}
			i--
		}
		bot.Pos = origNextPos
		posMap[origPos] = "."
		posMap[bot.Pos] = "@"
	} else {
		fmt.Println("Box can't be moved")
	}
	return posMap

}

func printMap(posMap map[Pos]string, width, height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(posMap[Pos{i, j}])
		}
		fmt.Println()
	}
}
