package main

import (
	"fmt"
	"strings"
)

func GenerateExtraEquations(equations []string) []string {
	var extraEquations []string
	for _, equation := range equations {
		numbers := strings.Split(equation, ": ")[1]
		fmt.Println("Numbers: ", numbers)
		newEquations := combineNumbers(strings.Split(numbers, " "))
		fmt.Println("New equations: ", newEquations)
	}
	return extraEquations
}

func combineNumbers(numbers []string) [][]string {
	var extraEquations [][]string
	for i, num := range numbers {
		tempSlice := numbers
		if i < len(numbers)-1 {
			fmt.Println("Num and next ", num, numbers[i+1])
			newNum := num + numbers[i+1]
			tempSlice[i] = newNum
			tempSlice = append(tempSlice[:i+1], tempSlice[i+2:]...)
			extraEquations = append(extraEquations, tempSlice)

		}
	}
	return extraEquations
}
