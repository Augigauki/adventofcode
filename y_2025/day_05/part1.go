package main

import (
	"strconv"
	"strings"
)

func findSafeIDs(freshRanges []string, ingredients []int) int {
	safeIDCount := 0
	for _, id := range ingredients {
		for _, freshID := range freshRanges {
			parts := strings.Split(freshID, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			if id >= start && id <= end {
				safeIDCount++
				break
			}
		}
	}

	return safeIDCount
}
