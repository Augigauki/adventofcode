package main

import "fmt"

func Part2(clawMachines []ClawMachine) {
	fmt.Println("Part 2")
	for i, machine := range clawMachines {
		clawMachines[i].Prize.line += 10000000000000
		clawMachines[i].Prize.char += 10000000000000
		fmt.Printf(
			"Machine %v: Button A: X+%v, Y+%v, Button B: X+%v, Y+%v, Prize: (%v, %v)\n",
			i+1,
			machine.A.lineIncr, machine.A.charIncr,
			machine.B.lineIncr, machine.B.charIncr,
			machine.Prize.line, machine.Prize.char,
		)
	}
	findMinimumCostDiophantine(clawMachines)
}

func solveDiophantine(A, B, Prize int) (bool, int, int) {
	gcd, x0, y0 := extendedGCD(A, B)
	if Prize%gcd != 0 {
		return false, 0, 0
	}

	x0 *= Prize / gcd
	y0 *= Prize / gcd

	return true, x0, y0
}

func extendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

func findSolutionPart2(machine ClawMachine) (int, int, int, bool) {
	minCost := int(^uint(0) >> 1) // Max int value
	var bestA, bestB int

	// Solve for X-axis
	canSolveX, aX, bX := solveDiophantine(machine.A.charIncr, machine.B.charIncr, machine.Prize.char)
	if !canSolveX {
		return 0, 0, 0, false
	}

	// Solve for Y-axis
	canSolveY, aY, bY := solveDiophantine(machine.A.lineIncr, machine.B.lineIncr, machine.Prize.line)
	if !canSolveY {
		return 0, 0, 0, false
	}

	// Combine solutions
	fmt.Printf("Diophantine solved for X: aX=%d, bX=%d (canSolve=%t)\n", aX, bX, canSolveX)
	fmt.Printf("Diophantine solved for Y: aY=%d, bY=%d (canSolve=%t)\n", aY, bY, canSolveY)

	a, b, valid := combineSolutions(machine, aX, bX, aY, bY)
	if !valid {
		return 0, 0, 0, false
	}

	// Minimize cost
	cost := Cost(a, b)
	if cost < minCost {
		minCost = cost
		bestA, bestB = a, b
	}

	return bestA, bestB, minCost, true
}

func combineSolutions(machine ClawMachine, aX, bX, aY, bY int) (int, int, bool) {

	gcdX := gcd(machine.A.charIncr, machine.B.charIncr)
	gcdY := gcd(machine.A.lineIncr, machine.B.lineIncr)

	periodX := machine.B.charIncr / gcdX
	periodY := machine.B.lineIncr / gcdY

	// Align aX and aY using CRT
	a, success := chineseRemainder(aX, periodX, aY, periodY)
	if !success || a < 0 {
		return 0, 0, false // No valid solution found
	}

	// Calculate b from the first equation
	b := (machine.Prize.char - machine.A.charIncr*a) / machine.B.charIncr
	fmt.Printf("Combining solutions: aX=%d, periodX=%d, aY=%d, periodY=%d\n", aX, periodX, aY, periodY)
	fmt.Printf("CRT result: a=%d (success=%t)\n", a, success)

	// Verify solution
	if machine.A.lineIncr*a+machine.B.lineIncr*b != machine.Prize.line {
		return 0, 0, false // Doesn't satisfy Y equation
	}

	return a, b, true
}

// Helper function to compute CRT
func chineseRemainder(a1, n1, a2, n2 int) (int, bool) {
	gcd, x1, _ := extendedGCD(n1, n2)
	if (a2-a1)%gcd != 0 {
		return 0, false // No solution exists
	}

	// Solve for k such that: a1 + k * n1 ≡ a2 (mod n2)
	k := (a2 - a1) / gcd * x1 % (n2 / gcd)

	// Compute the combined solution
	result := a1 + k*n1
	mod := n1 / gcd * n2

	// Ensure result is positive
	if result < 0 {
		result += mod
	}

	return result % mod, true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func findMinimumCostDiophantine(machines []ClawMachine) int {
	totalCost := 0
	prizesWon := 0

	for _, machine := range machines {
		bestA, bestB, cost, canWin := findSolutionPart2(machine)
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
