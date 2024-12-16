package main

import "fmt"

func part1(maze []string) {
	fmt.Println("Part 1")
	//fmt.Println("Maze:")
	racer := Reindeer{Pos{0, 0}, Pos{0, 0}, Directions[1], []Path{}}
	//fmt.Println("Maze: ", maze)
	for y, row := range maze {
		fmt.Println(row)
		//fmt.Println("y: ", y)
		for x, char := range row {
			//fmt.Println("x: ", x)
			//fmt.Println("Char:", char)
			if char == 'S' {
				//fmt.Println("Found start position")
				racer.startPos = Pos{x, y}
				racer.pos = Pos{x, y}
			}
		}
	}
	fmt.Println("Start position:", racer.startPos)
}

func move(racer *Reindeer) {

}
