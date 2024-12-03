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
	var memory string
	/* var allSums []int
	total := 0 */
	completeString := ""
	for scanner.Scan() {

		memory = scanner.Text()
		completeString += memory
		//sum, state := MulNumsDoDont(memory, true)
		//allSums = append(allSums, sum)
		//fmt.Println(memory)
	}
	/* for _, sum := range allSums {
		total += sum
	} */
	total := MulNumsDoDont(completeString)
	fmt.Println(total)
	//var sum int = MulNums(memory)
	//fmt.Println("Sum: ", sum)
}
