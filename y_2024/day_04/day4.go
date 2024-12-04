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
	var wordGrid [][]string
	for scanner.Scan() {
		var line string = scanner.Text()
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		wordGrid = append(wordGrid, row)
	}
	//fmt.Println(wordGrid)
	totalXmas := FindXmas(wordGrid)
	fmt.Println("Total XMAS: ", totalXmas)
}
