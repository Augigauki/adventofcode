package main

import "fmt"

func PlaceAntinodesHarmonically(antennaMap [][]string, poses []OtherPosSlice) {
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
					fmt.Printf("___Current Pos: %v, Next Pos: %v___\n", pos, nextPos)
					var prevAntinodes []Position
					var nextAntinodes []Position
					/* Checking if antinodes go left or right horizontally */
					fmt.Printf("nextPos.line(%v) - pos.line(%v): %v\n", nextPos.line, pos.line, nextPos.line-pos.line)
					if (nextPos.char - pos.char) > 0 {
						fmt.Printf("nextPos.char (%v) - pos.char (%v): %v\n", nextPos.char, pos.char, nextPos.char-pos.char)
						/* Direction is left */
						horizontalStep := nextPos.char - pos.char
						verticalStep := nextPos.line - pos.line
						fmt.Printf("Horizontal step: %v, Vertical step: %v\n", horizontalStep, verticalStep)
						antinodePos := pos
						for antinodePos.line >= 0 && antinodePos.char >= 0 {
							prevAntinodes = append(prevAntinodes, antinodePos)
							antinodePos.line = antinodePos.line - verticalStep
							antinodePos.char = antinodePos.char - horizontalStep
						}
						antinodePos = pos
						for antinodePos.line < height && antinodePos.char < width {
							nextAntinodes = append(nextAntinodes, antinodePos)
							antinodePos.line = antinodePos.line + verticalStep
							antinodePos.char = antinodePos.char + horizontalStep
						}

					} else {
						/* Direction is right */
						fmt.Println("Direction is right")
						fmt.Printf("nextPos.char (%v) - pos.char (%v): %v\n", nextPos.char, pos.char, nextPos.char-pos.char)
						horizontalStep := pos.char - nextPos.char
						verticalStep := nextPos.line - pos.line
						fmt.Printf("Horizontal step: %v, Vertical step: %v\n", horizontalStep, verticalStep)
						antinodePos := pos
						for antinodePos.line >= 0 && antinodePos.char >= 0 {
							prevAntinodes = append(prevAntinodes, antinodePos)
							antinodePos.line = antinodePos.line - verticalStep
							antinodePos.char = antinodePos.char + horizontalStep
						}
						antinodePos = pos
						for antinodePos.line < height && antinodePos.char < width {
							nextAntinodes = append(nextAntinodes, antinodePos)
							antinodePos.line = antinodePos.line + verticalStep
							antinodePos.char = antinodePos.char - horizontalStep
						}
					}
					j++
					fmt.Println("Prev antinodes: ", prevAntinodes)
					fmt.Println("Next antinodes: ", nextAntinodes)
					for _, antinode := range prevAntinodes {
						if antinode.char >= 0 && antinode.line >= 0 && antinode.char < width && antinode.line < height {
							antinodeMap[antinode.line][antinode.char] = "#"
						}
					}
					for _, antinode := range nextAntinodes {
						if antinode.char >= 0 && antinode.line >= 0 && antinode.char < width && antinode.line < height {
							antinodeMap[antinode.line][antinode.char] = "#"
						}
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
