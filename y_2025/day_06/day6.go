package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers [][]int
	var operators []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "+") || strings.Contains(line, "*") {
			operator := strings.Split(strings.TrimSpace(line), " ")
			for _, op := range operator {
				if op == "" {
					continue
				}
				operators = append(operators, op)
			}
		} else {
			numStrs := strings.Split(strings.TrimSpace(line), " ")
			var numRow []int
			for _, numStr := range numStrs {
				if numStr == "" {
					continue
				}
				num, _ := strconv.Atoi(numStr)
				numRow = append(numRow, num)
			}
			numbers = append(numbers, numRow)
		}
	}
	fmt.Println("Numbers:", numbers)
	fmt.Println("Operators:", operators)
	results := solveProblems(numbers, operators)
	fmt.Println("Results:", results)
	finalScore := 0
	for _, res := range results {
		finalScore += res
	}
	fmt.Println("Final Score:", finalScore)
}
