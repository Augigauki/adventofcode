package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func MulNumsDoDont(memory string) int {
	pattern := `mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`
	re := regexp.MustCompile(pattern)

	sum := 0
	matches := re.FindAllString(memory, -1)
	fmt.Println("Matches: ", matches)
	var multiples []int
	operate := true
	for i, match := range matches {

		fmt.Printf("Match %d: %s\n", i+1, matches[i])
		if match == "do()" {
			operate = true
		} else if match == "don't()" {
			operate = false
		} else {
			fmt.Println("Operate: ", operate)
			if operate {
				numPattern := `\d{1,3}`
				numRe := regexp.MustCompile(numPattern)
				nums := numRe.FindAllString(matches[i], -1)
				num1, _ := strconv.Atoi(nums[0])
				num2, _ := strconv.Atoi(nums[1])
				result := num1 * num2
				multiples = append(multiples, result)
			}
		}

	}
	fmt.Println("Multiples: ", multiples)
	for _, num := range multiples {
		sum += num
	}
	return sum
}
