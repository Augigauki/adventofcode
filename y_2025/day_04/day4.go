package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var paperMap [][]string
	for scanner.Scan() {
		line := scanner.Text()
		chars := []string{}
		for _, char := range line {
			chars = append(chars, string(char))
		}
		paperMap = append(paperMap, chars)
	}
	papers := removeAllPaper(paperMap)
	fmt.Println("All papers found:", papers)
}
