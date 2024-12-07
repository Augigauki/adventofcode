package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("oneline.txt")

	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var equations []string
	for scanner.Scan() {
		var line string = scanner.Text()
		equations = append(equations, line)
	}
	FindTrueEquations(equations)
}
