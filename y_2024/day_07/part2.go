package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateExtraEquations(equations []string) []string {
	//var extraEquations [][]string
	var newStrings []string
	for _, equation := range equations {
		stringNumbers := strings.Split(equation, ": ")
		//fmt.Printf("String numbers[0]: %v\n", stringNumbers[0])
		var numbers []int
		for _, stringNumber := range strings.Split(stringNumbers[1], " ") {
			number, err := strconv.Atoi(stringNumber)
			if err != nil {
				fmt.Println("Error converting string to int")
				return nil
			}
			numbers = append(numbers, number)
		}
		fmt.Println("Numbers: ", numbers)
		var result [][]int
		generateCombinations(numbers, []int{}, 0, &result)
		for _, res := range result {
			//fmt.Printf("Res: %v\n", res)
			numbersAsString := []string{}
			for _, num := range res {
				numbersAsString = append(numbersAsString, strconv.Itoa(num))
			}
			equationNumbers := strings.Join(numbersAsString, " ")
			//fmt.Println("Equation numbers: ", equationNumbers)
			equation := stringNumbers[0] + ": " + equationNumbers
			exists := false
			for _, newString := range newStrings {
				if newString == equation {
					exists = true
				}
			}
			if !exists {

				newStrings = append(newStrings, equation)
			}

		}
		//fmt.Println(newStrings)

		fmt.Println("Result: ", result)
		//extraEquations = append(extraEquations, result...)
		/* for _, res := range result {
		} */
		/* newEquations := combineNumbers(strings.Split(numbers, " "))
		fmt.Println("New equations: ", newEquations) */
	}
	return newStrings
}

// Function to generate all possible combinations
func generateCombinations(slice []int, current []int, index int, result *[][]int) {
	// If we've processed the entire slice, add the current combination to the result
	if index == len(slice) {
		*result = append(*result, append([]int(nil), current...))
		return
	}

	// Option 1: Add the current number as-is
	generateCombinations(slice, append(current, slice[index]), index+1, result)

	// Option 2: Combine the current number with the next (if possible)
	if index+1 < len(slice) {
		combined, _ := strconv.Atoi(fmt.Sprintf("%d%d", slice[index], slice[index+1]))
		generateCombinations(slice, append(current, combined), index+2, result)
	}

	// Option 3: Combine all remaining numbers into one (only at the top level)
	if index == 0 {
		combinedAll := 0
		for i := index; i < len(slice); i++ {
			combinedAll, _ = strconv.Atoi(fmt.Sprintf("%d%d", combinedAll, slice[i]))
		}
		// Add the fully combined result only if the slice has more than one number
		*result = append(*result, []int{combinedAll})
	}
}
