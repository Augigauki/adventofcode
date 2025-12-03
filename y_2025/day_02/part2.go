package main

import (
	"fmt"
	"strconv"
)

func findAllInvalidIDs(ids [][2]string) []string {
	invalidCount := 0
	invalids := []string{}
	for _, idRange := range ids {
		lower, upper := idRange[0], idRange[1]
		lowerInt, _ := strconv.Atoi(lower)
		upperInt, _ := strconv.Atoi(upper)
		for lowerInt <= upperInt {
			lower = strconv.Itoa(lowerInt)
			if isRepeatingPattern(lower) {
				fmt.Println("Invalid ID found:", lower)
				invalidCount++
				invalids = append(invalids, lower)
				lowerInt++
				continue
			}

			if len(lower) > 0 && lower[0] == '0' {
				fmt.Println("Invalid ID found:", lower)
				invalidCount++
				invalids = append(invalids, lower)
			}
			lowerInt++
		}

	}
	fmt.Println("Total invalid IDs found:", invalidCount)
	return invalids
}

func isRepeatingPattern(s string) bool {
	n := len(s)

	for patternLen := 1; patternLen <= n/2; patternLen++ {
		if n%patternLen == 0 {
			pattern := s[:patternLen]
			isRepeating := true

			for i := patternLen; i < n; i += patternLen {
				if s[i:i+patternLen] != pattern {
					isRepeating = false
					break

				}
			}
			if isRepeating {
				return true
			}
		}

	}
	return false
}
