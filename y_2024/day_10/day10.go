package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var topoMap [][]int
	for scanner.Scan() {
		var line string = scanner.Text()
		stringNums := strings.Split(line, "")
		chars := []int{}
		for _, char := range stringNums {
			num, err := strconv.Atoi(char)
			if err != nil {
				fmt.Println("Error converting to int")
				break
			}
			chars = append(chars, num)
		}
		topoMap = append(topoMap, chars)
	}

	FindTrailHeadScores(topoMap)
}
