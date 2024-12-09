package main

import (
	"fmt"
)

func MoveEntireFiles(diskMap string) {
	numbers := ConvertToNumberSlice(diskMap)
	sortedMap := SortDiskMap(numbers)
	expandedDiskMap := ExpandDiskMapToSlices(sortedMap)
	swappedMap := swapUntilDoneAgain(expandedDiskMap)
	//fmt.Println("Swapped disk map: ", swappedMap)
	checksum := CalculateNewChecksum(swappedMap)
	fmt.Println("New checksum: ", checksum)
}

func moveToFirstAvailableFreeSpace(diskMap [][]int, processedFiles map[int]bool) [][]int {
	lastFile, lastFileIndex := findLastFile(diskMap, processedFiles)

	if lastFile == nil {
		return diskMap // No unprocessed files left
	}

	freeSpaceIndex := findMatchingFreeSpaceIndex(diskMap, lastFile)

	if freeSpaceIndex != -1 && freeSpaceIndex < lastFileIndex {
		// Move the file to the identified free space
		//fmt.Printf("Moving file %v from index %d to free space at index %d\n", lastFile, lastFileIndex, freeSpaceIndex)

		// Replace `-1` values in the free space slice
		freeSpaceSlice := diskMap[freeSpaceIndex]
		fileSize := len(lastFile)
		inserted := 0
		for i := 0; i < len(freeSpaceSlice) && inserted < fileSize; i++ {
			if freeSpaceSlice[i] == -1 {
				freeSpaceSlice[i] = lastFile[inserted]
				inserted++
			}
		}

		// Mark the original file slice as free space (-1)
		for i := range diskMap[lastFileIndex] {
			diskMap[lastFileIndex][i] = -1
		}
	}

	// Mark this file as processed
	processedFiles[lastFileIndex] = true
	return diskMap
}

func findMatchingFreeSpaceIndex(diskMap [][]int, file []int) int {
	fileSize := len(file)
	for i, slice := range diskMap {
		freeSpaceCount := 0

		// Count the number of consecutive `-1` values in the slice
		for _, val := range slice {
			if val == -1 {
				freeSpaceCount++
			} else {
				freeSpaceCount = 0 // Reset if a non-free space is encountered
			}

			// Check if the file can fit in this free space
			if freeSpaceCount >= fileSize {
				return i
			}
		}
	}
	return -1 // No matching free space found
}

func findLastFile(diskMap [][]int, processedFiles map[int]bool) ([]int, int) {
	//fmt.Println("Finding last file...")
	var lastFile []int
	var index int
	for i := len(diskMap) - 1; i >= 0; i-- {
		if processedFiles[i] {
			continue
		}
		//fmt.Printf("Checking diskMap[%d]: %d\n", i, diskMap[i])
		justFile := true
		for j := 0; j < len(diskMap[i]); j++ {
			if diskMap[i][j] == -1 {
				justFile = false
				break
			}
		}
		if justFile {
			//fmt.Println("Found last file: ", diskMap[i])
			lastFile = diskMap[i]
			index = i
			return lastFile, index
		}
	}
	return lastFile, -1
}

func swapUntilDoneAgain(diskMap [][]int) [][]int {
	processedFiles := make(map[int]bool)
	fmt.Println("Swapping until done...")

	for {
		progress := false

		// Attempt to move the last file
		for i := 0; i < len(diskMap); i++ {
			if !processedFiles[i] {
				prevMap := fmt.Sprintf("%v", diskMap)

				// Process the current file
				diskMap = moveToFirstAvailableFreeSpace(diskMap, processedFiles)

				currMap := fmt.Sprintf("%v", diskMap)
				if prevMap != currMap {
					progress = true // Indicate that progress was made
				}
			}
		}

		// Check if progress was made or if all files are processed
		if !progress || len(processedFiles) == len(diskMap) {
			break
		}
	}
	return diskMap
}

func ExpandDiskMapToSlices(diskMap SortedMap) [][]int {
	//var diskMapString string
	var diskMapSlice [][]int
	for i := 0; i < len(diskMap.files); i++ {
		// Add file blocks
		var file []int
		for j := 0; j < diskMap.files[i]; j++ {
			file = append(file, i)
		}
		if len(file) > 0 {
			diskMapSlice = append(diskMapSlice, file)
		}
		file = nil
		// Add free space blocks
		var freeSpace []int
		if i < len(diskMap.freeSpace) {
			for j := 0; j < diskMap.freeSpace[i]; j++ {
				freeSpace = append(freeSpace, -1)
			}
			if len(freeSpace) > 0 {
				diskMapSlice = append(diskMapSlice, freeSpace)
			}
			freeSpace = nil
		}
	}
	return diskMapSlice
}

func CalculateNewChecksum(diskMap [][]int) int {
	var checksum int
	var nums []int

	for i := 0; i < len(diskMap); i++ {
		for j := 0; j < len(diskMap[i]); j++ {
			nums = append(nums, diskMap[i][j])
			if diskMap[i][j] != -1 {
			}
		}
	}
	//fmt.Println("Numbers: ", nums)
	fmt.Println("Calculating new checksum...")
	for i, num := range nums {
		if num != -1 {
			checksum += i * num

		}
	}
	return checksum
}
