package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var example = false

func main() {
	fileName := ""
	if example {
		fileName = "example.txt"
	} else {
		fileName = "input.txt"
	}
	patterns, designs := parseFile(fileName)

	possibleDesigns := 0
	fmt.Println("\nPart 1")
	for _, design := range designs {
		if canDesignBeMade(patterns, design, map[string]bool{}) {
			possibleDesigns++
		}
	}
	fmt.Println("Possible designs: ", possibleDesigns)

	fmt.Println("\nPart 2:")
	allPossibleWays := 0
	for _, design := range designs {
		fmt.Println("Checking with BFS for design:", design)
		ways := countWaysToMakeDesign(patterns, design, map[string]int{})
		fmt.Println("Ways to make design:", ways)
		allPossibleWays += ways
	}
	fmt.Println("All possible ways to make designs:", allPossibleWays)
}

func canDesignBeMade(patterns []string, design string, memo map[string]bool) bool {
	if result, ok := memo[design]; ok {
		return result
	}
	if design == "" {
		memo[design] = true
		return true
	}
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			remainingDesign := strings.TrimPrefix(design, pattern)
			if canDesignBeMade(patterns, remainingDesign, memo) {
				memo[design] = true
				return true
			}
		}
	}
	memo[design] = false
	return false
}

func countWaysToMakeDesign(patterns []string, design string, cache map[string]int) int {
	if ways, ok := cache[design]; ok {
		return ways
	}

	if design == "" {
		return 1 // Base case: one way to make an empty design
	}

	ways := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			remainingDesign := strings.TrimPrefix(design, pattern)
			ways += countWaysToMakeDesign(patterns, remainingDesign, cache)
		}
	}

	cache[design] = ways
	return ways
}

func parseFile(fileName string) ([]string, []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error reading file: ", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	patterns := []string{}
	designs := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			ptrns := strings.Split(line, ", ")
			patterns = append(patterns, ptrns...)

		} else if line == "" {
			continue
		} else {
			designs = append(designs, line)
		}

	}
	return patterns, designs
}
