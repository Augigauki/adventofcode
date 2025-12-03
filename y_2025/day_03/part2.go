package main

import (
	"fmt"
	"strconv"
)

func find12DigitJoltage(joltages []int) int {
	baseJoltage := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	index := 0
	battery := 0
	for i := 0; i < len(baseJoltage); i++ {
		//fmt.Println("Index:", index)
		battery, index = findLargestBattery(joltages, index, 12-i)
		baseJoltage[i] = battery
	}

	fmt.Println("12 digit Joltage:", baseJoltage)
	stringJoltage := ""
	for _, battery := range baseJoltage {
		stringBattery := strconv.Itoa(battery)
		stringJoltage += stringBattery
	}
	joltageInt, _ := strconv.Atoi(stringJoltage)
	return joltageInt
}

func findLargestBattery(batteries []int, index int, limit int) (int, int) {
	highest := 0
	for i := index; i <= (len(batteries) - limit); i++ {
		if batteries[i] > highest {
			highest = batteries[i]
			index = i
		}
	}
	//fmt.Println("Highest battery found:", highest, "at index:", index)
	return highest, index + 1
}
