package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	joltages := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		digits := []int{}
		for _, char := range line {
			digit := int(char - '0')
			digits = append(digits, digit)

		}
		highestJoltage := findLargestJoltage(digits)
		joltages = append(joltages, highestJoltage)

	}
	fmt.Println("Joltages:", joltages)
	combined := 0
	for _, joltage := range joltages {
		combined += joltage
	}
	fmt.Println("Combined Joltages:", combined)
}
