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
	var strNumbers [][]string
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
			strNums := []string{}
			for _, numStr := range numStrs {
				//fmt.Println("NumStr:", numStr)
				if numStr == "" {
					continue
				}
				strNums = append(strNums, numStr)
			}
			strNumbers = append(strNumbers, strNums)
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
	columns := [][]string{}
	for i := 0; i < len(strNumbers[0]); i++ {
		column := []string{}
		for j := 0; j < len(strNumbers); j++ {
			column = append(column, strNumbers[j][i])
		}
		columns = append(columns, column)
	}
	// No reversals needed - order doesn't affect final sum
	//fmt.Println("Columns:", columns)
	for i, row := range strNumbers {
		fmt.Printf("Row %d has %d columns\n", i, len(row))
	}
	results := rightToLeft(columns, operators)
	//fmt.Println("Results:", results)
	finalScore := 0
	for _, res := range results {
		finalScore += res
	}
	fmt.Println("Final Score:", finalScore)
	fmt.Println("Wrong Score:", "8378086516563")
}
