package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rules []string
	var updates []string
	for scanner.Scan() {
		var line string = scanner.Text()
		if strings.Contains(line, "|") {
			rules = append(rules, line)
		} else if strings.Contains(line, ",") {
			updates = append(updates, line)
		} else {
			continue
		}
	}
	FindUpdatesInCorrectOrder(rules, updates)

}
