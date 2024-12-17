package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"

	"os"
	"strconv"
	"strings"
)

type Comp struct {
	regA    int
	regB    int
	regC    int
	program []int
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
				value, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Fatal("Error converting Register A value: ", err)
				}
				com.regA = value
			}
		}
		if strings.HasPrefix(line, "Register B") {
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				value, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Fatal("Error converting Register B value: ", err)
				}
				com.regB = value
			}
		}
		if strings.HasPrefix(line, "Register C") {
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				value, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Fatal("Error converting Register C value: ", err)
				}
				com.regC = value
			}
		}
		if strings.HasPrefix(line, "Program:") {
			parts := strings.Split(line, ": ")
			for _, part := range parts[1] {
				if part == ',' {
					continue
				}
				value, err := strconv.Atoi(string(part))
				if err != nil {
					log.Fatal("Error converting Program value: ", err)
				}
				com.program = append(com.program, value)
			}
		}
	}
	fmt.Printf("Computer:\nRegister A: %v\nRegister B: %v\nRegister C: %v\nProgram: %v\n\n", com.regA, com.regB, com.regC, com.program)
	part1(com)
}

func part1(com Comp) {
	fmt.Println("Part 1")
	output := []int{}
	for i := 0; i < len(com.program); {
		switch com.program[i] {
		case 0:
			operand := getOperandValue(com, com.program[i+1])
			com.regA = adv(int64(com.regA), operand)
			i += 2
		case 1:
			//operand := getOperandValue(com, com.program[i+1])
			com.regB = bxl(com.regB, com.program[i+1])
			i += 2
		case 2:
			operand := getOperandValue(com, com.program[i+1])
			com.regB = bst(operand)
			i += 2
		case 3:
			if com.regA == 0 {
				i += 2
			} else {
				//fmt.Println("Jumping to: ", com.program[i+1])
				i = com.program[i+1]
			}
		case 4:
			com.regB = bxc(com.regB, com.regC)
			i += 2
		case 5:
			operand := getOperandValue(com, com.program[i+1])
			output = append(output, out(operand))
			i += 2
		case 6:
			operand := getOperandValue(com, com.program[i+1])
			com.regB = bdv(int64(com.regA), operand)
			i += 2
		case 7:
			operand := getOperandValue(com, com.program[i+1])
			com.regC = cdv(int64(com.regA), operand)
			i += 2
		}

	}
	fmt.Println("Output: ", output)
	outputStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output)), ","), "[]")
	fmt.Println("Output as string: ", outputStr)

}

func getOperandValue(com Comp, operand int) int {
	switch operand {
	case 0:
		return operand
	case 1:
		return operand
	case 2:
		return operand
	case 3:
		return operand
	case 4:
		return com.regA
	case 5:
		return com.regB
	case 6:
		return com.regC
	default:
		return operand
	}
}

func adv(a int64, combo int) int {
	if combo < 0 {
		fmt.Println("Error: combo cannot be negative")
		return 0
	}

	// Convert numerator to big.Int
	numerator := big.NewInt(a)

	// Calculate denominator as 2^combo
	denominator := new(big.Int).Lsh(big.NewInt(1), uint(combo))

	// Perform division: numerator / denominator
	result := new(big.Int).Div(numerator, denominator)

	// Safely convert result to int, checking for overflow
	if !result.IsInt64() {
		fmt.Println("Error: result exceeds int64 range")
		return 0
	}

	// Convert to int64 and then to int
	intResult := int(result.Int64())

	// Ensure intResult fits in int type safely
	if intResult > math.MaxInt || intResult < math.MinInt {
		fmt.Println("Error: result exceeds int range")
		return 0
	}

	return intResult
}

func bxl(b, operand int) int {
	return b ^ operand
}

func bst(combo int) int {
	return combo & 7
}

func jnz(rega, literal int) int {
	if rega != 0 {
		return literal
	}
	return -1
}

func bxc(b, c int) int {
	return b ^ c
}

func out(combo int) int {
	return combo & 7
}

func bdv(a int64, combo int) int {
	if combo < 0 {
		fmt.Println("Error: combo cannot be negative")
		return 0
	}

	// Convert numerator to big.Int
	numerator := big.NewInt(a)

	// Calculate denominator as 2^combo
	denominator := new(big.Int).Lsh(big.NewInt(1), uint(combo))

	// Perform division: numerator / denominator
	result := new(big.Int).Div(numerator, denominator)

	// Safely convert result to int, checking for overflow
	if !result.IsInt64() {
		fmt.Println("Error: result exceeds int64 range")
		return 0
	}

	// Convert to int64 and then to int
	intResult := int(result.Int64())

	// Ensure intResult fits in int type safely
	if intResult > math.MaxInt || intResult < math.MinInt {
		fmt.Println("Error: result exceeds int range")
		return 0
	}

	return intResult
}

func cdv(a int64, combo int) int {
	if combo < 0 {
		fmt.Println("Error: combo cannot be negative")
		return 0
	}

	// Convert numerator to big.Int
	numerator := big.NewInt(a)

	// Calculate denominator as 2^combo
	denominator := new(big.Int).Lsh(big.NewInt(1), uint(combo))

	// Perform division: numerator / denominator
	result := new(big.Int).Div(numerator, denominator)

	// Safely convert result to int, checking for overflow
	if !result.IsInt64() {
		fmt.Println("Error: result exceeds int64 range")
		return 0
	}

	// Convert to int64 and then to int
	intResult := int(result.Int64())

	// Ensure intResult fits in int type safely
	if intResult > math.MaxInt || intResult < math.MinInt {
		fmt.Println("Error: result exceeds int range")
		return 0
	}

	return intResult
}
