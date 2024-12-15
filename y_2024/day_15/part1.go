package main

import "fmt"

func Part1(goodsMap [][]string, bot Bot) {
	fmt.Println("Part 1")
	var PosMap = make(map[Pos]string)
	width := len(goodsMap[0])
	height := len(goodsMap)
	for i, line := range goodsMap {
		for j, char := range line {
			PosMap[Pos{i, j}] = string(char)
		}
	}
	MoveBot(PosMap, bot, width, height)
	//fmt.Println("PosMap: ", PosMap)
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
