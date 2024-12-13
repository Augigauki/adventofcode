package main

import (
	"fmt"
	"math/big"
)

type BigPos struct {
	line *big.Int
	char *big.Int
}

type BigButton struct {
	lineIncr *big.Int
	charIncr *big.Int
}

type BigClawMachine struct {
	A     BigButton
	B     BigButton
	Prize BigPos
	Pos   BigPos
}

func Part2(clawMachines []BigClawMachine) {
	fmt.Println("Part 2")
	for i, machine := range clawMachines {
		machine.Prize.line.Add(machine.Prize.line, big.NewInt(10000000000000))
		machine.Prize.char.Add(machine.Prize.char, big.NewInt(10000000000000))
		fmt.Printf(
			"Machine %v: Button A: X+%v, Y+%v, Button B: X+%v, Y+%v, Prize: (%v, %v)\n",
			i+1,
			machine.A.lineIncr, machine.A.charIncr,
			machine.B.lineIncr, machine.B.charIncr,
			machine.Prize.line, machine.Prize.char,
		)
		clawMachines[i] = machine
	}
	findMinimumCostDiophantine(clawMachines)
}

func findSolutionPart2(machine BigClawMachine) (*big.Int, *big.Int, *big.Int, bool) {
	a, b, canSolve := solveSystem(machine)
	if !canSolve {
		fmt.Println("No solution for this machine.")
		return nil, nil, nil, false
	}

	cost := BigCost(a, b)
	fmt.Printf("Solution found: a=%v, b=%v, cost=%v\n", a, b, cost)
	return a, b, cost, true
}

func solveSystem(machine BigClawMachine) (*big.Int, *big.Int, bool) {
	// Extract coefficients
	Ax, Bx, PrizeX := machine.A.charIncr, machine.B.charIncr, machine.Prize.char
	Ay, By, PrizeY := machine.A.lineIncr, machine.B.lineIncr, machine.Prize.line

	// Solve first equation for a in terms of b: a = (PrizeX - Bx * b) / Ax
	// Substitute into second equation to solve for b
	AxBy := new(big.Int).Mul(Ax, By)
	AyBx := new(big.Int).Mul(Ay, Bx)
	numerator := new(big.Int).Sub(new(big.Int).Mul(PrizeY, Ax), new(big.Int).Mul(PrizeX, Ay))
	denominator := new(big.Int).Sub(AxBy, AyBx)

	// Check if denominator divides numerator
	if new(big.Int).Mod(numerator, denominator).Cmp(big.NewInt(0)) != 0 {
		fmt.Println("No solution: numerator is not divisible by denominator")
		return nil, nil, false
	}

	// Compute b
	b := new(big.Int).Div(numerator, denominator)

	// Compute a from the first equation
	numeratorA := new(big.Int).Sub(PrizeX, new(big.Int).Mul(Bx, b))
	if new(big.Int).Mod(numeratorA, Ax).Cmp(big.NewInt(0)) != 0 {
		fmt.Println("No solution: numeratorA is not divisible by Ax")
		return nil, nil, false
	}
	a := new(big.Int).Div(numeratorA, Ax)

	// Verify solution
	yActual := new(big.Int).Add(new(big.Int).Mul(Ay, a), new(big.Int).Mul(By, b))
	xActual := new(big.Int).Add(new(big.Int).Mul(Ax, a), new(big.Int).Mul(Bx, b))
	if xActual.Cmp(PrizeX) != 0 || yActual.Cmp(PrizeY) != 0 {
		fmt.Printf("Validation failed: X_actual=%v, Y_actual=%v\n", xActual, yActual)
		return nil, nil, false
	}

	return a, b, true
}

func BigCost(a, b *big.Int) *big.Int {
	return new(big.Int).Add(new(big.Int).Mul(a, big.NewInt(3)), b)
}

func findMinimumCostDiophantine(machines []BigClawMachine) {
	totalCost := big.NewInt(0)
	prizesWon := 0

	for _, machine := range machines {
		bestA, bestB, cost, canWin := findSolutionPart2(machine)
		if canWin {
			fmt.Printf("Solution found for machine: A=%v, B=%v, Cost=%v\n", bestA, bestB, cost)
			totalCost.Add(totalCost, cost)
			prizesWon++
		} else {
			fmt.Println("No solution for this machine.")
		}
	}

	fmt.Printf("Won %d prizes for a total cost of %v tokens.\n", prizesWon, totalCost)
}
