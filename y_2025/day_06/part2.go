package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rightToLeft(problems [][]string, operators []string) []int {
	fmt.Println("Problems:", problems)
	results := []int{}
	for i := 0; i < len(operators); i++ {
		problem := problems[i]
		paddedProblem := []string{}
		result := 0
		operator := operators[i]
		fmt.Println("Operator:", operator)
		fmt.Println("Problem:", problem)
		numLength := 0
		//find the longest number in this column
		for _, num := range problem {
			if len(num) > numLength {
				numLength = len(num)
			}
		}
		//pad all numbers to the same length
		for _, num := range problem {

			//pad with zeroes based on operator
			for len(num) < numLength {
				if operator == "+" {
					num = num + "0"
				} else if operator == "*" {
					num = "0" + num
				}
			}
			//fmt.Println("Padded Number:", num)
			paddedProblem = append(paddedProblem, num)
		}
		fmt.Println("Padded Problem:", paddedProblem)
		//concatenate from right to left
		trimmedProblem := []int{}
		for j := numLength - 1; j >= 0; j-- {
			result := ""
			for k := 0; k < len(paddedProblem); k++ {
				result += string(paddedProblem[k][j])
			}
			trimmed := strings.Split(result, "0")
			for _, str := range trimmed {
				if str == "0" || str == "" {
					continue
				} else {
					num, _ := strconv.Atoi(str)
					trimmedProblem = append(trimmedProblem, num)
				}
			}
			fmt.Println("Concatenated Result at position", j, ":", trimmed)

		}
		switch operator {
		case "+":
			for _, val := range trimmedProblem {
				result += val
			}
		case "*":
			result = 1
			for _, val := range trimmedProblem {
				result *= val
			}
		default:
			fmt.Println("Unknown operator:", operator)
		}
		fmt.Println("Final Result for operator", operator, ":", result)
		results = append(results, result)
		fmt.Println("-----")

	}

	return results
}
