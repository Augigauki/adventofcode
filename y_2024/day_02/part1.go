package main

import (
	"fmt"
)

func findSafe(reports [][]int) {
	var safeReports [][]int
	for _, report := range reports {
		var safe bool = true
		var increasing bool = report[0] < report[1]
		fmt.Println("Increasing: ", increasing)
		for i, level := range report {
			if increasing {
				if i+1 == len(report) {
					fmt.Println("End of line")
					break
				}
				if i+1 < len(report) {
					fmt.Println("Checking: ", level, report[i+1])
					if level < report[i+1] {
						fmt.Println("Still increasing :)")
						if report[i+1]-level <= 3 {
							fmt.Println("Difference is less than 3")
						} else {
							fmt.Println("Unsafe! Breaking")
							safe = false
						}
					} else if level > report[i+1] {
						fmt.Println("Decreasing! Breaking")
						safe = false
						break

					} else {
						fmt.Println("Equal! Breaking")
						safe = false
						break
					}

				}
			}
			if !increasing {
				if i+1 == len(report) {
					fmt.Println("End of line")
					break
				}
				if i+1 < len(report) {
					fmt.Println("Checking: ", level, report[i+1])
					if level > report[i+1] {
						fmt.Println("Still decreasing :)")
						if level-report[i+1] <= 3 {
							fmt.Println("Difference is less than 3")
						} else {
							fmt.Println("Unsafe! Breaking")
							safe = false
						}
					} else if level < report[i+1] {

						fmt.Println("Increasing! Breaking")
						safe = false
						break
					} else {
						fmt.Println("Equal! Breaking")
						safe = false
						break
					}

				} /* else {
					fmt.Println("Unsafe! Breaking")
					safe = false
					break
				} */

			}
		}

		if safe {
			safeReports = append(safeReports, report)
		}
	}
	fmt.Println(safeReports)
	fmt.Println("Safe reports: ", len(safeReports))
}
