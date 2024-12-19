package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var example = true

func main() {
	fileName := ""
	if example {
		fileName = "example.txt"
	} else {
		fileName = "input.txt"
	}
	patterns, designs := parseFile(fileName)
	fmt.Println("Patterns:")
	for _, pattern := range patterns {
		fmt.Println(pattern)
	}
	fmt.Println("\nDesigns:")
	for _, design := range designs {
		fmt.Println(design)
	}
	possibleDesigns := 0
	for _, design := range designs {
		matching := getMatchingPatterns(patterns, design)
		canBeMade := canDesignBeMade(matching, design, map[string]bool{})
		if canBeMade {
			possibleDesigns++
		}
	}
	fmt.Println("Possible designs: ", possibleDesigns)
}

func getMatchingPatterns(patterns []string, design string) []string {
	matchingPatterns := []string{}
	for _, pattern := range patterns {
		if strings.Contains(design, pattern) {
			matchingPatterns = append(matchingPatterns, pattern)
		}
	}
	return matchingPatterns
}

func canDesignBeMade(patterns []string, design string, memo map[string]bool) bool {
	if result, ok := memo[design]; ok { // Check if result is already memoized
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
