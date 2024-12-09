package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type SortedMap struct {
	files     []int
	freeSpace []int
}

func MoveFiles(diskMap string) {
	//fmt.Println("Disk map: ", diskMap)
	numbers := convertToNumberSlice(diskMap)
	//isEven := len(diskMap)%2 == 0
	//fmt.Println("Is even: ", isEven)
	sortedMap := sortDiskMap(numbers)
	fmt.Println("::SORTED MAP::")
	fmt.Println("Files: ", sortedMap.files)
	fmt.Println("Free spaces: ", sortedMap.freeSpace)
	expandedDiskMap := expandDiskMap2(sortedMap)
	fmt.Println("Expanded disk map: ", expandedDiskMap)
	movedDiskMap := swapUntilDone(expandedDiskMap)
	isSwapped := validateDiskMap(movedDiskMap)
	if !isSwapped {
		fmt.Println("Disk map is invalid.")
		return
	}
	fmt.Println("Moved disk map: ", movedDiskMap)
	checksum := calculateChecksum(movedDiskMap)
	fmt.Println("Checksum: ", checksum)
}

func sortDiskMap(numbers []int) SortedMap {
	fmt.Println("Sorting disk map...")
	var files = []int{}
	var freeSpace = []int{}
	for i, num := range numbers {
		if i%2 == 0 {
			files = append(files, num)
		} else {
			freeSpace = append(freeSpace, num)
		}
	}
	return SortedMap{files, freeSpace}
}

/* func expandDiskMap(diskMap SortedMap) string {
	fmt.Println("Expanding disk map...")
	totalLength := len(diskMap.files) + len(diskMap.freeSpace)
	filesSlice := diskMap.files
	freeSpaceSlice := diskMap.freeSpace
	diskMapString := ""
	filesIndex := 0
	freeSpaceIndex := 0
	for i := 0; i < totalLength; i++ {
		for filesIterator := 0; filesIterator < filesSlice[filesIndex]; filesIterator++ {
			diskMapString += strconv.Itoa(filesIndex)
		}
		filesIndex++
		i++
		if i >= totalLength {
			break
		}
		for freeSpace := 0; freeSpace < freeSpaceSlice[freeSpaceIndex]; freeSpace++ {
			diskMapString += "."
		}
		freeSpaceIndex++
	}
	return diskMapString
} */

func expandDiskMap2(diskMap SortedMap) string {
	var diskMapString string
	for i := 0; i < len(diskMap.files); i++ {
		// Add file blocks
		for j := 0; j < diskMap.files[i]; j++ {
			if i > 9 {
				diskMapString += "-" + strconv.Itoa(i) + "-"
			} else {
				diskMapString += strconv.Itoa(i)
			}

		}
		// Add free space blocks
		if i < len(diskMap.freeSpace) {
			for j := 0; j < diskMap.freeSpace[i]; j++ {
				diskMapString += "."
			}
		}
	}
	return diskMapString
}

func swapFileAndFreeSpaceOptimized(diskMap string) string {
	runes := []rune(diskMap)
	freeSpaceIndex := 0
	fileIndex := len(runes) - 1

	//fmt.Println("Swapping file and free space...")
	for freeSpaceIndex < fileIndex {
		// Find the next free space from the left
		for freeSpaceIndex < len(runes) && runes[freeSpaceIndex] != '.' {
			freeSpaceIndex++
		}
		// Find the next file block from the right
		for fileIndex >= 0 && !unicode.IsDigit(runes[fileIndex]) {
			fileIndex--
		}
		// Swap if valid
		if freeSpaceIndex < fileIndex {
			runes[freeSpaceIndex], runes[fileIndex] = runes[fileIndex], runes[freeSpaceIndex]
			freeSpaceIndex++
			fileIndex--
		}
	}
	fmt.Println("Swapped disk map: ", string(runes))
	return string(runes)
}

func swapUntilDone(diskMap string) string {
	expanded := diskMap
	done := false
	fmt.Println("Swapping until done...")
	for !done {
		expanded = swapFileAndFreeSpaceOptimized(expanded)
		done = true
		for i := 0; i < len(expanded)-1; i++ {
			if expanded[i] == '.' && expanded[i+1] >= '0' && expanded[i+1] <= '9' {
				done = false
			}
		}
	}
	return expanded
}

func validateDiskMap(diskMap string) bool {
	seenFreeSpace := false
	for _, char := range diskMap {
		if char == '.' {
			seenFreeSpace = true
		} else if seenFreeSpace {
			return false
		}
	}
	return true
}

func calculateChecksum(diskMap string) int64 {
	var checksum int64 = 0
	fileID := int64(-1) // Initialize to -1 to detect new file IDs
	for i, char := range diskMap {
		if unicode.IsDigit(char) {
			if fileID == -1 || int64(char-'0') != fileID {
				fileID = int64(char - '0')
			}
			checksum += fileID * int64(i)
		}
	}
	return checksum
}

func convertToNumberSlice(diskMap string) []int {
	numbers := []int{}
	for _, char := range diskMap {
		text := string(char)
		number, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Error converting to int")
			break
		}
		numbers = append(numbers, number)
	}
	return numbers
}
