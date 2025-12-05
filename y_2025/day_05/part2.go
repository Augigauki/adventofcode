package main

import "fmt"

func findAllSafeIDs(freshRanges [][]int) int {
	sum := 0
	sortedRanges := sortRanges(freshRanges)
	fmt.Println("Sorted Ranges:", sortedRanges)
	mergedRanges := mergeRanges(sortedRanges)
	fmt.Println("Merged Ranges:", mergedRanges)
	for _, r := range mergedRanges {
		sum += (r[1] - r[0] + 1)
	}
	return sum
}

func sortRanges(freshRanges [][]int) [][]int {
	//sort ranges by start value
	for i := 0; i < len(freshRanges)-1; i++ {
		for j := 0; j < len(freshRanges)-i-1; j++ {
			if freshRanges[j][0] > freshRanges[j+1][0] {
				freshRanges[j], freshRanges[j+1] = freshRanges[j+1], freshRanges[j]
			}
		}
	}
	return freshRanges
}

func mergeRanges(ranges [][]int) [][]int {
	//combine ranges if next range starts before current range ends
	for i := 0; i < len(ranges)-1; i++ {
		currentRange := ranges[i]
		nextRange := ranges[i+1]
		if currentRange[1] >= nextRange[0]-1 {
			//merge ranges
			mergedRange := []int{currentRange[0], max(currentRange[1], nextRange[1])}
			ranges[i] = mergedRange
			//remove next range
			ranges = append(ranges[:i+1], ranges[i+2:]...)
			i-- //check the merged range again
		}
	}
	return ranges
}
