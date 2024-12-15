package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	Char, Line int
}

type Bot struct {
	Pos       Pos
	Dir       Direction
	Movements []Direction
}

type Direction struct {
	Name string
	Dir  []int
}

type Box struct {
	Pos      Pos
	GPSCoord int
}

type Wall struct {
	Pos Pos
}

var Directions = []Direction{{"N", []int{0, -1}}, {"E", []int{1, 0}}, {"S", []int{0, 1}}, {"W", []int{-1, 0}}}

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file: ", fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var goodsMap [][]string
	var movements string
	var botMovs []Direction
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			if line[0] == '#' {
				var newLine []string
				for _, char := range line {
					newLine = append(newLine, string(char))
				}
				goodsMap = append(goodsMap, newLine)
			} else {
				movements += line
			}
		}

	}

	fmt.Println("Goods map:")
	var botStartPos Pos
	for i, line := range goodsMap {
		fmt.Println(line)
		for j, char := range line {
			if char == "@" {
				botStartPos = Pos{i, j}
			}
		}
	}
	fmt.Println("Bot start position: ", botStartPos)

	for _, mov := range movements {
		newMov := Direction{}
		if mov == '^' {
			newMov = Directions[0]
		} else if mov == '>' {
			newMov = Directions[1]
		} else if mov == 'v' {
			newMov = Directions[2]
		} else if mov == '<' {
			newMov = Directions[3]
		}

		botMovs = append(botMovs, newMov)
	}
	bot := Bot{botStartPos, Directions[0], botMovs}
	Part1(goodsMap, bot)
	//fmt.Println("Movements: ", botMovs, " Length: ", len(movements))

}
