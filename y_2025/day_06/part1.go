package main

import "fmt"

func solveProblems(problems [][]int, operators []string) []int {
	fmt.Println("Operators length: ", len(operators))
	fmt.Println("Problems length: ", len(problems))
	results := []int{}
	for i := 0; i < len(operators); i++ {
		result := 0
		operator := operators[i]
		problem := []int{}
		for j := 0; j < len(problems); j++ {
			problem = append(problem, problems[j][i])
		}
		fmt.Println("Solving problem with operator", operator, "and values", problem)
		switch operator {
		case "+":
			for _, val := range problem {
				result += val
			}
		case "*":
			result = 1
			for _, val := range problem {
				result *= val
			}
		default:
			fmt.Println("Unknown operator:", operator)
		}
		results = append(results, result)
	}

	return results
}
