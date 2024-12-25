package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* AND gates output 1 if both inputs are 1; if either input is 0, these gates output 0.
OR gates output 1 if one or both inputs is 1; if both inputs are 0, these gates output 0.
XOR gates output 1 if the inputs are different; if the inputs are the same, these gates output 0. */

type Wire struct {
	name  string
	value int
}

type Gate struct {
	W1       string
	W2       string
	W3       string
	Operator string
}

type wireMap map[string]int

var example = true

func main() {
	fileName := ""
	if example {
		fileName = "input.txt"
	} else {
		fileName = "input.txt"
	}
	wires, gates := parseFile(fileName)

	simlulateCircuit(&wires, gates)
	for name, value := range wires {
		fmt.Printf("Wire: %s - Value: %v\n", name, value)
	}
	zwires := []string{}
	for _, gate := range gates {
		if strings.HasPrefix(gate.W3, "z") {
			zwires = append(zwires, gate.W3)
		}
	}
	fmt.Println("Z wires:")
	sort.Slice(zwires, func(i, j int) bool {
		return zwires[i] < zwires[j]
	})
	bit := ""
	for i := len(zwires) - 1; i >= 0; i-- {
		fmt.Printf("Wire: %s - Value: %v\n", zwires[i], zwires[i])
		bit += strconv.Itoa(wires[zwires[i]])
	}
	fmt.Println("Bit: ", bit)
	decimalValue, err := strconv.ParseInt(bit, 2, 64)
	if err != nil {
		log.Fatal("Error converting binary to decimal: ", err)
	}
	fmt.Println("Decimal Value: ", decimalValue)
}

func simlulateCircuit(wires *wireMap, gates []Gate) {
	fmt.Println("Simulating circuit...")
	fmt.Println("Gates:")
	for _, gate := range gates {
		fmt.Printf("Gate: %v\n", gate)
	}
	for {
		changed := false
		for i := range gates {
			gate := &gates[i]
			if (*wires)[gate.W3] == -1 && (*wires)[gate.W1] != -1 && (*wires)[gate.W2] != -1 {
				if gate.Operator == "AND" {
					(*wires)[gate.W3] = calcAND((*wires)[gate.W1], (*wires)[gate.W2])
				} else if gate.Operator == "OR" {
					(*wires)[gate.W3] = calcOR((*wires)[gate.W1], (*wires)[gate.W2])
				} else if gate.Operator == "XOR" {
					(*wires)[gate.W3] = calcXOR((*wires)[gate.W1], (*wires)[gate.W2])
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}

func calcAND(w1, w2 int) int {
	if w1 == 1 && w2 == 1 {
		return 1
	} else {
		return 0
	}
}

func calcOR(w1, w2 int) int {
	if w1 == 1 || w2 == 1 {
		return 1
	} else {
		return 0
	}
}

func calcXOR(w1, w2 int) int {
	if w1 != w2 {
		return 1
	} else {
		return 0
	}
}

func findOrCreateWire(wires *[]Wire, name string) Wire {
	for _, w := range *wires {
		if w.name == name {
			return w
		}
	}
	newWire := Wire{name: name, value: -1}
	*wires = append(*wires, newWire)
	return newWire
}

func parseFile(fileName string) (wireMap, []Gate) {
	wires := wireMap{}
	gates := []Gate{}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file: ", fileName)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ":") {
			wireParts := strings.Split(line, ": ")
			wireVal, _ := strconv.Atoi(wireParts[1])
			wire := Wire{name: wireParts[0], value: wireVal}
			wires[wire.name] = wire.value
			//wires = append(wires, wire)
		}
		if strings.Contains(line, "->") {
			parts := strings.Fields(line)
			w1Name, op, w2Name, _, w3Name := parts[0], parts[1], parts[2], parts[3], parts[4]
			_, ok := wires[w1Name]
			if !ok {
				wires[w1Name] = -1
			}
			_, ok = wires[w2Name]
			if !ok {
				wires[w2Name] = -1
			}
			_, ok = wires[w3Name]
			if !ok {
				wires[w3Name] = -1
			}
			//w1 := findOrCreateWire(&wires, w1Name)
			//w2 := findOrCreateWire(&wires, w2Name)
			//w3 := findOrCreateWire(&wires, w3Name)

			/* 	if op == "AND" {
				w3.value = calcAND(w1, w2)
			} else if op == "OR" {
				w3.value = calcOR(w1, w2)
			} else if op == "XOR" {
				w3.value = calcXOR(w1, w2)
			} */

			gate := Gate{W1: w1Name, Operator: op, W2: w2Name, W3: w3Name}
			gates = append(gates, gate)
		}
	}

	return wires, gates
}
