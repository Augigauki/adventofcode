package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	current := 50
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()
	//var zeroes []int
	var zeroClicks int
	scanner := bufio.NewScanner(file)
	//var left, right []int
	for scanner.Scan() {
		var line string = scanner.Text()
		var direction string = line[0:1]
		var distanceString string = line[1:]
		var distance int
		distance, err = strconv.Atoi(distanceString)
		if err != nil {
			fmt.Println("Error converting distance to int:", err)
			return
		}
		//current = turnDial(direction, distance, current)
		var zeroClicksForTurn int
		current, zeroClicksForTurn = countZeroes(direction, distance, current)
		zeroClicks += zeroClicksForTurn
		fmt.Printf("%s %d %d\n", direction, distance, current)
		/* if current == 0 {
			zeroes = append(zeroes, current)
		} */

	}
	//fmt.Println("Times at zero: ", len(zeroes))
	fmt.Println("Zero clicks: ", zeroClicks)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
