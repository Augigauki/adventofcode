package main

import (
	"fmt"
)

func getSimilarityScore(left, right []int) {
	var totalMatches []int
	for leftItem := range left {
		var matches int = 0
		var similarityScore int
		for rightItem := range right {
			if left[leftItem] == right[rightItem] {
				matches++
			}
		}
		similarityScore = left[leftItem] * matches
		if similarityScore > 0 {
			totalMatches = append(totalMatches, similarityScore)
		}
	}
	var sum int
	for item := range totalMatches {
		sum += totalMatches[item]
	}
	fmt.Println("Similarity Score: ", sum)
}