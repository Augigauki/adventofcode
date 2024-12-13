package main

import "fmt"

func Part1(clawMachines []ClawMachine) {
	fmt.Println("Part 1")
	for i, machine := range clawMachines {
		fmt.Printf(
			"Machine %v: Button A: X+%v, Y+%v, Button B: X+%v, Y+%v, Prize: (%v, %v)\n",
			i+1,
			machine.A.lineIncr, machine.A.charIncr,
			machine.B.lineIncr, machine.B.charIncr,
			machine.Prize.line, machine.Prize.char,
		)
	}
	FindMinimumCost(clawMachines)
}

func findSolution(machine ClawMachine) (int, int, int, bool) {
	minCost := int(^uint(0) >> 1) // Max int value
	var bestA, bestB int

	// Iterate through all possible values of a and b (both up to 100)
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			// Check if this combination of presses aligns the claw with the prize
			if machine.A.charIncr*a+machine.B.charIncr*b == machine.Prize.char &&
				machine.A.lineIncr*a+machine.B.lineIncr*b == machine.Prize.line {
				c := Cost(a, b)
				if c < minCost {
					minCost = c
					bestA = a
					bestB = b
				}
			}
		}
	}

	if minCost < int(^uint(0)>>1) { // If we found a valid solution
		return bestA, bestB, minCost, true
	}
	return 0, 0, 0, false
}

func Cost(a, b int) int {
	return 3*a + b
}

func FindMinimumCost(machines []ClawMachine) int {
	totalCost := 0
	prizesWon := 0

	for _, machine := range machines {
		bestA, bestB, cost, canWin := findSolution(machine)
		if canWin {
			fmt.Printf("Solution found for machine: A=%d, B=%d, Cost=%d\n", bestA, bestB, cost)
			totalCost += cost
			prizesWon++
		} else {
			fmt.Println("No solution for this machine.")
		}
	}

	fmt.Printf("Won %d prizes for a total cost of %d tokens.\n", prizesWon, totalCost)
	return totalCost
}
