package main

import (
	"fmt"
	"strings"
)

func FindUpdatesInCorrectOrder(rules []string, updates []string) {
	fmt.Println("Rules: ", rules)
	fmt.Println("Updates: ", updates)
	for _, update := range updates {
		fmt.Println("Update: ", update)
		pages := strings.Split(update, ",")

		for i, page := range pages {
			fmt.Println("Page: ", page)

			if i < len(pages)-1 {
				after := checkIfRuleIsMet(rules[i], page, pages[i+1])
				fmt.Println(after)
			}

		}
	}
}

func checkIfRuleIsMet(rule string, firstVal string, secondVal string) bool {
	fmt.Println("Rule: ", rule)
	ruleNums := strings.Split(rule, "|")
	if ruleNums[0] == firstVal && ruleNums[1] == secondVal {
		fmt.Println("Rule is met")
		return true
	}
	if ruleNums[0] == secondVal && ruleNums[1] == firstVal {
		fmt.Println("Rule is met")
		return true
	}
	return false
}
