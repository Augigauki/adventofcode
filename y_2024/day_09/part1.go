package main

import (
	"fmt"
	"strconv"
)

type SortedMap struct {
	files     []int
	freeSpace []int
}

func MoveFiles(diskMap string) {
	//fmt.Println("Disk map: ", diskMap)
	numbers := ConvertToNumberSlice(diskMap)
	//isEven := len(diskMap)%2 == 0
	//fmt.Println("Is even: ", isEven)
	sortedMap := SortDiskMap(numbers)
	fmt.Println("::SORTED MAP::")
	fmt.Println("Files: ", sortedMap.files)
	fmt.Println("Free spaces: ", sortedMap.freeSpace)
	expandedDiskMap := ExpandDiskMap(sortedMap)
	fmt.Println("Expanded disk map: ", expandedDiskMap)
	movedDiskMap := swapUntilDone(expandedDiskMap)
	isSwapped := ValidateDiskMap(movedDiskMap)
	if !isSwapped {
		fmt.Println("Disk map is invalid.")
		return
	}
	fmt.Println("Moved disk map: ", movedDiskMap)
	checksum := CalculateChecksum(movedDiskMap)
	fmt.Println("Checksum: ", checksum)
}

func SortDiskMap(numbers []int) SortedMap {
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

func ExpandDiskMap(diskMap SortedMap) []int {
	//var diskMapString string
	var diskMapSlice []int
	for i := 0; i < len(diskMap.files); i++ {
		// Add file blocks
		for j := 0; j < diskMap.files[i]; j++ {
			diskMapSlice = append(diskMapSlice, i)
			/* if i > 9 {
				diskMapString += "-" + strconv.Itoa(i) + "-"
			} else {
				diskMapString += strconv.Itoa(i)
			} */

		}
		// Add free space blocks
		if i < len(diskMap.freeSpace) {
			for j := 0; j < diskMap.freeSpace[i]; j++ {
				diskMapSlice = append(diskMapSlice, -1)
				//diskMapString += "."
			}
		}
	}
	return diskMapSlice
}

func swapFileAndFreeSpaceOptimized(diskMap []int) []int {
	//runes := []rune(diskMap)
	freeSpaceIndex := 0
	fileIndex := len(diskMap) - 1

	//fmt.Println("Swapping file and free space...")
	for freeSpaceIndex < fileIndex {
		// Find the next free space from the left
		for freeSpaceIndex < len(diskMap) && diskMap[freeSpaceIndex] != -1 {
			freeSpaceIndex++
		}
		// Find the next file block from the right
		for fileIndex >= 0 && diskMap[fileIndex] == -1 {
			fileIndex--
		}
		// Swap if valid
		if freeSpaceIndex < fileIndex {
			diskMap[freeSpaceIndex], diskMap[fileIndex] = diskMap[fileIndex], diskMap[freeSpaceIndex]
			freeSpaceIndex++
			fileIndex--
		}
	}
	//fmt.Println("Swapped disk map: ", diskMap)
	return diskMap
}

func swapUntilDone(diskMap []int) []int {
	expanded := diskMap
	done := false
	fmt.Println("Swapping until done...")
	for !done {
		expanded = swapFileAndFreeSpaceOptimized(expanded)
		done = true
		for i := 0; i < len(expanded)-1; i++ {
			if expanded[i] == -1 && expanded[i+1] >= 0 {
				done = false
			}
		}
	}
	return expanded
}

func ValidateDiskMap(diskMap []int) bool {
	seenFreeSpace := false
	for _, char := range diskMap {
		if char == -1 {
			seenFreeSpace = true
		} else if seenFreeSpace {
			return false
		}
	}
	return true
}

func CalculateChecksum(diskMap []int) int {
	var checksum int = 0

	//fileID := int64(-1) // Initialize to -1 to detect new file IDs
	/* for i, char := range diskMap {
		if char >= 0 {
			if fileID == -1 || int64(char-'0') != fileID {
				fileID = int64(char - '0')
			}
			checksum += fileID * int64(i)
		}
	} */
	fmt.Println("Calculating checksum...")
	for i, num := range diskMap {
		if num >= 0 {
			//fmt.Printf("i: %d, num: %d. Multiplied: %d\n", i, num, num*i)
			checksum += num * i
		}
	}
	return checksum
}

func ConvertToNumberSlice(diskMap string) []int {
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
