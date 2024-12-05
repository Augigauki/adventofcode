package main

import (
	"fmt"
	"strings"
)

func FixIncorrectUpdates(rules []string, updates []string) {
	incorrectUpdates := extractIncorrectUpdates(rules, updates)
	fixedUpdates := fixUpdates(rules, incorrectUpdates)
	fmt.Println("Fixed Updates: ", fixedUpdates)
	middleNumbers := []int{}
	for _, update := range fixedUpdates {
		if len(update) > 0 {
			middleIndex := len(update) / 2
			middleNumbers = append(middleNumbers, update[middleIndex])
		}
	}
	middleSum := 0
	for _, num := range middleNumbers {
		middleSum += num
	}
	fmt.Println("Sum of middle numbers: ", middleSum)

}

func fixPages(rules []string, pages []int) []int {
	//var fixedPages []int
	fmt.Printf("\nFIXING PAGES AFTER INITIAL RUN!\n\n")
	for i, page := range pages {
		fmt.Printf("\nPages: %v\n", pages)
		fmt.Printf("PAGE/i: %v\n", page)
		for j, otherPage := range pages {
			fmt.Printf("OTHER PAGE/j: %v\n", otherPage)
			correct := false
			swapped := false
			if j == i {
				continue
			}
			for _, rule := range rules {
				/* if j < i {
					if CheckIfRuleIsMet(rule, otherPage, page) {
						break
					}
				} */
				if j > i {
					if CheckIfRuleIsMet(rule, page, otherPage) {
						fmt.Println("Rule met: ", rule)
						correct = true
					} else if CheckIfRuleIsMet(rule, otherPage, page) {
						//metRules++
						pages[i], pages[j] = pages[j], pages[i]
						fmt.Printf("Swapped %v with %v and vice versa. Pages now is: %v\n", pages[i], pages[j], pages)
						swapped = true
						i--
						j--
						//correct = true

					}
				}
			}
			if correct || swapped {
				break
			}
		}
	}
	fixed := true
	for i, page := range pages {
		if !fixed {
			break
		}
		if i < len(pages)-1 {

			for _, rule := range rules {
				if CheckIfRuleIsMet(rule, page, pages[i+1]) {
					fixed = true
					break
				} else if page == pages[i+1] {

					break
				} else {
					fixed = false
				}
			}
		}
	}
	fmt.Println("Fixed: ", fixed)
	if fixed {
		return pages
	} else {
		fixPages(rules, pages)
	}
	return pages
}

func fixUpdates(rules []string, updates []string) [][]int {
	var fixedUpdates [][]int
	for _, update := range updates {
		//fmt.Println("Update: ", update)
		stringPages := strings.Split(update, ",")
		pages := ConvPagesToInts(stringPages)
		metRules := 0
		totalMetRules := 0
		//fmt.Printf("Required matches: %v\n", len(pages)*(len(pages)-1))
		for i, page := range pages {
			fmt.Printf("\n\nPages: %v\n ", pages)
			fmt.Printf("PAGE/i: %v\n", page)
			for j, otherPage := range pages {
				metRules = 0
				fmt.Printf("OTHER PAGE/j: %v\n", otherPage)
				correct := false
				swapped := false
				/* if j == i {
					//fmt.Printf("Page: %v is same as otherPage: %v, doing nothing\n", page, otherPage)
					if i == len(pages)-1 && j == len(pages)-1 {
						fmt.Printf("No rule met or swap found for page: %v and otherPage: %v\n", page, otherPage)
						fixedPages := fixPages(rules, pages)
						fmt.Println("Fixed Pages: ", fixedPages)
					}
					continue
				} */
				for _, rule := range rules {
					/* The page is before the current one */
					/* if j < i {
						//fmt.Printf("Page: %v, Other Page: %v\n", page, otherPage)
						//fmt.Printf("Checking rule (opposite) %v with page combination: %v,%v\n", rule, page, otherPage)
						if CheckIfRuleIsMet(rule, page, otherPage) {
							correct = true
							fmt.Printf("Rule met: %v with page combination: %v,%v\n", rule, page, otherPage)
							metRules++
							break
						} else if CheckIfRuleIsMet(rule, otherPage, page) {
							//metRules++
							pages[i], pages[j] = pages[j], pages[i]
							fmt.Printf("Swapped %v with %v and vice versa. Pages now is: %v\n", pages[i], pages[j], pages)
							swapped = true
							//correct = true
							break
						}
					} */
					/* The page is after the current one */
					if j > i {

						fmt.Printf("Checking rule %v with page combination: %v,%v\n", rule, page, otherPage)
						if CheckIfRuleIsMet(rule, page, otherPage) {
							metRules++
							correct = true
							break
						} else if CheckIfRuleIsMet(rule, otherPage, page) {
							//metRules++
							pages[i], pages[j] = pages[j], pages[i]
							fmt.Printf("Swapped %v with %v and vice versa. Pages now is: %v\n", pages[i], pages[j], pages)
							swapped = true
							//correct = true
							break
						}
					}

				}
				fmt.Printf("Correct: %v, Swapped: %v\n", correct, swapped)
				if swapped {
					break
				}
				if !correct && !swapped && i == len(pages)-1 && j == len(pages)-1 {
					fmt.Printf("No rule met or swap found for last page: %v and last otherPage: %v\n", page, otherPage)
					fmt.Printf("Sending pages to fixPages: %v\n", pages)
					fixedPages := fixPages(rules, pages)
					fmt.Println("Fixed Pages: ", fixedPages)
					fixedUpdates = append(fixedUpdates, fixedPages)

				} else {

				}

				totalMetRules += metRules
			}
			/* if totalMetRules == requiredMetRules {
				fixedUpdates = append(fixedUpdates, update)
			} */

		}

	}
	//fmt.Println("Fixed Updates: ", fixedUpdates)
	return fixedUpdates
}

func extractIncorrectUpdates(rules []string, updates []string) []string {
	var incorrectUpdates []string
	for _, update := range updates {
		//fmt.Println("Update: ", update)
		stringPages := strings.Split(update, ",")
		pages := ConvPagesToInts(stringPages)
		addedToIncorrect := false
		metRules := 0
		totalMetRules := 0
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
					if !addedToIncorrect {
						incorrectUpdates = append(incorrectUpdates, update)
						addedToIncorrect = true
					}

					break
				}

				totalMetRules += metRules
			}

		}

	}
	fmt.Println("Incorrect Updates: ", incorrectUpdates)
	return incorrectUpdates
}
