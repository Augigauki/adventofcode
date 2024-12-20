package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction struct {
	x, y int
}

type Pos struct {
	x, y int
}

var example = false

var Directions = []Direction{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

var startPos = Pos{0, 0}
var endPos = Pos{0, 0}
var width, height = 0, 0

func main() {
	part1()

}

/* Old code from first attempts */
func findCheats(path []Pos, racetrack map[Pos]string) map[int]int {
	cheatSavings := make(map[int]int)
	length := len(path)
	for i := 0; i < 1; i++ {
		cheatPos := path[i]
		//fmt.Println("Checking cheats for pos: ", cheatPos)
		for _, dir := range Directions {
			newPos := Pos{cheatPos.x + dir.x, cheatPos.y + dir.y}
			if newPos.x < 0 || newPos.x >= width || newPos.y < 0 || newPos.y >= height {
				continue
			}
			if racetrack[newPos] == "#" {
				for _, dir2 := range Directions {
					newPos2 := Pos{newPos.x + dir2.x, newPos.y + dir2.y}
					if newPos2.x < 0 || newPos2.x >= width || newPos2.y < 0 || newPos2.y >= height {
						continue
					}
					if racetrack[newPos2] == "#" {
						if newPos2 == newPos {
							continue
						}
						for _, dir3 := range Directions {
							newPos3 := Pos{newPos2.x + dir3.x, newPos2.y + dir3.y}
							if newPos3.x < 0 || newPos3.x >= width || newPos3.y < 0 || newPos3.y >= height {
								continue
							}
							if racetrack[newPos3] == "O" {
								if newPos3 == cheatPos {
									continue
								}
								//fmt.Println("New pos3: ", newPos3)
								pathPos := findPositionInPath(path, newPos3)
								//fmt.Println("Path pos: ", pathPos)
								newLength := getCheatScore(length, pathPos)
								saved := length - newLength - 2
								//fmt.Printf("Saved: length-1 (%v) - i (%v) - steps(%v)\n", length-1, pathPos, 2)
								cheatSavings[saved]++
								//fmt.Println("Found cheat!")
								continue
							}
						}
					}
					if racetrack[newPos2] == "O" {
						if newPos2 == cheatPos {
							continue
						}
						//fmt.Println("New pos2: ", newPos2)
						pathPos := findPositionInPath(path, newPos2)
						newLength := getCheatScore(length, pathPos)
						saved := length - newLength - 1
						//fmt.Printf("Saved: length-1 (%v) - i (%v) - steps(%v)\n", length-1, pathPos, 2)
						cheatSavings[saved]++
						//fmt.Println("Found cheat!")
						continue
					}
				}
			}

		}
	}
	return cheatSavings
}

func getCheatScore(endIndex, index int) int {
	return endIndex - index
}

func isWithinBounds(pos Pos) bool {
	return pos.x >= 0 && pos.x < width && pos.y >= 0 && pos.y < height
}

func isPathPosition(path []Pos, pos Pos) bool {
	for _, p := range path {
		if p == pos {
			return true
		}
	}
	return false
}

func findPositionInPath(path []Pos, pos Pos) int {
	for i, p := range path {
		if p == pos {
			return i
		}
	}
	return -1
}

func printMap(memSpace map[Pos]string) {
	counter := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if memSpace[Pos{x, y}] == "." {
				counter++
			}
			fmt.Print(memSpace[Pos{x, y}])
		}
		fmt.Println()
	}
	fmt.Println("Total number of empty spaces: ", counter)
}

func parseFile(fileName string) map[Pos]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error reading file: ", err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	memSpace := map[Pos]string{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)
		for x, char := range line {
			if char == 'S' {
				startPos = Pos{x, y}
				memSpace[Pos{x, y}] = "."
			} else if char == 'E' {
				endPos = Pos{x, y}
				memSpace[Pos{x, y}] = "."
			} else {
				memSpace[Pos{x, y}] = string(char)
			}
		}
		y++
	}
	height = y
	return memSpace
}
