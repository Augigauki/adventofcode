package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CheckWholeUpdate(rules []string, updates []string) {

	var correctUpdates []string
	for _, update := range updates {
		fmt.Println("Update: ", update)
		stringPages := strings.Split(update, ",")
		pages := ConvPagesToInts(stringPages)
		metRules := 0
		totalMetRules := 0
		requiredMetRules := len(pages) * (len(pages) - 1)
		//fmt.Printf("Required matches: %v\n", len(pages)*(len(pages)-1))
		for i, page := range pages {
			//fmt.Printf("PAGE/i: %v\n", page)
			for j, otherPage := range pages {
				metRules = 0
				//fmt.Printf("OTHER PAGE/j: %v\n", otherPage)
				correct := false
				if j == i {
					//fmt.Printf("Page: %v is same as otherPage: %v, doing nothing\n", page, otherPage)
					continue
				}
				for _, rule := range rules {
					if j < i {
						//fmt.Printf("Page: %v, Other Page: %v\n", page, otherPage)
						//fmt.Printf("Checking rule (opposite) %v with page combination: %v,%v\n", rule, page, otherPage)
						if CheckIfRuleIsMet(rule, otherPage, page) {
							metRules++
							correct = true
							break
						}
					}
					if j > i {
						//fmt.Printf("Checking rule %v with page combination: %v,%v\n", rule, page, otherPage)
						if CheckIfRuleIsMet(rule, page, otherPage) {
							metRules++
							correct = true
							break
						}
					}

				}
				if !correct {
					//fmt.Printf("No rule met for page: %v and otherPage: %v\n", page, otherPage)
					break
				}

				totalMetRules += metRules
			}
			//fmt.Printf("Total Met Rules: %v\n", totalMetRules)
			if totalMetRules == requiredMetRules {
				correctUpdates = append(correctUpdates, update)
			}
		}
	}
	fmt.Println("Correct Updates: ", correctUpdates)
	middleNumbers := []int{}
	for _, update := range correctUpdates {
		updateNumsAsStrings := strings.Split(update, ",")
		updateNums := ConvPagesToInts(updateNumsAsStrings)
		if len(updateNums) > 0 {
			middleIndex := len(updateNums) / 2
			middleNumbers = append(middleNumbers, updateNums[middleIndex])
		}
	}
	middlesSum := 0
	for _, num := range middleNumbers {
		middlesSum += num
	}
	fmt.Printf("Sum of middle numbers: %v\n", middlesSum)

}

func CheckIfRuleIsMet(rule string, left int, right int) bool {
	ruleNums := ConvRuleToInts(rule)
	if ruleNums[0] == left && ruleNums[1] == right {
		//fmt.Printf("Rule is met: %v, left: %v, right: %v\n", rule, left, right)
		return true
	}
	return false
}

func ConvRuleToInts(rule string) []int {
	ruleNumsStrings := strings.Split(rule, "|")
	ruleNums := make([]int, len(ruleNumsStrings))
	for i, str := range ruleNumsStrings {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			continue
		}
		ruleNums[i] = num
	}
	return ruleNums
}

func ConvPagesToInts(pages []string) []int {
	pageNums := make([]int, len(pages))
	for i, page := range pages {
		num, err := strconv.Atoi(page)
		if err != nil {
			fmt.Printf("Error converting page to int: %v\n", err)
			continue
		}
		pageNums[i] = num
	}
	return pageNums
}
