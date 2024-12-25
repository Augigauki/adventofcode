package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var example = true

type Computer struct {
	Code        string
	ConnectedTo []string
}

type Compmap map[string][]string

type LANParty struct {
	Computers []string
}

func main() {
	fileName := ""
	if example {
		fileName = "example.txt"
	} else {
		fileName = "input.txt"
	}
	computers := parseFile(fileName)
	fmt.Println("Computers:")
	for code, computer := range computers {
		fmt.Printf("Code: %s - Connected to: %v\n", code, computer)
	}
	lanParties := findLANParties(computers)
	fmt.Println("LAN Parties:")
	for _, party := range lanParties {
		fmt.Println(party)
	}
	tPartiesCount := countLANPartiesWithTComputer(lanParties)
	fmt.Println("Number of LAN parties with a computer starting with 't': ", tPartiesCount)
}

func countLANPartiesWithTComputer(lanParties []LANParty) int {
	count := 0
	for _, party := range lanParties {
		for _, computer := range party.Computers {
			if strings.HasPrefix(computer, "t") {
				count++
				break
			}
		}
	}
	return count
}

func findLANParties(computers Compmap) []LANParty {
	lanParties := []LANParty{}

	for comp1, connections1 := range computers {
		for _, comp2 := range connections1 {
			for _, comp3 := range computers[comp2] {
				if comp1 != comp3 && comp3 != comp2 && contains(computers[comp1], comp3) {
					// Check if this LAN party already exists
					exists := false
					for _, party := range lanParties {
						if contains(party.Computers, comp1) && contains(party.Computers, comp2) && contains(party.Computers, comp3) {
							exists = true
							break
						}
					}
					if !exists {
						lanParties = append(lanParties, LANParty{Computers: []string{comp1, comp2, comp3}})
					}
				}
			}
		}
	}
	for i := range lanParties {
		sort.Strings(lanParties[i].Computers)
	}
	sort.Slice(lanParties, func(i, j int) bool {
		return strings.Join(lanParties[i].Computers, "") < strings.Join(lanParties[j].Computers, "")
	})
	return lanParties
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func parseFile(fileName string) Compmap {
	computerMap := make(Compmap)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error reading file: ", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		code := strings.Split(line, "-")[0]
		remote := strings.Split(line, "-")[1]
		fmt.Printf("Code: %s - Remote: %s\n", code, remote)

		if _, ok := computerMap[code]; ok {
			foundConnection := false
			for _, c := range computerMap[code] {
				if c == remote {
					foundConnection = true
					break
				}
			}
			if !foundConnection {
				computerMap[code] = append(computerMap[code], remote)
			}
		} else {
			computerMap[code] = []string{remote}
		}

		if _, ok := computerMap[remote]; ok {
			foundConnection := false
			for _, c := range computerMap[remote] {
				if c == code {
					foundConnection = true
					break
				}
			}
			if !foundConnection {
				computerMap[remote] = append(computerMap[remote], code)
			}
		} else {
			computerMap[remote] = []string{code}
		}

	}

	return computerMap
}
