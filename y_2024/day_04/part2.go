package main

func FindMasMas(input [][]string) int {
	totalMaxMas := 0
	for i, row := range input {
		for j, char := range row {
			if char == "A" {
				diag1, diag2 := false, false
				downRight, downLeft, upRight, upLeft := "", "", "", ""
				/* Checking diagonal down-right */
				if j < len(row)-1 && i < len(input)-1 {
					if input[i+1][j+1] == "S" || input[i+1][j+1] == "M" {
						downRight = input[i+1][j+1]
					}
				}
				/* Checking diagonal down-left */
				if j > 0 && i < len(input)-1 {
					if input[i+1][j-1] == "S" || input[i+1][j-1] == "M" {
						downLeft = input[i+1][j-1]
					}
				}
				/* Checking diagonal up-right */
				if j < len(row)-1 && i > 0 {
					if input[i-1][j+1] == "S" || input[i-1][j+1] == "M" {
						upRight = input[i-1][j+1]
					}
				}
				/* Checking diagonal up-left */
				if j > 0 && i > 0 {
					if input[i-1][j-1] == "S" || input[i-1][j-1] == "M" {
						upLeft = input[i-1][j-1]
					}
				}
				diag1 = (upLeft == "M" && downRight == "S") || (upLeft == "S" && downRight == "M")
				diag2 = (upRight == "M" && downLeft == "S") || (upRight == "S" && downLeft == "M")
				if diag1 && diag2 {
					totalMaxMas++
				}

			}
		}
	}
	return totalMaxMas
}
