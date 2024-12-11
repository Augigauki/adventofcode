package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}
	stones := strings.Fields(input)
	intStones := make([]int, len(stones))
	for i, stone := range stones {
		var err error
		intStones[i], err = strconv.Atoi(stone)
		if err != nil {
			log.Fatalf("Error converting stone to int: %v", stone)
		}
	}

	//newStones := CountStones(input, 6)
	stonesCount := CountIntStones(intStones, 75)
	//fmt.Println("New stones count:", newStones)
	fmt.Println("Stones count:", stonesCount)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
