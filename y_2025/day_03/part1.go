package main

import (
	"fmt"
	"strconv"
)

func findLargestJoltage(joltages []int) int {
	//fmt.Println("Joltages:", joltages)
	highest := 0
	secondHighest := 0
	firstIndex := 0

	for i := 0; i < (len(joltages) - 1); i++ {
		if joltages[i] > highest {
			highest = joltages[i]
			firstIndex = i
		}
	}
	for i := firstIndex + 1; i < len(joltages); i++ {
		if joltages[i] > secondHighest {
			secondHighest = joltages[i]
		}
	}
	combined := strconv.Itoa(highest) + strconv.Itoa(secondHighest)
	combinedInt, _ := strconv.Atoi(combined)
	fmt.Println("Highest:", highest, "Second Highest:", secondHighest, "Combined:", combinedInt)
	return combinedInt
}
