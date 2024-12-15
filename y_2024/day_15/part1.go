package main

import "fmt"

func check2(room []string, y, x1, x2, dir int, objects *[][2]int) bool {
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
}

func move2(room []string, dir int, objects [][2]int) []string {
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
}

func run2(room []string, sx, sy int, inst byte) ([]string, int, int) {
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
}

func countSumBoxesCoordsX2(room []string) int {
	sum := 0
	for y, line := range room {
		for x := 0; x < len(line); x++ {
			if line[x] == '[' {
				sum += x + y*100
			}
		}
	}
	return sum
}
