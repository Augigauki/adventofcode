package main

import (
	"fmt"
	"strconv"
)

type Pos struct {
	x, y int
}

type Robot struct {
	P  Pos
	V  Pos
	id int
}

func Part1(space [][]string, robots []Robot) int {
	/* fmt.Println("Space:")
	for _, line := range space {
		fmt.Println(line)
	} */
	var robotMap = make(map[Pos]int)
	width := len(space[0])
	height := len(space)
	//testRobot := Robot{Pos{2, 4}, Pos{2, -3}}
	//testRobots := []Robot{testRobot}
	//fmt.Println("Test robot: ", testRobot)
	/* fmt.Println("Robots:")
	for i, robot := range robots {
		fmt.Printf("Robot %v at position %v with velocity %v\n", i+1, robot.P, robot.V)
		robotMap[robot.P] = robotMap[robot.P] + 1
	} */
	placeRobots(space, robotMap, robots, 100)

	fmt.Println("Robot map:")
	var ne, nw, se, sw, safetyFactor int
	//fmt.Println("Width / 2: ", width/2, " Height / 2: ", height/2)
	for pos, value := range robotMap {
		if value > 0 {
			//fmt.Println("Position: ", pos, " Robots: ", value)
			if pos.x < width/2 && pos.y < height/2 {
				nw += value
			} else if pos.x > width/2 && pos.y < height/2 {
				ne += value
			} else if pos.x < width/2 && pos.y > height/2 {
				sw += value
			} else if pos.x > width/2 && pos.y > height/2 {
				se += value
			}
		}

	}
	safetyFactor = ne * nw * se * sw
	fmt.Println("Safety factor: ", safetyFactor)
	//placeRobots(space, robotMap, testRobots)

	//placeRobots(space, testRobots)
	return 0
}

func placeRobots(space [][]string, robotMap map[Pos]int, robots []Robot, limit int) {
	for i := 0; i < 10000000; i++ {
		for j := range robots {
			robot := &robots[j]
			//fmt.Println("Robot position: ", robot.P)
			robotMap[robot.P] = robotMap[robot.P] - 1
			if robotMap[robot.P] < 0 {
				robotMap[robot.P] = 0
			}
			space[robot.P.y][robot.P.x] = "."
			*robot = getNextPos(robot, space)
			//fmt.Println("Robot new position: ", robot.P)
			robotMap[robot.P] = robotMap[robot.P] + 1
			space[robot.P.y][robot.P.x] = strconv.Itoa(robotMap[robot.P])
		}
		counter := 0
		fmt.Println("Second: ", i, ". Checking if all robots are alone...")
		for _, v := range robotMap {
			if v == 1 {
				counter++
			}
		}
		fmt.Println("robots.length: ", len(robots), "Counter: ", counter)
		if counter == len(robots) {
			fmt.Println("All robots are alone at second: ", i+1, "!")
			break
		}
	}
	/* fmt.Println("Space after robots:")
	for _, line := range space {
		fmt.Println(line)
	} */
}

func getNextPos(bot *Robot, space [][]string) Robot {
	width := len(space[0])
	height := len(space)
	newX := bot.P.x + bot.V.x
	newY := bot.P.y + bot.V.y
	//fmt.Println("New x: ", newX, " New y: ", newY)
	if newX < 0 {
		newX = width + (bot.P.x + bot.V.x)
	}
	if newY < 0 {
		//fmt.Printf("Y under zero! %v. Getting new y: %v (height: %v, bot.P.y: %v, bot.V.y: %v)\n", newY, height-(bot.P.y+bot.V.y), height, bot.P.y, bot.V.y)
		newY = height + (bot.P.y + bot.V.y)
	}
	if newX >= width {
		newX = (bot.P.x + bot.V.x) - width
	}
	if newY >= height {
		newY = (bot.P.y + bot.V.y) - height
	}
	bot.P = Pos{newX, newY}
	return *bot
}
