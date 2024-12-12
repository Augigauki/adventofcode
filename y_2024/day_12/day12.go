package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/* type Pos struct {
	line int
	char int
}

type GardenMapPos struct {
	value string
	pos   Pos
} */

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file: ", fileName)
	}
	defer file.Close()

	var gardenMap [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gardenMap = append(gardenMap, strings.Split(scanner.Text(), ""))
	}
	//CalcFenceCost(gardenMap)
	CalcFenceCostWithDiscount(gardenMap)

}
