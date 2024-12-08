package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// This is a placeholder for the main function
	file, err := os.Open("example2.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var antennaMap [][]string
	for scanner.Scan() {
		var line string = scanner.Text()
		chars := []string{}
		for _, char := range line {
			chars = append(chars, string(char))
		}
		antennaMap = append(antennaMap, chars)
	}
	FindAntinodes(antennaMap)
}
