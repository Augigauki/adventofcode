package main

import (
	"fmt"
)

type Position struct {
	line int
	char int
}

type OtherPosSlice struct {
	antenna string
	poses   []Position
}

func FindAntinodes(antennaMap [][]string) {
	fmt.Println("::ANTENNA MAP::")
	for _, line := range antennaMap {
		fmt.Println(line)
	}
	var otherPosSlice = OtherPosSlice{}
	allOtherPoses := []OtherPosSlice{}
	for i, line := range antennaMap {
		for j, char := range line {
			if char != "." {
				currChar := char
				currPos := Position{i, j}
				if len(allOtherPoses) == 0 {
					otherPosSlice = FindAllMatchingAntenna(currPos, antennaMap, currChar)
					allOtherPoses = append(allOtherPoses, otherPosSlice)
				}
				posInSlice := false
				for i = 0; i < len(allOtherPoses); i++ {
					if allOtherPoses[i].antenna == currChar {
						posInSlice = true
						break
					}
				}
				if !posInSlice {
					otherPosSlice = FindAllMatchingAntenna(currPos, antennaMap, currChar)
					allOtherPoses = append(allOtherPoses, otherPosSlice)
				}
			}
		}
	}
	fmt.Println("\n::ALL ANTENNA POSITIONS::")
	for _, otherPosSlice := range allOtherPoses {
		fmt.Println(otherPosSlice)
	}
	PlaceAntinodes(antennaMap, allOtherPoses)
}

func FindAllMatchingAntenna(currPos Position, antennaMap [][]string, currChar string) OtherPosSlice {
	matchingPoses := []Position{}
	for i, line := range antennaMap {
		for j, char := range line {
			if char == currChar {
				matchingPoses = append(matchingPoses, Position{i, j})
			}
		}
	}
	return OtherPosSlice{currChar, matchingPoses}
}

func PlaceAntinodes(antennaMap [][]string, poses []OtherPosSlice) {
	antinodeMap := antennaMap
	width := len(antennaMap[0])
	height := len(antennaMap)
	for _, poslist := range poses {
		fmt.Println(poslist.antenna)
		for i, pos := range poslist.poses {
			if i < len(poslist.poses)-1 {
				checking := true
				j := i
				for checking {
					if j == len(poslist.poses)-1 {
						checking = false
						break
					}
					nextPos := poslist.poses[j+1]
					fmt.Printf("Current position: %v, Next position: %v\n", pos, nextPos)
					var firstAntinode Position
					var secondAntinode Position
					if (nextPos.line - pos.line) > 0 {
						fmt.Println("Antinode should be placed upwards")
						/* Checking if horizontal direction is left */
						if (nextPos.char - pos.char) > 0 {
							firstAntinode = Position{pos.line - (nextPos.line - pos.line), pos.char - (nextPos.char - pos.char)}
							secondAntinode = Position{nextPos.line + (nextPos.line - pos.line), nextPos.char + (nextPos.char - pos.char)}
						} else {
							firstAntinode = Position{pos.line - (nextPos.line - pos.line), pos.char + (pos.char - nextPos.char)}
							secondAntinode = Position{nextPos.line + (nextPos.line - pos.line), nextPos.char - (pos.char - nextPos.char)}
						}
					} else {
						fmt.Println("Antinode should be placed downwards")
						/* Checking if antinode should be placed diagonally down left */
						if (nextPos.char - pos.char) > 0 {
							firstAntinode = Position{pos.line + (pos.line - nextPos.line), pos.char - (nextPos.char - pos.char)}
							/* Antinode should be placed diagonally down right */
						} else {
							firstAntinode = Position{pos.line + (pos.line - nextPos.line), pos.char + (pos.char - nextPos.char)}
						}
					}
					j++
					//antiNodepos = Position{nextPos.line - pos.line, nextPos.char - pos.char}
					fmt.Println("First Antinode position: ", firstAntinode)
					fmt.Println("Second Antinode position: ", secondAntinode)
					canPlaceFirst, canPlaceSecond := true, true
					if firstAntinode.char < 0 || firstAntinode.line < 0 || firstAntinode.char >= width || firstAntinode.line >= height {
						fmt.Println("First Antinode is out of bounds!")
						canPlaceFirst = false
					}
					//fmt.Printf("len(antinodeMap[0]): %v, len(antinodeMap): %v\n", len(antinodeMap[0]), len(antinodeMap))
					if secondAntinode.char >= width || secondAntinode.line >= height || secondAntinode.char < 0 || secondAntinode.line < 0 {
						fmt.Println("Second antinode is out of bounds!")
						canPlaceSecond = false
					}
					if canPlaceFirst {
						antinodeMap[firstAntinode.line][firstAntinode.char] = "#"
					}
					if canPlaceSecond {
						antinodeMap[secondAntinode.line][secondAntinode.char] = "#"
					}
				}

			}
		}
	}
	fmt.Println("\n::ANTINODE MAP::")
	var totalAntinodes int = 0
	for _, line := range antinodeMap {
		fmt.Println(line)
		for _, char := range line {
			if char == "#" {
				totalAntinodes++
			}
		}
	}
	fmt.Println("Total antinodes: ", totalAntinodes)
}
