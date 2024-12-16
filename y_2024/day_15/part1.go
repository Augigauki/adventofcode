package main

import "fmt"

/* func check2(room []string, y, x1, x2, dir int, objects *[][2]int) bool {
	*objects = append(*objects, [2]int{x1, y}) // Add the current box part to the list
	fmt.Println("Objects:", objects)

	if room[y+dir][x1] == '#' || room[y+dir][x2] == '#' {
		return false // Movement is blocked
	}
	if room[y+dir][x1] == '.' && room[y+dir][x2] == '.' {
		return true // Movement is possible
	}
	// Check for recursive box connections
	var results []bool
	if room[y+dir][x1] == '[' {
		results = append(results, check(room, y+dir, x1, x2, dir, objects))
	}
	if room[y+dir][x1] == ']' {
		results = append(results, check(room, y+dir, x1-1, x2-1, dir, objects))
	}
	if room[y+dir][x2] == '[' {
		results = append(results, check(room, y+dir, x1+1, x2+1, dir, objects))
	}

	// If any part is blocked, the entire move is blocked
	for _, res := range results {
		if !res {
			return false
		}
	}
	return true
} */

/* func move2(room []string, dir int, objects [][2]int) []string {
	fmt.Println("Objects:", objects)

	// Remove old box positions
	for _, obj := range objects {
		x, y := obj[0], obj[1]
		room[y] = room[y][:x] + ".." + room[y][x+2:]
	}

	// Add boxes in the new positions
	for _, obj := range objects {
		x, y := obj[0], obj[1]
		room[y+dir] = room[y+dir][:x] + "[]" + room[y+dir][x+2:]
	}

	return room
} */

/* func run2(room []string, sx, sy int, inst byte) ([]string, int, int) {
	var objects [][2]int // To track parts of wide boxes
	var x, y int         // New robot position after move

	switch inst {
	case '^': // Move up
		switch room[sy-1][sx] {
		case '.': // Empty space
			x, y = sx, sy-1
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy-1] = room[sy-1][:sx] + "@" + room[sy-1][sx+1:]
		case '[': // Wide box
			if check(room, sy-1, sx, sx+1, -1, &objects) {
				room = move(room, -1, objects)
				room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
				y--
				room[sy] = room[sy][:sx] + "@" + room[sy][sx+1:]
			}
		case ']': // Wide box (right part)
			if check(room, sy-1, sx-1, sx, -1, &objects) {
				room = move(room, -1, objects)
				room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
				y--
				room[sy] = room[sy][:sx] + "@" + room[sy][sx+1:]
			}
		}
	case 'v': // Move down
		switch room[sy+1][sx] {
		case '.':
			x, y = sx, sy+1
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy+1] = room[sy+1][:sx] + "@" + room[sy+1][sx+1:]
		case '[':
			if check(room, sy+1, sx, sx+1, +1, &objects) {
				room = move(room, +1, objects)
				room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
				y++
				room[sy] = room[sy][:sx] + "@" + room[sy][sx+1:]
			}
		case ']':
			if check(room, sy+1, sx-1, sx, +1, &objects) {
				room = move(room, +1, objects)
				room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
				y++
				room[sy] = room[sy][:sx] + "@" + room[sy][sx+1:]
			}
		}
	case '<': // Move left
		switch room[sy][sx-1] {
		case '.':
			x, y = sx-1, sy
			room[sy] = room[sy][:sx-1] + "@" + room[sy][sx:]
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
		case ']':
			for i := sx - 1; i > 0; i-- {
				if room[sy][i] == '#' {
					break
				}
				if room[sy][i] == '.' {
					x, y = sx-1, sy
					room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
					room[sy] = room[sy][:sx-1] + "@" + room[sy][sx:]
					for j := i; j < sx-1; j += 2 {
						room[sy] = room[sy][:j] + "[" + room[sy][j+1:]
						room[sy] = room[sy][:j+1] + "]" + room[sy][j+2:]
					}
					break
				}
			}
		}
	case '>': // Move right
		switch room[sy][sx+1] {
		case '.':
			x, y = sx+1, sy
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy] = room[sy][:sx+1] + "@" + room[sy][sx+2:]
		case '[':
			for i := sx + 1; i < len(room[sy]); i++ {
				if room[sy][i] == '#' {
					break
				}
				if room[sy][i] == '.' {
					x, y = sx+1, sy
					room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
					room[sy] = room[sy][:sx+1] + "@" + room[sy][sx+2:]
					for j := i; j > sx+1; j -= 2 {
						room[sy] = room[sy][:j] + "]" + room[sy][j+1:]
						room[sy] = room[sy][:j-1] + "[" + room[sy][j:]
					}
					break
				}
			}
		}
	}
	return room, x, y
} */

