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

	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		var rangeBounds [][2]string
		for _, r := range ranges {
			bounds := strings.Split(r, "-")
			var lower, upper string = bounds[0], bounds[1]
			rangeBounds = append(rangeBounds, [2]string{lower, upper})
		}
		//fmt.Println(rangeBounds)
		invalids := findAllInvalidIDs(rangeBounds)
		sumInvalids := 0
		for _, id := range invalids {
			num, _ := strconv.Atoi(id)
			sumInvalids += num
		}
		fmt.Println("Invalid IDs sum:", sumInvalids)
	}
}
