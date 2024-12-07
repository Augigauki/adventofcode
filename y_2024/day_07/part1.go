package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FindTrueEquations(equations []string) {
	var trueEquations []int
	for _, equation := range equations {
		foundTrue := false
		fmt.Println("Equation: ", equation)
		splitTestValueAndNumbers := strings.Split(equation, ": ")
		testValue, err := strconv.Atoi(splitTestValueAndNumbers[0])
		if err != nil {
			fmt.Println("Error converting string to int")
			return
		}
		stringNumbers := strings.Split(splitTestValueAndNumbers[1], " ")
		var numbers []int
		for _, stringNumber := range stringNumbers {
			number, err := strconv.Atoi(stringNumber)
			if err != nil {
				fmt.Println("Error converting string to int")
				return
			}
			numbers = append(numbers, number)
		}
		for _, num := range numbers {
			if num == testValue {
				fmt.Println("Test value is in numbers")
				trueEquations = append(trueEquations, testValue)
				foundTrue = true
				break
			}
		}
		if foundTrue {
			continue
		}
		trueEquation := TestNumbers(testValue, numbers)
		if trueEquation {
			trueEquations = append(trueEquations, testValue)
		}

	}
	fmt.Println("True equations: ", trueEquations)
	var sum int
	for _, trueEquation := range trueEquations {
		sum += trueEquation
	}
	fmt.Println("Sum of true equations: ", sum)
}

func TestNumbers(testValue int, numbers []int) bool {
	return evaluateCombinations(numbers, 1, numbers[0], strconv.Itoa(numbers[0]), testValue)
}

// Function to evaluate all possible combinations
func evaluateCombinations(numbers []int, currentIndex int, currentResult int, currentExpression string, testValue int) bool {
	// Base case: If we've reached the end of the slice
	if currentIndex == len(numbers) {
		fmt.Printf("Result: %d, Expression: %s\n", currentResult, currentExpression)
		if currentResult == testValue {
			fmt.Println("Found a valid expression!")
			return true
		}
		return false
	}

	// Try addition
	if evaluateCombinations(numbers, currentIndex+1, currentResult+numbers[currentIndex],
		fmt.Sprintf("%s + %d", currentExpression, numbers[currentIndex]), testValue) {
		return true
	}

	// Try multiplication
	if evaluateCombinations(numbers, currentIndex+1, currentResult*numbers[currentIndex],
		fmt.Sprintf("%s * %d", currentExpression, numbers[currentIndex]), testValue) {
		return true
	}

	// Try concatenation (|| operator)
	// Concatenate currentResult and numbers[currentIndex] as strings
	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentResult, numbers[currentIndex]))
	if evaluateCombinations(numbers, currentIndex+1, concatenated,
		fmt.Sprintf("%s || %d", currentExpression, numbers[currentIndex]), testValue) {
		return true
	}

	return false
}
