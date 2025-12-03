package main

import (
	"fmt"
	"strconv"
)

func findInvalidIDs(ids [][2]string) []string {
	invalidCount := 0
	invalids := []string{}
	fmt.Println("ids", ids)
	for _, idRange := range ids {
		lower, upper := idRange[0], idRange[1]
		lowerInt, _ := strconv.Atoi(lower)
		upperInt, _ := strconv.Atoi(upper)
		for lowerInt <= upperInt {
			lower = strconv.Itoa(lowerInt)
			//fmt.Println("Checking ID:", lower)
			mid := len(lower) / 2
			firstHalf := lower[:mid]
			secondHalf := lower[mid:]
			if firstHalf == secondHalf {
				fmt.Println("Invalid ID found:", lower)
				invalidCount++
				invalids = append(invalids, lower)

			}
			if len(firstHalf) > 0 && firstHalf[0] == '0' {
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
