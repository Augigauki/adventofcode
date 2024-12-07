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
	var equations []string
	for scanner.Scan() {
		var line string = scanner.Text()
		equations = append(equations, line)
	}
	extraEquations := GenerateExtraEquations(equations)
	fmt.Println("Extra equations:")
	for _, eq := range extraEquations {
		fmt.Println(eq)
	}
	/* Part 1 */
	FindTrueEquations(equations)
	/* Part 2 */
	//FindTrueEquations(extraEquations)
}
