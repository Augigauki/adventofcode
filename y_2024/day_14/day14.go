package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file: ", fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	robots := []Robot{}
	for i, line := range input {
		bot := parseRobot(line)
		bot.id = i
		robots = append(robots, bot)
	}
	exampleMap := make([][]string, 7)
	for i := range exampleMap {
		exampleMap[i] = make([]string, 11)
		for j := range exampleMap[i] {
			exampleMap[i][j] = "."
		}
	}

	realMap := make([][]string, 103)
	for i := range realMap {
		realMap[i] = make([]string, 101)
		for j := range realMap[i] {
			realMap[i][j] = "."
		}
	}
	Part1(realMap, robots)
}

func parseRobot(line string) Robot {
	var robot Robot
	args := strings.Split(line, " ")
	//fmt.Println(args)
	for _, arg := range args {
		if strings.Contains(arg, "p") {
			vars := strings.Split(strings.Split(arg, "=")[1], ",")
			//fmt.Println("P vars: ", vars)
			intVars := []int{}
			for _, v := range vars {
				num, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal("Error converting string to int")
				}
				intVars = append(intVars, num)
			}
			robot.P = Pos{intVars[0], intVars[1]}
		}
		if strings.Contains(arg, "v") {
			vars := strings.Split(strings.Split(arg, "=")[1], ",")
			//fmt.Println("V vars: ", vars)
			intVars := []int{}
			for _, v := range vars {
				num, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal("Error converting string to int")
				}
				intVars = append(intVars, num)
			}
			robot.V = Pos{intVars[0], intVars[1]}
		}
	}

	return robot
}
