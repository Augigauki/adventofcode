package main

import (
	"fmt"
	"sort"
)

func getTotalDiffs(left, right []int) {
	sort.Ints(left)
	sort.Ints(right)
	var diffs []int
	for item := range left {
		var diff int
			if(left[item] > right[item]){
				diff = left[item] - right[item]
			} else {
				diff = right[item] - left[item]
			}
			diffs = append(diffs, diff)
	}
	
	var total int
	for item := range diffs {
		total += diffs[item]
	}
	fmt.Println("Total: ", total)
}