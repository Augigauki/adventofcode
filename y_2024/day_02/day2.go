package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var levels [][]int
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
		//fmt.Println(trimmed)
		levels = append(levels, trimmed)
	}
	fmt.Println(levels)
	findSafe(levels)
}