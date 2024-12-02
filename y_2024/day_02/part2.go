package main

import (
	"fmt"
)

var safeReports [][]int

func findSafeWithExtraTolerance(reports [][]int) {
	for _, report := range reports {
		isReportSafe := checkReport(report)
		if !isReportSafe {

			for i := 0; i < len(report); i++ {
				testReport := make([]int, 0, len(report)-1)
				testReport = append(testReport, report[:i]...)
				testReport = append(testReport, report[i+1:]...)
				fmt.Printf("Excluding index %d: %v\n", i, testReport)
				testSafe := checkReport(testReport)
				if testSafe {
					fmt.Println("Safe report: ", testReport)
					safeReports = append(safeReports, testReport)
					break
				}
			}
		} else {
			safeReports = append(safeReports, report)
		}
	}
	fmt.Println(safeReports)
	fmt.Println("Safe reports: ", len(safeReports))
}

func checkReport(report []int) bool {
	fmt.Println("Checking report: ", report)
	var safe bool
	var increasing bool = report[0] < report[1]
	if increasing {
		fmt.Println(":::INCREASING:::")
	} else {
		fmt.Println(":::DECREASING:::")
	}
	for i, level := range report {

		if increasing {

			if i+1 == len(report) {
				fmt.Println("End of line")
				safe = true
				break
			}
			if i+1 < len(report) {
				fmt.Println("Checking: ", level, report[i+1])
				if level < report[i+1] {
					if report[i+1]-level <= 3 {
						fmt.Println("Difference is less than 3")
						continue
					} else {
						fmt.Println("Difference is greater than 3! Breaking")
						safe = false
						break
					}
				} else if level > report[i+1] {
					fmt.Println("Decreasing when it should be increasing! Breaking")
					safe = false
					break
				} else {
					fmt.Println("Equal! Breaking")
					safe = false
					break
				}
			}
		}
		/* :::DECREASING::: */
		if !increasing {
			if i+1 == len(report) {
				fmt.Println("End of line")
				safe = true
				break
			}
			if i+1 < len(report) {
				fmt.Println("Checking: ", level, report[i+1])
				if level > report[i+1] {
					//fmt.Println("Still decreasing :)")
					if level-report[i+1] <= 3 {
						fmt.Println("Difference is less than 3")
					} else {
						fmt.Println("Difference is greater than 3! Breaking")
						safe = false
						break
					}
				} else if level < report[i+1] {
					fmt.Println("Increasing when it should be decreasing! Breaking")
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
	return safe
}
