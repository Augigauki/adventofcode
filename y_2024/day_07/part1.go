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

func evaluateCombinations(numbers []int, currentIndex int, currentResult int, currentExpression string, testValue int) bool {
	result := false
	// If we've reached the end of the slice, print the result and expression
	if currentIndex == len(numbers) {
		fmt.Printf("Result: %d, Expression: %s\n", currentResult, currentExpression)
		if currentResult == testValue {
			fmt.Println("Found a valid expression!")
			result = true
			return result
		}
		return result

	}
	if result {
		return result
	}
	// Add the current number
	testPlus := evaluateCombinations(numbers, currentIndex+1, currentResult+numbers[currentIndex],
		fmt.Sprintf("%s + %d", currentExpression, numbers[currentIndex]), testValue)

	if testPlus {
		result = true
		return result
	}
	// Multiply the current number
	testMultiply := evaluateCombinations(numbers, currentIndex+1, currentResult*numbers[currentIndex],
		fmt.Sprintf("%s * %d", currentExpression, numbers[currentIndex]), testValue)

	if testMultiply {
		result = true
		return result
	}
	return result
}