func CheckIfWideBoxCanBeMoved(posMap map[Pos]string, dir Direction, bot *Bot) map[Pos]string {
	fmt.Println("Checking if wide box can be moved...")

	// Starting position
	origPos := bot.Pos
	nextPos := Pos{bot.Pos.Char + dir.Dir[1], bot.Pos.Line + dir.Dir[0]} // Position where the robot is pushing
	boxesMap := make(map[Pos]string)

	// ** Handle North-South (Vertical) Movement **
	if dir.Name == "N" || dir.Name == "S" {
		line := nextPos.Line
		for {
			// Register all box parts in the current line
			currentLine := make(map[Pos]string)
			for left := nextPos.Char; posMap[Pos{left, line}] == "[" || posMap[Pos{left, line}] == "]"; left-- {
				currentLine[Pos{left, line}] = posMap[Pos{left, line}]
			}
			for right := nextPos.Char + 1; posMap[Pos{right, line}] == "[" || posMap[Pos{right, line}] == "]"; right++ {
				currentLine[Pos{right, line}] = posMap[Pos{right, line}]
			}

			// If the line is blocked or empty, stop
			isBlocked := false
			for pos := range currentLine {
				belowPos := Pos{pos.Char, line + dir.Dir[0]}
				if posMap[belowPos] == "#" || posMap[belowPos] != "." {
					isBlocked = true
				}
			}
			if isBlocked || len(currentLine) == 0 {
				break
			}

			// Add current line to the boxesMap
			for pos, char := range currentLine {
				boxesMap[pos] = char
			}

			// Move to the next line
			line += dir.Dir[0]
		}
	}

	// ** Handle East-West (Horizontal) Movement **
	if dir.Name == "W" || dir.Name == "E" {
		for posMap[nextPos] == "[" || posMap[nextPos] == "]" {
			boxesMap[nextPos] = posMap[nextPos]
			nextPos = Pos{nextPos.Char + dir.Dir[1], nextPos.Line + dir.Dir[0]} // Traverse horizontally
		}
	}

	fmt.Println("Boxes map: ", boxesMap)

	// Check if all new positions for the box parts are clear
	canMove := true
	newBoxPositions := make(map[Pos]string)
	for boxPos, value := range boxesMap {
		newPos := Pos{boxPos.Char + dir.Dir[1], boxPos.Line + dir.Dir[0]}
		if posMap[newPos] != "." {
			canMove = false
			fmt.Printf("Obstacle detected at %v\n", newPos)
			break
		}
		newBoxPositions[newPos] = value
	}

	if canMove {
		fmt.Println("Wide boxes can be moved!")
		// Clear old positions
		for boxPos := range boxesMap {
			posMap[boxPos] = "."
		}
		// Update new positions
		for pos, char := range newBoxPositions {
			posMap[pos] = char
		}
		// Move the robot
		posMap[origPos] = "."
		bot.Pos = Pos{bot.Pos.Char + dir.Dir[1], bot.Pos.Line + dir.Dir[0]}
		posMap[bot.Pos] = "@"
	} else {
		fmt.Println("Wide box cannot be moved. Obstacle detected.")
	}

	return posMap
}

func CheckIfWideBoxCanBeMoved2(posMap map[Pos]string, dir Direction, bot *Bot) map[Pos]string {
	fmt.Println("Checking if wide box can be moved...")
	origPos := bot.Pos
	nextPos := Pos{bot.Pos.Char + dir.Dir[1], bot.Pos.Line + dir.Dir[0]}
	origNextPos := nextPos
	boxesMap := make(map[Pos]string)
	if dir.Name == "N" {
		line := nextPos.Char
		fmt.Println("Line: ", line)
		boxesMap[nextPos] = posMap[nextPos]
		leftPos := Pos{nextPos.Char - 1, nextPos.Line}
		rightPos := Pos{nextPos.Char + 1, nextPos.Line}
		for posMap[leftPos] == "[" || posMap[leftPos] == "]" {
			leftPos = Pos{leftPos.Char - 1, leftPos.Line}
			boxesMap[leftPos] = posMap[leftPos]
		}
		for posMap[rightPos] == "[" || posMap[rightPos] == "]" {
			rightPos = Pos{rightPos.Char + 1, rightPos.Line}
			boxesMap[rightPos] = posMap[rightPos]
		}
		line--
		/* for !checkNextLine(posMap, boxesMap, line) {
			boxesMap = findNextLineWideBoxes(posMap, boxesMap)
			line--
		} */
		fmt.Println("Boxes map: ", boxesMap)

	}
	if dir.Name == "W" || dir.Name == "E" {
		for posMap[nextPos] == "[" || posMap[nextPos] == "]" {
			nextPos = Pos{nextPos.Char + dir.Dir[1], nextPos.Line + dir.Dir[0]}
			boxesMap[nextPos] = posMap[nextPos]
		}
	}
	for pos, value := range boxesMap {
		fmt.Println("Pos: ", pos, " value: ", value)
	}
	if posMap[nextPos] == "." {
		fmt.Println("Wide boxes can be moved!")
		i := len(boxesMap)
		for boxPos := range boxesMap {
			if i%2 == 0 {
				posMap[boxPos] = "]"
			} else {
				posMap[boxPos] = "["
			}
			i--
		}
		bot.Pos = origNextPos
		posMap[origPos] = "."
		posMap[bot.Pos] = "@"
	}
	return posMap
}

/* func findNextLineWideBoxes(posMap map[Pos]string, boxesMap map[Pos]string) map[Pos]string {
	fmt.Println("Finding next line wide boxes...")
	//nextPos := Pos{botPos.Char + dir.Dir[1], botPos.Line + dir.Dir[0]}

	//boxesMap[nextPos] = posMap[nextPos]
	for pos := range boxesMap {
		leftPos := Pos{pos.Char - 1, pos.Line}
		rightPos := Pos{pos.Char + 1, pos.Line}
		for posMap[leftPos] == "[" || posMap[leftPos] == "]" {
			boxesMap[leftPos] = posMap[leftPos]
			leftPos = Pos{leftPos.Char - 1, leftPos.Line}
		}
		for posMap[rightPos] == "[" || posMap[rightPos] == "]" {
			boxesMap[rightPos] = posMap[rightPos]
			rightPos = Pos{rightPos.Char + 1, rightPos.Line}
		}
	}

	return boxesMap
} */
