package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var example = false

func main() {
	fileName := ""
	if example {
		fileName = "example.txt"
	} else {
		fileName = "input.txt"
	}
	secrets := parseFile(fileName)
	fmt.Println("Secrets:")
	for _, secret := range secrets {
		fmt.Println(secret)
	}
	sum := 0
	for _, secret := range secrets {
		for i := 0; i < 2000; i++ {

			secret = calcNextSecret(secret)
			//fmt.Printf("Next secret: ", secret)
			if i == 2000-1 {
				sum += secret
				fmt.Println("2000th secret: ", secret)
			}
		}
	}
	fmt.Println("Sum of 2000th secrets: ", sum)
}

func calcNextSecret(secret int) int {
	secret = mixSecret(secret, secret*64)
	secret = mixSecret(secret, secret/32)
	secret = mixSecret(secret, secret*2048)
	return secret
}

func mixSecret(secret, value int) int {
	newSecret := secret ^ value
	newSecret = pruneNumber(newSecret)
	return newSecret
}

func pruneNumber(number int) int {
	newNumber := number % 16777216
	return newNumber
}

func parseFile(fileName string) []int {
	secrets := []int{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error reading file: ", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Error parsing line: ", err)
		}
		secrets = append(secrets, num)
	}

	return secrets
}
