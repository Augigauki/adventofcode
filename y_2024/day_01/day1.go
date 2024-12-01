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
	var left, right []int
	for scanner.Scan(){
		var line string = scanner.Text()
		parts := strings.Split(line, " ")
			var trimmed []int
			for _, item := range parts {
				item = strings.TrimSpace(item)
				if num, err := strconv.Atoi(item); err == nil {
					trimmed = append(trimmed, num)
				}
			}
			left = append(left, trimmed[0])
			right = append(right, trimmed[1])
	}
	
	getTotalDiffs(left, right)
	getSimilarityScore(left, right)
	

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
