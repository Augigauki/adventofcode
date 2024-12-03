package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func MulNums(memory string) int {
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)

	sum := 0
	matches := re.FindAllString(memory, -1)
	fmt.Println("Matches: ", matches)
	var multiples []int
	for i := range matches {
		fmt.Printf("Match %d: %s\n", i+1, matches[i])
		numPattern := `\d{1,3}`
		numRe := regexp.MustCompile(numPattern)
		nums := numRe.FindAllString(matches[i], -1)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		result := num1 * num2
		multiples = append(multiples, result)
	}
	fmt.Println("Multiples: ", multiples)
	for _, num := range multiples {
		sum += num
	}
	return sum
}
