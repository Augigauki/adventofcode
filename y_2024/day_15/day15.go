package main

import (
	"bytes"
	"fmt"
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

// Expand the map for wide boxes
func expandMap(room []string) []string {
	var expanded []string
	for _, line := range room {
		var buffer bytes.Buffer
		for _, char := range line {
			switch char {
			case '.':
				buffer.WriteString("..")
			case '#':
				buffer.WriteString("##")
			case '@':
				buffer.WriteString("@.")
			case 'O':
				buffer.WriteString("[]")
			}
		}
		expanded = append(expanded, buffer.String())
	}
	return expanded
}

// Recursive function to validate if a wide box can be moved
func check(room []string, y, x1, x2, dir int, objects *[][2]int) bool {
	*objects = append(*objects, [2]int{x1, y}) // Add the current box part to the list
	if room[y+dir][x1] == '#' || room[y+dir][x2] == '#' {
		return false // Blocked by a wall
	}
	if room[y+dir][x1] == '.' && room[y+dir][x2] == '.' {
		return true // Movement is possible
	}

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

	for _, res := range results {
		if !res {
			return false
		}
	}
	return true
}

// Move boxes based on the validated positions
func move(room []string, dir int, objects [][2]int) []string {
	// Remove boxes from old positions
	for _, obj := range objects {
		x, y := obj[0], obj[1]
		room[y] = room[y][:x] + ".." + room[y][x+2:]
	}

	// Add boxes to the new positions
	for _, obj := range objects {
		x, y := obj[0], obj[1]
		room[y+dir] = room[y+dir][:x] + "[]" + room[y+dir][x+2:]
	}

	return room
}

// Main logic to move the robot and boxes
func run(room []string, sx, sy int, inst byte) ([]string, int, int) {
	var objects [][2]int
	switch inst {
	case '^': // Move up
		switch room[sy-1][sx] {
		case '.':
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy-1] = room[sy-1][:sx] + "@" + room[sy-1][sx+1:]
			sy--
		case '[':
			if check(room, sy-1, sx, sx+1, -1, &objects) {
				room = move(room, -1, objects)
				room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
				sy--
				room[sy] = room[sy][:sx] + "@" + room[sy][sx+1:]
			}
		case ']':
			if check(room, sy-1, sx-1, sx, -1, &objects) {
				room = move(room, -1, objects)
				room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
				sy--
				room[sy] = room[sy][:sx] + "@" + room[sy][sx+1:]
			}
		}
	case 'v': // Move down
		switch room[sy+1][sx] {
		case '.':
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy+1] = room[sy+1][:sx] + "@" + room[sy+1][sx+1:]
			sy++
		case '[':
			if check(room, sy+1, sx, sx+1, +1, &objects) {
				room = move(room, +1, objects)
				room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
				sy++
				room[sy] = room[sy][:sx] + "@" + room[sy][sx+1:]
			}
		case ']':
			if check(room, sy+1, sx-1, sx, +1, &objects) {
				room = move(room, +1, objects)
				room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
				sy++
				room[sy] = room[sy][:sx] + "@" + room[sy][sx+1:]
			}
		}
	case '<': // Move left
		switch room[sy][sx-1] {
		case '.':
			room[sy] = room[sy][:sx-1] + "@" + room[sy][sx:]
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			sx--
		case ']':
			for i := sx - 1; i > 0; i-- {
				if room[sy][i] == '#' {
					break
				}
				if room[sy][i] == '.' {
					room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
					room[sy] = room[sy][:sx-1] + "@" + room[sy][sx:]
					for j := i; j < sx-1; j += 2 {
						room[sy] = room[sy][:j] + "[" + room[sy][j+1:]
						room[sy] = room[sy][:j+1] + "]" + room[sy][j+2:]
					}
					sx--
					break
				}
			}
		}
	case '>': // Move right
		switch room[sy][sx+1] {
		case '.':
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy] = room[sy][:sx+1] + "@" + room[sy][sx+2:]
			sx++
		case '[':
			for i := sx + 1; i < len(room[sy]); i++ {
				if room[sy][i] == '#' {
					break
				}
				if room[sy][i] == '.' {
					room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
					room[sy] = room[sy][:sx+1] + "@" + room[sy][sx+2:]
					for j := i; j > sx+1; j -= 2 {
						room[sy] = room[sy][:j] + "]" + room[sy][j+1:]
						room[sy] = room[sy][:j-1] + "[" + room[sy][j:]
					}
					sx++
					break
				}
			}
		}
	}
	return room, sx, sy
}

// Calculate GPS sum for boxes
func countSumBoxesCoords(room []string) int {
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

// Main function
func main() {
	// Read input file
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := bytes.Split(file, []byte("\n"))
	var room []string
	var instructions string

	// Parse the input
	for _, line := range lines {
		str := string(line)
		if len(str) == 0 {
			continue
		}
		if str[0] == '#' {
			room = append(room, str)
		} else {
			instructions += str
		}
	}

	// Expand the map for wide boxes
	room = expandMap(room)

	// Find initial robot position
	var sx, sy int
	for y, line := range room {
		for x, char := range line {
			if char == '@' {
				sx, sy = x, y
				break
			}
		}
	}

	// Process the instructions
	for _, inst := range instructions {
		room, sx, sy = run(room, sx, sy, byte(inst))
	}

	// Calculate and print the GPS sum
	fmt.Println("Sum of GPS coordinates:", countSumBoxesCoords(room))
}
