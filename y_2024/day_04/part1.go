package main

import (
	"fmt"
)

func FindXmas(input [][]string) int {
	//fmt.Println(input)
	totalXmas := 0
	for i, row := range input {
		for j, char := range row {
			if char == "X" {
				/* Checking forward */
				if j < len(row)-3 {
					if row[j+1] == "M" && row[j+2] == "A" && row[j+3] == "S" {
						fmt.Println("XMAS forward found at: ", i+1, j+1)
						totalXmas++
					}
				}
				/* Checking backward */
				if j > 2 {
					if row[j-1] == "M" && row[j-2] == "A" && row[j-3] == "S" {
						fmt.Println("XMAS backward found at: ", i+1, j+1)
						totalXmas++
					}
				}
				/* Checking down */
				if i < len(input)-3 {
					if input[i+1][j] == "M" && input[i+2][j] == "A" && input[i+3][j] == "S" {
						fmt.Println("XMAS down found at: ", i+1, j+1)
						totalXmas++
					}
				}
				/* Checking up */
				if i > 2 {
					if input[i-1][j] == "M" && input[i-2][j] == "A" && input[i-3][j] == "S" {
						fmt.Println("XMAS up found at: ", i+1, j+1)
						totalXmas++
					}
				}
				/* Checking diagonal down-right */
				if j < len(row)-3 && i < len(input)-3 {
					if input[i+1][j+1] == "M" && input[i+2][j+2] == "A" && input[i+3][j+3] == "S" {
						fmt.Println("XMAS diagonal down-right found at: ", i+1, j+1)
						totalXmas++
					}
				}
				/* Checking diagonal down-left */
				if j > 2 && i < len(input)-3 {
					if input[i+1][j-1] == "M" && input[i+2][j-2] == "A" && input[i+3][j-3] == "S" {
						fmt.Println("XMAS diagonal down-left found at: ", i+1, j+1)
						totalXmas++
					}
				}
				/* Checking diagonal up-right */
				if j < len(row)-3 && i > 2 {
					if input[i-1][j+1] == "M" && input[i-2][j+2] == "A" && input[i-3][j+3] == "S" {
						fmt.Println("XMAS diagonal up-right found at: ", i+1, j+1)
						totalXmas++
					}
				}
				/* Checking diagonal up-left */
				if j > 2 && i > 2 {
					if input[i-1][j-1] == "M" && input[i-2][j-2] == "A" && input[i-3][j-3] == "S" {
						fmt.Println("XMAS diagonal up-left found at: ", i+1, j+1)
						totalXmas++
					}
				}
			}
		}
	}
	return totalXmas
}
