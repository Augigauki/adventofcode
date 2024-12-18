package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"

	"os"
	"strconv"
	"strings"
)

type Comp struct {
	regA    int64
	regB    int64
	regC    int64
	program []int64
}

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	com := Comp{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Register A") {
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				value, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					log.Fatal("Error converting Register A value: ", err)
				}
				com.regA = value
			}
		}
		if strings.HasPrefix(line, "Register B") {
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				value, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					log.Fatal("Error converting Register B value: ", err)
				}
				com.regB = value
			}
		}
		if strings.HasPrefix(line, "Register C") {
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				value, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					log.Fatal("Error converting Register C value: ", err)
				}
				com.regC = value
			}
		}
		if strings.HasPrefix(line, "Program:") {
			parts := strings.Split(line, ": ")
			for _, part := range strings.Split(parts[1], ",") { // Correct splitting
				value, err := strconv.ParseInt(part, 10, 64)
				if err != nil {
					log.Fatal("Error converting Program value: ", err)
				}
				com.program = append(com.program, value)
			}
		}

	}
	fmt.Printf("Computer:\nRegister A: %v\nRegister B: %v\nRegister C: %v\nProgram: %v\n\n", com.regA, com.regB, com.regC, com.program)
	output := part1(com)
	fmt.Println("Output: ", output)
	smallestAReg := findRegisterA(com.program)
	fmt.Println("Smallest Register A: ", smallestAReg)

}

func part1(com Comp) []int64 {
	fmt.Println("Part 1")
	output := []int64{}
	for i := 0; i < len(com.program); {
		switch com.program[i] {
		case 0:
			operand := getOperandValue(com, com.program[i+1])
			com.regA = adv(com.regA, operand)
			i += 2
		case 1:
			com.regB = bxl(com.regB, com.program[i+1])
			i += 2
		case 2:
			operand := getOperandValue(com, com.program[i+1])
			com.regB = bst(operand)
			i += 2
		case 3:
			if com.regA != 0 {
				jumpTo := int(com.program[i+1])
				if jumpTo >= len(com.program) {
					return output // Halt if jump goes beyond program bounds
				}
				i = jumpTo
			} else {
				i += 2
			}
		case 4:
			com.regB = bxc(com.regB, com.regC)
			i += 2
		case 5:
			operand := getOperandValue(com, com.program[i+1])
			val := out(operand)
			output = append(output, val)
			i += 2
		case 6:
			operand := getOperandValue(com, com.program[i+1])
			com.regB = adv(com.regA, operand)
			i += 2
		case 7:
			operand := getOperandValue(com, com.program[i+1])
			com.regC = adv(com.regA, operand)
			i += 2
		default:
			fmt.Println("Invaling instruction, returning output i guess")
			return output
		}

	}
	fmt.Println("End of program - output: ", output)
	return output
}

func findRegisterA(program []int64) int64 {
	return findRegisterARecursive(program, 0, 0)
}

func findRegisterARecursive(program []int64, a int64, i int) int64 {
	tempCom := Comp{regA: a, regB: 0, regC: 0, program: program}
	output := part1(tempCom)

	if compareSlices(output, program) { // Base case: full match
		return a
	}

	if i == 0 || compareSlices(output, program[len(program)-i:]) { // Recursive step
		for n := int64(0); n < 8; n++ {
			result := findRegisterARecursive(program, 8*a+n, i+1)
			if result != -1 { // Found a valid value
				return result
			}
		}
	}

	return -1 // No valid value found in this branch
}

// Compare program outputs to expected outputs
func compareSlices(outputs []int64, program []int64) bool {
	if len(outputs) != len(program) {
		return false
	}
	for i := 0; i < len(outputs); i++ {
		if outputs[i] != program[i] {
			return false
		}
	}
	return true
}

func getOperandValue(com Comp, operand int64) int64 {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return com.regA
	case 5:
		return com.regB
	case 6:
		return com.regC
	default:
		fmt.Println("Invalid operand value")
		return 0
	}
}

func adv(a int64, combo int64) int64 {
	if combo < 0 {
		fmt.Println("Error: combo cannot be negative")
		return 0
	}

	numerator := big.NewInt(a)

	denominator := new(big.Int).Lsh(big.NewInt(1), uint(combo))

	result := new(big.Int).Div(numerator, denominator)

	if !result.IsInt64() {
		fmt.Println("Error: result exceeds int64 range")
		return 0
	}
	return result.Int64()

}

func bxl(b, operand int64) int64 {
	return b ^ operand
}

func bst(combo int64) int64 {
	return combo & 7
}

func bxc(b, c int64) int64 {
	return b ^ c
}

func out(combo int64) int64 {
	return combo & 7
}
