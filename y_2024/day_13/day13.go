package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "example.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file: ", fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//var clawMachines []ClawMachine
	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	machinesLines := strings.Split(strings.Join(input, "\n"), "\n\n")
	var clawMachines []ClawMachine
	for _, machine := range machinesLines {
		clawMachine, err := parseClawMachine(machine)
		if err != nil {
			log.Fatal("Error parsing machine: ", err)
		}
		clawMachines = append(clawMachines, clawMachine)
	}

	/* fmt.Println("Claw Machines:")
	for _, cm := range clawMachines {
		fmt.Printf("%+v\n", cm)
	} */
	//Part1(clawMachines)
	Part2(clawMachines)
}

func parseClawMachine(machine string) (ClawMachine, error) {
	lines := strings.Split(machine, "\n")
	var clawMachine ClawMachine

	for _, line := range lines {
		if strings.HasPrefix(line, "Button A:") {
			button, err := parseButton(line)
			if err != nil {
				log.Fatal("Error parsing button: ", err)
				return clawMachine, err
			}
			clawMachine.A = button
		} else if strings.HasPrefix(line, "Button B:") {
			button, err := parseButton(line)
			if err != nil {
				log.Fatal("Error parsing button: ", err)
				return clawMachine, err
			}
			clawMachine.B = button
		} else if strings.HasPrefix(line, "Prize:") {
			pos, err := parsePos(line)
			if err != nil {
				log.Fatal("Error parsing pos: ", err)
				return clawMachine, err
			}
			clawMachine.Prize = pos
		}
		clawMachine.Pos = Pos{0, 0}
	}
	return clawMachine, nil
}

func parseButton(line string) (Button, error) {
	var button Button
	parts := strings.Split(strings.Split(line, ": ")[1], ", ")
	for _, part := range parts {
		if strings.Contains(part, "X+") {
			value, err := strconv.Atoi(strings.TrimPrefix(part, "X+"))
			if err != nil {
				return button, err
			}
			button.lineIncr = value
		} else if strings.Contains(part, "Y+") {
			value, err := strconv.Atoi(strings.TrimPrefix(part, "Y+"))
			if err != nil {
				return button, err
			}
			button.charIncr = value
		}
	}
	return button, nil
}

func parsePos(line string) (Pos, error) {
	var pos Pos
	parts := strings.Split(line, ", ")
	for _, part := range parts {
		if strings.Contains(part, "X=") {
			value, err := strconv.Atoi(strings.TrimPrefix(part, "Prize: X="))
			if err != nil {
				return pos, err
			}
			pos.line = value
		} else if strings.Contains(part, "Y=") {
			value, err := strconv.Atoi(strings.TrimPrefix(part, "Y="))
			if err != nil {
				return pos, err
			}
			pos.char = value
		}
	}
	return pos, nil
}
