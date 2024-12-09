package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("short.txt")

	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	if len(lines) != 1 {
		fmt.Println("Invalid input: expected one line of disk map")
		return
	}
	diskMap := lines[0]
	/* for scanner.Scan() {
		diskMap = scanner.Text()
	} */
	//fmt.Println("Disk map: ", diskMap)
	MoveFiles(diskMap)
}
