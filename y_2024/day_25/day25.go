package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var example = false

func main() {
	fileName := ""
	if example {
		fileName = "example.txt"
	} else {
		fileName = "input.txt"
	}
	locks, keys := parseFile(fileName)
	fmt.Println("Locks: ", len(locks))
	fmt.Println("Keys: ", len(keys))

	lockCols := [][]int{}
	fmt.Println("Locks:")
	for _, lock := range locks {
		//fmt.Printf("Lock %d:\n%s\n", i+1, lock)
		colCounts := countColumns(lock, true)
		fmt.Println("Lock column counts: ", colCounts)
		lockCols = append(lockCols, colCounts)
	}
	keyCols := [][]int{}
	fmt.Println("Keys:")
	for _, key := range keys {
		//fmt.Printf("Key %d:\n%s\n", i+1, key)
		colCounts := countColumns(key, false)
		fmt.Println("Key column counts: ", colCounts)
		keyCols = append(keyCols, colCounts)
	}
	fittingKeys := findFittingKeys(lockCols, keyCols)
	fmt.Println("Fitting keys: ", fittingKeys)

}

func findFittingKeys(locks, keys [][]int) int {
	fitting := 0
	for _, lock := range locks {
		for _, key := range keys {
			fit := true
			for k, count := range lock {
				//fmt.Printf("Lock: %d - Key: %d\n", count, key[k])
				if 5-count < key[k] {
					fit = false
					break
				}
			}
			if fit {
				fitting++
			}
		}
	}
	return fitting
}

func countColumns(object string, isLock bool) []int {
	lines := strings.Split(object, "\n")
	numCols := len(lines[0])
	colCounts := make([]int, numCols)

	startRow := 0
	if isLock {
		startRow = 1
	}
	endRow := len(lines)
	if !isLock {
		endRow--
	}
	for row := startRow; row < endRow; row++ {
		for col, char := range lines[row] {
			if char == '#' {
				colCounts[col]++
			}
		}
	}
	return colCounts
}

func parseFile(fileName string) ([]string, []string) {
	locks := []string{}
	keys := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file")
		return nil, nil
	}
	defer file.Close()

	currentObject := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentObject) > 0 {
				object := strings.Join(currentObject, "\n")
				if strings.HasPrefix(currentObject[0], "#####") {
					locks = append(locks, object)
				} else if strings.HasPrefix(currentObject[0], ".....") {
					keys = append(keys, object)
				}
			}
			currentObject = nil
		} else {
			currentObject = append(currentObject, line)
		}
	}
	if len(currentObject) > 0 {
		object := strings.Join(currentObject, "\n")
		if strings.HasPrefix(currentObject[0], "#####") {
			locks = append(locks, object)
		} else if strings.HasPrefix(currentObject[0], ".....") {
			keys = append(keys, object)
		}
	}

	return locks, keys

}
