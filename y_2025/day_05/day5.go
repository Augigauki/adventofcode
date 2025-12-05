package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var freshRanges []string
	var ingredientIDs []int
	var intRanges [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			intRanges = append(intRanges, []int{start, end})
			freshRanges = append(freshRanges, line)
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				continue
			}

			ingredientIDs = append(ingredientIDs, id)
		}
	}

	/* fmt.Println("Ingredient IDs:", ingredientIDs)
	validCount := findSafeIDs(freshRanges, ingredientIDs)
	fmt.Println("Number of valid ingredient IDs:", validCount) */
	fmt.Println("Int ranges:", intRanges)
	allSafeIDs := findAllSafeIDs(intRanges)
	fmt.Println("All safe ingredient IDs:", allSafeIDs)

}
